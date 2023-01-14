package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sergiobarria/dev-camper-api/controllers"
)

type BootcampsRouter struct {
	bootcampsController controllers.BootcampsController
}

func NewBootcampsRouter(controller controllers.BootcampsController) BootcampsRouter {
	return BootcampsRouter{
		bootcampsController: controller,
	}
}

func (router *BootcampsRouter) BootcampsRoutes(e *echo.Group) {
	r := e.Group("/bootcamps")

	r.GET("", router.bootcampsController.GetAllBootcamps)
	r.GET("/:id", router.bootcampsController.GetBootcampById)
	r.POST("", router.bootcampsController.CreateBootcamp)
	r.PUT("/:id", router.bootcampsController.UpdateBootcamp)
	r.DELETE("/:id", router.bootcampsController.DeleteBootcamp)
}
