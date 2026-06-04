$response = Invoke-WebRequest -Uri "http://localhost:8090/LicenseManagement/ActivateLicense"
$response.StatusCode
$response.Headers
$response.Content

Get-ChildItem "C:\Program Files (x86)\Microsoft\Edge\Application\" -ErrorAction SilentlyContinue
Get-ChildItem "C:\Program Files\Microsoft\Edge\Application\" -ErrorAction SilentlyContinue


# Step 1 - Login and capture the session cookie

$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession

Invoke-WebRequest -Uri "http://localhost:8090/account/login" `
  -Method POST `
  -Body @{ username = "admin"; password = "yourpassword" } `
  -SessionVariable "session"

# Step 2 - Hit the license endpoint with the session

Invoke-WebRequest -Uri "http://localhost:8090/api/license/getlicenseinfo" `
  -WebSession $session | Select-Object -ExpandProperty Content

$cred = New-Object System.Management.Automation.PSCredential("admin", (ConvertTo-SecureString "password" -AsPlainText -Force))

Invoke-WebRequest -Uri "http://localhost:8090/api/license/getlicenseinfo" `
  -Credential $cred | Select-Object -ExpandProperty Content