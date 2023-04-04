import * as http from 'http'

import chalk from 'chalk'
import dotenv from 'dotenv'
import config from 'config'

import { app } from './app'
import { logger } from './lib'

dotenv.config()

let server: http.Server
const PORT = config.get<number>('PORT')
const NODE_ENV = config.get<string>('NODE_ENV')

async function main(): Promise<void> {
    server = http.createServer(app)

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

void main()
