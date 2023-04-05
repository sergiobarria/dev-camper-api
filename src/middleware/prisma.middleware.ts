import type { Prisma } from '@prisma/client'
import slugify from 'slugify'

// This file includes all the middleware functions applied to the prisma client
type PrismaNextFunc = (params: Prisma.MiddlewareParams) => Promise<any>

export async function bootcampMiddleware(
    params: Prisma.MiddlewareParams,
    next: PrismaNextFunc
): Promise<void> {
    const collection = 'bootcamps'
    // Actions BEFORE creating a new bootcamp 👇🏼
    if (params.model === collection && params.action === 'create') {
        const { name } = params.args.data
        params.args.data.slug = slugify(name, { lower: true })
    }

    const result = await next(params)

    // Actions AFTER creating a new bootcamp 👇🏼
    if (params.model === collection && params.action === 'create') {
        // Do something with the result
    }
    return result
}
