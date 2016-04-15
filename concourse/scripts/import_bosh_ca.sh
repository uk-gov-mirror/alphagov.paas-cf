#!/bin/sh
set -eu

echo "Adding bosh-CA to root certificates"
tar -xf bosh-CA/bosh-CA.tar.gz -C /usr/local/share/ca-certificates bosh-CA.crt
update-ca-certificates
