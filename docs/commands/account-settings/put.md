## linodectl account-settings put

Update account settings

### Synopsis

Updates your account settings. For a Longview subscription plan, see [Update a Longview plan](https://techdocs.akamai.com/linode-api/reference/put-longview-plan). **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_account_settings` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write` **CLI** ```shell linode-cli account settings-update \ --network_helper false ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl account-settings put [flags]
```

### Options

```
      --api-version v4                                  __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --backups-enabled true                            Account-wide backups default. If true, all Linodes created will automatically be enrolled in the Backups service. If `false`, Linodes will not be enrolled by default, but may still be enrolled on creation or later.
      --data string                                     Request body JSON, @file, or @- for stdin
  -h, --help                                            help for put
      --interfaces-for-new-linodes legacy_config_only   Defines if new Linodes can use legacy configuration interfaces: - legacy_config_only. All new Linodes need to use legacy configuration interfaces. Previously created Linodes with Linode Interfaces can still exist. Linodes using legacy configuration interfaces can't be upgraded to use Linode Interfaces. - `legacy_config_default_but_linode_allowed`. New Linodes can use legacy configuration interfaces or Linode Interfaces, depending on the `interface_generation` setting specified when creating the Linode. By default, new Linodes use legacy configuration interfaces unless otherwise specified. Linodes that use legacy configuration interfaces can upgrade to Linode Interfaces. This is the default setting for existing accounts. - `linode_default_but_legacy_config_allowed`. New Linodes can use legacy configuration interfaces or Linode Interfaces, depending on the `interface_generation` setting specified when creating the Linode. By default, new Linodes use Linode Interfaces unless otherwise specified. Linodes that use legacy configuration interfaces can upgrade to Linode interfaces. This is the default setting for new accounts. - `linode_only`. All new Linodes need to use Linode Interfaces. Previously created Linodes with legacy configuration profile interfaces can still exist if they were created under a previous setting. Linodes using legacy configuration interfaces can be upgraded to Linode Interfaces.
      --longview-subscription null                      __Read-only__ The Longview Pro tier you are currently subscribed to. The value must be a [Longview subscription](https://techdocs.akamai.com/linode-api/reference/get-longview-subscriptions) ID or null for Longview Free.
      --maintenance-policy string                       Defines the default maintenance policy for new Linodes created on this account. Review [maintenance policy](https://techdocs.akamai.com/cloud-computing/docs/host-maintenance-policy) documentation for more details.
      --managed string                                  __Read-only__ Our 24/7 incident response service. This robust, multi-homed monitoring system distributes monitoring checks to ensure that your servers remain online and available at all times. Linode Managed can monitor any service or software stack reachable over TCP or HTTP. Once you add a service to Linode Managed, we'll monitor it for connectivity, response, and total request time.
      --network-helper string                           Enables network helper across all users by default for new Linodes and Linode Configs.
      --object-storage string                           __Read-only__ A string describing the status of this account's Object Storage service enrollment. (default "disabled")
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

* [linodectl account-settings](./index.md)	 - account-settings operations

###### Auto generated by spf13/cobra on 31-Aug-2026
