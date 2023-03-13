import type { FastifyInstance } from 'fastify'

import {
    getBootcampsHandler,
    getBootcampHandler,
    createBootcampHandler,
    updateBootcampHandler,
    deleteBootcampHandler
} from './bootcamps.controller'

/**
 * @desc: Bootcamps routes
 * @route: /bootcamps
 */
export async function bootcampsRouter(app: FastifyInstance): Promise<void> {
    // Get all bootcamps
    app.get('/', getBootcampsHandler)

    // Get single bootcamp
    app.get('/:id', getBootcampHandler)

    // Create new bootcamp
    app.post('/', createBootcampHandler)

    // Update bootcamp
    app.patch('/:id', updateBootcampHandler)

    // Delete bootcamp
    app.delete('/:id', deleteBootcampHandler)
}
