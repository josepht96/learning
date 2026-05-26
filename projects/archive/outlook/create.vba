Sub EnsureFoldersExist()
    Dim outlookApp As Outlook.Application
    Dim namespace As Outlook.Namespace
    Dim inbox As Outlook.Folder
    Dim test1Folder As Outlook.Folder
    Dim test2Folder As Outlook.Folder
    Dim folderExists As Boolean
    
    Set outlookApp = New Outlook.Application
    Set namespace = outlookApp.GetNamespace("MAPI")
    Set inbox = namespace.GetDefaultFolder(olFolderInbox)
    
    folderExists = False
    
    ' Check if "test1" exists
    On Error Resume Next
    Set test1Folder = inbox.Folders("test1")
    On Error GoTo 0
    If test1Folder Is Nothing Then
        inbox.Folders.Add "test1", olFolderInbox
        folderExists = True
    End If
    
    ' Check if "test2" exists
    On Error Resume Next
    Set test2Folder = inbox.Folders("test2")
    On Error GoTo 0
    If test2Folder Is Nothing Then
        inbox.Folders.Add "test2", olFolderInbox
        folderExists = True
    End If
    
    ' Exit if both folders exist
    If Not folderExists Then Exit Sub
    
    MsgBox "Folders created successfully!", vbInformation
End Sub