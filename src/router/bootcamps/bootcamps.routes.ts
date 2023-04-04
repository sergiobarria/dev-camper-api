import express from 'express'

import * as bootcampsController from './bootcamps.controller'

const router = express.Router()

router.route('/').get(bootcampsController.getBootcamps).post(bootcampsController.createBootcamp)
router
    .route('/:id')
    .get(bootcampsController.getBootcamp)
    .put(bootcampsController.updateBootcamp)
    .delete(bootcampsController.deleteBootcamp)

export { router as bootcampsRouter }
