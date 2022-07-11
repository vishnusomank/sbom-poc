package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishnusomank/sbom-poc/src/syft"
)

type ImageInput struct {
	ImageName string `json:"image" binding:"required"`
	Version   string `json:"version" binding:"required"`
}

//AddRegistry - Add Registry Credential
func AddImage(c *gin.Context) {

	var input ImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	syft.StartScan(input.ImageName, input.Version)
	c.JSON(http.StatusOK, gin.H{"Submitted": input.ImageName + ":" + input.Version})

}

func ShowData(c *gin.Context) {

	var input ImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	syft.ShowResult(input.ImageName, input.Version)
	c.JSON(http.StatusOK, gin.H{})

}
