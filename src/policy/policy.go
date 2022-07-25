package policy

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"github.com/vishnusomank/sbom-poc/models"
)

type SystemPolicy struct {
	Version     string      `yaml:"apiVersion"`
	KindVal     string      `yaml:"kind"`
	MetadataVal MetadataVal `yaml:"metadata"`
	SpecVal     SpecVal     `yaml:"spec"`
}
type MetadataVal struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}
type SpecVal struct {
	Severity int         `yaml:"severity"`
	Skeleton PolSkeleton `yaml:"process"`

	Action string `yaml:"action"`
}
type PolSkeleton struct {
	MatchPath Matchpath `yaml:"matchPaths"`
}
type Matchpath struct {
	Path string `yaml:"path"`
}

func PolicySearch(imageName, imageVersion string, file []byte, value gjson.Result, id int) {
	numValue, _ := strconv.Atoi(value.Raw)
	for i := 1; i <= numValue; i++ {
		dataVal := gjson.Get(string(file), "matches."+strconv.Itoa(i)+".artifact.name")
		PolicyCreate(imageName, imageVersion, dataVal.String(), id)
	}

}

func PolicyCreate(imageName, imageVersion string, datastring string, id int) {

	var binarypathdb []models.BinaryPathDB
	count := models.BINARYPATHDB.Where("binary_name = ?", datastring).Find(&binarypathdb)
	for i := 0; i < int(count.RowsAffected); i++ {

		pol := SystemPolicy{
			Version: "security.kubearmor.com/v1",
			KindVal: "KubeArmorPolicy",
			MetadataVal: MetadataVal{
				Name:      "sbom-policy-for-" + imageName + "-" + imageVersion + "-" + datastring + "-" + strconv.Itoa(i),
				Namespace: "default",
			},
			SpecVal: SpecVal{
				Severity: 5,
				Skeleton: PolSkeleton{
					MatchPath: Matchpath{
						Path: binarypathdb[i].BinaryPath,
					},
				},
				Action: "Audit",
			},
		}

		jsonData, err := json.Marshal(&pol)
		if err != nil {
			fmt.Printf("[%s][%s] Error while Marshaling. %v\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)
		}

		policydb := models.PolicyDB{CVEId: "null", PolicyData: string(jsonData)}

		models.POLICYDB.Create(&policydb)

		if err := models.POLICYDB.Last(&policydb).Error; err != nil {
			log.Panic(err)
		}

		sbompolicy := models.SBOMPolicy{SbomID: id, PolicyID: int(policydb.ID)}
		models.SBOMPOLICYDB.Create(&sbompolicy)
	}

}
