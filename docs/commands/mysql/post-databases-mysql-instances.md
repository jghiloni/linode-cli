## linodectl mysql post-databases-mysql-instances

Create or restore a MySQL Managed Database

### Synopsis

**Provision a MySQL Managed Database** Use this operation to create a new MySQL Managed Database. - New instances can take 10 to 15 minutes to deploy. - When you create a new MySQL Managed Database, our partner [Aiven](https://aiven.io/docs/platform/concepts/cloud-security#data-encryption) automatically enables disk encryption on each cluster. - All Managed Databases include automatic, daily backups. Up to seven backups are automatically stored for each Managed Database, providing restore points for each day of the past week. - All Managed Databases include automatic updates, which apply security patches to the underlying operating system of the MySQL Managed Database. Configure the maintenance window for these updates with the [Update a managed MySQL database](https://techdocs.akamai.com/linode-api/reference/put-databases-mysql-instance) operation. - If your database cluster is configured with a single node, downtime occurs during maintenance updates. You should adjust the window to match a time that's the least disruptive to your application and users. Also consider upgrading to a [high availability](https://techdocs.akamai.com/cloud-computing/docs/aiven-database-clusters#high-availability) plan to avoid any maintenance downtime. - Major upgrades are optional until the service reaches end of service, and can be done in place. - A successful request triggers a `database_create` [event](https://techdocs.akamai.com/linode-api/reference/get-events). **Beta** **Virtual Private Cloud (VPC) support** You can create a MySQL Managed Database in a VPC using the `private_network` object in the request. Talk to your Akamai account team for more details. > blue-book > > Currently, VPC subnets associated with Managed Database instances don't automatically block outbound connections outside the subnet. To limit network exposure, you should configure Cloud Firewall rules to explicitly deny outbound connections beyond the intended subnet. For more details on configuring rules, see the [Cloud Firewall](https://techdocs.akamai.com/cloud-computing/docs/cloud-firewall) documentation. **Restore a MySQL Managed Database** Include the `fork` object in the request to target a backed-up database. The target MySQL database's status can be `active`, `degraded`, or `failed`. > blue-book > > Restoring from a backup creates a second running cluster, which incurs billing. Delete the first cluster after the restore is complete, to avoid this billing. > thumbs-up There's a tutorial > > We offer an example API workflow you can follow to [restore a Managed Database backup](https://techdocs.akamai.com/linode-api/reference/restore-a-managed-database-backup). **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Roles: `account_database_creator` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `databases:read_write` **CLI** ```shell linode-cli databases mysql-create \ --label example-db1 \ --region us-east \ --type g6-dedicated-2 \ --cluster_size 3 \ --engine mysql/8.0.26 \ --engine_config.binlog_retention_period 600 \ --engine_config.mysql.connect_timeout 10 \ --engine_config.mysql.default_time_zone +03:00 \ --ssl_connection true \ --allow_list 203.0.113.1 \ --allow_list 192.0.1.0/24 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl mysql post-databases-mysql-instances [flags]
```

### Options

```
      --allow-list 0.0.0.0/0                           Controls access to the Managed Database. - Individually included IP addresses or CIDR ranges can access the Managed Database while all other sources are blocked. - A standalone value of 0.0.0.0/0 allows all IP addresses access to the Managed Database. - An empty array (`[]`) blocks all public and private connections to the Managed Database.
      --api-version v4                                 __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --cluster-size 1                                 The number of Linode instance nodes deployed to the Managed Database, from 1 to `3`. Consider these points when setting a `cluster_size`: - Choose `3` nodes to create a high availability cluster that consists of one `primary` node and two `standby` nodes. - A `2` node cluster is only available with a dedicated plan. It consists of one `primary` node and one `standby` node. (default "1")
      --data string                                    Request body JSON, @file, or @- for stdin
      --engine string                                  The Managed Database engine in engine/version format.
      --engine-config.binlog-retention-period binlog   The minimum amount of time in seconds to keep binlog entries before deletion. This may be extended for services that require `binlog` entries for longer than the default, for example if using the MySQL Debezium Kafka connector.
      --engine-config.mysql string                     MySQL-specific advanced configuration parameters.
      --fork.restore-time string                       A specific database timestamp to restore from.
      --fork.source id                                 The unique instance id for the database to fork from. Run the [List Managed Databases](https://techdocs.akamai.com/linode-api/reference/get-databases-instances) operation and store the unique id for the target Managed Database.
  -h, --help                                           help for post-databases-mysql-instances
      --label engine                                   __Filterable__ A name used to identify the Managed Database. This needs to be unique per Managed Database engine type. For example, you could use `database_1`, `database_2`, and `database_3` for three unique MySQL Managed Databases. You can also use these same names for three unique PostgreSQL Managed Databases. However, you can't repeat any of these names for either engine type.
      --private-network.public-access true             Set to true to allow clients outside of the VPC to connect to the database using a public IP address. Defaults to `false`, where only nodes within the specified `vpc_id` can access the Managed Database cluster. > blue-book > > If your Managed Database is also configured using an `allow_list`, only IP addresses set in it can access that database, even if this object is set to `true`. (default "false")
      --private-network.subnet-id vpc_id               If the vpc_id includes multiple subnets, specify the `subnet_id` you want to use to control access to the database. Use the [List VPCs](https://techdocs.akamai.com/linode-api/reference/get-vpcs) operation to find and store the `id` of the relevant subnet object. > blue-book > > A VPC needs at least one subnet to assign a Manage Database instance.
      --private-network.vpc-id id                      The unique identifier of the VPC you want to use to enable private access to the Managed Database. Run the [List VPCs](https://techdocs.akamai.com/linode-api/reference/get-vpcs) operation and store the id for the applicable VPC.
      --region string                                  __Filterable__ The unique identifier for the [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the Managed Database lives.
      --ssl-connection true                            Currently required to be true. Whether to require SSL credentials to establish a connection to the Managed Database. Run the [Get managed MySQL database credentials](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-credentials) operation for access information. (default "true")
      --type string                                    __Filterable__ The Linode Instance type used by the Managed Database for its nodes.
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

* [linodectl mysql](./index.md)	 - mysql operations

###### Auto generated by spf13/cobra on 31-Aug-2026
