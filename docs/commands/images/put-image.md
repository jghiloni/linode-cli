## linodectl images put-image

Update an image

### Synopsis

Updates a private image. > blue-book > > You can't update the `regions` with this operation. Use the [Replicate an image](https://techdocs.akamai.com/linode-api/reference/post-replicate-image) operation to modify the existing regions for your image. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `update_image` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `images:read_write` **CLI** ```shell linode-cli images update private/12345 \ --label "My gold image" \ --description "The detailed description \ of my image." ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl images put-image [flags]
```

### Options

```
      --api-version v4            __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --capabilities cloud-init   __Read-only__ A list of the possible capabilities of this image. - cloud-init. The image supports the cloud-init multi-distribution method with our [Metadata service](https://techdocs.akamai.com/cloud-computing/docs/overview-of-the-metadata-service). This only applies to public images. - `distributed-sites`. Whether the image can be used in distributed compute regions. Compared to a core compute region, distributed compute regions offer limited functionality, but they're globally distributed. Your image can be geographically closer to you, potentially letting you deploy it quicker. See [Regions and images](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-images) for complete details.
      --created string            __Read-only__ When this image was created.
      --created-by linode         __Read-only__ The name of the user who created this image, or linode for public images.
      --data string               Request body JSON, @file, or @- for stdin
      --deprecated true           __Filterable__, __Read-only__ A true value indicates a deprecated image. Only public images can be deprecated.
      --description string        A detailed description of this image.
      --eol null                  __Read-only__ The time of the public image's planned removal from service. This is null for private images.
      --expiry null               __Read-only__ Only images created automatically from a deleted compute instance (type=automatic) expire. This is null for private images.
  -h, --help                      help for put-image
      --id string                 __Read-only__ The unique identifier for each image.
      --image-id string           The unique identifier assigned to the image after creation.
      --is-public true            __Filterable__, __Read-only__ A true value if the image is a public distribution image. A `false` value indicates private, account-specific images.
      --is-shared true            __Filterable__, __Read-only__ A true value for shared private images. `none` for images shared within a group.
      --label string              __Filterable__ A short description of the image.
      --regions regions           __Read-only__ Details on the regions where this image is stored. See [Regions and images](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-images) for full details on support for regions.
      --size string               __Filterable__, __Read-only__ The minimum size in MB this image needs to deploy.
      --status available          __Filterable__, __Read-only__ The current status of the image. Possible values are available, `creating`, and `pending_upload`. > blue-book > > The `+order_by` and `+order` operators are not available when [filtering](https://techdocs.akamai.com/linode-api/reference/filtering-and-sorting) on this key.
      --tags string               __Filterable__ Tags used for organizational purposes. A tag can be from 3 to 100 characters long, and an image can have a maximum of 500 total tags.
      --total-size regions        __Read-only__ The total size in bytes of all instances of this image, in all regions. > blue-book > > This object is empty for existing images. It's intended for use with future functionality.
      --type manual               __Filterable__, __Read-only__ How the image was created. Create a manual image at any time. An `automatic` image is created automatically from a deleted compute instance.
      --updated string            __Read-only__ When this image was last updated.
      --vendor null               __Filterable__, __Read-only__ The upstream distribution vendor. This is null for private images.
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

* [linodectl images](./index.md)	 - images operations

###### Auto generated by spf13/cobra on 31-Aug-2026
