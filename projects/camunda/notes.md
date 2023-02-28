Removing port on zeebe_address does nothing
Removiong zeebe_address altogeher triggers error
Making zeebe_address wrong does nothing


This error occurs when Tasklist cannot communicate with zeebe-gateway
WARN 1 --- [ault-executor-3] i.c.z.c.i.ZeebeCallCredentials           : The request's security level does not guarantee that the credentials will be confidential.

This error occurs when there is an issue with keycloak communication
INFO 1 --- [           main] i.c.t.z.PartitionHolder                  : Partition ids can't be fetched from Zeebe. Try next round (2).