import type { FastifyInstance } from 'fastify'

import { bootcampsRouter } from '@/resources/bootcamps/bootcamps.routes'

export async function router(app: FastifyInstance): Promise<void> {
    // Register routes resources here 👇🏼
    void app.register(bootcampsRouter, { prefix: '/bootcamps' })

    // Health check route
    app.get('/healthcheck', function () {
        return { status: 'OK' }
    })
}
