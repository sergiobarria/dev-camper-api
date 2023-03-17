import { z } from 'zod'
import { buildJsonSchemas } from 'fastify-zod'

export const baseBootcampSchema = {
    name: z.string().min(1).max(50),
    slug: z.string().optional(),
    description: z.string().min(1).max(500),
    website: z.string().url().optional(),
    phone: z.string().max(20).optional(),
    email: z.string().email().optional(),
    address: z.string().min(1).optional(),
    location: z
        .object({
            type: z.enum(['Point']),
            coordinates: z.array(z.number()),
            formattedAddress: z.string().optional(),
            street: z.string().optional(),
            city: z.string().optional(),
            state: z.string().optional(),
            zipcode: z.string().optional(),
            country: z.string().optional()
        })
        .optional(),
    careers: z.array(
        z.enum([
            'Full-Stack',
            'Web Development',
            'Mobile Development',
            'UI/UX Design',
            'Coding for Kids',
            'Cybersecurity',
            'Data Science',
            'Business',
            'Other',
            'Cloud Infrastructure',
            'Cloud Architecture'
        ])
    ),
    averageRating: z.number().min(1).max(10).optional(),
    averageCost: z.number().optional(),
    photo: z.string().optional(),
    housing: z.boolean().optional(),
    jobAssistance: z.boolean().optional(),
    jobGuarantee: z.boolean().optional(),
    acceptGi: z.boolean().optional()
}

export const createBootcampSchema = z.object({
    ...baseBootcampSchema
})

export const createBootcampResponse = z.object({
    success: z.boolean(),
    data: z.object({
        _id: z.string(),
        ...baseBootcampSchema
    })
})

export const getBootcampSchema = z.object({
    params: z.object({ id: z.string() })
})

export const bootcampSchema = z.object({
    body: z.object({ ...baseBootcampSchema }),
    params: z.object({ id: z.string() })
})

export const getBootcampsResponse = z.object({
    success: z.boolean(),
    count: z.number(),
    data: z.array(
        z.object({
            _id: z.string(),
            ...baseBootcampSchema
        })
    )
})

export const getBootcampResponse = z.object({
    success: z.boolean(),
    data: z.object({
        _id: z.string(),
        ...baseBootcampSchema
    })
})

export type CreateBootcampBody = z.infer<typeof createBootcampSchema>
export type BootcampInputs = z.infer<typeof bootcampSchema>

export const { schemas: bootcampSchemas, $ref } = buildJsonSchemas({
    createBootcampSchema,
    createBootcampResponse,
    getBootcampsResponse,
    getBootcampResponse,
    getBootcampSchema
})
