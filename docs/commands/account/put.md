## linodectl account put

Update your account

### Synopsis

Updates contact and billing information related to your account. If you exclude any properties from the request, the operation leaves them unchanged. > blue-book > > When updating an account's `country` to `US`, you'll get an error if the account's `zip` is not a valid US zip code. __Parent and child accounts__ In the context of the [parent and child accounts](https://techdocs.akamai.com/cloud-computing/docs/parent-and-child-accounts-for-akamai-partners) feature: - You can't change the company for a parent account. - Child account users can't run this operation. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_account` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write` **CLI** ```shell linode-cli account update \ --address_1 "123 Main St." \ --address_2 "Suite 101" \ --city Philadelphia \ --company My Company \ LLC \ --country US \ --email jsmith@mycompany.com \ --first_name John \ --last_name Smith \ --phone 555-555-1212 \ --state PA \ --tax_id ATU99999999 \ --zip 19102 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl account put [flags]
```

### Options

```
      --active-promotions string            
      --active-since string                 __Read-only__ The date and time the account was activated.
      --address-1 string                    The first line of this account's billing address.
      --address-2 string                    The second line of this account's billing address.
      --api-version v4                      __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --balance string                      __Read-only__ This account's balance, in US dollars.
      --balance-uninvoiced string           __Read-only__ This account's current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate.
      --billing-source akamai               __Read-only__ The source of service charges for this account. Accounts that are associated with Akamai-specific customers return a value of akamai. All other accounts return a value of `linode`.
      --capabilities string                 __Read-only__ The Akamai Cloud Computing services your account supports.
      --city address                        The city for this account's address.
      --company <                           The company name assigned to this account. This value can't include the characters, < `>` `(` `)` `"` `=`.
      --country address                     The two-letter ISO 3166 country code for this account's address.
      --credit-card.expiry credit_card      The expiration month and year of the credit_card.
      --credit-card.last-four credit_card   The last four digits of the credit_card assigned to this account.
      --data string                         Request body JSON, @file, or @- for stdin
      --email string                        The email address of the person assigned to this account.
      --euuid string                        __Read-only__ An external unique identifier for this account.
      --first-name <                        The first name of the person assigned to this account. This value can't include the characters, < `>` `(` `)` `"` `=`.
  -h, --help                                help for put
      --last-name <                         The last name of the person assigned to this account. This value can't include the characters, < `>` `(` `)` `"` `=`.
      --phone string                        The phone number assigned to this account.
      --state address                       The state or province for the address set for your account, if applicable. - If the `address` is in the United States (US) or Canada (CA), this is the two-letter ISO 3166 code for the state or province. - If it's a US military `address`, this is the abbreviation for that territory. This includes `AA` for Armed Forces Americas (excluding Canada), `AE` for Armed Forces Africa, Europe, Middle East, and Canada, or `AP` for Armed Forces Pacific. - If outside the US or CA, this is the province associated with the account's `address`.
      --tax-id country                      The tax identification number (TIN) assigned to this account, used for tax calculations. A TIN is set by the national authorities in your country, based on your `address_1`, and it may be named differently between countries. Set to an empty string (`""`) if a TIN doesn't apply or for countries that don't collect tax. > blue-book > > This value is externally validated. If the validation is successful, a `tax_id_valid` [event](https://techdocs.akamai.com/linode-api/reference/get-events) is triggered. If unsuccessful, a `tax_id_invalid` event is triggered and an error response is issued for an operation that included it.
      --zip address                         The zip code for this account's address. - It can only contain ASCII letters, numbers, and dashes (`-`). - It can't contain more than nine letter or number characters.
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

* [linodectl account](./index.md)	 - account operations

###### Auto generated by spf13/cobra on 31-Aug-2026
