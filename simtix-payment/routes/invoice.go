package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simtix/handlers"
	"simtix/models"
)

type InvoiceRoutes struct {
	//logger         lib.Logger
	//userController controllers.UserController
	//authMiddleware middlewares.JWTAuthMiddleware
	invoiceHandler handlers.InvoiceHandler
}

func (s InvoiceRoutes) Setup(engine *gin.RouterGroup) {
	log.Print("Setting up routes")
	invoiceRoutes := engine.Group("invoice")
	invoiceRoutes.GET(
		"/", func(context *gin.Context) {
			response := models.Response{
				Code: http.StatusOK,
				Body: models.ResponseBody{
					Message: "Success",
				},
			}
			context.JSON(response.Code, response.Body)
		},
	)
	invoiceRoutes.POST("/", s.invoiceHandler.PostInvoice)
}

func NewInvoiceRoute(handler handlers.InvoiceHandler) *InvoiceRoutes {
	return &InvoiceRoutes{
		invoiceHandler: handler,
	}
}
