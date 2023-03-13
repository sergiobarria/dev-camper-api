import * as path from 'path'

import fastify, { type FastifyInstance } from 'fastify'
import config from 'config'

export const loggerConfig = {
    development: {
        transport: {
            target: 'pino-pretty',
            options: {
                translateTime: 'HH:MM:ss Z',
                ignore: 'pid,hostname'
            }
        }
    },
    production: true
}

export function createServer(): FastifyInstance {
    const env = config.get<string>('NODE_ENV')

    const app = fastify({
        logger: loggerConfig[env as keyof typeof loggerConfig] ?? loggerConfig.development
    })

    // Register plugins here 👇🏼

    // Register JSON schemas here 👇🏼

    // Register routes here 👇🏼

    // Custom error handler here 👇🏼

    return app
}
