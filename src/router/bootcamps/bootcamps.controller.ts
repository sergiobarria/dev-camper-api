import type { Request, Response } from 'express'
import asyncHandler from 'express-async-handler'
import httpStatus from 'http-status'

export const createBootcamp = asyncHandler(async (req: Request, res: Response) => {
    res.status(httpStatus.CREATED).json({
        success: true,
        message: 'Create new bootcamp',
    })
})

export const getBootcamps = asyncHandler(async (req: Request, res: Response) => {
    res.status(httpStatus.OK).json({
        success: true,
        message: 'Get all bootcamps',
    })
})

export const getBootcamp = asyncHandler(async (req: Request, res: Response) => {
    res.status(httpStatus.OK).json({
        success: true,
        message: `Get bootcamp ${req.params.id}`,
    })
})

export const updateBootcamp = asyncHandler(async (req: Request, res: Response) => {
    res.status(httpStatus.OK).json({
        success: true,
        message: `Update bootcamp ${req.params.id}`,
    })
})

export const deleteBootcamp = asyncHandler(async (req: Request, res: Response) => {
    res.status(httpStatus.OK).json({
        success: true,
        message: `Delete bootcamp ${req.params.id}`,
    })
})
