package routes

import (
 "CaitsCurates/backend/controller"

 "github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
 r := gin.Default()
 mido := r.Group("/caitscurates")
 {
  mido.GET("products", controller.GetProducts)
  mido.POST("products", controller.CreateProducts)
  mido.GET("config", controller.Config)
  mido.POST("create-payment-intent", controller.HandleCreatePaymentIntent)

 }
 return r
}