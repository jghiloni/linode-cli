## linodectl linode-interfaces post-upgrade

Upgrade to Linode interfaces

### Synopsis

Automatically upgrades all legacy config interfaces of a single configuration profile to Linode interfaces. A Linode interface is directly associated with the Linode, rather than being tied to a configuration profile. Firewalls are applied to the Linode interface, not directly to the Linode itself. > red-exclamation-mark This upgrade is irreversible > > Once you upgrade a Linode to use Linode interfaces, you can't use legacy config interfaces. This means you can no longer use the Linode with any Linode products that require private IPs, such as NodeBalancer. You can use `dry_run` to preview the upgrade. Before upgrading interfaces, you can check the new Linode interface configuration by performing a dry run, where you set `dry_run` to `true` or omit it. A `dry_run` runs the upgrade logic and returns a JSON representation of what the interface configuration will be after the upgrade without committing any changes. When you run this operation with `dry_run` set to `false`, the following occurs: - It creates matching interfaces on the Linode based on the interfaces present on the `config_id`. - All firewalls are removed from the Linode. Any firewalls that were originally attached to the Linode are now applied to the public and VPC interfaces. Firewalls are not applied to VLAN interfaces. If the Linode has no firewalls attached, then default firewalls are not used. - If no legacy config interfaces are defined (`legacy_config`) in the `config_id`, a public interface is created using the public IPv4 assigned to the Linode. The same is the case if the Linode has no config defined. - For public interfaces, the Linode's current MAC address and SLAAC address are conserved. The MAC address won't change. - It deletes all legacy config interfaces from all configs. - It returns the list of interfaces for the Linode. Requirements: - The `config_id` for the legacy config interfaces can't use a public interface private IPv4 address. - The Linode needs a MAC address in the database if it's IPv6 enabled. If it doesn't, an error message tells you what to do. - The Linode must be in a region that supports Linode interfaces. Run [Get a region](https://techdocs.akamai.com/linode-api/reference/get-region). - Your account must allow creation of Linodes with Linode interfaces, run [Get account settings](https://techdocs.akamai.com/linode-api/reference/get-account-settings). - If the Linode has a user with a non-standard username, it can't be upgraded. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes interfaces-upgrade $linodeId ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl linode-interfaces post-upgrade [flags]
```

### Options

```
      --api-version v4        __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --config-id config_id   The Linode's config_id. Only one `config_id` can be specified: - If there are no legacy configuration interfaces or configurations defined in the config_id, a public interface is created using the Linode's automatically assigned public IPv4 address. - If a config_id is not provided and the Linode has only one configuration, the upgrade automatically uses that config_id. - If the Linode has multiple configurations and a config_id is not specified, an error is returned. (default "null")
      --data string           Request body JSON, @file, or @- for stdin
      --dry-run dry_run       Before you upgrade interfaces, you can preview the new Linode interface by performing a dry_run: - Either omit `dry_run` or set it to `true` to simulate the upgrade process. The response data shows what the Linode interface will look like after the upgrade, but without committing any changes. Since the interface doesn't exist yet, the interface `id` value is 0. - If `dry_run` is set to `false`, the Linode undergoes the actual upgrade, but note you need to first shut down the Linode. (default "true")
  -h, --help                  help for post-upgrade
      --linode-id id          The id of the Linode.
```

### Options inherited from parent commands

```
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

* [linodectl linode-interfaces](./index.md)	 - linode-interfaces operations

###### Auto generated by spf13/cobra on 31-Aug-2026
