package routes

import (
	"golang-pdf-generator/app/controllers/pdf"
	"net/http"

	"github.com/labstack/echo"
)

// Init initialize application
func Init(server *echo.Echo) {

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PDF Generator!")
	})

	pdf := pdf.New()

	server.POST("/pdf", pdf.GeneratePdf)
	server.GET("/pdf/:filename", pdf.GetPdf)

	// Results strictly for testing purpose
	// server.POST("/result", results.SubmitResults)
	// server.GET("/result/:id", results.ViewResults)
	// server.GET("/result/:id/pdf", results.GetResultsPdf)
}
