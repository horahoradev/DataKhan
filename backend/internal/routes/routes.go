package routes

type RouteHandler struct {
}

func NewRouteHandler() {
	return RouteHandler{}
}

func SetupRoutes(e *echo.Echo, cfg *config.Config) {
	r := NewRouteHandler()

	e.POST("/view/", r.handleView)
	e.POST("/event/", r.handleEvent)
}
