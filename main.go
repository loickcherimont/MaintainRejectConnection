package main

import (
	"github.com/loickcherimont/MaintainRejectConnection/handler"
	"github.com/loickcherimont/MaintainRejectConnection/router"

	"fmt"
)

var (
	r = router.GetRouter()
)

func main() {
	// Load templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/loginpage", handler.LoginPageHandler)
	r.GET("/internalpage", handler.InternalPageHandler)

	r.POST("/login", handler.LoginHandler)
	r.POST("/logout", handler.LogoutHandler)

	if err := r.Run(":9000"); err != nil {
		fmt.Errorf("error: %v", err)
		return
	}
}