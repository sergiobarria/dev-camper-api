import express, { type Request, type Response } from 'express'

import { bootcampsRouter } from './bootcamps/bootcamps.routes'

const router = express.Router()

router.get('/healthcheck', (_: Request, res: Response) => {
    res.status(200).json({
        succes: true,
        message: 'server is up and running',
        details: {
            version: '1.0.0',
            name: 'DevCamper API',
            uptime: process.uptime(),
            memory: process.memoryUsage(),
        },
    })
})

router.use('/bootcamps', bootcampsRouter)

export { router }
