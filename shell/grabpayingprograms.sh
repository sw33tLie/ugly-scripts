#!/usr/bin/env bash

#to be used as a cronjob for your automation needs 

BOUNTY_TARGETS="/home/$(whoami)/bounty-targets-data"
BLACKLIST="/home/$(whoami)/blacklist.txt"
HACKERONE="$BOUNTY_TARGETS/data/hackerone_data.json"
BUGCROWD="$BOUNTY_TARGETS/data/bugcrowd_data.json"
INTIGRITI="$BOUNTY_TARGETS/data/intigriti_data.json"
YESWEHACK="$BOUNTY_TARGETS/data/yeswehack_data.json"
DOMAINS="/home/$(whoami)/domains.txt"
#git pull 
update_wildcards() {
        if [ -d "$BOUNTY_TARGETS" ]; then
                echo -n "Updating bounty-targets-data... "
                (cd "$BOUNTY_TARGETS" && git pull) >/dev/null 2>/dev/null
                echo "Done"
        else
                echo -n "bounty-targets-data not found. exiting"
                exit
        fi
}
#jq and fuckery
grabwildcardsthatpay() {
    echo "Ching Ching $ $ "
    jq -r '.[].targets.in_scope[] | select(.eligible_for_bounty==true and .asset_type == "URL") | .asset_identifier' < "$HACKERONE" | tee /tmp/domains.mixed
    jq -r '.[] | select(.max_payout!=0) | .targets.in_scope[] | select(.type=="website testing") | .target' < "$BUGCROWD"  | tee -a /tmp/domains.mixed
    jq -r '.[] | select(.max_bounty!=0.0) | .targets.in_scope[] | select(.type == "url") | .endpoint' < "$INTIGRITI" | tee -a /tmp/domains.mixed
    jq -r '.[] | select(.max_bounty!=0.0) | .targets.in_scope[] | select(.type == "web-application") | .target' < "$YESWEHACK" | tee -a /tmp/domains.mixed
    egrep '^\*\.' /tmp/domains.mixed | egrep -v '\.\*$' | sed -e 's/^\*\.//g'| tee /tmp/domains.noasterisk
    rm -rf /tmp/domains.mixed
}

clear_garbage() {
        echo "Cleaning up garbage!"
grep -vf "$BLACKLIST" /tmp/domains.noasterisk > /tmp/wildcard.ready
        mv /tmp/wildcard.ready "$DOMAINS"
echo "Done."
}
main() {
    update_wildcards
    grabwildcardsthatpay
    clear_garbage
}
main $@
