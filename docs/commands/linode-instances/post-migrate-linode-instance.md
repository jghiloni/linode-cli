## linodectl linode-instances post-migrate-linode-instance

Migrate a Linode

### Synopsis

Use this operation to move a Linode. There are two types of migration you can perform: * **Migrate to a new data center**. This applies when you want to move your Linode from one data center to another. You need to include the desired `region` for the new data center in the request. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to see all available regions and their feature capabilities. Also, see our [pricing page](https://www.akamai.com/cloud/pricing) for region-specific pricing, which applies after migration is complete. * **Initiate a pending host migration**. This applies to a migration that Akamai has scheduled for your Linode. Run the [List notifications](https://techdocs.akamai.com/linode-api/reference/get-notifications) operation to see any pending host migrations. When the migration begins, the API shuts down the Linode if it's not already off. If the migration initiated the shutdown, the API reboots the Linode when completed. Consider these points before starting any migration: - You can't migrate a Linode to a new data center when that Linode has a `vpc` purpose, legacy configuration interface or a VPC Linode interface. Either of these situations can only be migrated within the same data center. See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations. - Only next generation network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of each data center. If a VLAN is attached to your Linode and you try to migrate or clone it to a non-NGN data center, the operation won't initiate. If a Linode can't be migrated or cloned because of an incompatibility, you'll be prompted to select a different data center or contact support. - NGN data centers don't support IPv6 `/116` pools or IP Failover. If you have these features enabled on your Linode and try to migrate to an NGN data center, the migration won't initiate. If a Linode can't be migrated because of an incompatibility, you'll be prompted to select a different data center or contact support. See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations. - If the Linode is using a public Linode interface, the destination region needs to also support Linode interfaces. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to see their feature capabilities. - After migrating to a different data center, Linode public interfaces retain the same number of IP addresses, but the IP addresses themselves change. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `migrate_linode` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes migrate 123 \ --region us-central \ --placement_group.id 528 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl linode-instances post-migrate-linode-instance [flags]
```

### Options

```
      --api-version v4                            __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --data string                               Request body JSON, @file, or @- for stdin
  -h, --help                                      help for post-migrate-linode-instance
      --linode-id string                          ID of the Linode to migrate.
      --placement-group.id string                 The placement group's ID. You need to provide it for all operations that affect it.
      --region linode_migrate_datacenter_create   The region to which the Linode will be migrated. Must be a valid region slug. A list of regions can be viewed by running the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation. A cross data center migration will cancel a pending migration that has not yet been initiated. A cross data center migration will initiate a linode_migrate_datacenter_create event.
      --type warm                                 Type of migration used in moving to a new host or Linode type. warm: the Linode will not power down until the migration is complete. Warm migrations are not available for DC migrations. `cold`: the Linode will be powered down and migrated. When the migration is complete, the Linode will be powered on. (default "cold")
      --upgrade true                              When initiating a cross DC migration, setting this value to true will also ensure that the Linode is upgraded to the latest generation of hardware that corresponds to your Linode's Type, if any free upgrades are available for it. If no free upgrades are available, and this value is set to `true`, then the endpoint will return a 400 error code and the migration will not be performed. If the data center set in the `region` field does not allow upgrades, then the endpoint will return a 400 error code and the migration will not be performed. (default "false")
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

* [linodectl linode-instances](./index.md)	 - linode-instances operations

###### Auto generated by spf13/cobra on 31-Aug-2026
