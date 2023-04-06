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

export const getBootcamps = asyncHandler(
    async (req: Request<any, any, any, GetBootcampsQueryType>, res: Response) => {
        // const { fields, sort, limit = 0, page, name, housing } = req.query
        // console.log({ fields, sort, limit, page, name, housing })

        const bootcamps = await services.findMany()

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
