# Databases overivew

## Indexes

Columns that are frequently searched can be stored in an index which allows for faster look ups. Assume you had a table with an id field that people often searched by. To find the proper id, an unindexed db would potentially need to check every row. An index takes the ids, sorts them in a binary tree - allowing for log2 lookup time. Binary search much faster. Takes up a lot of space so don't use them everywhere. Also not that helpful for some data (like gender or something) that may only have a few values.

## Sharding

## Replication

## Failover

## CAP

### Consistency patterns

Weak consistency - writes are updated across databases on a best effort basis.
Usually pretty fast, but best used in situations where losing some data is ok (video games, phone calls). Data may never reach replicas

Eventual consistency - Makes guarantees about eventually delivering the data
to replica databases, but does Read requests may not get the latest. Very good for systems that are not time critcal but eventual consistency is good (DNS, many websites).

Strong consistency - All replicas see the latest write data. Good for systems where data consistency is critical (transactions).

### Availability Patterns

Failover - With active-passive fail-over, heartbeats are sent between the active and the passive server on standby. If the heartbeat is interrupted, the passive server takes over the active's IP address and resumes service.
The length of downtime is determined by whether the passive server is already running in 'hot' standby or whether it needs to start up from 'cold' standby. Only the active server handles traffic.
Failover adds complexity to the system because heart beat and failover must be managed. Data can be lost if the master server goes down before the failover takes over write operations.

Replication - Sequence vs parallel. A services availability is dependent upon the availability of its individual components. If two components have an uptme of 99.9%, then total availability is only 99.8.
