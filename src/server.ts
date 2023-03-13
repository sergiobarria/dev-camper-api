import dotenv from 'dotenv'
import config from 'config'

import { createServer } from './app'
import { connectToMongoDB } from './utils'

dotenv.config()

const app = createServer()

async function main(): Promise<void> {
    const PORT = config.get<number>('PORT')

    // Connect to MongoDB
    await connectToMongoDB(app)

    try {
        app.listen({ port: PORT }, (err, address) => {
            if (err !== null) {
                app.log.error(err)
                process.exit(1)
            }
        })
    } catch (error) {
        app.log.error(error)
        process.exit(1)
    }
}

const signals = ['SIGINT', 'SIGTERM'] as const

signals.forEach((signal) => {
    process.on(signal, () => {
        void app.close().then(() => {
            process.exit(0)
        })
    })
})

// start the server
void main()
