import express from 'express'

import * as bcController from './bootcamps.controller'
import { validate } from '@/middleware'
import { createBootcampSchema, getBootcampSchema, updateBootcampSchema } from './bootcamps.schemas'

const router = express.Router()

router.route('/radius/:zipcode/:distance').get(bcController.getBootcampsInRadius)

router
    .route('/')
    .get(bcController.getBootcamps)
    .post(validate(createBootcampSchema), bcController.createBootcamp)
router
    .route('/:id')
    .get(validate(getBootcampSchema), bcController.getBootcamp)
    .patch(validate(updateBootcampSchema), bcController.updateBootcamp)
    .delete(bcController.deleteBootcamp)

export { router as bootcampsRouter }
