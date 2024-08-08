package Controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"demo_for_gin/Binding"
)

func BindingJsonHandler(c *gin.Context) {
	var bindingJsonData Binding.JsonBindingStruct
	err := c.BindJSON(&bindingJsonData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding Json Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Binding Json Successfully: %+v", &bindingJsonData)})
}
func BindingFormHandler(c *gin.Context) {
	var bindingFormData Binding.FormBindingStruct
	err := c.MustBindWith(&bindingFormData, binding.FormMultipart)
	// err := c.ShouldBind(&bindingFormData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding Form Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Binding Form Successfully: %+v", &bindingFormData)})
}
func ValidateJsonHandler(c *gin.Context) {
	var bindingJsonData Binding.JsonBindingStruct
	err := c.BindJSON(&bindingJsonData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding Json Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Binding Json Successfully: %+v", &bindingJsonData)})
}
