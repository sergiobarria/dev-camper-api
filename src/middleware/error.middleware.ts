import type { Request, Response, NextFunction } from 'express'
import httpStatus from 'http-status'
import config from 'config'

import { type APIError, logger } from '@/lib'

export const globalErrorHandler = (
    err: APIError,
    req: Request,
    res: Response,
    next: NextFunction
): void => {
    let error = { ...err }
    error.message = err.message
    const ENV = config.get<string>('NODE_ENV')

    if (ENV === 'development') {
        error.stack = err.stack
        error.statusCode = err.statusCode ?? httpStatus.INTERNAL_SERVER_ERROR
        res.status(error.statusCode).json({
            success: false,
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
