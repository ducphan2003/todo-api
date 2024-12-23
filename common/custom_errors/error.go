package custom_errors

import "todo-api/common"

func MissingProperty(property string) error {
	return common.NewCustomError(nil, "Missing property: "+property, "ERR_MISSING_PROPERTY")
}

func InvalidProperty(property string) error {
	return common.NewCustomError(nil, "Invalid property: "+property, "ERR_INVALID_PROPERTY")
}
