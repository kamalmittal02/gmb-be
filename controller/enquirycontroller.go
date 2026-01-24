package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/request"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/service"
	"net/http"
)

type EnquiryControllerI interface {
	CreateEnquiry(c *gin.Context)
}

type EnquiryController struct {
	EnquiryService service.EnquiryServiceI
}

func NewEnquiryController(service service.EnquiryServiceI) EnquiryControllerI {
	return &EnquiryController{EnquiryService: service}
}

func (eqc *EnquiryController) CreateEnquiry(c *gin.Context) {
	var enquiry request.Enquiry

	// Bind JSON input to the Enquiry struct
	if err := c.ShouldBindJSON(&enquiry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Validate required fields
	if enquiry.Name == "" || enquiry.Phone == "" || enquiry.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (name, email, phone, message) are required"})
		return
	}

	// Validate email format
	if enquiry.Email != "" && !IsValidEmail(enquiry.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Validate phone number format
	if !IsValidPhone(enquiry.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number format"})
		return
	}

	// Call the service to create the enquiry
	err := eqc.EnquiryService.CreateEnquiry(c.Request.Context(), enquiry)
	if err != nil {
		fmt.Printf("error creating enquiry: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create enquiry"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Enquiry created successfully"})
}
