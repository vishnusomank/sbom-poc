package controller

import (
	"fmt"
	"time"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
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
	Skeleton PolSkeleton `yaml:"skeleton"`
	Action   string      `yaml:"action"`
}
type PolSkeleton struct {
	DummyData1 string `yaml:"dummydata1"`
	DummyData2 string `yaml:"dummydata2"`
	DummyDataN string `yaml:"dummydataN"`
}

func PolicyCreate(imageName, imageVersion string) string {
	pol := SystemPolicy{
		Version: "security.kubearmor.com/v1",
		KindVal: "KubeArmorPolicy",
		MetadataVal: MetadataVal{
			Name:      "sbom-policy-for-"+imageName+"-"+imageVersion,
			Namespace: "default",
		},
		SpecVal: SpecVal{
			Severity: 5,
			Skeleton: PolSkeleton{
				DummyData1: "value1",
				DummyData2: "value2",
				DummyDataN: "valueN",
			},
			Action: "Audit",
		},
	}

	yamlData, err := yaml.Marshal(&pol)

	if err != nil {
		fmt.Printf("[%s][%s] Error while Marshaling. %v\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"),err)
	}
	return string(yamlData) // yamlData will be in bytes. So converting it to string.
}
