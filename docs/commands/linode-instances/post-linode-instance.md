## linodectl linode-instances post-linode-instance

Create a Linode

### Synopsis

Creates a new Linode instance on your account. This incurs a charge on your account. - There are several types of Linode you can choose from. Each provides a different size, processing potential, and cost. Run the [List Linode types](https://techdocs.akamai.com/linode-api/reference/get-linode-types) operation for information on each available type. - You can create a Linode in any available regions. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to see what's available. - You need to provide at least one authentication mechanism for disk access: a `root_pass`, `authorized_keys`, or `authorized_users`, with the latter two offering SSH protection. - To fight spam, Akamai restricts outbound connections on ports 25, 465, and 587 on all Linodes. For more information, see our guide on [Running a Mail Server](https://www.linode.com/docs/guides/running-a-mail-server/). - To add or modify `tags` on a Linode, your account [user](https://techdocs.akamai.com/linode-api/reference/get-users) needs `read-write` access for **Linode instances** and **Tags** operations. Talk to your local account administrator about access management. There are several ways to create a Linode, using the objects available in this operation: **Use an Image** This includes using a Linode public image distribution or a private image that you created, based on another Linode. > thumbs-up There are tutorials > > We offer example API workflows you can follow for both [public image distributions](https://techdocs.akamai.com/linode-api/reference/create-a-linode-using-a-public-image) as well as [private images](https://techdocs.akamai.com/linode-api/reference/create-a-linode-using-a-private-image). **Use cloud-init** Our cloud-init solution with [Metadata](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/) automates system configuration and software installation by providing a base-64 encoded [cloud-config](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata-cloud-config/) file. - You need a compatible image. To check for compatibility, run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation and check for `cloud-init` under `capabilities`. - This also requires a compatible region. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation and look for `Metadata` under `capabilities` for a specific `region`. **Use a StackScript** You can use a StackScript to automate deployment of new systems, through a customized script. > thumbs-up There's a tutorial > > We offer an example API workflow you can follow to create a new Linode [using a StackScript](https://techdocs.akamai.com/linode-api/reference/create-a-linode-using-a-stackscript). **Use a Linode backup** You can create a backup of an existing Linode and restore it to a new one. > thumbs-up There's a tutorial > > We offer an example API workflow you can follow to create a new Linode [using a backup](https://techdocs.akamai.com/linode-api/reference/create-a-linode-using-a-backup). **Attached to a private VLAN** - See the `interfaces` object in the Body Params section for details. - For more information, see our guide on [Getting Started with VLANs](https://www.linode.com/docs/products/networking/vlans/get-started/). **Create an empty Linode** - Use this operation to create a new one. - The Linode will remain `offline` and you need to manually start it. Run the [Boot a Linode](https://techdocs.akamai.com/linode-api/reference/post-boot-linode-instance) operation. - Manually create the Linode's [disks](https://techdocs.akamai.com/linode-api/reference/post-add-linode-disk) and [configuration profiles](https://techdocs.akamai.com/linode-api/reference/post-add-linode-config) (legacy) or [configuration profile interfaces](https://techdocs.akamai.com/linode-api/reference/post-linode-config-interface). > blue-book > > You should only use this method for advanced use cases. **Linodes and interfaces** Depending on your [account settings](https://techdocs.akamai.com/linode-api/reference/get-account-settings), you can choose between legacy configuration interfaces or Linode interfaces when creating a Linode. Only one type of interface is allowed per Linode. The `interface_generation` field lets you select one interface type for new Linodes when both legacy and Linode interfaces options are available on your account. If a Linode is configured with a Linode interface, legacy configuration interfaces can no longer be used on that Linode. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `create_linode` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `linodes:read_write` **CLI** ```shell linode-cli linodes create \ --label linode123 \ --root_pass aComplex@Password \ --booted true \ --stackscript_id 10079 \ --stackscript_data '{"gh_username": "linode"}' \ --region us-east \ --disk_encryption enabled\ --placement_group.id 528 \ --type g6-standard-2 \ --authorized_keys "ssh-rsa AAAA_valid_public_ssh_key_123456785== user@their-computer" \ --authorized_users "myUser" \ --authorized_users "secondaryUser" \ --metadata.user_data "I2Nsb3VkLWNvbmZpZw==" \ --firewall_id 9000 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl linode-instances post-linode-instance [flags]
```

### Options

```
      --api-version v4   __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --data string      Request body JSON, @file, or @- for stdin
  -h, --help             help for post-linode-instance
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
