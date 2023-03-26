import type { FastifyRequest, FastifyReply } from 'fastify'
import httpStatus from 'http-status'

import { Bootcamp, BootcampModel, IBootcamp } from '@/models/bootcamp.model'
import type { BootcampInputs, BootcampWithinRadiusInputs } from './bootcamps.schemas'
import { APIError, geocoder } from '@/utils'
import { APIFeatures } from '@/utils/apiFeatures'
import { Document } from 'mongoose'

/**
 * @desc: Get all bootcamps
 * @route: GET /bootcamps
 * @access: Public
 */
export async function getBootcampsHandler(
    request: FastifyRequest<{ Querystring: Record<string, any> }>,
    reply: FastifyReply
): Promise<void> {
    const queryStr = request.query
    console.log('queryStr', queryStr)

    const features = new APIFeatures<IBootcamp>(Bootcamp, queryStr)
    const bootcamps = await features.query

    return await reply.code(httpStatus.OK).send({
        success: true,
        count: bootcamps.length,
        data: bootcamps
    })
}

/**
 * @desc: Get single bootcamp
 * @route: GET /bootcamps/:id
 * @access: Public
 */
export async function getBootcampHandler(
    request: FastifyRequest<{ Params: BootcampInputs['params'] }>,
    reply: FastifyReply
): Promise<void> {
    const { id } = request.params

    const bootcamp = await Bootcamp.findById(id)

    if (bootcamp === null) {
        throw APIError.notFound(`Bootcamp not found with id of ${id}`)
    }

    return await reply.code(httpStatus.OK).send({
        success: true,
        data: bootcamp
    })
}

/**
 * @desc: Create new bootcamp
 * @route: POST /bootcamps
 * @access: Private
 */
export async function createBootcampHandler(
    request: FastifyRequest<{ Body: BootcampInputs['body'] }>,
    reply: FastifyReply
): Promise<void> {
    const { body } = request

    // NOTE: I'm passing the body directly to the model, because I'm already validating the body with a Zod schema in the route handler
    const bootcamp = await Bootcamp.create(body)

    return await reply.code(httpStatus.CREATED).send({
        success: true,
        data: bootcamp
    })
}

/**
 * @desc: Update bootcamp
 * @route: PATCH /bootcamps/:id
 * @access: Private
 */
export async function updateBootcampHandler(
    request: FastifyRequest<{ Body: Partial<BootcampInputs['body']>; Params: BootcampInputs['params'] }>,
    reply: FastifyReply
): Promise<void> {
    const { id } = request.params
    const { body } = request

    const updatedBootcamp = await Bootcamp.findByIdAndUpdate(id, body, { new: true, runValidators: true })

    if (updatedBootcamp === null) {
        throw APIError.notFound(`Bootcamp not found with id of ${id}`)
    }

    return await reply.code(httpStatus.OK).send({
        success: true,
        data: updatedBootcamp
    })
}

/**
 * @desc: Delete bootcamp
 * @route: DELETE /bootcamps/:id
 * @access: Private
 */
export async function deleteBootcampHandler(
    request: FastifyRequest<{ Params: BootcampInputs['params'] }>,
    reply: FastifyReply
): Promise<void> {
    const { id } = request.params

    const bootcamp = await Bootcamp.findByIdAndDelete(id)

    if (bootcamp === null) {
        throw APIError.notFound(`Bootcamp not found with id of ${id}`)
    }

    return await reply.code(httpStatus.OK).send({
        success: true,
        data: null
    })
}

/**
 * @desc: Get bootcamps within a radius
 * @route: GET /bootcamps/radius/:zipcode/:distance
 * @access: Public
 */
export async function getBootcampsInRadiusHandler(
    request: FastifyRequest<{ Params: BootcampWithinRadiusInputs['params'] }>,
    reply: FastifyReply
): Promise<void> {
    const { zipcode, distance } = request.params

    // Get lat/lng from geocoder
    const loc = await geocoder.geocode(zipcode)
    const { latitude, longitude } = loc[0]

    // Calc radius using radians (divide distance by radius of Earth)
    // Earth Radius = 3,963 mi / 6,378 km
    const radius = Number(distance) / 3963

    const bootcamps = await Bootcamp.find({
        location: { $geoWithin: { $centerSphere: [[longitude, latitude], radius] } }
    })

    return await reply.code(httpStatus.OK).send({
        success: true,
        count: bootcamps.length,
        data: bootcamps
    })
}
