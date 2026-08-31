## linodectl linode-instances post-clone-linode-instance

Clone a Linode

### Synopsis

Clone your Linode's existing disks, configuration profiles, and interfaces to another Linode on your account. Consider these points before cloning a Linode: - Cloning to a new Linode incurs a charge on your account. - If cloning to an existing Linode, any actions currently running or queued need to finish before you can clone to it. - You can run up to five concurrent clone operations, from a single Linode. If you exceed this limit, you'll receive an HTTP 400 error. - For Linodes using Linode interfaces, the region where you want the clone to live also needs to support Linode interfaces. You can run the [GET a region](https://techdocs.akamai.com/linode-api/reference/get-region) operation to see what's supported. Your user's [account settings](https://techdocs.akamai.com/linode-api/reference/get-account-settings) also need ot be set to allow creation of Linodes with Linode interfaces. - Any [tags](https://techdocs.akamai.com/linode-api/reference/get-tags) you have set on your source Linode will be cloned to the target Linode. - If your source Linode is protected with a resource lock, the lock is _not included_ on the clone. You'll need to [add a new resource lock](https://techdocs.akamai.com/linode-api/reference/post-resource-lock) if you want to protect the clone. - If the target Linode uses Metadata, with `has_user_data` set to `true`, you need to create the clone with `metadata.user_data` in the request. - If the target Linode has a `vpc` interface on its active legacy configuration profile, and it includes a `1:1 NAT`, the resulting clone is configured with an `any` 1:1 NAT. See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations. - Only next generation network (NGN) data centers (regions) support VLANs. If a VLAN is attached to your Linode and you try to clone it to a non-NGN region, the clone won't start. If a Linode can't be cloned because of an incompatibility, you're prompted to select a different region or contact support. See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide for more specifications and limitations. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `clone_linode` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes clone 123 \ --linode_id 124 \ --region us-east \ --type g6-standard-2 \ --label cloned-linode \ --backups_enabled true \ --placement_group.id 528 \ --disks 25674 \ --configs 23456 \ --private_ip true \ --metadata.user_data I2Nsb3VkLWNvbmZpZw== ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl linode-instances post-clone-linode-instance [flags]
```

### Options

```
      --api-version v4                 __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --backups-enabled true           If this field is set to true, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. Pricing is included in the response from [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types). - Can only be included when cloning to a new Linode.
      --configs configs                An array of configuration profile IDs. - If the configs parameter __is not provided__, then __all configuration profiles and their associated disks will be cloned__ from the source Linode. Any disks specified by the `disks` parameter will also be cloned. - __If an empty array is provided__ for the `configs` parameter, then __no configuration profiles (nor their associated disks) will be cloned__ from the source Linode. Any disks specified by the `disks` parameter will still be cloned. - __If a non-empty array is provided__ for the `configs` parameter, then __the configuration profiles specified in the array (and their associated disks) will be cloned__ from the source Linode. Any disks specified by the `disks` parameter will also be cloned.
      --data string                    Request body JSON, @file, or @- for stdin
      --disks disks                    An array of disk IDs. - If the disks parameter __is not provided__, then __no extra disks will be cloned__ from the source Linode. All disks associated with the configuration profiles specified by the `configs` parameter will still be cloned. - __If an empty array is provided__ for the `disks` parameter, then __no extra disks will be cloned__ from the source Linode. All disks associated with the configuration profiles specified by the `configs` parameter will still be cloned. - __If a non-empty array is provided__ for the `disks` parameter, then __the disks specified in the array will be cloned__ from the source Linode, in addition to any disks associated with the configuration profiles specified by the `configs` parameter.
      --group string                   __Deprecated__ A label used to group Linodes for display. Linodes are not required to have a group.
  -h, --help                           help for post-clone-linode-instance
      --label linode                   The label to assign this Linode when cloning to a new Linode. - Can only be provided when cloning to a new Linode. - Defaults to linode.
      --linode-id string               ID of the Linode to clone.
      --maintenance-policy string      Defines the maintenance policy for the new Linode. If you don't provide it, the new Linode inherits the maintenance policy from the original Linode. Review [maintenance policy](https://techdocs.akamai.com/cloud-computing/docs/host-maintenance-policy) documentation for more details.
      --metadata.user-data user_data   Base64-encoded [cloud-config](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata-cloud-config/) data. Consider these points when including user_data: - You can't update an existing Linode to modify this information. You need to either [clone](https://techdocs.akamai.com/linode-api/reference/post-clone-linode-instance) or [rebuild](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) an existing Linode and apply new `user_data`. - You can't include this field when cloning to an existing Linode. - Unencoded data can't exceed 65535 bytes when encoded.
      --placement-group.id string      The placement group's ID. You need to provide it for all operations that affect it.
      --private-ip true                If true, the created Linode will have private networking enabled and assigned a private IPv4 address. - Can only be provided when cloning to a new Linode.
      --region string                  This is the Region where the Linode will be deployed. To view all available Regions you can deploy to, run [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions). - Region can only be provided and is required when cloning to a new Linode.
      --type specs                     A Linode's Type determines what resources are available to it, including disk space, memory, and virtual cpus. The amounts available to a specific Linode are returned as specs on the Linode object. To view all available Linode Types you can deploy with, run [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types). - Type can only be provided and is required when cloning to a new Linode.
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
