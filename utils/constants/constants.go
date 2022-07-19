package constants

import "github.com/jinzhu/gorm"

//constants for Request Body
const (
	InvalidReqBody  = "Invalid Request Body"
	RequiredReqBody = "Required Request Body Missing"
	WrongReqBody    = "Wrong Request Body : "
)

//constants for Error Messages
const (
	Err400        = "err-400"
	Err500        = "err-500"
	ErrInAddImage = "Error in Add Image"
)

//constants for URL Path
const (
	ADD_IMAGE      = "add-image"
	GET_OP         = "show-data"
	GET_ALL_IMAGES = "all_scanned-images"
	GET_IMAGE      = "scanned-image"
	Get_VULN       = "get-vuln"
	GET_POL        = "get-policy"
)

const (
	ACTIVE = "Active"
	DELETE = "Delete"
)

var DB *gorm.DB
var POLICYDB *gorm.DB
var SBOMPOLICYDB *gorm.DB

var PolDummyData = `{ "apiVersion": "security.kubearmor.com/v1", "kind": "KubeArmorPolicy", "metadata": { "name": "ksp-cve-2022-29458", "namespace": "default" }, "spec": { "tags": [ "CVE", "CVE-2022-29458", "ksp" ], "message": "Accessed passwd and/or shadow file from frontend", "selector": { "matchLabels": { "app": "frontend" } }, "file": { "severity": 10, "matchPaths": [ { "path": "/etc/passwd" }, { "path": "/etc/shadow", "fromSource": [ { "path": "/src/server" } ] } ], "action": "Audit" } } }`
