package pdf

import (
	"fmt"
	"golang-pdf-generator/app/helpers"
	"golang-pdf-generator/app/models"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
)

type (
	// IController controller interface
	IController interface {
		GeneratePdf(c echo.Context) error
		GetPdf(c echo.Context) error
	}

	// Controller controller type
	Controller struct {
		generator helpers.IPdfGenerator
	}
)

var extRegexp = regexp.MustCompile(`(?i).pdf$`)

// GeneratePdf generates pdf ans save it to file system
func (controller *Controller) GeneratePdf(c echo.Context) error {
	input := new(models.Input)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	filename := extRegexp.ReplaceAllString(input.Filename, "") + ".pdf"

	if err := controller.generator.HTMLToPdf(input.Body, filename); err != nil {
		return helpers.InternalServerError(c, err)
	}

	return c.JSONBlob(http.StatusOK, []byte(fmt.Sprintf(`{"filename" : "%s"}`, filename)))
}

// GetPdf get pdf by filename
func (controller *Controller) GetPdf(c echo.Context) error {
	filename := c.Param("filename")

	filePath, err := controller.generator.Find(filename)
	if err != nil {
		return helpers.HTTPNotFound(c)
	}

	return c.Attachment(filePath, filename)
}

// New controller constructor
func New() IController {
	controller := &Controller{helpers.NewPdfGenerator()}

	return controller
}
