# Grafana Login
Log in to Grafana using Golang!
Tip: add this to your automation pipeline and check for default credentials (admin:admin)

Example:
```
$ go run . -U admin -P admin -h https://grafana.example.com:3000
Response Status: 200 OK
Response Headers: map[Cache-Control:[no-cache] Content-Length:[23] Content-Type:[application/json] Date:[Sat, 08 Aug 2020 11:54:23 GMT] Expires:[-1] Pragma:[no-cache] Set-Cookie:[grafana_session=SECRET_SESSION_TOKEN_AAAAAAAAAAA; Path=/; Max-Age=2595600; HttpOnly; SameSite=Lax] X-Content-Type-Options:[nosniff] X-Frame-Options:[deny] X-Xss-Protection:[1; mode=block]]
Response Body: {"message":"Logged in"}

```