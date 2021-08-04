package registry

//Sends registration to registry server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//RegisterService func
func RegisterService(reg Registration) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(reg)
	fmt.Printf("Json obj: %v\n", buffer)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "application/json", buffer)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service: %v", res.StatusCode)
	}
	return nil
}

//ShutDownService func
func ShutDownService(ServiceURL string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL,
		bytes.NewBuffer([]byte(ServiceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service: %v", res.StatusCode)
	}
	return err
}
