import fastify, { type FastifyInstance } from 'fastify'
import config from 'config'

import { router } from './router'
import { bootcampSchemas } from './resources/bootcamps/bootcamps.schemas'

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
    for (const schema of [...bootcampSchemas]) {
        app.addSchema(schema)
    }

    // Register routes here 👇🏼
    void app.register(router, { prefix: '/api/v1' })

    // Custom error handler here 👇🏼

    return app
}
