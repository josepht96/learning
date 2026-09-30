# RDS MySQL 8.0 → 8.4 Migration Runbook (Sparx Enterprise Architect)

## Overview

This runbook covers migrating the Sparx Enterprise Architect (EA) repository databases from Amazon RDS for MySQL Community 8.0 (currently 8.0.42) to 8.4.

| Environment | RDS instances | Consumers |
|---|---|---|
| Test | 2 × MySQL 8.0.42 | Sparx EA / Pro Cloud Server on Windows servers |
| Preprod | 2 × MySQL 8.0.42 | Sparx EA / Pro Cloud Server on Windows servers |

Complete the full procedure in **Test** first, record timings and issues, then repeat in **Preprod**.

## Approach

The RDS instances are provisioned through **AWS Service Catalog**, which offers MySQL 8.4 at launch. The migration therefore uses a side-by-side approach:

1. Provision a **new 8.4 instance from the Service Catalog**.
2. Stop Sparx, then **dump and load** the EA schema from the old instance into the new one.
3. **Repoint** the Sparx connections to the new endpoint and validate.
4. Keep the old 8.0 instance untouched as the rollback target until the migration is accepted.

### Why this approach

- The new instance stays **catalog-managed** (no CloudFormation drift).
- Avoids in-place upgrade risks in the catalog template (`AllowMajorVersionUpgrade`, parameter groups hardcoded to the `mysql8.0` family).
- **Rollback is a connection change only**; the old instance is never modified.
- EA repositories are small and lightly used, so the downtime for a logical dump and load is short.

### Alternatives considered

| Option | Why not chosen |
|---|---|
| RDS Blue/Green deployment | Green is read-only, so write testing is limited before cutover; interacts with a catalog-managed stack. |
| In-place upgrade via RDS console | Causes drift from the Service Catalog / CloudFormation stack; rollback requires snapshot restore. |
| Update the provisioned product to 8.4 | Depends on template support for major version upgrades; risk of instance replacement if properties change. |
| Snapshot restore in the RDS console | Resulting instance is not catalog-managed; a restore cannot change the major version. |

## Key Changes in MySQL 8.4

- **Authentication:** `caching_sha2_password` is the default. `mysql_native_password` is deprecated and disabled by default in community MySQL 8.4. Older MySQL ODBC drivers (5.x) do not support `caching_sha2_password`.
- **Parameter groups:** 8.4 requires a `mysql8.4`-family parameter group. `default_authentication_plugin` is replaced by `authentication_policy`.
- **Replication syntax:** `MASTER`/`SLAVE` statements and variables are removed in favor of `SOURCE`/`REPLICA`.
- **InnoDB defaults:** several defaults changed (e.g. `innodb_adaptive_hash_index`, `innodb_change_buffering`). Monitor performance after cutover.

## Prerequisites

- [ ] Confirm whether the two instances in each environment are independent databases or a primary/replica pair. If a replica exists, recreate it from the new instance after cutover.
- [ ] Confirm the Sparx EA and Pro Cloud Server versions support MySQL 8.4 (check Sparx documentation).
- [ ] Identify any other consumers of these databases (reporting, backup jobs, monitoring) that will also need repointing.
- [ ] Confirm a migration host with network access to both old and new instances. MySQL Workbench on the Windows servers includes the required client tools:
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysqldump.exe`
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysql.exe`
- [ ] Schedule a maintenance window and notify EA users.

## Runbook

Repeat for each RDS instance.

### Phase 1: Preparation (before the window)

**1. Provision the new instance.**
Launch a new product from the Service Catalog with engine version **8.4.x**. Match the old instance's settings: instance class, storage, subnets, and security groups (the Windows servers must be able to reach it).

**2. Update the ODBC driver on the Windows servers.**
Install or upgrade MySQL Connector/ODBC to **8.x or later**, matching EA's bitness (32-bit EA requires the 32-bit driver and DSN). Confirm Sparx still connects to the current 8.0 instances with the new driver.

**3. Record the schema and user grants on the old instance.**

```sql
SHOW DATABASES;
SHOW GRANTS FOR 'ea_user'@'%';
SELECT user, host, plugin FROM mysql.user;
```

**4. Create the EA user on the new instance** using the grants from step 3:

```sql
CREATE USER 'ea_user'@'%' IDENTIFIED WITH caching_sha2_password BY '<password>';
GRANT <privileges> ON ea_repo.* TO 'ea_user'@'%';
```

**5. Record baseline row counts on the old instance:**

```sql
SELECT 't_object' AS tbl, COUNT(*) FROM ea_repo.t_object
UNION ALL SELECT 't_package', COUNT(*) FROM ea_repo.t_package
UNION ALL SELECT 't_diagram', COUNT(*) FROM ea_repo.t_diagram
UNION ALL SELECT 't_connector', COUNT(*) FROM ea_repo.t_connector;
```

### Phase 2: Migration (during the window)

**6. Stop Sparx services.**
Stop Pro Cloud Server and any EA services. Ask users to close EA. Confirm no active sessions on the old instance:

```sql
SHOW PROCESSLIST;
```

**7. Take a manual snapshot** of the old instance (e.g. `ea-<env>-pre-84-migration`).

**8. Dump the EA schema from the old instance.**

> **Note:** Do not use `>` or `<` redirection in PowerShell. Windows PowerShell writes UTF-16 with `>`, which corrupts the dump, and does not support `<`. Use `--result-file` and `source` as shown.

```powershell
cd "C:\Program Files\MySQL\MySQL Workbench 8.0 CE"

.\mysqldump.exe -h <old-endpoint> -u admin -p `
  --databases ea_repo --single-transaction --routines --triggers --events `
  --set-gtid-purged=OFF --result-file=C:\temp\ea_repo.sql
```

If mysqldump reports that `--set-gtid-purged` is invalid because GTID is not enabled, remove that option.

**9. Load the dump into the new instance.**

```powershell
.\mysql.exe -h <new-endpoint> -u admin -p -e "source C:\temp\ea_repo.sql"
```

**10. Verify row counts** on the new instance using the query from step 5 and compare with the baseline.

#### GUI alternative (MySQL Workbench)

- **Export:** connect to the old instance → *Server → Data Export* → select the EA schema → check *Dump Stored Procedures and Functions*, *Dump Events*, *Dump Triggers* → *Export to Self-Contained File* → check *Include Create Schema* → *Start Export*. If a GTID error occurs, set `set-gtid-purged` to `OFF` under *Advanced Options*.
- **Import:** connect to the new instance → *Server → Data Import* → *Import from Self-Contained File* → select the file → *Start Import*.

Workbench may warn about an unsupported server version when connecting to 8.4; this can be dismissed.

### Phase 3: Cutover and validation

**11. Repoint Sparx.**
Update the ODBC DSNs and the Pro Cloud Server connection configuration on the Windows servers to the new endpoint.

**12. Start services and validate.**

- [ ] Start Pro Cloud Server and EA services.
- [ ] Open each repository in EA.
- [ ] Make and save a test edit.
- [ ] Run *Project Integrity Check*.
- [ ] Test WebEA, if used.
- [ ] Check the new instance's RDS error log for authentication or other errors.

## Rollback

The old 8.0 instance is not modified during the migration.

1. Stop Pro Cloud Server and EA services.
2. Revert the ODBC DSNs and Pro Cloud Server configuration to the old endpoint.
3. Start services and confirm EA opens the repository.

Any changes made in EA on the new instance after cutover will not exist on the old instance.

## Post-Migration Cleanup

After a soak period with no issues:

- [ ] Review RDS connection metrics on the old instance to confirm no remaining clients.
- [ ] Take a final snapshot of the old instance.
- [ ] Terminate the old provisioned product **through Service Catalog** (not the RDS console) so the CloudFormation stack is removed.
- [ ] Recreate any read replica from the new instance, if applicable.
- [ ] Update documentation with the new endpoints.

> The old instances continue to incur RDS Extended Support charges for MySQL 8.0 until they are deleted.

## Execution Log

| Environment | Instance | Date | Downtime | Dump/load duration | Issues | Outcome |
|---|---|---|---|---|---|---|
| Test | | | | | | |
| Test | | | | | | |
| Preprod | | | | | | |
| Preprod | | | | | | |