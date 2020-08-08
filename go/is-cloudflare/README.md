# Is Cloudflare
Check if an IP is owned by Cloudflare.

Example:
```
$ go run main.go -ip 12.34.56.78
false
```

This is useful when doing bug bounties as running a full port scan (from 1 to 65535) on Cloudflare's servers is not a good idea.

Read more [here](https://support.cloudflare.com/hc/en-us/articles/200169156-Identifying-network-ports-compatible-with-Cloudflare-s-proxy).