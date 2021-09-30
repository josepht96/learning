$connectTestResult = Test-NetConnection -ComputerName testfileshare19.file.core.windows.net -Port 445
if ($connectTestResult.TcpTestSucceeded) {
    # Save the password so the drive will persist on reboot
    cmd.exe /C "cmdkey /add:`"testfileshare19.file.core.windows.net`" /user:`"localhost\testfileshare19`" /pass:`"DAcnDq59AC28TyiVgOsNdDA/nTfo/FWPKjQdgAgCyIDclfp+DnqzOBd9exaipXUCz4oFQ0ur2DhXqGiAOz5iSQ==`""
    # Mount the drive
    New-PSDrive -Name Z -PSProvider FileSystem -Root "\\testfileshare19.file.core.windows.net\testfilehsare2" -Persist
} else {
    Write-Error -Message "Unable to reach the Azure storage account via port 445. Check to make sure your organization or ISP is not blocking port 445, or use Azure P2S VPN, Azure S2S VPN, or Express Route to tunnel SMB traffic over a different port."
}