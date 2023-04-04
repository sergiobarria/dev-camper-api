import express, { type Request, type Response } from 'express'

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

export { router }
