## linodectl profile put

Update a profile

### Synopsis

Update information in your profile. __Parent and child accounts__ In the context of the [parent and child accounts](https://techdocs.akamai.com/cloud-computing/docs/parent-and-child-accounts-for-akamai-partners) feature, delegate users on a child account can't change their email address. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_profile` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write` **CLI** ```shell linode-cli profile update \ --email example-user@gmail.com \ --timezone US/Eastern \ --email_notifications true \ --list_auth_method keys_only \ --two_factor_auth true \ --restricted false ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl profile put [flags]
```

### Options

```
      --api-version v4                     __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --authentication-type password       __Read-only__ This account's Cloud Manager authentication type. You choose an authentication type in Cloud Manager and Akamai authorizes it when you log into your account. Authentication types include your user's password (in conjunction with your username), or the name of your identity provider, such as GitHub. Here are some examples: - If a user has never used third-party authentication, the authentication type will be password. - If a user is using third-party authentication, the name of their identity provider is used for the authentication type, for example, `github`. - If a user has used third-party authentication and has since revoked it, the authentication type is `password`.
      --authorized-keys lish_auth_method   Your user can use these SSH Keys to access Lish. This value is ignored if lish_auth_method is `disabled`.
      --data string                        Request body JSON, @file, or @- for stdin
      --email string                       Your email address. We use this address for Akamai Cloud Computing-related communication.
      --email-notifications true           When set to true, you will receive email notifications about account activity. When `false`, you may still receive business-critical communications through email.
  -h, --help                               help for put
      --ip-whitelist-enabled true          __Deprecated__ When set to true, your user logins are only allowed from whitelisted IPs. This setting is deprecated, and can't be enabled. If you disable this setting, you won't be able to re-enable it.
      --lish-auth-method keys_only         The authentication methods that you can use when connecting to the [Linode Shell (Lish)](https://www.linode.com/docs/guides/lish/). - keys_only is the most secure if you intend to use Lish. - `disabled` is recommended if you don't want to use Lish. - If this account's Cloud Manager authentication type is set to a third-party authentication method, you can't use `password_keys` as your Lish authentication method. Run the [Get a profile](https://techdocs.akamai.com/linode-api/reference/get-profile) operation to view your account's Cloud Manager `authentication_type` field.
      --referrals.code string              __Read-only__ Your referral code. If others use this when signing up for Linode, you receive an account credit.
      --referrals.completed string         __Read-only__ The number of completed sign-ups that used your referral code.
      --referrals.credit string            __Read-only__ Your referral program account credit in US dollars.
      --referrals.pending string           __Read-only__ The number of pending sign-ups that used your referral code. Akamai gives you credit for these sign-ups once they've completed.
      --referrals.total string             __Read-only__ The number of users who have signed up with your referral code.
      --referrals.url string               __Read-only__ The referral URL that Akamai uses to direct others to sign up for Akamai Cloud Computing with your referral code.
      --restricted true                    When set to true, there are restrictions on what your user can access on your account. Run [List grants](https://techdocs.akamai.com/linode-api/reference/get-profile-grants) to get details on what entities and actions you can access and perform.
      --timezone string                    The time zone you want to display for your Linode assets. This API doesn't directly use this time zone. It's provided for the benefit of clients such as the Akamai Cloud Manager and other clients built on the API. All times returned by the API are in UTC.
      --two-factor-auth true               When set to true, a login from an untrusted computer requires two-factor authentication. You also need to run [Create a two factor secret](https://techdocs.akamai.com/linode-api/reference/post-tfa-enable) to enable two-factor authentication.
      --uid string                         __Read-only__ Your unique ID in our system. This value will never change, and can safely be used to identify your user.
      --username string                    __Read-only__ Your username, used for logging in to our system.
      --verified-phone-number null         __Read-only__ The phone number verified for this profile with the [Verify a phone number](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number-verify) operation. Displayed as null if the profile doesn't have a verified phone number.
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

* [linodectl profile](./index.md)	 - profile operations

###### Auto generated by spf13/cobra on 31-Aug-2026
