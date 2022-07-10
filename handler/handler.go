package handler

import (
	"net/http"

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

	c.Redirect(http.StatusMovedPermanently, "/internalpage")

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

	c.Redirect(http.StatusMovedPermanently, "/loginpage")
}

// Check if user is authenticated
// if not display to user a message of FORBIDDEN ACCESS
// else go directly to internal page
func InternalPageHandler(c *gin.Context) {
	session, err := store.Get(c.Request, "cookie-name")

	if err != nil {
		panic(err)
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(c.Writer, "FORBIDDEN ACCESS", http.StatusForbidden)
		return
	}

	c.HTML(http.StatusOK, "internal.page.tmpl", nil)
}