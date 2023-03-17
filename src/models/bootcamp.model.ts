import { Schema, model, type Model, type Types } from 'mongoose'
import validator from 'validator'

export interface IBootcamp {
    name: string
    slug: string
    description: string
    website: string
    phone: string
    email: string
    address: string
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

const BootcampSchema = new Schema<IBootcamp, BootcampModel, IBootcampMethods>(
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
        address: {
            type: String,
            required: [true, 'Please add an address']
        },
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

export const Bootcamp = model<IBootcamp, BootcampModel>('Bootcamp', BootcampSchema)
