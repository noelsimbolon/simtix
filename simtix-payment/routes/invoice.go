package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simtix/handlers"
	"simtix/models"
	"simtix/utils/logger"
)

type InvoiceRoutes struct {
	//logger         lib.Logger
	//userController controllers.UserController
	//authMiddleware middlewares.JWTAuthMiddleware
	invoiceHandler handlers.InvoiceHandler
}

func (s InvoiceRoutes) Setup(engine *gin.RouterGroup) {
	logger.Log.Info("Setting up routes")
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
