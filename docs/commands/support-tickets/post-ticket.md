## linodectl support-tickets post-ticket

Open a support ticket

### Synopsis

Open a support ticket. A ticket can only target a single, specific entity. For example, for an issue with a specific Linode, open a ticket and target it using its `linode_id`. Leave all other entities out of the request or set them to `null`. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `account:read_write` **CLI** ```shell linode-cli tickets create \ --description "I'm having trouble setting the root password on my Linode. I tried following the instructions but something is not working and I'm not sure what I'm doing wrong. Can you please help me figure out how I can reset it?" \ --linode_id 123 \ --summary "Having trouble resetting root password on my Linode" ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl support-tickets post-ticket [flags]
```

### Options

```
      --api-version v4         __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --bucket label           The name of an Object Storage bucket entity for this ticket. Run the [List Object Storage buckets](https://techdocs.akamai.com/linode-api/reference/get-object-storage-buckets) operation and store the label for the target bucket. You also need to provide the specific `region` where the bucket is located.
      --data string            Request body JSON, @file, or @- for stdin
      --database-id id         The ID of the Managed Database entity for the ticket. Run the [List Managed Databases](https://techdocs.akamai.com/linode-api/reference/get-databases-instances) operation and store the id for the target database.
      --description string     The full details of the issue or question.
      --domain-id id           The ID of the domain entity for the ticket. Run the [List domains](https://techdocs.akamai.com/linode-api/reference/get-domains) operation and store the id for the target domain.
      --firewall-id id         The ID of the Firewall entity for the ticket. Run the [List a Linode's firewalls](https://techdocs.akamai.com/linode-api/reference/get-linode-firewalls) operation and store the id for the target Linode firewall.
  -h, --help                   help for post-ticket
      --linode-id id           The ID of the Linode entity for the ticket. Run the [List Linodes](https://techdocs.akamai.com/linode-api/reference/get-linode-instances) operation and store the id for the target Linode.
      --lkecluster-id id       The ID of the Linode Kubernetes Engine (LKE) cluster entity for the ticket. Run the [List Kubernetes clusters](https://techdocs.akamai.com/linode-api/reference/get-lke-clusters) operation and store the id for the target LKE cluster.
      --longviewclient-id id   The ID of the Longview client entity for the ticket. Run the [List Longview clients](https://techdocs.akamai.com/linode-api/reference/get-longview-clients) operation and store the id for the target client.
      --managed-issue true     Whether this ticket is related to a [managed service](https://www.linode.com/products/managed/). If true, the following constraints apply: - You can't provide an entity, such as a `linode_id` or `bucket` with this request. - Your account needs a managed service [enabled](https://techdocs.akamai.com/linode-api/reference/post-enable-managed-service). (default "false")
      --nodebalancer-id id     The ID of the NodeBalancer entity for the ticket. Run the [List NodeBalancers](https://techdocs.akamai.com/linode-api/reference/get-node-balancers) operation and store the id for the target NodeBalancer.
      --region vlan            The ID of the [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where this ticket's target entity resides. This only applies to tickets for a vlan or an Object Storage `bucket`. > blue-book > > Set this to the `clusterId` for a legacy Object Storage `bucket`.
      --severity 1             The severity of the issue. A value of 1 indicates a major issue, `2` indicates a moderate priority issue, and `3` is a low priority issue. Your account may not have access to set this value. Talk to your Akamai account team for more details. (default "null")
      --summary string         The summary or title for this support ticket.
      --vlan id                The label of the VLAN entity for the ticket. Run the [List VLANs](https://techdocs.akamai.com/linode-api/reference/get-vlans) operation and store the id for the target VLAN. You also need to provide the specific `region` where the VLAN is located.
      --volume-id id           The ID of the volume entity for the ticket. Run the [List volumes](https://techdocs.akamai.com/linode-api/reference/get-volumes) operation and store the id for the target volume.
      --vpc-id id              The ID of the VPC entity for the ticket. Run the [List VPCs](https://techdocs.akamai.com/linode-api/reference/get-vpcs) operation and store the id for the target VPC.
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

* [linodectl support-tickets](./index.md)	 - support-tickets operations

###### Auto generated by spf13/cobra on 31-Aug-2026
