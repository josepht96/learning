# Sparx EA Database Migration: RDS MySQL 8.0 → 8.4

## Overview

Migrates the Sparx Enterprise Architect (EA) repository database from the old RDS MySQL 8.0.42 instance to a new Service Catalog–provisioned RDS MySQL 8.4 instance using a logical dump and load. The old instance is not modified and remains the rollback target.

| Item | Value |
|---|---|
| Old instance endpoint | `<old-ea-endpoint>` |
| New instance endpoint | `<new-ea-endpoint>` |
| Schema(s) | `<schema>` |
| Database user | `admin` (confirm in pre-check 2) |
| Migration host | Windows server running Pro Cloud Server |

**Consumers:** Pro Cloud Server connects to the EA database through the MySQL ODBC driver. Prolaborate reaches the EA models through Pro Cloud Server, so it is affected while Pro Cloud Server is down.

## Background and Known Differences

- **Authentication plugin:** the new 8.4 instance uses `caching_sha2_password`. The MySQL ODBC driver used by Pro Cloud Server must be version 8.x or newer; 5.x drivers cannot authenticate.
- **Tools:** MySQL Workbench 8.0 may crash when connecting to MySQL 8.4. All steps below use the command-line tools bundled with Workbench:
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysqldump.exe`
  - `C:\Program Files\MySQL\MySQL Workbench 8.0 CE\mysql.exe`
- **PowerShell redirection:** do not use `>` or `<`. Windows PowerShell writes UTF-16 with `>`, which corrupts the dump, and does not support `<`. Use `--result-file` and `source` as shown below.

To open an interactive SQL prompt against either instance:

```powershell
cd "C:\Program Files\MySQL\MySQL Workbench 8.0 CE"
.\mysql.exe -h <endpoint> -u admin -p
```

## Pre-Checks (no downtime)

### 1. Identify the schema(s) (old instance)

```sql
SHOW DATABASES;
```

Ignore `mysql`, `information_schema`, `performance_schema`, and `sys`. There may be more than one EA repository; note every remaining schema.

### 2. Confirm how Pro Cloud Server connects

In the **Pro Cloud Server Configuration Client**, open the EA repository's **Database Manager** entry and record:

- Whether it uses an **ODBC DSN** or a **connection string**, and the host it points to.
- The **username**.

If the user is `admin`, no user needs to be created. If it is another user:

1. On the old instance, record its grants:
   ```sql
   SHOW GRANTS FOR '<user>'@'<host>';
   ```
2. On the new instance, create it with the same password and apply the grants:
   ```sql
   CREATE USER '<user>'@'<host>' IDENTIFIED WITH caching_sha2_password BY '<existing password>';
   GRANT <privileges> ON `<schema>`.* TO '<user>'@'<host>';
   ```

### 3. Check the ODBC driver version

Open the ODBC Data Source Administrator:

- 64-bit: `C:\Windows\System32\odbcad32.exe`
- 32-bit: `C:\Windows\SysWOW64\odbcad32.exe`

Find the DSN used by Pro Cloud Server (**System DSN** tab) and the driver it references (**Drivers** tab). It must be **MySQL ODBC 8.x or newer**. If it is 5.x, install a current MySQL Connector/ODBC matching Pro Cloud Server's bitness before the window.

### 4. Check definers (old instance)

```sql
SELECT 'routine' AS type, ROUTINE_NAME AS name, DEFINER FROM information_schema.ROUTINES WHERE ROUTINE_SCHEMA = '<schema>'
UNION ALL
SELECT 'trigger', TRIGGER_NAME, DEFINER FROM information_schema.TRIGGERS WHERE TRIGGER_SCHEMA = '<schema>'
UNION ALL
SELECT 'event', EVENT_NAME, DEFINER FROM information_schema.EVENTS WHERE EVENT_SCHEMA = '<schema>'
UNION ALL
SELECT 'view', TABLE_NAME, DEFINER FROM information_schema.VIEWS WHERE TABLE_SCHEMA = '<schema>';
```

EA repositories usually return nothing. Any definer other than `admin` must exist on the new instance before loading.

### 5. Record baseline counts (old instance)

```sql
SELECT 't_object' AS tbl, COUNT(*) FROM <schema>.t_object
UNION ALL SELECT 't_package',   COUNT(*) FROM <schema>.t_package
UNION ALL SELECT 't_diagram',   COUNT(*) FROM <schema>.t_diagram
UNION ALL SELECT 't_connector', COUNT(*) FROM <schema>.t_connector
UNION ALL SELECT 't_attribute', COUNT(*) FROM <schema>.t_attribute
UNION ALL SELECT 't_operation', COUNT(*) FROM <schema>.t_operation;
```

Save the output for comparison. Use exact `COUNT(*)`, not `information_schema.TABLES.TABLE_ROWS`, which is only an estimate.

## Migration (maintenance window)

### 1. Stop everything that touches the EA database

- Stop the **Pro Cloud Server** Windows service.
- Stop Prolaborate (IIS site and app pool) or notify users that EA models will be unavailable in Prolaborate during the window.
- Ask anyone using EA directly to close it.

Confirm nothing is still connected on the old instance:

```sql
SELECT id, user, host, db, command FROM information_schema.PROCESSLIST WHERE db = '<schema>';
```

### 2. Snapshot the old instance

In the RDS console, take a manual snapshot of the old EA instance (e.g. `ea-<env>-pre-84-migration`) and wait until it shows **Available**.

### 3. Dump from the old instance

Create `C:\temp` if it does not exist. List every EA schema after `--databases` if there is more than one.

```powershell
cd "C:\Program Files\MySQL\MySQL Workbench 8.0 CE"

.\mysqldump.exe -h <old-ea-endpoint> -u admin -p `
  --databases <schema> --single-transaction --routines --triggers --events `
  --set-gtid-purged=OFF --result-file=C:\temp\ea.sql
```

Confirm `C:\temp\ea.sql` exists and is not empty. If mysqldump reports that `--set-gtid-purged` is invalid because GTID is not enabled, remove that option.

### 4. Load into the new instance

```powershell
.\mysql.exe -h <new-ea-endpoint> -u admin -p -e "source C:\temp\ea.sql"
```

No output means success. If errors appear, stop and resolve them before continuing.

### 5. Verify

Run the baseline count query from pre-check 5 on the **new** instance. All counts must match exactly.

### 6. Repoint Pro Cloud Server

Update the Pro Cloud Server connection to `<new-ea-endpoint>`:

- **ODBC DSN:** edit the DSN's server field in the ODBC Data Source Administrator (matching bitness).
- **Connection string:** edit the host in the Database Manager entry.

Update any other EA DSNs on the Windows servers that point to the old endpoint. If a DSN is shared, check what else uses it.

### 7. Start and validate

Start the Pro Cloud Server service, then:

- [ ] The Database Manager entry shows as connected in the Pro Cloud Server Configuration Client.
- [ ] Open the repository in EA.
- [ ] Make and save a test edit.
- [ ] Run **Project Integrity Check**.
- [ ] Start Prolaborate if stopped, and confirm it can browse the EA models.
- [ ] Test WebEA, if used.

## Troubleshooting

### ODBC authentication error

An error mentioning `caching_sha2_password` (e.g. "authentication plugin cannot be loaded") means the ODBC driver is too old. Install a current MySQL Connector/ODBC matching Pro Cloud Server's bitness, update the DSN to use it, and retry.

## Rollback

The old instance is not modified during the migration.

1. Stop the Pro Cloud Server service.
2. Point the DSN or connection string back to `<old-ea-endpoint>`.
3. Start Pro Cloud Server and confirm the repository opens in EA.

Any changes made in EA after cutover will not exist on the old instance.

## Post-Migration

After a soak period with no issues:

- [ ] Check the old instance for remaining connections (processlist query above, or RDS connection metrics).
- [ ] Take a final snapshot of the old instance.
- [ ] Terminate the old provisioned product **through Service Catalog**.

## Execution Log

| Environment | Date | Dump size | Downtime | Issues | Outcome |
|---|---|---|---|---|---|
| Test | | | | | |
| Preprod | | | | | |