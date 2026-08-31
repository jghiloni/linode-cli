## linodectl users put-user

Update a user

### Synopsis

Update information about a user on your account, including its restricted status. When setting a user to `restricted`, the API sets no grants for it. You need to set grants so that user can access things on the account. > construction RBAC Identity and Access > > If Identity and Access is enabled on your account, the *restricted* users have the `account_admin` role and *restricted* users have specific roles assigned to them. To update the role assignment, run the [Update a user's access level](https://techdocs.akamai.com/linode-api/reference/put-iam-users-role-permissions) operation instead. To learn more, see [Identity and Access for Akamai Cloud](https://techdocs.akamai.com/cloud-computing/docs/identity-and-access-cm). __Parent and child accounts__ In the context of the [parent and child accounts](https://techdocs.akamai.com/cloud-computing/docs/parent-and-child-accounts-for-akamai-partners) feature, on a child account, you can't edit the `username` or `email` values for delegate users. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_user` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write` **CLI** ```shell linode-cli users update example_user \ --username example_user \ --email example@linode.com \ --restricted true ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl users put-user [flags]
```

### Options

```
      --api-version v4                     __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --data string                        Request body JSON, @file, or @- for stdin
      --email string                       This user's email address. Akamai uses this address for account management communications.
  -h, --help                               help for put-user
      --last-login.login-datetime string   __Read-only__ The date and time of this user's most recent login attempt.
      --last-login.status string           __Read-only__ The result of this user's most recent login attempt.
      --password-created null              __Read-only__ When this user's current password was created. You initially create a password during the account sign-up process, and you can update it using the [Reset Password](https://login.linode.com/forgot/password) webpage. Returned as null if this user doesn't have a password set.
      --restricted false                   Defines user access to account. If false, the user is assigned the `account_admin` role. If `true`, the user's access is limited to selected roles.
      --ssh-keys authorized_users          __Read-only__ A list of the labels for SSH keys added by this user. Users can add keys with the [Add an SSH key](https://techdocs.akamai.com/linode-api/reference/post-add-ssh-key) operation. These keys are deployed when this user is included in the authorized_users field of the following requests: - [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) - [Rebuild a Linode](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) - [Create a disk](https://techdocs.akamai.com/linode-api/reference/post-add-linode-disk)
      --tfa-enabled string                 __Read-only__ Whether this user has Two Factor Authentication (TFA) enabled. Run the [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) operation to enable TFA.
      --username string                    The username to look up.
      --verified-phone-number null         __Read-only__ The [verified](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) phone number for this user profile. Returned as null if the user doesn't have a verified phone number.
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
