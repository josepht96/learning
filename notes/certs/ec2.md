General Purpose — M and T families

T series (t3, t4g) — burstable, cheapest option. Good for low-traffic workloads that occasionally spike. Credit-based CPU model means sustained high CPU will throttle you
M series (m6i, m7i) — balanced CPU/memory ratio, the workhorse. What you'd use for most standard workloads

Compute Optimized — C family

C series (c6i, c7i) — higher CPU to memory ratio, more vCPUs for the cost. Good for compute-heavy workloads like batch processing, gaming servers, high-traffic web servers

Memory Optimized — R and X families

R series (r6i, r7i) — high memory to CPU ratio. Good for in-memory databases, caching, real-time analytics
X series (x2iedn) — extreme memory, up to several TB. SAP HANA, large in-memory databases

Storage Optimized — I and D families

I series (i4i) — high IOPS NVMe local storage. Good for NoSQL databases, data warehousing
D series — dense HDD storage, high throughput. Hadoop/distributed file systems

Accelerated Computing — P and G families

P series (p4, p5) — GPU instances for ML training
G series (g4, g5) — GPU for inference and graphics workloads

m  6  i  .  2xlarge
↑  ↑  ↑      ↑
|  |  |      size
|  |  processor (i=Intel, a=AMD, g=Graviton/ARM)
|  generation
family