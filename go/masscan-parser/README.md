# Masscan Parser
This script can be used to parse Masscan's json output and print all IP:port combinations.

Example:
```
$ sudo masscan 12.34.56.78/16 -p1-65535 --rate 100000 --output-format json --output-filename masscan-output.json
```
Then run:
```
$ go run main.go -f masscan-output.json
12.34.0.14:80
12.34.0.14:8443
12.34.0.36:22
...
```
