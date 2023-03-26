import type { FastifyInstance } from 'fastify'

import {
    getBootcampsHandler,
    getBootcampHandler,
    createBootcampHandler,
    updateBootcampHandler,
    deleteBootcampHandler,
    getBootcampsInRadiusHandler
} from './bootcamps.controller'
import { $ref } from './bootcamps.schemas'

/**
 * @desc: Bootcamps routes
 * @route: /bootcamps
 */
export async function bootcampsRouter(app: FastifyInstance): Promise<void> {
    // Get all bootcamps
    app.get(
        '/',
        {
            schema: {
                response: {
                    200: $ref('getBootcampsResponse')
                }
            }
        },
        getBootcampsHandler
    )

    // Get single bootcamp
    app.get(
        '/:id',
        {
            schema: {
                response: {
                    200: $ref('getBootcampResponse')
                }
            }
        },
        getBootcampHandler
    )

    // Create new bootcamp
    app.post(
        '/',
        {
            schema: {
                response: {
                    201: $ref('createBootcampResponse')
                },
                body: $ref('createBootcampSchema')
            }
        },
        createBootcampHandler
    )

    // Update bootcamp
    app.patch('/:id', updateBootcampHandler)

    // Delete bootcamp
    app.delete('/:id', deleteBootcampHandler)

    // Get bootcamps within a radius
    app.get('/radius/:zipcode/:distance', getBootcampsInRadiusHandler)
}
