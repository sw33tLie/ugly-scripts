# Shodan Favicon Hash
This script can be used to calculate the Shodan favicon hash for any website.
Credits to @sshell_ for the original code.

Example:
```
$ go run main.go -u https://github.com/favicon.ico
1848946384
```

You can now search for hosts having that favicon using this Shodan query:
```
http.favicon.hash:1848946384
```
