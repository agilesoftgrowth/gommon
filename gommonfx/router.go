package gommonfx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinRouter(controllers []Controller) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	for _, controller := range controllers {
		for _, route := range controller.Routes() {
			router.Handle(route.Method, route.Pattern, route.Handlers...)
		}
	}

	return router
}

func AsRouter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(http.Handler)),
		fx.ParamTags(`group:"controllers"`),
	)
}
