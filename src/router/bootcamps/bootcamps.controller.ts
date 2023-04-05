import type { NextFunction, Request, Response } from 'express'
import asyncHandler from 'express-async-handler'
import httpStatus from 'http-status'

import { APIError, prisma } from '@/lib'
import type { CreateBootcampType, GetBootcampType } from './bootcamps.schemas'

export const createBootcamp = asyncHandler(
    async (req: Request<any, any, CreateBootcampType>, res: Response, next: NextFunction) => {
        const record = await prisma.bootcamps.findFirst({
            where: { name: req.body.name },
        })

        if (record !== null) {
            return next(APIError.conflict('bootcamp with that name already exists'))
        }

        const bootcamp = await prisma.bootcamps.create({
            data: { ...req.body },
        })

        res.status(httpStatus.CREATED).json({
            success: true,
            message: 'bootcamp created',
            data: bootcamp,
        })
    }
)

export const getBootcamps = asyncHandler(async (req: Request, res: Response) => {
    const bootcamps = await prisma.bootcamps.findMany()

    res.status(httpStatus.OK).json({
        success: true,
        message: 'bootcamps retrieved',
        results: bootcamps.length,
        data: bootcamps,
    })
})

export const getBootcamp = asyncHandler(
    async (req: Request<GetBootcampType>, res: Response, next: NextFunction) => {
        const { id } = req.params

        const bootcamp = await prisma.bootcamps.findUnique({
            where: { id },
        })
        if (bootcamp === null) {
            next(APIError.notFound('bootcamp not found'))
            return
        }

        res.status(httpStatus.OK).json({
            success: true,
            data: bootcamp,
        })
    }
)

export const updateBootcamp = asyncHandler(
    async (req: Request<GetBootcampType, any, Partial<CreateBootcampType>>, res: Response) => {
        const { id } = req.params

        const bootcamp = await prisma.bootcamps.update({
            where: { id },
            data: { ...req.body },
        })
        res.status(httpStatus.OK).json({
            success: true,
            data: bootcamp,
        })
    }
)

export const deleteBootcamp = asyncHandler(async (req: Request<GetBootcampType>, res: Response) => {
    const { id } = req.params

    await prisma.bootcamps.delete({
        where: { id },
    })

    res.status(httpStatus.OK).json({
        success: true,
        message: 'bootcamp deleted',
    })
})
