package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/dtos"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/injector"
)

func SetupRouter(r *gin.Engine, config *dtos.Config, di *injector.Injector) {
	api := r.Group("gmb-sweets-services/api/v1")
	api.POST("/enquiry", di.EnquiryController.CreateEnquiry)
	api.GET("/enquiries", di.EnquiryController.GetEnquiry)
}
