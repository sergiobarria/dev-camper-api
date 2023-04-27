import type { NextFunction, Request, Response } from 'express'
import asyncHandler from 'express-async-handler'
import httpStatus from 'http-status'

import { APIError, geocoder, prisma } from '@/lib'
import * as services from './bootcamps.services'
import type {
    CreateBootcampType,
    GetBootcampType,
    GetBootcampsInRadiusType,
    GetBootcampsQueryType,
} from './bootcamps.schemas'
import { type Prisma } from '@prisma/client'

export const createBootcamp = asyncHandler(
    async (req: Request<any, any, CreateBootcampType>, res: Response, next: NextFunction) => {
        const record = await services.findFirst({ name: req.body.name }) // check if bootcamp already exists

        if (record !== null) {
            return next(APIError.conflict('bootcamp with that name already exists'))
        }

        const bootcamp = await services.createOne(req.body)

        res.status(httpStatus.CREATED).json({
            success: true,
            message: 'bootcamp created',
            data: bootcamp,
        })
    }
)

function mapQueryToWhere(query: GetBootcampsQueryType): Prisma.bootcampsWhereInput {
    const where: Prisma.bootcampsWhereInput = {}
    where.AND = []

    for (const [key, value] of Object.entries(query)) {
        if (value !== undefined) {
            switch (key) {
                case 'name':
                    where.AND.push({ name: { contains: query.name } })
                    break
                case 'housing':
                case 'jobAssistance':
                case 'jobGuarantee':
                case 'acceptGi':
                    where.AND.push({ [key]: value === 'true' })
                    break
                default:
                    break
            }
        }
    }

    return where
}

export const getBootcamps = asyncHandler(
    async (req: Request<any, any, any, GetBootcampsQueryType>, res: Response) => {
        console.log({ query: req.query })
        const where = mapQueryToWhere(req.query)
        console.log({ where })

        const bootcamps = await services.findMany({
            where,
        })

        res.status(httpStatus.OK).json({
            success: true,
            message: 'bootcamps retrieved',
            results: bootcamps.length,
            data: bootcamps,
        })
    }
)

export const getBootcamp = asyncHandler(
    async (req: Request<GetBootcampType>, res: Response, next: NextFunction) => {
        const bootcamp = await services.findOneById(req.params.id)

        if (bootcamp === null) {
            return next(APIError.notFound('bootcamp not found'))
        }

        res.status(httpStatus.OK).json({
            success: true,
            data: bootcamp,
        })
    }
)

export const updateBootcamp = asyncHandler(
    async (
        req: Request<GetBootcampType, any, Partial<CreateBootcampType>>,
        res: Response,
        next: NextFunction
    ) => {
        const bootcamp = services.updateOne(req.params.id, req.body)

        if (bootcamp === null) {
            return next(APIError.notFound('bootcamp not found'))
        }

        res.status(httpStatus.OK).json({
            success: true,
            data: bootcamp,
        })
    }
)

export const deleteBootcamp = asyncHandler(async (req: Request<GetBootcampType>, res: Response) => {
    await services.deleteOne(req.params.id)

    res.status(httpStatus.OK).json({
        success: true,
        message: 'bootcamp deleted',
    })
})

/**
 * @desc    Get bootcamps within a radius
 * @route   GET /api/v1/bootcamps/radius/:zipcode/:distance/:unit
 */
export const getBootcampsInRadius = asyncHandler(
    async (req: Request<GetBootcampsInRadiusType>, res: Response) => {
        const { zipcode, distance, unit = 'mi' } = req.params

        // Get lat/lng from geocoder
        const loc = await geocoder.geocode(zipcode)
        const { latitude, longitude } = loc[0]

        // Calc radius using radians
        // Divide dist by radius of Earth - Earth Radius = 3,963 mi / 6,378 km
        const radius = unit === 'mi' ? Number(distance) / 3963 : Number(distance) / 6378

        const bootcamps = await prisma.bootcamps.findRaw({
            filter: {
                location: {
                    $geoWithin: {
                        $centerSphere: [[longitude, latitude], radius] as [number[], number],
                    },
                },
            },
        })

        res.status(httpStatus.OK).json({
            success: true,
            message: 'bootcamps retrieved',
            results: bootcamps.length,
            data: bootcamps,
        })
    }
)
