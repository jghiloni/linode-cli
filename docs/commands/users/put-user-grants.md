## linodectl users put-user-grants

Update a user's grants

### Synopsis

__Deprecated__ Update the grants for a [restricted](https://techdocs.akamai.com/linode-api/reference/post-user) user. This can be used to give a user access to new entities or actions, or take access away. Omit a grant object from the request to keep its current setting. > construction RBAC Identity and Access > > Grants are now replaced with RBAC Identity and Access. Use the [Update a user's access level](https://techdocs.akamai.com/linode-api/reference/put-iam-users-role-permissions) operation instead. To learn more, see [Identity and Access for Akamai Cloud](https://techdocs.akamai.com/cloud-computing/docs/identity-and-access-cm). **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_user_grants` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write`

```
linodectl users put-user-grants [flags]
```

### Options

```
      --api-version v4                        __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --data string                           Request body JSON, @file, or @- for stdin
      --database string                       The grants this user has for individual Managed Databases on this account.
      --domain string                         The grants this user has for individual domains on this account.
      --firewall string                       The grants this user has for individual firewalls on this account.
      --global.account-access restricted      The level of access this user has to account-level actions, like billing information and user management. > blue-book > > A restricted user can't be used to manage users, even if this is set to `read-write`. Only unrestricted users can manage other users on an account. __Parent and child accounts__ In the context of the [parent and child accounts](https://techdocs.akamai.com/cloud-computing/docs/parent-and-child-accounts-for-akamai-partners) feature, use the [Update a user's access level](https://techdocs.akamai.com/linode-api/reference/put-iam-users-role-permissions) to manage delegate user's access.
      --global.add-databases string           Whether this user can add Managed Databases on the account.
      --global.add-domains string             Whether this user can add domains on the account.
      --global.add-firewalls string           Whether this user can add Firewalls on the account.
      --global.add-images string              Whether this user can create images from disks on your Linodes, on the account.
      --global.add-linodes string             Whether this user can create Linodes.
      --global.add-longview string            Whether this user can create Longview clients and view the current plan.
      --global.add-nodebalancers string       Whether this user can add NodeBalancers on the account.
      --global.add-stackscripts string        Whether this user can add StackScripts on the account.
      --global.add-volumes string             Whether this user can add volumes on the account.
      --global.add-vpcs string                Whether this user can add Virtual Private Clouds (VPCs) on the account.
      --global.cancel-account string          Whether this user can cancel the entire account.
      --global.child-account-access null      __Deprecated__ In order to provide a parent account user access to a child account, [add them to an account delegation](https://techdocs.akamai.com/linode-api/reference/put-iam-delegation-child-account-users). This is null for all non-parent accounts.
      --global.longview-subscription string   Whether this user can manage your account's Longview subscription.
  -h, --help                                  help for put-user-grants
      --image string                          The grants this user has for individual images on this account.
      --linode string                         The grants this user has for individual Linodes on this account.
      --longview string                       The grants this user has for individual Longview Clients on this account.
      --nodebalancer string                   The grants this user has for individual NodeBalancers on this account.
      --stackscript string                    The grants this User has for individual StackScripts on this account.
      --username string                       The username to look up.
      --volume string                         The grants this user has individual Block Storage Volumes on this account.
      --vpc string                            The grants this user has individual Virtual Private Clouds (VPCs) on this account.
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

* [linodectl users](./index.md)	 - users operations

###### Auto generated by spf13/cobra on 31-Aug-2026
