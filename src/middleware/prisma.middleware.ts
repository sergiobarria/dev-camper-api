import type { Prisma } from '@prisma/client'
import slugify from 'slugify'

import { geocoder } from '../lib'

// This file includes all the middleware functions applied to the prisma client
type PrismaNextFunc = (params: Prisma.MiddlewareParams) => Promise<any>

export async function bootcampMiddleware(
    params: Prisma.MiddlewareParams,
    next: PrismaNextFunc
): Promise<void> {
    const collection = 'bootcamps'
    // Actions BEFORE creating a new bootcamp 👇🏼
    if (params.model === collection && params.action === 'create') {
        // -> Add a slug to the bootcamp name
        const { name } = params.args.data
        params.args.data.slug = slugify(name, { lower: true })

        // -> Generate geolocation data from address
        const loc = await geocoder.geocode(params.args.data.address)
        const { latitude, longitude } = loc[0]
        params.args.data.location = {
            type: 'Point',
            coordinates: [longitude, latitude],
            formattedAddress: loc[0].formattedAddress,
            street: loc[0].streetName,
            city: loc[0].city,
            state: loc[0].stateCode,
            zipcode: loc[0].zipcode,
            country: loc[0].countryCode,
        }

        // -> Remove address from the args
        delete params.args.data.address
    }

    const result = await next(params)

    // Actions AFTER creating a new bootcamp 👇🏼
    if (params.model === collection && params.action === 'create') {
        // Do something with the result
    }
    return result
}
