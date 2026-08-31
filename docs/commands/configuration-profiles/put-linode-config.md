## linodectl configuration-profiles put-linode-config

Update a configuration profile

### Synopsis

Updates a configuration profile. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_linode_config_profile` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes config-update 123 23456 \ --kernel "linode/latest-64bit" \ --comments "This is my main Config" \ --memory_limit 2048 \ --run_level default \ --virt_mode paravirt \ --helpers.updatedb_disabled true \ --helpers.distro true \ --helpers.modules_dep true \ --helpers.network true \ --helpers.devtmpfs_automount false \ --label "My Config" \ --devices.sda.disk_id 123456 \ --devices.sdb.disk_id 123457 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl configuration-profiles put-linode-config [flags]
```

### Options

```
      --api-version v4                    __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --comments string                   Optional field for arbitrary user comments on this configuration.
      --config-id id                      The id of the Configuration Profile.
      --data string                       Request body JSON, @file, or @- for stdin
      --devices.sda disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdb disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdc disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdd disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sde disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdf disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdg disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
      --devices.sdh disk_id               Device can be either a Disk or Volume identified by disk_id or `volume_id`. Only one type per slot allowed. Can be `null`. Devices mapped from _sde_ through _sdh_ are unavailable in `fullvirt` virt_mode.
  -h, --help                              help for put-linode-config
      --helpers.devtmpfs-automount /dev   Populates the /dev directory early during boot without `udev`. Defaults to `false`. (default "false")
      --helpers.distro inittab            Helps maintain correct inittab or `upstart` console device.
      --helpers.modules-dep string        Creates a modules dependency file for the kernel you run.
      --helpers.network true              Set to true to automatically configure static networking. The `network` option applies only to legacy configuration profile interfaces and does not apply to [Linode interfaces](https://techdocs.akamai.com/linode-api/reference/post-linode-interface).
      --helpers.updatedb-disabled true    Set to true to disable the `updatedb` cron job to avoid disk thrashing.
      --id string                         __Read-only__ The ID of this Config.
      --interfaces interfaces             interfaces is applicable only to legacy configuration profiles and does not apply to [Linode interfaces](https://techdocs.akamai.com/linode-api/reference/post-linode-interface). From one to three network interfaces to add to this Linode's configuration profile. The position in the array determines which of the Linode's network interfaces is configured: - First [0]: `eth0` - Second [1]: `eth1` - Third [2]: `eth2` When updating a Linode's legacy interfaces, _each interface must be redefined_. An empty `interfaces` array results in a default `public` type interface configuration only. If no public Interface is configured, public IP addresses are still assigned to the Linode but will not be usable without manual configuration. > blue-book > > Changes to Linode Interface configurations can be enabled by rebooting the Linode. `vpc` details See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations. `vlan` details - Only Next Generation Network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of data center regions. If a VLAN is attached to your Linode and you attempt to migrate or clone it to a non-NGN data center, the migration or cloning will not initiate. If a Linode cannot be migrated or cloned because of an incompatibility, you will be prompted to select a different data center or contact support. - See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations.
      --kernel linode/latest-64bit        The ID of the kernel used to boot a Linode. Run the [List kernels](https://techdocs.akamai.com/linode-api/reference/get-kernels) operation to see all available kernels. Here are some commonly used kernels: - linode/latest-64bit. This is the default, our latest kernel at the time of an instance boot or reboot. - `linode/grub2`. The upstream distribution-supplied kernel that's installed on the primary disk, or a custom kernel if installed. - `linode/direct-disk`. The master boot record (MBR) of the primary disk or root device. Use this in place of a Linux kernel. (default "linode/latest-64bit")
      --label string                      __Filterable__ The name of the configuration for display in Akamai Cloud Manager.
      --linode-id id                      The id of the Linode.
      --memory-limit string               Defaults to the total RAM of the Linode.
      --root-device /dev/sda              The root device to boot. > blue-book - If you leave this empty or set an invalid value, the root device defaults to /dev/sda. - If you specify a device at the root device location and it's not mounted, the Linode won't boot until a device is mounted.
      --run-level default                 Defines the state of your Linode after booting. Defaults to default.
      --virt-mode paravirt                Controls the virtualization mode. Defaults to paravirt. - `paravirt` is suitable for most cases. Linodes running in `paravirt` mode share some qualities with the host, ultimately making it run faster since there is less transition between it and the host. - `fullvirt` affords more customization, but is slower because 100% of the VM is virtualized.
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

* [linodectl configuration-profiles](./index.md)	 - configuration-profiles operations

###### Auto generated by spf13/cobra on 31-Aug-2026
