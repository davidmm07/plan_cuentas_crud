package appmessagemanager

// InsertErrorMessage ... returns error message for DB insertion error.
func InsertErrorMessage() string {
	return "DB Insert Error"
}

// ParamsErrorMessage ... returns error message for Params convertion.
func ParamsErrorMessage() string {
	return "Some params cannot be converted"
}

// DeleteErrorMessage ... returns error message for DB insertion error.
func DeleteErrorMessage() string {
	return "DB Delete Error"
}

//AprobarpresupuestoErrorMessage
func AprobarPresupuestoErrorMessage() string {
	return "Error in AprobarPresupuesto function"
}