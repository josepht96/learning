package queries

//Tenant struct
type Tenant struct {
	ID   string
	Name string
}

//IdnTenant struct
type IdnTenant struct {
	ID          string
	TenantID    string
	Environment string
}
