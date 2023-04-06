import * as fs from 'fs'
import * as path from 'path'

import { PrismaClient } from '@prisma/client'
import chalk from 'chalk'

import { bootcampMiddleware } from '../src/middleware/prisma.middleware'

const prisma = new PrismaClient()

const bootcamps = JSON.parse(
    fs.readFileSync(path.join(__dirname, '../_data/bootcamps.json'), 'utf-8')
)

async function main() {
    prisma.$use(bootcampMiddleware)

    console.log(chalk.greenBright.bold.underline('⇨ 🌱 Seeding database...'))
    console.log(chalk.yellowBright.bold.underline('⇨ 🗑️ Deleting old bootcamps...'))
    await prisma.bootcamps.deleteMany()

    for (const b of bootcamps) {
        await prisma.bootcamps.create({ data: { ...b } })
        console.log(chalk.greenBright.bold(`⇨ ✅ Created bootcamp: ${b.name}`))
    }
}

main()
    .catch((e: any) => {
        console.error(chalk.redBright.bold.underline(`❌ Database seeding error: ${e.message}`))
        process.exit(1)
    })
    .finally(async () => {
        await prisma.$disconnect()
    })
