## linodectl linode-instances put-linode-instance

Update a Linode

### Synopsis

Updates a Linode. > blue-book Updating `tags`? > > If you include the `tags` object in this operation to modify the Linode's tags, you need to include existing `tags` to maintain them. You can run the [List tagged objects](https://techdocs.akamai.com/linode-api/reference/get-tagged-objects) operation to view your Linode, based on its `label` and store its current `tags`. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_linode` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes update 7833080 \ --label linode123 \ --backups.schedule.day "Saturday" \ --backups.schedule.window "W22" \ --alerts.cpu 180 \ --alerts.network_in 10 \ --alerts.network_out 10 \ --alerts.transfer_quota 80 \ --alerts.io 10000 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl linode-instances put-linode-instance [flags]
```

### Options

```
      --alerts.cpu 180                                              The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we'll send you an alert. Your Linode's total CPU capacity is represented as 100%, multiplied by its number of cores. For example, a two-core Linode's CPU capacity is represented as 200%. If you want to be alerted at 90% of a two-core Linode's CPU capacity, set the alert value to 180. The default value is 90% multiplied by the number of cores. If the value is set to `0` (zero), the alert is disabled.
      --alerts.io 0                                                 The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we'll send you an alert. If set to 0 (zero), this alert is disabled.
      --alerts.network-in 0                                         The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
      --alerts.network-out 0                                        The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we'll send you an alert. If this is set to 0 (zero), the alert is disabled.
      --alerts.transfer-quota 0                                     The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we'll alert you. If this is set to 0 (zero), the alert is disabled.
      --api-version v4                                              __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --backups.available string                                    __Read-only__ Whether backups taken for this Linode are available for restoration. Backups undergoing maintenance are not available for restoration.
      --backups.enabled string                                      __Read-only__ If this Linode has the Backup service enabled. To enable backups, run [Enable backups](https://techdocs.akamai.com/linode-api/reference/post-enable-backups).
      --backups.last-successful null                                __Read-only__ The last successful backup time. Displayed as null if there was no previous backup.
      --backups.schedule string                                     
      --capabilities string                                         __Limited availability__, __Read-only__ A list of capabilities this Linode supports.
      --created string                                              __Read-only__ When this Linode was created.
      --data string                                                 Request body JSON, @file, or @- for stdin
      --disk-encryption null                                        __Read-only__ Indicates the local disk encryption setting for this Linode. If the Linode is part of an LKE cluster, the value is null. (default "enabled")
      --group string                                                __Deprecated__, __Filterable__ The group label for this Linode.
      --has-user-data user_data                                     __Read-only__ Whether this Linode was provisioned with user_data provided via the Metadata service. See the [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) description for more information on Metadata.
  -h, --help                                                        help for put-linode-instance
      --host-uuid string                                            __Read-only__ The Linode's host machine identifier.
      --hypervisor string                                           __Read-only__ The virtualization software powering this Linode.
      --id string                                                   __Filterable__, __Read-only__ This Linode's unique identifier, which you need to provide for all operations impacting this Linode.
      --image string                                                
      --interface-generation linode                                 __Filterable__, __Read-only__ Indicates if the Linode is configured to use Linode interfaces (linode) or legacy configuration profile interfaces (`legacy_config`).
      --ipv4 string                                                 __Filterable__, __Read-only__ This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to [Open a support ticket](https://techdocs.akamai.com/linode-api/reference/post-ticket) to get additional IPv4 addresses. IPv4 addresses may be reassigned between your Linodes, or shared with other Linodes. See the [networking](https://techdocs.akamai.com/linode-api/reference/post-firewalls) operations for details.
      --ipv6 null                                                   __Read-only__ This Linode's IPv6 SLAAC address. This address is specific to a Linode, and may not be shared. If the Linode has not been assigned an IPv6 address, the return value will be null.
      --label -                                                     __Filterable__ Provides a name for the Linode. If not provided, the API generates one for it. A Linode label has some constraints: - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (-), underscores (`_`), or periods (`.`). - It can't contain two consecutive hyphens (`--`), underscores (`__`) or periods (`..`).
      --linode-id string                                            ID of the Linode to look up.
      --lke-cluster-id string                                       __Read-only__ The ID of the Kubernetes cluster if the Linode is part of cluster.
      --maintenance-policy string                                   The maintenance policy configured by the user for this Linode. Review [maintenance policy](https://techdocs.akamai.com/cloud-computing/docs/host-maintenance-policy) documentation for more details.
      --placement-group.id string                                   The placement group's ID. You need to provide it for all operations that affect it.
      --placement-group.label -                                     __Filterable__ The unique name set for the placement group. A label has these constraints: - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (-), underscores (`_`), or periods (`.`).
      --placement-group.placement-group-policy strict               How requests to add future Linodes to your placement group are handled, and whether it remains compliant: - strict. Don't assign a new Linode if it breaks the grouped-together or spread-apart model set by the `placement_group_type`. Use this to ensure the placement group stays compliant (`is_compliant: true`). - `flexible`. Assign a new Linode, even if it breaks the grouped-together or spread-apart model set by the `placement_group_type`. This makes the group non-compliant (`is_compliant: false`). You need to wait for Akamai to move the offending Linode to make it compliant again, once the necessary capacity is available in the region. Offers flexibility to add future Linodes if compliance isn't an immediate concern. <<LB>> > blue-book > > In rare cases, non-compliance can occur with a `strict` placement group if Akamai needs to failover or migrate your Linodes for maintenance. Fixing non-compliance for a `strict` placement group is prioritized over a `flexible` group.
      --placement-group.placement-group-type placement_group_type   __Filterable__, __Read-only__ How Linodes are distributed in your placement group. A placement_group_type using anti-affinity (`anti_affinity:local`) places Linodes in separate hosts, but still in the same region. This best supports the spread-apart model for high availability. A `placement_group_type` using affinity places Linodes physically close together, possibly on the same host. This supports the grouped-together model for low-latency. > blue-book > > Currently, only `anti_affinity:local` is available for `placement_group_type`.
      --region region                                               __Filterable__, __Read-only__ The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where you've deployed the Linode. A region can only be changed by [migrating to a new data center](https://techdocs.akamai.com/linode-api/reference/post-migrate-linode-instance).
      --site-type core                                              __Read-only__ The Linode region's site type. A core region indicates a traditional cloud computing [region](https://www.linode.com/docs/products/platform/get-started/guides/choose-a-data-center/#product-availability) that offers all compute services. A `distributed` region indicates sites that are globally dispersed to be closer to end users and workloads. These regions offer limited services.
      --specs.disk image                                            __Read-only__ The amount of storage space, in MB, this Linode has access to. A typical Linode divides this space between a primary disk with an image deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an `image` through [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance). While this configuration is suitable for most use cases, if you need finer control over your Linode's disks, see the [List disks](https://techdocs.akamai.com/linode-api/reference/get-linode-disks) operations.
      --specs.gpus string                                           __Read-only__ The number of GPUs this Linode has access to.
      --specs.memory string                                         __Read-only__ The amount of RAM, in MB, this Linode has access to. Typically, a Linode boots with all of its available RAM, but this can be configured in a config profile. See the [List config profiles](https://techdocs.akamai.com/linode-api/reference/get-linode-configs) operation for more information.
      --specs.transfer string                                       __Read-only__ The amount of network transfer this Linode is allotted each month.
      --specs.vcpus string                                          __Read-only__ The number of VCPUs this Linode has access to.
      --status stopped                                              __Read-only__ A brief description of the Linode's current state. This value can change without direct action from you. For example, when a Linode goes into maintenance mode, its status is stopped. Status is generally self-explanatory, based on its name. - `busy` indicates you've assigned the Linode to a [placement group](https://techdocs.akamai.com/cloud-computing/docs/work-with-placement-groups), but the Linode is currently booting. Once the boot completes, the API completes the assignment and updates the Linode's `status` accordingly. - `provisioning` indicates that the API is applying operating system or Marketplace applications on the Linode. - `billing_suspension` indicates that payment is past due on the Linode, so we've suspended its use.
      --tags string                                                 __Filterable__ Tags to help you organize your content.
      --type string                                                 __Read-only__ The [type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) that this Linode was deployed with. To change a Linode's type, use [Resize a Linode](https://techdocs.akamai.com/linode-api/reference/post-resize-linode-instance).
      --updated string                                              __Read-only__ When this Linode was last updated.
      --watchdog-enabled string                                     The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and reboots it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie gives up if there have been more than 5 boot jobs issued within 15 minutes.
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
