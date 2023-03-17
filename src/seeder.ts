import * as path from 'path'
import * as fs from 'fs'

import * as mongoose from 'mongoose'
import * as dotenv from 'dotenv'
import chalk from 'chalk'
import config from 'config'

import { Bootcamp } from './models/bootcamp.model'

dotenv.config()

// Get config vars
const MONGO_URI = config.get<string>('MONGO_URI')

// Connect to MongoDB
mongoose
    .connect(MONGO_URI, {})
    .then(() => {
        console.log(chalk.greenBright.bold.underline('MongoDB connected successfully 🚀'))
    })
    .catch((error) => {
        console.log(chalk.redBright.bold('Error connecting to MongoDB 💥'), error)
    })

// Read JSON files
const bootcamps = JSON.parse(fs.readFileSync(path.join(__dirname, '_data', 'bootcamps.json'), 'utf-8'))
console.log(chalk.yellowBright.bold.underline('Bootcamps data loaded successfully 🚀'))

async function importData(): Promise<void> {
    try {
        // delete all data from the database
        await Bootcamp.deleteMany()

        // import data into the database
        await Bootcamp.create(bootcamps)
        console.log(chalk.yellowBright.bold.underline('Data imported successfully 🚀'))
        process.exit(0)
    } catch (error) {
        console.log(chalk.redBright.bold('Error importing data 💥'), error)
        process.exit(1)
    }
}

async function deleteData(): Promise<void> {
    try {
        // delete all data from the database
        await Bootcamp.deleteMany()
        console.log(chalk.yellowBright.bold.underline('Data deleted successfully 🚀'))
        process.exit(0)
    } catch (error) {
        console.log(chalk.redBright.bold('Error deleting data 💥'), error)
        process.exit(1)
    }
}

if (process.argv[2] === '-i') {
    void importData()
}

if (process.argv[2] === '-d') {
    void deleteData()
}
