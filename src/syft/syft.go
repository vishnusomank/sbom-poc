package syft

import (
	"log"
	"os/exec"
)

func StartScan(imageName, version string) {
	cmd := exec.Command("/bin/sh", "-c", "syft "+imageName+":"+version+" -o syft-json="+imageName+"_"+version+".json")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("SBOM generation failed : " + err.Error())
	}

}

func ShowResult(imageName, version string) {

}
