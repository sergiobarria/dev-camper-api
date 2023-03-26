import type { Document, FilterQuery, Model, Query } from 'mongoose'

/**
 * @desc: Takes an object and returns a new object with the same properties but with the values of the properties being the values of the properties of the object passed in
 */
export function removeFieldsFromObject(obj: Record<string, string>, fields: string[]): Record<string, string> {
    const objMap = new Map(Object.entries(obj))

    fields.forEach((fielf) => objMap.delete(fielf))

    return Object.fromEntries(objMap)
}

/**
 * @desc: Takes and object and returns a new object with only the specified fields
 * @param {object} obj - The object to filter
 * @param {string[]} allowedFields - The fields to keep
 * @returns {object} - The filtered object
 */
export function filterObj(obj: Record<string, string>, ...allowedFields: string[]): Record<string, string> {
    const newObj: Record<string, string> = {}

    Object.keys(obj).forEach((el) => {
        if (allowedFields.includes(el)) newObj[el] = obj[el]
    })

    return newObj
}

interface QueryString {
    sort?: string
    limit?: string
    page?: string
    fields?: string
    [key: string]: any
}

/**
 * @desc: This class contains all the methods to filter, sort, and paginate resources
 * @param: {Object} query - The mongoose query object
 * @param: {Object} queryStr - The query string from the request object
 */
export class APIFeatures<T extends Document, U = Record<string, any>> {
    model: Model<T>
    query: Query<T[] & U, T, unknown>
    queryStr: QueryString

    constructor(model: Model<T>, queryStr: QueryString) {
        this.model = model
        this.queryStr = queryStr
        this.query = model.find() as Query<T[] & U, T, unknown>
    }

    /**
     * @desc: Filter the query
     * @return: {Object} - The query object
     * @example: /api/v1/bootcamps?housing=true&careers[]=Web%20Development&careers[]=Mobile%20Development
     */
    filter(): APIFeatures<T, FilterQuery<T> & U> {
        let queryObj = { ...this.queryStr }
        const excludedFields = ['sort', 'limit', 'page', 'fields']

        queryObj = removeFieldsFromObject(queryObj, excludedFields)

        // Create query string and add $ to operators like $gt, $gte, etc. so that mongoose can understand them
        let queryStr = JSON.stringify(queryObj)
        queryStr = queryStr.replace(/\b(gte|gt|lte|lt)\b/g, (match) => `$${match}`)

        // Find resource
        this.query = this.model.find(JSON.parse(queryStr)) as Query<T[] & FilterQuery<T> & U, T, unknown>

        return this as APIFeatures<T, FilterQuery<T> & U>
    }
}
