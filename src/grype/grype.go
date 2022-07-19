package grype

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/vishnusomank/sbom-poc/models"
)

func StartGrype(imageName, version string, id int, sbom *models.SBOM) {

	cmd := exec.Command("/bin/sh", "-c", "grype sbom:./"+imageName+"_"+version+".json  -o json > grype-"+imageName+"_"+version+".json")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}
	fileContent, err := os.Open("grype-" + imageName + "_" + version + ".json")

	if err != nil {
		log.Fatal(err)
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	if err := models.DB.Model(&sbom).Where("id = ?", id).Update("Vulnerability", string(byteResult)).Error; err != nil {
		fmt.Printf("[%s][%s] Unable to update SBOM value\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"))

		return
	}

	var policydb []models.PolicyDB
	count := models.POLICYDB.Find(&policydb)

	file, _ := ioutil.ReadFile("grype-" + imageName + "_" + version + ".json")
	value := gjson.Get(string(file), "matches.#")
	numValue, _ := strconv.Atoi(value.Raw)

	for i := 1; i <= numValue; i++ {
		polcount := 0
		dataVal := gjson.Get(string(file), "matches."+strconv.Itoa(i)+".vulnerability.id")
		println("data value", dataVal.String())
		for j := 1; j <= int(count.RowsAffected); j++ {
			if err := models.POLICYDB.Where("ID = ?", j).First(&policydb).Error; err != nil {
				fmt.Printf("[%s][%s] Failed to retrieve data %v\n", color.RedString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)
				return
			}
			if strings.Contains(policydb[0].CVEId, dataVal.String()) && polcount == 0 {
				sbompolicy := models.SBOMPolicy{SbomID: id, PolicyID: j}
				models.SBOMPOLICYDB.Create(&sbompolicy)
				polcount = 1
			}

		}

	}

	cmd = exec.Command("/bin/sh", "-c", "rm "+imageName+"_"+version+".json"+" grype-"+imageName+"_"+version+".json")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("File deletion failed : " + err.Error())
	}

}
