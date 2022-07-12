package grype

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func StartGrype(imageName, version string) []byte {

	cmd := exec.Command("/bin/sh", "-c", "grype sbom:./"+imageName+"_"+version+".json  -o json > grype-"+imageName+"_"+version+".json")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}
	fileContent, err := os.Open("grype-" + imageName + "_" + version + ".json")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	cmd = exec.Command("/bin/sh", "-c", "rm "+imageName+"_"+version+".json"+" grype-"+imageName+"_"+version+".json")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}
	return byteResult
}
