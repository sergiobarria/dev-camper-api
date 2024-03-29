import * as http from 'http'

import chalk from 'chalk'
import dotenv from 'dotenv'
import config from 'config'

import { app } from './app'
import { logger, prisma } from './lib'

dotenv.config()

let server: http.Server
const PORT = config.get<number>('PORT')
const NODE_ENV = config.get<string>('NODE_ENV')

// Handle uncaught exceptions globally
process.on('uncaughtException', err => {
    logger.error('Uncaught exception:', err)
    process.exit(1)
})

// Handle unhandled promise rejections globally
process.on('unhandledRejection', err => {
    logger.error('Unhandled rejection:', err)
    process.exit(1)
})

async function main(): Promise<void> {
    server = http.createServer(app)

    // connect to database
    await prisma.$connect().finally(() => {
        logger.info(chalk.greenBright.bold.underline('⇨ 💾 Connected to mongodb database'))
    })

    try {
        server.listen(PORT, () => {
            logger.info(
                chalk.greenBright.bold.underline(
                    `⇨ 🚀 Server running in ${NODE_ENV} mode on port ${PORT}`
                )
            )
        })
    } catch (err: any) {
        logger.error(chalk.redBright.bold.underline(`❌ Server error: ${err.message}`))
        process.exit(1)
    }
}

function shutdown(): void {
    logger.info(chalk.magentaBright.bold.underline('⇨ 🔴 Shutting down server...'))
    void server.close()
    prisma.$disconnect().finally(() => {
        process.exit(0)
    })
}

process.on('SIGTERM', shutdown)
process.on('SIGINT', shutdown)

void main()
