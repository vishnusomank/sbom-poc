package constants

//constants for Request Body
const (
	InvalidReqBody  = "Invalid Request Body"
	RequiredReqBody = "Required Request Body Missing"
	WrongReqBody    = "Wrong Request Body : "
)

//constants for Error Messages
const (
	Err400                        = "err-400"
	Err500                        = "err-500"
	ErrInAddRegistry              = "Error in Add Registry"
	ERR_IN_GET_REGISTRY_TYPE      = "Error in Get Registry Type"
	ERR_IN_GET_LIST_OF_REGISTRY   = "Error in Get List of Registry"
	ERR_IN_EDIT_REGISTRY          = "Error in Edit Registry"
	ERR_IN_CHANGE_STATUS_REGISTRY = "Error in Change Status Registry"
	ERR_IN_DELETE_REGISTRY        = "Error in Delete Registry"
	ERR_IN_GET_REGISTRY           = "Error in Get Registry"
	ERR_IN_GET_DASHBOARD          = "Error in Get Dashboard"
	ERR_IN_GET_LIST_OF_REGIONS    = "Error in Get List of Regions"
)

//constants for URL Path
const (
	ADD_REGISTRY_PATH           = "add-registry"
	GET_REGISTRY_TYPE_PATH      = "get-registry-type"
	GET_LIST_OF_REGISTRY_PATH   = "get-list-of-registry"
	EDIT_REGISTRY_PATH          = "edit-registry"
	CHANGE_STATUS_REGISTRY_PATH = "change-status-registry"
	DELETE_REGISTRY_PATH        = "delete-registry"
	GET_REGISTRY_PATH           = "get-registry"
	GET_DASHBOARD_PATH          = "get-dashboard"
	GET_LIST_OF_REGIONS_PATH    = "get-list-of-regions"
)

const (
	ACTIVE = "Active"
	DELETE = "Delete"
)
