$response = Invoke-WebRequest -Uri "http://localhost:8090/LicenseManagement/ActivateLicense"
$response.StatusCode
$response.Headers
$response.Content