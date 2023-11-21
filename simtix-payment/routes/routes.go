package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Routes []Route

var Module = fx.Options(
	fx.Provide(NewInvoiceRoute),
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup(engine *gin.RouterGroup)
}

func NewRoutes(invoiceRoutes *InvoiceRoutes) *Routes {
	return &Routes{
		invoiceRoutes,
	}
}

func (r Routes) Setup(engine *gin.RouterGroup) {
	for _, route := range r {
		route.Setup(engine)
	}
}
