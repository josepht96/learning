eventHubs = [
  {
    name                 = "audit",
    namespaceNameVersion = "001",
    partition_count      = 1,
    retention            = 1,
    authorization_name   = "authorization",
    listen               = true,
    send                 = true
    manage               = false
  },
  {
    name                 = "datasync",
    namespaceNameVersion = "001",
    partition_count      = 1,
    retention            = 1,
    authorization_name   = "authorization",
    listen               = true,
    send                 = true
    manage               = false
  }
]
