# GO ENS RESTFUL
<br />

## Deploy to docker

> make sure you already install docker and docker compose

- https://docs.docker.com/install/
- https://docs.docker.com/compose/install/ 
<br />

> then run the following command:

``` bash
$ dockerd
$ docker-compose up
```

<br><br>

## Ropsten version

<details>
<summary>RESOLVE</summary>

###### *Resolves an ENS name in to an Etheruem address. This will return an error if the name is not found or otherwise 0*

> POST: http://localhost:8081/ens/resolve

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth"
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "unregistered name"
}

# OR

{
    "status": "failed",
    "error_msg": "no address"
}

# OR

{
    "status": "failed",
    "error_msg": "could not parse address"
}

# OR

{
    "status": "failed",
    "error_msg": "401 Unauthorized basic auth failure\n"
}

# OR

{
    "status": "failed",
    "error_msg": "Post https://mainnet.infura.io/v3/ee73b34e48944bb083fd7d32dcfbcce8: context deadline exceeded"
}
```
</details>
<hr /><br>

<details>
<summary>REVERSE RESOLVE</summary>

###### *Resolves an address in to an ENS name. This will return an error if the name is not found or otherwise 0*

> POST: http://localhost:8081/ens/reverse-resolve

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"address": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "no resolution"
}

# OR

{
    "status": "failed",
    "error_msg": "0x1234... has no reverse lookup"
}

# OR

{
    "status": "failed",
    "error_msg": "Post https://mainnet.infura.io/v3/ee73b34e48944bb083fd7d32dcfbcce8: context deadline exceeded"
}
```
</details>
<hr /><br>

<details>
<summary>SET OWNER</summary>

###### *Sets the ownership of a domain*

> POST: http://localhost:8081/ens/set-owner

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"registrant": "0x1234...",
	"to_address": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "registrant": "0x1234...",
        "to_address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}
```
</details>
<hr /><br>

<details>
<summary>CREATE SUBDOMAIN</summary>

###### *Creates a subdomain on the name*

> POST: http://localhost:8081/ens/create-subdomain

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"subdomain": "iam",
	"registrant": "0x1234...",
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "subdomain": "iam",
        "registrant": "0x1234...",
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}
```
</details>
<hr /><br>

<details>
<summary>SET SUBDOMAIN OWNER</summary>

###### *Sets the owner for a subdomain of a name*

> POST: http://localhost:8081/ens/set-subdomain-owner

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"subdomain": "iam",
	"registrant": "0x1234...",
	"to_address": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "subdomain": "iam",
        "registrant": "0x1234...",
        "to_address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}
```
</details>
<hr /><br>

<details>
<summary>SET RESOLVER</summary>

###### *Sets the resolver for a name*

> POST: http://localhost:8081/ens/set-resolver

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"registrant": "0x1234...",
	"to_address": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "registrant": "0x1234...",
        "to_address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}
```
</details>
<hr /><br>

<details>
<summary>SET ADDRESS</summary>

###### *Sets the Ethereum address of the domain*

> POST: http://localhost:8081/ens/set-address

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"registrant": "0x1234...",
	"to_address": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "registrant": "0x1234...",
        "to_address": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "no resolver"
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}

# OR

{
    "status": "failed",
    "error_msg": "failed to estimate gas needed: gas required exceeds allowance (8000000) or always failing transaction"
}
```
</details>
<hr /><br>

<details>
<summary>SET CONTENT HASH</summary>

###### *Sets the content hash of the domain*

> POST: http://localhost:8081/ens/set-contenthash

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "iam.nawikartini.eth",
	"registrant": "0xe0456F670f74d2219f72BC7f0caCBBde9337F40e",
	"contenthash": [71,107,98, ...]
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "registrant": "0x1234...",
        "contenthash": [71,107,98, ...]
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "no resolver"
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}

# OR

{
    "status": "failed",
    "error_msg": "failed to estimate gas needed: gas required exceeds allowance (8000000) or always failing transaction"
}
```
</details>
<hr /><br>

<details>
<summary>SET ABI</summary>

###### *Sets the ABI associated with a name*

> POST: http://localhost:8081/ens/set-abi

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
    "abi_name": "Sample",
	"abi": '[{"constant":true,"inputs":...}]"',
	"abi_content_type": 1
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "abi_name": "Sample",
        "abi": '[{"constant":true,"inputs":...}]"',
        "abi_content_type": 1
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "no resolver"
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}

# OR

{
    "status": "failed",
    "error_msg": "failed to estimate gas needed: gas required exceeds allowance (8000000) or always failing transaction"
}
```
</details>
<hr /><br>

<details>
<summary>SET TEXT</summary>

###### *Sets the text associated with a name*

> POST: http://localhost:8081/ens/set-text

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"text_name": "Sample",
	"text_value": "Hello, world"
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "text_name": "Sample",
        "text_value": "Hello, world"
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "no resolver"
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}

# OR

{
    "status": "failed",
    "error_msg": "failed to estimate gas needed: gas required exceeds allowance (8000000) or always failing transaction"
}
```
</details>
<hr /><br>

<details>
<summary>SET NAME</summary>

###### *Sets the name*

> POST: http://localhost:8081/ens/set-name

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"domain_name": "nawikartini.eth",
	"registrant": "0x1234..."
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "nawikartini.eth",
        "registrant": "0x1234..."
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "Failed to obtain private key for 0x1234..."
}

# OR

{
    "status": "failed",
    "error_msg": "value resulted in fractional number of Wei"
}

# OR

{
    "status": "failed",
    "error_msg": "failed to estimate gas needed: gas required exceeds allowance (8000000) or always failing transaction"
}
```
</details>
<hr /><br>

</details>
