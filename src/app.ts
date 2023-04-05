import express, { type Request, type Response, type NextFunction } from 'express'
import config from 'config'

import { router } from './router'
import { globalErrorHandler, morganMiddleware, bootcampMiddleware } from './middleware'
import { APIError, prisma } from './lib'

export const app = express()
const env = config.get('NODE_ENV')

// ===== Apply middlewares 👇🏼 =====
app.use(express.json())

if (env === 'development') {
    app.use(morganMiddleware)
}

// ===== Apply routes 👇🏼 =====
app.use('/api/v1', router)

app.all('*', (req: Request, _: Response, next: NextFunction) => {
    next(APIError.notFound(`Can't find ${req.originalUrl} on this server!`))
})

// ===== Apply error handlers 👇🏼 =====
app.use(globalErrorHandler)

// ===== Register Prisma middleware 👇🏼 =====
prisma.$use(bootcampMiddleware)
