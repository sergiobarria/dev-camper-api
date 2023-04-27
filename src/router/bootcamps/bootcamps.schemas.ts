import { z } from 'zod'

const careers = z.enum([
    'Web Development',
    'Full-Stack',
    'Mobile Development',
    'UI/UX Design',
    'Data Science',
    'Coding for Kids',
    'Cibersecurity',
    'Cloud Infrastructure',
    'Cloud Architecture',
    'Other',
])

export const bootcampBase = {
    body: z.object({
        name: z.string({
            required_error: 'name is required',
        }),
        description: z.string({
            required_error: 'description is required',
        }),
        website: z
            .string({
                required_error: 'website is required',
            })
            .url({
                message: 'website must be a valid URL',
            })
            .transform(value => value.toLowerCase())
            .transform(value => value.trim()),
        phone: z.string({
            required_error: 'phone is required',
        }),
        email: z
            .string({
                required_error: 'email is required',
            })
            .email({
                message: 'email must be a valid email address',
            })
            .transform(value => value.toLowerCase())
            .transform(value => value.trim()),
        address: z.string({
            required_error: 'address is required',
        }),
        careers: z.array(careers).nonempty({
            message: 'careers must contain at least one career',
        }),
        housing: z.boolean().default(false),
        jobAssistance: z.boolean().default(false),
        jobGuarantee: z.boolean().default(false),
        acceptGi: z.boolean().default(false),
    }),
}

export const createBootcampSchema = z.object({
    ...bootcampBase,
})

export const updateBootcampSchema = z.object({
    body: z.object({
        name: z.string().optional(),
        description: z.string().optional(),
        website: z
            .string()
            .url()
            .transform(value => value.toLowerCase())
            .transform(value => value.trim())
            .optional(),
        phone: z.string().optional(),
        email: z
            .string()
            .email()
            .transform(value => value.toLowerCase())
            .transform(value => value.trim())
            .optional(),
        address: z.string().optional(),
        careers: z.array(careers).optional(),
        housing: z.boolean().optional(),
        jobAssistance: z.boolean().optional(),
        jobGuarantee: z.boolean().optional(),
        acceptGi: z.boolean().optional(),
    }),
})

export const getBootcampSchema = z.object({
    params: z.object({
        id: z.string({
            required_error: 'id is required',
        }),
    }),
})

export const getBootcampsSchema = z.object({
    query: z.object({
        page: z.string().optional(),
        limit: z.string().optional(),
        sort: z.string().optional(),
        fields: z.string().optional(),
        name: z.string().optional(),
        housing: z.boolean().optional(),
        jobAssistance: z.boolean().optional(),
        jobGuarantee: z.boolean().optional(),
        acceptGi: z.boolean().optional(),
        location: z.object({
            state: z.string().optional(),
        }),
    }),
})

export const getBootcampsInRadiusSchema = z.object({
    params: z.object({
        zipcode: z.string({
            required_error: 'zipcode is required',
        }),
        distance: z.string({
            required_error: 'distance is required',
        }),
        unit: z.string().optional(),
    }),
})

export type CreateBootcampType = z.infer<typeof createBootcampSchema>['body']
export type GetBootcampType = z.infer<typeof getBootcampSchema>['params']
export type UpdateBootcampType = z.infer<typeof updateBootcampSchema>['body']
export type GetBootcampsQueryType = z.infer<typeof getBootcampsSchema>['query']
export type GetBootcampsInRadiusType = z.infer<typeof getBootcampsInRadiusSchema>['params']
