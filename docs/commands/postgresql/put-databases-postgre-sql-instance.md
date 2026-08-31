## linodectl postgresql put-databases-postgre-sql-instance

Update a PostgreSQL Managed Database

### Synopsis

Make changes to an existing PostgreSQL Managed Database. - The database's status needs to be `active`. - New values set in the `allow_list` overwrite existing values. To keep existing values, run the [List PostgreSQL Managed Databases](https://techdocs.akamai.com/linode-api/reference/get-databases-postgre-sql-instances) operation, store the `allow_list` addresses from the response, and include them with any new addresses in this operation. - Updates to your `allow_list` may take a short period of time to complete, making this operation inappropriate for rapid successive updates. - Also allows resizing the database cluster to a larger one. Clusters can't be resized to smaller plans. - All Managed Databases include automatic updates, which apply security patches to the underlying operating system of the Managed PostgreSQL Database. Use the `updates` object in this operation to modify the maintenance window for these updates. - If your database cluster is configured with a single node, downtime occurs during maintenance updates. Use the `updates` object to adjust the window to match a time that's the least disruptive to your application and users. Also consider upgrading to a [high availability](https://techdocs.akamai.com/cloud-computing/docs/aiven-database-clusters#high-availability) plan to avoid any maintenance downtime. - Major upgrades are optional until the service reaches end of service, and can be done in place. - You can't update `engine_config` advanced parameter settings for a suspended database. You'll need to [resume](https://techdocs.akamai.com/linode-api/reference/resume-databases-postgre-sql-instance) it first. - A successful request triggers a `database_update` [event](https://techdocs.akamai.com/linode-api/reference/get-events). - **Beta**. You can update an existing PostgreSQL Managed Database to move it to a Virtual Private Cloud (VPC) using the `private_network` object in the request. This support is in beta. Talk to your Akamai account team for more details. > blue-book > > Currently, VPC subnets associated with Managed Database instances don't automatically block outbound connections outside the subnet. To limit network exposure, you should configure Cloud Firewall rules to explicitly deny outbound connections beyond the intended subnet. For more details on configuring rules, see the [Cloud Firewall](https://techdocs.akamai.com/cloud-computing/docs/cloud-firewall) documentation. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Roles: `database_admin` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `databases:read_write` **CLI** ```shell linode-cli databases postgresql-update 123 \ --label example-db \ --allow_list 203.0.113.1 \ --allow_list 192.0.1.0/24 \ --type g6-standard-1 \ --updates.frequency weekly \ --updates.duration 3 \ --updates.hour_of_day 12 \ --updates.day_of_week 4 \ ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl postgresql put-databases-postgre-sql-instance [flags]
```

### Options

```
      --allow-list 0.0.0.0/0                                     Controls access to the Managed Database. - Individually included IP addresses or CIDR ranges can access the Managed Database while all other sources are blocked. - A standalone value of 0.0.0.0/0 allows all IP addresses access to the Managed Database. - An empty array (`[]`) blocks all public and private connections to the Managed Database.
      --api-version v4                                           __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --data string                                              Request body JSON, @file, or @- for stdin
      --engine-config.pg string                                  PostgreSQL-specific advanced configuration parameters.
      --engine-config.pg-stat-monitor-enable pg_stat_monitor     Enable the pg_stat_monitor extension. When this extension is enabled, PostgreSQL restarts the cluster it's in. Additionally, `pg_stat_statements` results for utility commands are unreliable.
      --engine-config.pglookout string                           Parameter used to apply PGLookout settings.
      --engine-config.shared-buffers-percentage shared_buffers   Percentage of total RAM that the database server uses for shared memory buffers. Valid range is 20-60 (float), which corresponds to 20% - 60%. This setting adjusts the shared_buffers configuration value.
      --engine-config.work-mem string                            Sets the maximum amount of memory in MB to be used by a query operation, such as a sort or hash table, before writing to temporary disk files. Default is 1MB + 0.075% of total RAM, up to 32 MB.
  -h, --help                                                     help for put-databases-postgre-sql-instance
      --label engine                                             __Filterable__ A name used to identify the Managed Database. This needs to be unique per Managed Database engine type. For example, you could use `database_1`, `database_2`, and `database_3` for three unique MySQL Managed Databases. You can also use these same names for three unique PostgreSQL Managed Databases. However, you can't repeat any of these names for either engine type.
      --postgresql-instance-id id                                The unique identifier for a PostgreSQL Managed Database. Run the [List PostgreSQL Managed Databases](https://techdocs.akamai.com/linode-api/reference/get-databases-postgre-sql-instances) operation and store the id for the desired one.
      --private-network.public-access true                       Set to true to allow clients outside of the VPC to connect to the database using a public IP address. Defaults to `false`, where only nodes within the specified `vpc_id` can access the Managed Database cluster. > blue-book > > If your Managed Database is also configured using an `allow_list`, only IP addresses set in it can access that database, even if this object is set to `true`. (default "false")
      --private-network.subnet-id vpc_id                         If the vpc_id includes subnets, you can include the one you want to limit access to the database. From the [List VPCs](https://techdocs.akamai.com/linode-api/reference/get-vpcs) operation, store the `id` for the applicable `subnets` object.
      --private-network.vpc-id id                                The unique identifier of the VPC you want to use to enable private access to the Managed Database. Run the [List VPCs](https://techdocs.akamai.com/linode-api/reference/get-vpcs) operation and store the id for the applicable VPC.
      --type price                                               Request re-sizing of your cluster to a Linode Type with more disk space. For example, you could request a Linode Type that uses a higher plan. - Needs to be a Linode Type with more disk space than your current Linode. - Resizing to a larger Linode Type can accrue additional cost. Review the price output from the [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types) operation for more information. - You can't update the `allow_list` and set a new `type` in the same request. - Any active updates to your cluster need to complete before you can request a resize. The reverse is also true: An active resizing needs to complete before you can perform any other update.
      --updates.day-of-week 1                                    The numeric reference for the day of the week to perform maintenance. 1 is Monday, `2` is Tuesday, through to `7` which is Sunday.
      --updates.duration string                                  The maximum maintenance window time in hours.
      --updates.frequency weekly                                 How frequently maintenance occurs. Currently can only be weekly. (default "weekly")
      --updates.hour-of-day string                               The hour to begin maintenance based in UTC time.
      --updates.pending string                                   __Read-only__ An array of pending updates.
      --version string                                           __Filterable__ The Managed Database engine version.
```

### Options inherited from parent commands

```
      --dry-run            Print HTTP request without sending
      --format string      Output format: json, pretty, yaml, jsonl, table, csv, raw (default "json")
      --max-retries int    Max retry attempts for 429/5xx errors (0 = no retry)
      --page-limit int     Max pages to fetch (auto-detects Link, cursor, offset, page-number schemes)
      --profile string     Configuration profile name
      --stream             Stream response line-by-line (SSE / NDJSON)
      --template string    Go template string for custom output formatting
      --transform string   GJSON expression to filter/transform JSON output
      --verbose            Log HTTP request/response details to stderr
```

### SEE ALSO

* [linodectl postgresql](./index.md)	 - postgresql operations

###### Auto generated by spf13/cobra on 31-Aug-2026
