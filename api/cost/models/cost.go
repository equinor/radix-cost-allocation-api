package models

// Cost details of cost
// swagger:model Cost
type Cost struct {
	// ApplicationName the name of the application
	//
	// required: false
	// example: radix-canary-golang
	ApplicationName string `json:"name"`

	// ApplicationOwner of the application (email). Can be a single person or a shared group email
	//
	// required: false
	ApplicationOwner string `json:"owner"`

	// ApplicationCreator of the application (user principle name).
	//
	// required: false
	ApplicationCreator string `json:"creator"`
}
