package condition

import (
	"fmt"
	"io"
	"net/http"

	"github.com/josepht96/learning/go/serviceMonitor/cmd/api/models"
)

var r *models.FhirResource = &models.FhirResource{
	ResourceName: "Condition",
	ApiRequests:  nil,
}

func Get() (string, error) {
	resStr := fmt.Sprintf("%v route hit...\nuse %v/<number> to access api call",
		r.ResourceName, r.ResourceName)
	return resStr, nil
}

func GetID(id int) (string, error) {
	resStr := fmt.Sprintf("%v id: %d\n", r.ResourceName, id)
	resp, err := http.Get("http://localhost:4041/endpoint")
	if err != nil {
		fmt.Printf("There was` an err: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("There was` an err: %v\n", err)
	}
	fmt.Println(string(body))
	return resStr, nil
}
