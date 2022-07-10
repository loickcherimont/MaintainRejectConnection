package handler

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// Generate login page
func LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.page.tmpl", nil)
}

// Set user as authenticated
// Connect to internal page
func LoginHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "cookie-name")
	if err != nil {
		panic(err)
	}

	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)

	c.HTML(http.StatusOK, "internal.page.tmpl", nil)
	return

}

// Revoke users authentication
// redirect user to login page
func LogoutHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "cookie-name")
	if err != nil {
		panic(err)
	}

	session.Values["authenticated"] = false
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/loginpage")
}

// Check if user is authenticated
// if not user stay on login page
// else go directly to internal page
// TODO: CHANGE URI AUTOMATICALLY WHEN USER CONNECTION IS REJECTED
func InternalPageHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "cookie-name")

	if err != nil {
		log.Println(err.Error())
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.HTML(http.StatusForbidden, "login.page.tmpl", nil)
		return
	}

	c.HTML(http.StatusOK, "internal.page.tmpl", nil)
	return
}