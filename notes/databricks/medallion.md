# Bronze, Silver, Gold Data Architecture

This is a **medallion architecture** pattern commonly used in Databricks (and data lakehouses generally) to organize data by quality and refinement level.

## Bronze Layer (Raw/Landing)

**Purpose**: Store raw data exactly as ingested from source systems

**Characteristics**:

- Data arrives in original format (JSON, CSV, Parquet, etc.)
- Minimal or no transformation
- Includes all fields, even ones you might not need
- May contain duplicates, malformed records, or bad data
- Append-only - never delete original data
- Often includes metadata like ingestion timestamp, source system ID

**Example Use Cases**:

- Landing zone for API responses
- Database CDC (change data capture) streams
- Log files from applications
- IoT sensor data dumps

**Why Keep It**:

- Audit trail - can always go back to source truth
- Re-process data if transformation logic changes
- Debug data quality issues
- Compliance/regulatory requirements

```python
# Bronze example - minimal transformation
raw_events = (
    spark.readStream
    .format("cloudFiles")
    .option("cloudFiles.format", "json")
    .load("/source/events/")
    .withColumn("ingestion_time", current_timestamp())
    .writeStream
    .format("delta")
    .outputMode("append")
    .table("bronze.raw_events")
)
```

## Silver Layer (Cleaned/Validated)

**Purpose**: Cleaned, deduplicated, validated data ready for business logic

**Characteristics**:

- Schema enforcement and validation
- Data type conversions and standardization
- Deduplication
- Null handling and data quality checks
- Basic filtering (remove obviously bad records)
- Still fairly granular - not heavily aggregated
- May join multiple bronze sources

**Example Transformations**:

- Parse JSON into structured columns
- Convert string dates to timestamps
- Standardize country codes, phone formats
- Remove test/dev data
- Handle late-arriving data
- Enforce business rules (valid email format, positive quantities)

**Why It Exists**:

- Separates data quality concerns from business logic
- Reusable foundation for multiple downstream use cases
- Faster queries than bronze (better file organization, fewer nulls)
- Still detailed enough for most analytics

```python
# Silver example - cleaning and validation
silver_events = (
    spark.readStream
    .table("bronze.raw_events")
    .filter(col("user_id").isNotNull())
    .withColumn("event_date", to_date(col("timestamp")))
    .dropDuplicates(["event_id"])
    .withColumn("amount", col("amount").cast("decimal(10,2)"))
    .filter(col("amount") > 0)
    .writeStream
    .format("delta")
    .outputMode("append")
    .option("mergeSchema", "true")
    .table("silver.events")
)
```

## Gold Layer (Business/Aggregated)

**Purpose**: Highly refined, aggregated data optimized for specific business use cases

**Characteristics**:

- Aggregations, calculations, and business metrics
- Denormalized for query performance
- Often dimensions and fact tables (star schema)
- Filtered to specific business domains
- May combine multiple silver tables
- Optimized for BI tools and dashboards

**Example Transformations**:

- Daily/monthly revenue rollups
- Customer lifetime value calculations
- Cohort analysis tables
- Product recommendation features
- Executive dashboards
- ML feature stores

**Why It Exists**:

- Fast queries for end users (pre-aggregated)
- Business-friendly column names and structure
- Different gold tables for different teams/use cases
- Isolation - analytics users don't hit raw data

```python
# Gold example - business aggregations
daily_revenue = (
    spark.read.table("silver.orders")
    .join(spark.read.table("silver.customers"), "customer_id")
    .groupBy("order_date", "customer_segment")
    .agg(
        sum("order_total").alias("total_revenue"),
        count("order_id").alias("order_count"),
        countDistinct("customer_id").alias("unique_customers")
    )
    .write
    .format("delta")
    .mode("overwrite")
    .partitionBy("order_date")
    .table("gold.daily_revenue_by_segment")
)
```

## Flow Between Layers

```
Source Systems → Bronze (raw) → Silver (cleaned) → Gold (aggregated)
```

Each layer is typically implemented as Delta Live Tables pipelines or scheduled jobs that read from the previous layer.

## Key Principles

**Incremental Processing**: Each layer processes only new/changed data, not full reloads

**Idempotency**: Re-running the same transformation produces the same result

**Schema Evolution**: Handle new fields without breaking existing pipelines

**Data Quality**: Each layer validates and improves quality progressively

**Separation of Concerns**:

- Bronze = ingestion
- Silver = data quality and standardization  
- Gold = business logic and optimization

## Real World Example

**E-commerce Order Pipeline**:

**Bronze**: Raw JSON from order API, includes system fields, duplicates from retries

```
{"order_id": "123", "user": "456", "items": [...], "meta": {...}}
```

**Silver**: Parsed, validated orders with customer and product references

```
order_id | customer_id | order_date | item_count | subtotal | tax | total
```

**Gold**: Daily sales dashboard table

```
date | region | category | orders | revenue | avg_order_value | returning_customers_pct
```

Different gold tables might exist for finance (revenue analysis), marketing (customer segments), and operations (fulfillment metrics) - all built from the same silver layer.

## Common Variations

Some teams use **Platinum** for even more specific use cases (ML features, executive reports)

Some teams skip silver for simple pipelines (bronze → gold directly)

Some teams have multiple silver layers (silver_raw, silver_curated)

The key is consistency - pick a pattern and stick with it across your organization.
