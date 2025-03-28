package api

import (
	"github.com/gin-gonic/gin"
)

func Serve(addr string, startbanner bool, debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use()
	r.GET("/:bucket/", ListObjectsV2Handler)
	r.Run(addr)
}
