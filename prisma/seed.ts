import { PrismaClient } from '@prisma/client'
import chalk from 'chalk'

const prisma = new PrismaClient()

async function main() {
    console.log(chalk.greenBright.bold.underline('⇨ 🌱 Seeding database...'))

    const bootcamp = await prisma.bootcamps.create({
        data: {
            name: 'Devworks Bootcamp',
            description: 'The best coding bootcamp in the country',
        },
    })

    console.log(chalk.greenBright.bold(`⇨ ✅ Created bootcamp: ${bootcamp.name}`))
    // console.log(chalk.greenBright.bold.underline('⇨ 🌱 Creating courses...'))
}

main()
    .catch((e: any) => {
        console.error(chalk.redBright.bold.underline(`❌ Database seeding error: ${e.message}`))
        process.exit(1)
    })
    .finally(async () => {
        await prisma.$disconnect()
    })
