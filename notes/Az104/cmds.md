# Creating new user
Install-Module -Name AzureAD
Connect-AzureAD

$PasswordProfile = New-Object -TypeName Microsoft.Open.AzureAD.Model.PasswordProfile
$PasswordProfile.Password = "some password"
$PasswordProfile.EnforceChangePasswordPolicy = $true

New-AzureADUser -DisplayName "name" -PasswordProfile  $PasswordProfile
-UserPrincpal etc...

az ad user create --display-name "name" \
    --password ""
    --user-principal-name ""
    --force-change-password-next-login 
    --output table


# Adding member to group
Install-Module -Name AzureAD
Connect-AzureAD

New-AzureADGroup -Description "" -DisplayName "" `
 -MailEnabled $fasle -SecurityEnabled $true -MailNickName ""

 Add-AzureADGroupMember -ObjectId "" -RefObjectId ""

az ad group create --display-name Sales --mail-nickname Sales
az ad group member check --group Sales --member-id ""
az ad 
az ad group member add -group Sales --member-id ""

# Adding VM login
az vm extension set \
 --publisher Microsoft.Azure.ActiveDirectory
 --name AADLoginForWindows \
 --resource-group "" \
 --vm-name ""


# Bulk update users
Import-Module AzureAD
Connect-AzureAD 

$adGroupId = "<Azure AD Group Id here>"
$users = Get-AzureADGroupMember -ObjectId $adGroupId

foreach ($u in $users) 
{
    Write-Host $u.DisplayName
    Set-AzureADUser -ObjectId $u.Mail -Department "<New Value to update here>"
}

# Assigning a role
New-AzRoleAssignment -SignInName "email" `
    -RoleDefinitionName "Reader"
    -ResourceGroupName "test"

# Create new role
$role = Get-AzRoleDefinition "VM contributor"
$role.id = $null
$role.Name = "VM Reader"
$role.Actions.Add("Microsoft.Compute/*/read")
$role.AssignableScopes.Add("sub/sdasd")
New-AzRoleDefinition -Role $role

# Set tag on resource group and sub resources
Get-AzResource -ResourceGroupName "" | Set-AzResource Tag @{"" = ""}

# Azure policy creation
$rg = Get-AzResourceGroup -Name ""
$definition = Get-AzPolicyDefinition | Where-Object { $_.Properties.DisplayName -eq ''}
New-AzPolicyAssignment -Name "" -DisplayName "" -Scope $rg.ResourceId -PolicyDefinition $definition

# File upload
az storage file upload \ 
    --account-name ""
    --account-key ""
    --share-name ""
    --source "./test.jpg"
    --path "testDest/test.jpg"