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


  # Try common OIDC token endpoints
curl.exe -v -X POST http://localhost:8090/connect/token `
  -d "grant_type=password&username=admin&password=yourpassword&client_id=Prolaborate-spa"


$body = @{
    grant_type = 'password'
    username   = 'admin'
    password   = 'yourpassword'
    client_id  = 'Prolaborate-spa'
}

$response = Invoke-WebRequest -Uri 'http://localhost:8090/connect/token' `
    -Method POST `
    -Body $body

$token = ($response.Content | ConvertFrom-Json).access_token


Invoke-WebRequest -Uri 'http://localhost:8090/api/license/getlicenseinfo' `
    -Headers @{ Authorization = "Bearer $token" } | Select-Object -ExpandProperty Content


# Windows Machine GUID (most common for licensing)
Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Cryptography" -Name MachineGuid

# Product ID
Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion" -Name ProductId

# BIOS/Hardware UUID
Get-WmiObject Win32_ComputerSystemProduct | Select-Object UUID

# Motherboard serial
Get-WmiObject Win32_BaseBoard | Select-Object SerialNumber

# Volume serial of C drive
Get-WmiObject Win32_LogicalDisk -Filter "DeviceID='C:'" | Select-Object VolumeSerialNumber