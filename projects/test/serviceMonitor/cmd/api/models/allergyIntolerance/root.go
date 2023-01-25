package allergyIntolerance

import (
	"fmt"

	"github.com/josepht96/learning/go/serviceMonitor/cmd/api/models"
)

var r *models.FhirResource = &models.FhirResource{
	ResourceName: "allergyintolerance",
	ApiRequests:  nil,
}

func Get() (string, error) {
	resStr := fmt.Sprintf("%v route hit...\nuse %v/<number> to access api call",
		r.ResourceName, r.ResourceName)
	return resStr, nil
}

func GetID(id int) (string, error) {
	resStr := fmt.Sprintf("%v id: %d\n", r.ResourceName, id)
	return resStr, nil
}
