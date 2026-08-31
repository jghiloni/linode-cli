## linodectl configurations post-rebuild-node-balancer-config

Rebuild a config

### Synopsis

Rebuilds a NodeBalancer config and its nodes that you have permission to modify. Use this operation to update a NodeBalancer's config and nodes with a single request. You can't rebuild a NodeBalancer config if it has a `cannot_delete_with_subresources` lock. Only account administrators can remove locks using the [Delete a resource lock](https://techdocs.akamai.com/linode-api/reference/delete-resource-lock) operation. > construction > > You can configure UDP on the same NodeBalancer that also uses TCP, HTTP, or HTTPS, but only when managing it through the API. If UDP is configured and you make changes to the TCP, HTTP or HTTPS settings in Cloud Manager, the existing UDP configuration will be overwritten. This is because Cloud Manager doesn't currently support UDP. **Permissions and scopes** To call this operation, you need permissions, based on the model you're using: - **Identity and access permissions**. Your user needs a role with these permissions. [Learn more](https://techdocs.akamai.com/cloud-computing/docs/identity-access-cm-available-roles). - Permissions: `rebuild_nodebalancer_config` - **OAuth scopes**. Your user needs these scopes assigned. [Learn more](https://techdocs.akamai.com/linode-api/reference/get-started#oauth). - Scopes: `nodebalancers:read_write` **CLI** *HTTPS* ```shell linode-cli nodebalancers config-rebuild \ 12345 4567 \ --port 443 \ --protocol https \ --algorithm roundrobin \ --stickiness http_cookie \ --check http_body \ --check_interval 90 \ --check_timeout 10 \ --check_attempts 3 \ --check_path "/test" \ --check_body "it works" \ --check_passive true \ --proxy_protocol "none" \ --ssl_cert "-----BEGIN CERTIFICATE----- CERTIFICATE_INFORMATION -----END CERTIFICATE-----" \ --ssl_key "-----BEGIN PRIVATE KEY----- PRIVATE_KEY_INFORMATION -----END PRIVATE KEY-----" \ --cipher_suite recommended \ --nodes.label "node1" --nodes.address "192.168.210.120:80" --nodes.mode "accept" --nodes.weight 50 \ --nodes '[{"address":"192.168.210.122:80","label":"node2","weight":50,"mode":"accept"}]' \ --nodes '[{"address":"10.0.0.45:80","label":"vpc-node","weight":10,"mode":"accept","subnet_id:1"}]' ``` *UDP* ```shell linode-cli nodebalancers config-rebuild \ 12345 4567 \ --port 80 \ --protocol udp \ --algorithm ring_hash \ --udp_check_port 80 \ --nodes.label "node1" --nodes.address "192.168.210.120:80" --nodes.mode "accept" --nodes.weight 50 \ --nodes '[{"address":"192.168.210.122:80","label":"node2","weight":50}]' \ --nodes '[{"address":"10.0.0.45:80","label":"vpc-node","weight":10,"mode":"accept","subnet_id:1"}]' ``` *TCP* ```shell linode-cli nodebalancers config-rebuild \ 12345 4567 \ --port 80 \ --protocol tcp \ --algorithm roundrobin \ --stickiness none \ --proxy_protocol "v2" --nodes.label "node1" --nodes.address "192.168.210.120:80" --nodes.mode "accept" --nodes.weight 50 \ --nodes '[{"address":"192.168.210.122:80","label":"node2","weight":50,"mode":"accept"}]' \ --nodes '[{"address":"10.0.0.45:80","label":"vpc-node","weight":10,"mode":"accept","subnet_id:1"}]' ``` *HTTP* ```shell linode-cli nodebalancers config-rebuild \ 12345 4567 \ --port 440 \ --protocol http \ --algorithm roundrobin \ --stickiness none \ --check http_body \ --check_interval 90 \ --check_timeout 10 \ --check_attempts 3 \ --check_path "/test" \ --check_body "it works" \ --nodes.label "node1" --nodes.address "192.168.210.120:80" --nodes.mode "accept" --nodes.weight 50 \ --nodes '[{"address":"192.168.210.122:80","label":"node2","weight":50,"mode":"accept"}]' \ --nodes '[{"address":"10.0.0.45:80","label":"vpc-node","weight":10,"mode":"accept","subnet_id:1"}]' ``` [Learn more](https://techdocs.akamai.com/cloud-computing/docs/getting-started-with-the-linode-cli)

```
linodectl configurations post-rebuild-node-balancer-config [flags]
```

### Options

```
      --api-version v4            __Enum__ Call either the v4 URL, or `v4beta` for operations still in Beta.
      --config-id string          The ID of the Config to access.
      --data string               Request body JSON, @file, or @- for stdin
  -h, --help                      help for post-rebuild-node-balancer-config
      --node-balancer-id string   The ID of the NodeBalancer.
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

* [linodectl configurations](./index.md)	 - configurations operations

###### Auto generated by spf13/cobra on 31-Aug-2026
