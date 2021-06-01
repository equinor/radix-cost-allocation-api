package models

// Whitelist contains list of apps that should not be part of cost distribution
type Whitelist struct {
	// List is the list of apps
	//
	// required: true
	List []string `json:"whiteList"`
}
