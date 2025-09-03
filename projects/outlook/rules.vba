Sub CreateEmailRule()
    Dim olApp As Outlook.Application
    Dim olRules As Outlook.Rules
    Dim olRule As Outlook.Rule
    Dim olInbox As Outlook.Folder
    Dim olMoveToFolder As Outlook.Folder
    Dim olCondition As Outlook.RuleCondition
    Dim olAction As Outlook.RuleAction
    
    ' Initialize Outlook application
    Set olApp = Outlook.Application
    Set olRules = olApp.Session.DefaultStore.GetRules()
    
    ' Get Inbox and target folder
    Set olInbox = olApp.Session.GetDefaultFolder(olFolderInbox)
    Set olMoveToFolder = olInbox.Folders("Test456")
    
    ' Create new rule
    Set olRule = olRules.Create("Move Test Emails", olRuleReceive)
    
    ' Condition: From test123@test.com
    Set olCondition = olRule.Conditions.SenderAddress
    If Not olCondition Is Nothing Then
        olCondition.Enabled = True
        olCondition.Address = Array("test123@test.com")
    End If
    
    ' Condition: Subject contains "test123"
    Set olCondition = olRule.Conditions.Subject
    If Not olCondition Is Nothing Then
        olCondition.Enabled = True
        olCondition.Text = Array("test123")
    End If
    
    ' Action: Move to folder Test456
    Set olAction = olRule.Actions.MoveToFolder
    If Not olAction Is Nothing Then
        olAction.Enabled = True
        Set olAction.Folder = olMoveToFolder
    End If
    
    ' Save and update rules
    olRules.Save
    
    ' Cleanup
    Set olAction = Nothing
    Set olCondition = Nothing
    Set olRule = Nothing
    Set olMoveToFolder = Nothing
    Set olInbox = Nothing
    Set olRules = Nothing
    Set olApp = Nothing
    
    MsgBox "Rule created successfully!", vbInformation
End Sub
