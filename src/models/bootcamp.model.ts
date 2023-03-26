import { Schema, model, type Model, type Types, type CallbackError, type Document } from 'mongoose'
import validator from 'validator'
import slugify from 'slugify'

import { geocoder } from '@/utils'

export interface IBootcamp extends Document {
    name: string
    slug: string
    description: string
    website: string
    phone: string
    email: string
    address?: string
    location: {
        type: string
        coordinates: Types.Array<number> // suggested by mongoose to use Types.Array instead of number[]
        formattedAddress: string
        street: string
        city: string
        state: string
        zipcode: string
        country: string
    }
    careers: Types.Array<string>
    averageRating: number
    averageCost: number
    photo: string
    housing: boolean
    jobAssistance: boolean
    jobGuarantee: boolean
    acceptGi: boolean
}

export interface IBootcampMethods {} // eslint-disable-line @typescript-eslint/no-empty-interface

export type BootcampModel = Model<IBootcamp, unknown, IBootcampMethods>

const bootcampSchema = new Schema<IBootcamp, BootcampModel, IBootcampMethods>(
    {
        name: {
            type: String,
            required: [true, 'Please add a name'],
            unique: true,
            trim: true,
            maxlength: [50, 'Name can not be more than 50 characters']
        },
        slug: String,
        description: {
            type: String,
            required: [true, 'Please add a description'],
            maxlength: [500, 'Description can not be more than 500 characters']
        },
        website: {
            type: String,
            validator: [validator.isURL, 'Please use a valid URL with HTTP or HTTPS']
        },
        phone: {
            type: String,
            maxlength: [20, 'Phone number can not be longer than 20 characters']
        },
        email: {
            type: String,
            validator: [validator.isEmail, 'Please add a valid email']
        },
        address: String,
        location: {
            type: {
                type: String,
                enum: ['Point']
                // required: true
            },
            coordinates: {
                type: [Number],
                // required: true,
                index: '2dsphere'
            },
            formattedAddress: String,
            street: String,
            city: String,
            state: String,
            zipcode: String,
            country: String
        },
        careers: {
            type: [String],
            required: true,
            enum: [
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
            ]
        },
        averageRating: {
            type: Number,
            min: [1, 'Rating must be at least 1'],
            max: [10, 'Rating must can not be more than 10']
        },
        averageCost: Number,
        photo: {
            type: String,
            default: 'no-photo.jpg'
        },
        housing: {
            type: Boolean,
            default: false
        },
        jobAssistance: {
            type: Boolean,
            default: false
        },
        jobGuarantee: {
            type: Boolean,
            default: false
        },
        acceptGi: {
            type: Boolean,
            default: false
        }
    },
    { timestamps: true, versionKey: false }
)

// Create bootcamp slug from the name
bootcampSchema.pre(/save/, function (this: IBootcamp, next: (error?: CallbackError) => void) {
    this.slug = slugify(this.name, { lower: true })
    next()
})

// Geocode & create location field
bootcampSchema.pre(/save/, async function (this: IBootcamp, next: (error?: CallbackError) => void) {
    const loc = await geocoder.geocode(this.address as string)

    this.location = {
        type: 'Point',
        coordinates: [loc[0].longitude, loc[0].latitude] as Types.Array<number>,
        formattedAddress: loc[0].formattedAddress as string,
        street: loc[0].streetName as string,
        city: loc[0].city as string,
        state: loc[0].stateCode as string,
        zipcode: loc[0].zipcode as string,
        country: loc[0].countryCode as string
    }

    // Do not save address in DB
    this.address = undefined

    next()
})

export const Bootcamp = model<IBootcamp, BootcampModel>('Bootcamp', bootcampSchema)
