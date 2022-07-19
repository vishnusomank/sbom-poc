package syft

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/vishnusomank/sbom-poc/models"
	"github.com/vishnusomank/sbom-poc/src/grype"
)

func StartScan(imageName, version string, id int, sbom *models.SBOM) {
	cmd := exec.Command("/bin/sh", "-c", "syft "+imageName+":"+version+" --scope all-layers -o syft-json="+imageName+"_"+version+".json")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}
	fileContent, err := os.Open(imageName + "_" + version + ".json")

	if err != nil {
		log.Fatal(err)
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	if err := models.DB.Model(&sbom).Where("id = ?", id).Update("Value", string(byteResult)).Error; err != nil {
		fmt.Printf("[%s][%s] Unable to update SBOM value\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"))

		return
	}

	go grype.StartGrype(imageName, version, id, sbom)

}
