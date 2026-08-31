## linodectl domains put-domain

Update a domain

### Synopsis

Update information about a domain in DNS Manager. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Roles: `domain_admin` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `domains:read_write` **CLI** ```shell linode-cli domains update 1234 \ --retry_sec 7200 \ --ttl_sec 300 ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl domains put-domain [flags]
```

### Options

```
      --api-version v4       __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --axfr-ips string      The list of IPs that may perform a zone transfer for this domain. The total combined length of all data within this array cannot exceed 1000 characters. > blue-book > > This is potentially dangerous, and should be set to an empty list unless you intend to use it.
      --data string          Request body JSON, @file, or @- for stdin
      --description string   A description for this domain. This is for display purposes only.
      --domain string        __Filterable__ The domain this domain represents. domain labels cannot be longer than 63 characters and must conform to [RFC1035](https://tools.ietf.org/html/rfc1035). domains must be unique on Linode's platform, including across different Linode accounts; there cannot be two domains representing the same domain.
      --domain-id string     The ID of the Domain to access.
      --expire-sec string    The amount of time in seconds that may pass before this domain is no longer authoritative. - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 1209600. (default "0")
      --group string         __Deprecated__, __Filterable__ The group this domain belongs to. This is for display purposes only.
  -h, --help                 help for put-domain
      --id string            __Read-only__ This domain's unique ID.
      --master-ips type      The IP addresses representing the master DNS for this domain. At least one value is required for type slave domains. The total combined length of all data within this array cannot exceed 1000 characters.
      --refresh-sec string   The amount of time in seconds before this domain should be refreshed. - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 14400. (default "0")
      --retry-sec string     The interval, in seconds, at which a failed refresh should be retried. - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 14400. (default "0")
      --soa-email type       Start of Authority email address. This is required for type master domains.
      --status string        Used to control whether this domain is currently being rendered. (default "active")
      --tags string          __Filterable__ An array of tags applied to this object. Tags are for organizational purposes only.
      --ttl-sec string       "Time to Live" - the amount of time in seconds that this domain's records may be cached by resolvers or other domain servers. - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 86400. (default "0")
      --type master          Whether this domain represents the authoritative source of information for the domain it describes (master), or whether it is a read-only copy of a master (`slave`).
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

* [linodectl domains](./index.md)	 - domains operations

###### Auto generated by spf13/cobra on 31-Aug-2026
