package app

import (
	"camereye_backend_test/app/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", handler.MainHandler)

	r.GET("/login", handler.LoginHandler)

	r.GET("/home", handler.HomeHandler)

	r.GET("/device-list", handler.TailscaleDevicesHandler)

	r.GET("/device-install-list", handler.InstallListHandler)

	r.Static("/front", ".")
}
