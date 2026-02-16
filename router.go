package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/dtos"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/injector"
)

func SetupRouter(r *gin.Engine, config *dtos.Config, di *injector.Injector) {
	api := r.Group("/gmb-sweets-services/api/v1")
	
	// Handle OPTIONS preflight requests
	api.OPTIONS("/enquiry", func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Max-Age", "86400")
		c.Status(204)
	})
	
	api.POST("/enquiry", di.EnquiryController.CreateEnquiry)
	//api.GET("/enquiries", di.EnquiryController.GetEnquiry)
}
