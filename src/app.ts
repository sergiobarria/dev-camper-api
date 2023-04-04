import express from 'express'

import { router } from './router'
import { morganMiddleware } from './middleware'

export const app = express()

// ===== Apply middlewares 👇🏼 =====
app.use(express.json())

if (process.env.NODE_ENV === 'development') {
    app.use(morganMiddleware)
}

// ===== Apply routes 👇🏼 =====
app.use('/api/v1', router)

// ===== Apply error handlers 👇🏼 =====
