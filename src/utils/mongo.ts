import * as mongoose from 'mongoose'
import type { FastifyInstance } from 'fastify'
import config from 'config'
import chalk from 'chalk'

export async function connectToMongoDB(app: FastifyInstance): Promise<void> {
    const db = config.get<string>('MONGO_URI')

    try {
        await mongoose.connect(db, {})
        app.log.info(chalk.yellowBright.bold.underline('MongoDB connected successfully 🚀'))
    } catch (error) {
        app.log.error(chalk.redBright.bold('Error connecting to MongoDB 💥'), error)
    }
}
