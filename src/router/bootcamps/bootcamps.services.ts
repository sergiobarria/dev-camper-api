import { prisma } from '@/lib'
import type { Prisma, bootcamps } from '@prisma/client'
import type { CreateBootcampType } from './bootcamps.schemas'

export const createOne = async (data: CreateBootcampType): Promise<bootcamps> => {
    const record = await prisma.bootcamps.create({ data })

    return record
}

export const findFirst = async (filter: Prisma.bootcampsWhereInput): Promise<bootcamps | null> => {
    const record = await prisma.bootcamps.findFirst({ where: filter })

    return record
}

export const findMany = async (filter?: Prisma.bootcampsFindManyArgs): Promise<bootcamps[]> => {
    const records = await prisma.bootcamps.findMany(filter)

    return records
}

export const findOneById = async (id: string): Promise<bootcamps | null> => {
    const record = await prisma.bootcamps.findUnique({ where: { id } })

    return record
}

export const updateOne = async (
    id: string,
    data: Partial<CreateBootcampType>
): Promise<bootcamps | null> => {
    const record = await prisma.bootcamps.update({ where: { id }, data })

    return record
}

export const deleteOne = async (id: string): Promise<void> => {
    await prisma.bootcamps.delete({ where: { id } })
}
