import type { Request, Response, NextFunction } from 'express'
import { PrismaClientValidationError } from '@prisma/client/runtime/library'
import httpStatus from 'http-status'
import config from 'config'

import { type APIError, logger } from '@/lib'

export const globalErrorHandler = (
    err: APIError,
    req: Request,
    res: Response,
    next: NextFunction
): void => {
    const error = { ...err }
    error.message = err.message
    const ENV = config.get<string>('NODE_ENV')

    if (err instanceof PrismaClientValidationError) {
        error.statusCode = httpStatus.BAD_REQUEST
        // The follwoing will parse the prisma error message and find all the missing fields
        // and return them in a string.
        // e.g. 'Argument name for data.name is missing, Argument email for data.email is missing'
        error.message =
            err.stack
                ?.match(/Argument\s\w+\sfor\sdata\.\w+\sis\smissing/)
                ?.filter(text => text.startsWith('Argument'))
                .join(', ') ?? err.message
    }

    if (ENV === 'development') {
        error.stack = err.stack
        error.statusCode = err.statusCode ?? httpStatus.INTERNAL_SERVER_ERROR
        res.status(error.statusCode).json({
            success: false,
            message: error.message,
            error,
        })
    } else {
        logger.error('💥 ERROR PROD: ', err)
        res.status(error.statusCode).json({
            success: false,
            name: error.name,
            message: error.message,
            status: error.statusCode,
        })
    }
}
