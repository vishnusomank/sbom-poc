package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vishnusomank/sbom-poc/models"
	"github.com/vishnusomank/sbom-poc/src/syft"
	"github.com/vishnusomank/sbom-poc/utils/constants"
)

type ImageInput struct {
	ImageName string `json:"image" binding:"required"`
	Version   string `json:"version" binding:"required"`
}

func PopulateDummyPol() {

	policydb := models.PolicyDB{CVEId: "CVE-2022-29458", PolicyData: constants.PolDummyData}

	models.POLICYDB.Create(&policydb)

}

func getLastID() int {

	var sbom models.SBOM

	if err := models.DB.Last(&sbom).Error; err != nil {
		log.Panic(err)
	}
	return int(sbom.ID)

}

//AddRegistry - Add Registry Credential
func AddImage(c *gin.Context) {

	var input ImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sbom := models.SBOM{ImageName: input.ImageName, Version: input.Version, Value: "null", Vulnerability: "null"}

	models.DB.Create(&sbom)

	id := getLastID()

	go syft.StartScan(input.ImageName, input.Version, id, &sbom)

	c.JSON(http.StatusOK, gin.H{"Submitted": input.ImageName + ":" + input.Version})

}

func GetAllScannedImages(c *gin.Context) {
	var sbom []models.SBOM
	count := models.DB.Find(&sbom)
	if int(count.RowsAffected) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{"Computation in progress": "Please wait"})
		c.String(204, "\n")

	} else {
		for i := 0; i < int(count.RowsAffected); i++ {
			c.JSON(http.StatusOK, gin.H{"ID": sbom[i].ID, "IMAGE NAME": sbom[i].ImageName, "IMAGE VERSION": sbom[i].Version})
			c.String(200, "\n")
		}
		c.String(200, "Total Records loaded = "+strconv.FormatInt(count.RowsAffected, 10))
	}
}

func GetScannedImage(c *gin.Context) {
	var sbom models.SBOM
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sbom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	if sbom.Value == "null" {
		c.JSON(http.StatusNoContent, gin.H{"Computation in progress": "Please wait"})
		c.String(204, "\n")

	} else {

		c.JSON(http.StatusOK, gin.H{"ID": sbom.ID, "IMAGE NAME": sbom.ImageName, "IMAGE VERSION": sbom.Version})
		c.String(200, "\n")

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(sbom.Value), &jsonMap)

		c.IndentedJSON(http.StatusOK, gin.H{"SBOM": jsonMap})
	}

}

func GetVulnFromImage(c *gin.Context) {
	var sbom models.SBOM
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sbom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	if sbom.Value == "null" {
		c.JSON(http.StatusOK, gin.H{"Computation in progress": "Please wait"})
		c.String(200, "\n")

	} else {

		c.JSON(http.StatusOK, gin.H{"ID": sbom.ID, "IMAGE NAME": sbom.ImageName, "IMAGE VERSION": sbom.Version})
		c.String(200, "\n")

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(sbom.Vulnerability), &jsonMap)
		//c.String(1, sbom.Vulnerability)
		c.IndentedJSON(http.StatusOK, gin.H{"Vulnerabilities": jsonMap})
	}

}

func GetPolicyForImage(c *gin.Context) {
	var sbompolicy []models.SBOMPolicy
	var test []models.SBOMPolicy
	var policy models.PolicyDB

	count := models.SBOMPOLICYDB.Where("sbom_id = ?", c.Param("id")).Find(&test)

	fmt.Println(count.RowsAffected)

	if err := models.SBOMPOLICYDB.Where("sbom_id = ?", c.Param("id")).Find(&sbompolicy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Record not found!"})
		return
	}

	if sbompolicy[1].SbomID == 0 {
		c.JSON(http.StatusOK, gin.H{"Computation in progress": "Please wait"})
		c.String(200, "\n")

	} else {
		for i := 1; i <= int(count.RowsAffected); i++ {
			/*
				if err := models.POLICYDB.Where("id = ?", sbompolicy[i].PolicyID).First(&policy).Error; err != nil {
					c.String(200, "\n")
					c.JSON(http.StatusBadRequest, gin.H{"ID": sbompolicy[i].PolicyID, "Error": "Record not found!"})
					c.String(200, "\n")
					//return
				}
			*/
			models.POLICYDB.Where("id = ?", sbompolicy[i].PolicyID).First(&policy)

			var jsonMap map[string]interface{}
			json.Unmarshal([]byte(policy.PolicyData), &jsonMap)
			c.JSON(http.StatusOK, gin.H{"ID": sbompolicy[i].SbomID, "PolicyID": sbompolicy[i].PolicyID})
			c.String(200, "\n")
			c.IndentedJSON(http.StatusOK, gin.H{"Policy": jsonMap})
		}
	}

}
