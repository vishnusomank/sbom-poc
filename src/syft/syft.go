package syft

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func StartScan(imageName, version string) []byte {
	cmd := exec.Command("/bin/sh", "-c", "syft "+imageName+":"+version+" --scope all-layers -o syft-json="+imageName+"_"+version+".json")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}
	fileContent, err := os.Open(imageName + "_" + version + ".json")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)
	cmd = exec.Command("/bin/sh", "-c", "rm "+imageName+"_"+version+".json")
	_, err = cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}

	return byteResult

}
