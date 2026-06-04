$response = Invoke-WebRequest -Uri "http://localhost:8090/LicenseManagement/ActivateLicense"
$response.StatusCode
$response.Headers
$response.Content

Get-ChildItem "C:\Program Files (x86)\Microsoft\Edge\Application\" -ErrorAction SilentlyContinue
Get-ChildItem "C:\Program Files\Microsoft\Edge\Application\" -ErrorAction SilentlyContinue