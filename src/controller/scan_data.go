package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vishnusomank/sbom-poc/models"
	"github.com/vishnusomank/sbom-poc/src/grype"
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

	byteval := syft.StartScan(input.ImageName, input.Version)
	grypeVal := grype.StartGrype(input.ImageName, input.Version)
	sbom := models.SBOM{ImageName: input.ImageName, Version: input.Version, Value: string(byteval), Vulnerability: string(grypeVal)}

	models.DB.Create(&sbom)
	c.JSON(http.StatusOK, gin.H{"Submitted": input.ImageName + ":" + input.Version})

}

func GetAllScannedImages(c *gin.Context) {
	var sbom []models.SBOM
	count := models.DB.Find(&sbom)
	for i := 0; i < int(count.RowsAffected); i++ {
		c.JSON(http.StatusOK, gin.H{"ID": sbom[i].ID, "IMAGE NAME": sbom[i].ImageName, "IMAGE VERSION": sbom[i].Version})
		c.String(1, "\n")
	}
	c.String(1, "Total Records loaded = "+strconv.FormatInt(count.RowsAffected, 10))

}

func GetScannedImage(c *gin.Context) {
	var sbom models.SBOM
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sbom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": sbom.ID, "IMAGE NAME": sbom.ImageName, "IMAGE VERSION": sbom.Version})
	c.String(1, "\n")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(sbom.Value), &jsonMap)

	c.IndentedJSON(http.StatusOK, gin.H{"SBOM": jsonMap})

}

func GetVulnFromImage(c *gin.Context) {
	var sbom models.SBOM
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sbom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": sbom.ID, "IMAGE NAME": sbom.ImageName, "IMAGE VERSION": sbom.Version})
	c.String(1, "\n")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(sbom.Vulnerability), &jsonMap)
	//c.String(1, sbom.Vulnerability)
	c.IndentedJSON(http.StatusOK, gin.H{"Vulnerabilities": jsonMap})

}

func GetPolicyForImage(c *gin.Context) {
	var sbom models.SBOM
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sbom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": sbom.ID, "IMAGE NAME": sbom.ImageName, "IMAGE VERSION": sbom.Version})
	c.String(1, "\n")

	polVal := PolicyCreate(sbom.ImageName,sbom.Version)
	c.String(1, polVal)

}
