package main

import (
	"io"
	"myproject/config"
	"myproject/routes"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render method that conforms to echo.Renderer interface
func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Render the template with the given data
	return tr.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// loading .env file to get url & port
	// if err := godotenv.Load(); err != nil {
	// log.Fatalf("Error in loading .env file : %v", err)
	//}
	// port := os.Getenv("PORT")

	e := echo.New()

	// Set the custom renderer for Echo
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.New("").ParseGlob("templates/*.html")), // Parse all HTML templates in the templates directory
	}

	// Use session middleware with a custom cookie store
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("Akash@1234"))))

	// Connect to the database
	config.ConnectDB()

	// Setup all the routes for the app
	routes.Routes(e)

	// Start the server on port 1234
	e.Logger.Fatal(e.Start(":1234"))
}
