package main

import (
	"context"
	"log"
	"os"
	"net/http"
	"html/template"
	"io"
	
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/slack"
	
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/joho/godotenv"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var oauth2conf *oauth2.Config
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while load .env file")
	}
	
	oauth2conf = &oauth2.Config{
		ClientID:	  os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_ID_SECRET"),
		Scopes:		  []string{"users.profile:write"},
		Endpoint: slack.Endpoint,
		RedirectURL: os.Getenv("REDIRECT_URL"),
	}
	
	e := echo.New()
	
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/", index)
	e.GET("/redirect", redirect)
	e.GET("/login", login)
	
	e.Start("localhost:8888")
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func login(c echo.Context) error {
	var url = oauth2conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return c.Redirect(http.StatusFound, url)
}

func redirect(c echo.Context) error {
	var ctx = context.Background()
	
	code := c.QueryParam("code")
	tok, err := oauth2conf.Exchange(ctx, code)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.login.html", map[string]interface{}{
			"error_message": err,
		})
	}
	
	//client := oauth2conf.Client(ctx, tok)
	return c.Render(http.StatusOK, "token.html", map[string]interface{}{
		"token": tok.AccessToken,
	})
}
