import httpStatus from 'http-status'

enum ErrorStatus {
    Fail = 'fail',
    Error = 'error',
}

export class APIError extends Error {
    public status: ErrorStatus
    public statusCode: number
    public isOperational: boolean

    constructor(
        message: string,
        status: ErrorStatus = ErrorStatus.Error,
        statusCode: number = httpStatus.INTERNAL_SERVER_ERROR,
        isOperational: boolean = true
    ) {
        super(message)

        this.name = this.constructor.name
        this.status = status
        this.statusCode = statusCode
        this.isOperational = isOperational

        Error.captureStackTrace(this)
    }

    static badRequest(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.BAD_REQUEST)
    }

    static unauthorized(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.UNAUTHORIZED)
    }

    static forbidden(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.FORBIDDEN)
    }

    static notFound(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.NOT_FOUND)
    }

    static conflict(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.CONFLICT)
    }

    static tooMany(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.TOO_MANY_REQUESTS)
    }

    static internal(message: string): APIError {
        return new APIError(message, ErrorStatus.Error, httpStatus.INTERNAL_SERVER_ERROR)
    }

    static notImplemented(message: string): APIError {
        return new APIError(message, ErrorStatus.Error, httpStatus.NOT_IMPLEMENTED)
    }

    static badGateway(message: string): APIError {
        return new APIError(message, ErrorStatus.Error, httpStatus.BAD_GATEWAY)
    }

    static serviceUnavailable(message: string): APIError {
        return new APIError(message, ErrorStatus.Error, httpStatus.SERVICE_UNAVAILABLE)
    }

    static gatewayTimeout(message: string): APIError {
        return new APIError(message, ErrorStatus.Error, httpStatus.GATEWAY_TIMEOUT)
    }

    static dbValidationError(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.UNPROCESSABLE_ENTITY)
    }

    static jwtExpired(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.UNAUTHORIZED)
    }

    static jwtInvalid(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.UNAUTHORIZED)
    }

    static jwtMissing(message: string): APIError {
        return new APIError(message, ErrorStatus.Fail, httpStatus.UNAUTHORIZED)
    }
}
