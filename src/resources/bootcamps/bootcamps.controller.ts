import type { FastifyRequest, FastifyReply } from 'fastify'

/**
 * @desc: Get all bootcamps
 * @route: GET /bootcamps
 * @access: Public
 */
export async function getBootcampsHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    return await reply.send({ message: 'Get all bootcamps' })
}

/**
 * @desc: Get single bootcamp
 * @route: GET /bootcamps/:id
 * @access: Public
 */
export async function getBootcampHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    return await reply.send({ message: 'Get single bootcamp' })
}

/**
 * @desc: Create new bootcamp
 * @route: POST /bootcamps
 * @access: Private
 */
export async function createBootcampHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    return await reply.send({ message: 'Create new bootcamp' })
}

/**
 * @desc: Update bootcamp
 * @route: PATCH /bootcamps/:id
 * @access: Private
 */
export async function updateBootcampHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    return await reply.send({ message: 'Update bootcamp' })
}

/**
 * @desc: Delete bootcamp
 * @route: DELETE /bootcamps/:id
 * @access: Private
 */
export async function deleteBootcampHandler(request: FastifyRequest, reply: FastifyReply): Promise<void> {
    return await reply.send({ message: 'Delete bootcamp' })
}
