# Prolaborate Database Migration: RDS MySQL 8.0 → 8.4

## Overview

Migrates the Prolaborate database from the old RDS MySQL 8.0.42 instance to a new Service Catalog–provisioned RDS MySQL 8.4 instance using a logical dump and load. The old instance is not modified and remains the rollback target.

| Item | Value |
|---|---|
| Old instance endpoint | `<old-endpoint>` |
| New instance endpoint | `<new-endpoint>` |
| Schema | `<schema>` |
| Database user | `admin` (same username and password on both instances) |
| Migration host | Windows server running Prolaborate |

**Downtime:** Prolaborate only. EA and Pro Cloud Server use a separate database and can remain running.

## Background and Known Differences

- **Authentication plugin:** the old instance's `admin` user uses `mysql_native_password`; the new instance uses `caching_sha2_password`. Prolaborate's bundled MySQL library must support `caching_sha2_password` (see Pre-Checks).
- **Tools:** MySQL Workbench 8.0 may crash when connecting to MySQL 8.4. All steps below use the command-line tools bundled with Workbench, which work independently of the GUI:
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysqldump.exe`
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysql.exe`
- **PowerShell redirection:** do not use `>` or `<`. Windows PowerShell writes UTF-16 with `>`, which corrupts the dump, and does not support `<`. Use `--result-file` and `source` as shown below.

To open an interactive SQL prompt against either instance:

```powershell
cd "C:\Program Files\MySQL\MySQL Workbench 8.0 CE"
.\mysql.exe -h <endpoint> -u admin -p
```

## Pre-Checks (no downtime)

### 1. Identify the schema (old instance)

```sql
SHOW DATABASES;
```

Ignore `mysql`, `information_schema`, `performance_schema`, and `sys`.

### 2. Check definers (old instance)

```sql
SELECT 'routine' AS type, ROUTINE_NAME AS name, DEFINER FROM information_schema.ROUTINES WHERE ROUTINE_SCHEMA = '<schema>'
UNION ALL
SELECT 'trigger', TRIGGER_NAME, DEFINER FROM information_schema.TRIGGERS WHERE TRIGGER_SCHEMA = '<schema>'
UNION ALL
SELECT 'event', EVENT_NAME, DEFINER FROM information_schema.EVENTS WHERE EVENT_SCHEMA = '<schema>'
UNION ALL
SELECT 'view', TABLE_NAME, DEFINER FROM information_schema.VIEWS WHERE TABLE_SCHEMA = '<schema>';
```

Any definer other than `admin` must exist on the new instance before loading.

### 3. Check authentication plugins (both instances)

```sql
SELECT user, host, plugin FROM mysql.user WHERE user = 'admin';
```

### 4. Check fallback availability (new instance)

```sql
SELECT PLUGIN_NAME, PLUGIN_STATUS FROM INFORMATION_SCHEMA.PLUGINS
WHERE PLUGIN_NAME = 'mysql_native_password';
```

If `ACTIVE`, a `mysql_native_password` fallback user can be created if needed (see Troubleshooting).

### 5. Check Prolaborate's MySQL library (optional)

In the Prolaborate install or IIS site folder (usually `bin`), check the file version of the MySQL library (right-click → Properties → Details):

- `MySql.Data.dll`: 8.x or later supports `caching_sha2_password`; 6.x does not.
- `MySqlConnector.dll`: recent versions support it.

If Prolaborate's connection string disables TLS (`SslMode=None` or `SslMode=Disabled`), `caching_sha2_password` may also require `AllowPublicKeyRetrieval=true`.

### 6. Record baseline table sizes (old instance)

```sql
SELECT TABLE_NAME, TABLE_ROWS FROM information_schema.TABLES
WHERE TABLE_SCHEMA = '<schema>' ORDER BY TABLE_ROWS DESC LIMIT 10;
```

`TABLE_ROWS` is an estimate; use it to choose tables for exact `COUNT(*)` comparison in step 5 of the migration.

## Migration (maintenance window)

### 1. Stop Prolaborate

In IIS Manager, stop the Prolaborate site and its application pool, or in an administrator PowerShell:

```powershell
Import-Module WebAdministration
Stop-Website -Name "<Prolaborate site name>"
Stop-WebAppPool -Name "<Prolaborate app pool>"
```

### 2. Snapshot the old instance

In the RDS console, take a manual snapshot of the old Prolaborate instance (e.g. `prolaborate-<env>-pre-84-migration`) and wait until it shows **Available**.

### 3. Dump from the old instance

Create `C:\temp` if it does not exist, then:

```powershell
cd "C:\Program Files\MySQL\MySQL Workbench 8.0 CE"

.\mysqldump.exe -h <old-endpoint> -u admin -p `
  --databases <schema> --single-transaction --routines --triggers --events `
  --set-gtid-purged=OFF --result-file=C:\temp\prolaborate.sql
```

Confirm `C:\temp\prolaborate.sql` exists and is not empty. If mysqldump reports that `--set-gtid-purged` is invalid because GTID is not enabled, remove that option.

### 4. Load into the new instance

```powershell
.\mysql.exe -h <new-endpoint> -u admin -p -e "source C:\temp\prolaborate.sql"
```

No output means success. If errors appear, stop and resolve them before continuing.

### 5. Verify

On **both** instances, compare:

```sql
SELECT COUNT(*) FROM information_schema.TABLES WHERE TABLE_SCHEMA = '<schema>';
SELECT COUNT(*) FROM information_schema.ROUTINES WHERE ROUTINE_SCHEMA = '<schema>';
```

For the largest tables identified in pre-check 6, compare exact counts on both instances:

```sql
SELECT COUNT(*) FROM `<schema>`.`<table>`;
```

### 6. Repoint Prolaborate

1. Back up Prolaborate's database connection config file (a copy in the same folder is fine).
2. Change only the hostname to `<new-endpoint>`. Username and password are unchanged.

### 7. Start and validate

```powershell
Start-WebAppPool -Name "<Prolaborate app pool>"
Start-Website -Name "<Prolaborate site name>"
```

- [ ] Log in to Prolaborate.
- [ ] Dashboards load.
- [ ] Users and permissions are present.
- [ ] Repository connections to EA (via Pro Cloud Server) work.

## Troubleshooting

Check for errors on the Prolaborate page, in Prolaborate's logs, or in **Event Viewer → Windows Logs → Application**.

### Authentication failure

Indicates Prolaborate's MySQL library does not support `caching_sha2_password`. Options:

- **Fallback user** (only if `mysql_native_password` is `ACTIVE` per pre-check 4), on the new instance:
  ```sql
  CREATE USER 'prolaborate'@'%' IDENTIFIED WITH mysql_native_password BY '<password>';
  GRANT ALL PRIVILEGES ON `<schema>`.* TO 'prolaborate'@'%';
  ```
  Update Prolaborate's config to use this user. This is a stopgap only: `mysql_native_password` is removed in MySQL 9.x. The long-term fix is a Prolaborate version whose MySQL library supports `caching_sha2_password`.
- **Roll back** (below) until Prolaborate is updated.

## Rollback

The old instance is not modified during the migration.

1. Stop the Prolaborate site and app pool.
2. Restore the backed-up config file.
3. Start the app pool and site, and confirm Prolaborate works against the old instance.

Any changes made in Prolaborate after cutover will not exist on the old instance.

## Post-Migration

After a soak period with no issues:

- [ ] Confirm no remaining connections to the old instance (RDS connection metrics).
- [ ] Take a final snapshot of the old instance.
- [ ] Terminate the old provisioned product **through Service Catalog**.
- [ ] Consider moving Prolaborate from the `admin` master user to a dedicated least-privilege user.

## Execution Log

| Environment | Date | Dump size | Downtime | Issues | Outcome |
|---|---|---|---|---|---|
| Test | | | | | |
| Preprod | | | | | |