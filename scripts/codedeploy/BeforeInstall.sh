#!/bin/bash
set -eux pipefail

# remove prev program 
# TODO: Needed?
# sudo rm -f /usr/sbin/httpd
# sudo rm -f /usr/lib/systemd/system/httpd.service

# shellcheck disable=SC2046
sudo chmod 0444 $(find /srv -type f)
# clean
sudo rm /srv/assets.zip
