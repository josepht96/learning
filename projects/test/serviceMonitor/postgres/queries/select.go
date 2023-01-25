package queries

import (
	"database/sql"
	"fmt"
)

//SelectIdnTenants func
func SelectIdnTenants(db *sql.DB) {
	queryStr := "SELECT id, tenant_id, environment FROM public.idntenants;"
	rows, err := db.Query(queryStr)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var idnTenantArr []IdnTenant
	for rows.Next() {
		var t IdnTenant
		err = rows.Scan(&t.ID, &t.TenantID, &t.Environment)
		if err != nil {
			// handle this error
			panic(err)
		}
		idnTenantArr = append(idnTenantArr, t)
		fmt.Println(t)
	}
}

//SelectTenants func
func SelectTenants(db *sql.DB) {
	queryStr := "SELECT id, name FROM public.tenants"
	rows, err := db.Query(queryStr)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var tenantArr []Tenant
	for rows.Next() {
		var t Tenant
		err = rows.Scan(&t.ID, &t.Name)
		if err != nil {
			// handle this error
			panic(err)
		}
		tenantArr = append(tenantArr, t)
		fmt.Println(t)
	}
}
