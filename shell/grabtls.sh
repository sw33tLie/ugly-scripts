#!/usr/bin/env bash
#usage: grabtls target.com for example: grabtls hackerone.com
URL=$1
PAT=$(echo $URL | sed 's/\./\\./g' | sed -e 's/^/\\./')
main(){
    curl -s http://tls.bufferover.run/dns\?\q\=$URL | jq -jr | egrep $PAT | awk -F'"' '{print $2}' | column -s ',' -t
}
main $@
