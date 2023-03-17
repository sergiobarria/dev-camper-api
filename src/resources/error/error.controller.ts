import type { FastifyError, FastifyInstance, FastifyReply, FastifyRequest } from 'fastify'
import httpStatus from 'http-status'
import config from 'config'

interface ErrorReply {
    fieldName?: string
    statusCode: number
    error: string
    message: string
}

type ErrorTypes = Record<string, (fieldName?: string) => ErrorReply>

const errors: ErrorTypes = {
    MongoServerError: (value) => ({
        statusCode: httpStatus.CONFLICT,
        error: 'MongoServerError',
        message: value !== undefined ? `Duplicate key: ${value}` : 'Duplicate key'
    }),
    CastError: (value) => ({
        statusCode: httpStatus.BAD_REQUEST,
        error: 'CastError',
        message: value !== undefined ? `Invalid ${value}` : 'Invalid value'
    }),
    ValidationError: (value) => ({
        statusCode: httpStatus.BAD_REQUEST,
        error: 'ValidationError',
        message: value !== undefined ? `Invalid ${value}` : 'Invalid value'
    }),
    JsonWebTokenError: () => ({
        statusCode: httpStatus.UNAUTHORIZED,
        error: 'JsonWebTokenError',
        message: 'Invalid token'
    }),
    TokenExpiredError: () => ({
        statusCode: httpStatus.UNAUTHORIZED,
        error: 'TokenExpiredError',
        message: 'Token expired'
    }),
    default: () => ({
        statusCode: httpStatus.INTERNAL_SERVER_ERROR,
        error: 'InternalError',
        message: 'Internal error'
    })
}

function handleDuplicateKeyError(err: FastifyError): ErrorReply {
    const fieldName = err.message.match(/(["'])(\\?.)*?\1/)?.[0] ?? ''
    const error = errors[err.name](fieldName)
    return error
}

function handleValidationError(err: FastifyError): ErrorReply {
    const fieldName = err.message.match(/(["'])(\\?.)*?\1/)?.[0] ?? ''
    const error = errors[err.name](fieldName)
    return error
}

function handleCastError(err: FastifyError): ErrorReply {
    const fieldName = err.message.match(/(["'])(\\?.)*?\1/)?.[0] ?? ''
    const error = errors[err.name](fieldName)
    return error
}

export async function globalErrorHandler(
    this: FastifyInstance,
    error: FastifyError,
    _: FastifyRequest,
    reply: FastifyReply
): Promise<void> {
    const env = config.get<string>('NODE_ENV')
    let err = { ...error }
    let mappedError = errors.default()
    err.message = error.message
    err.statusCode = error.statusCode ?? httpStatus.INTERNAL_SERVER_ERROR

    if (error.name === 'MongoServerError') {
        mappedError = handleDuplicateKeyError(error)
        err = { ...err, ...mappedError }
    }

    if (error.name === 'ValidationError') {
        mappedError = handleValidationError(error)
        err = { ...err, ...mappedError }
    }

    if (error.name === 'CastError') {
        mappedError = handleCastError(error)
        err = { ...err, ...mappedError }
    }

    if (env === 'development') {
        err.stack = error.stack
        return await reply.code(err.statusCode as number).send(err)
    } else {
        return await reply.code(err.statusCode as number).send({
            name: err.name,
            message: err.message,
            statusCode: err.statusCode
        })
    }
}
