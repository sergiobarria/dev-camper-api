import type { FastifyRequest, FastifyReply } from 'fastify'
import httpStatus from 'http-status'

import { Bootcamp } from '@/models/bootcamp.model'
import type { BootcampInputs } from './bootcamps.schemas'
import { APIError } from '@/utils'

/**
 * @desc: Get all bootcamps
 * @route: GET /bootcamps
 * @access: Public
 */
export async function getBootcampsHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    const bootcamps = await Bootcamp.find()

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
        data: {}
    })
}
