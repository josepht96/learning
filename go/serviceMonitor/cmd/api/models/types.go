package models

//Condition struct
type FhirResource struct {
	ResourceName string
	ApiRequests  []ApiRequest
}

//ApiRequest struct
type ApiRequest struct {
	ID    int
	Route string
}
