package gommonfx

import "github.com/gin-gonic/gin"

type Route struct {
	Method   string
	Pattern  string
	Handlers []gin.HandlerFunc
}
