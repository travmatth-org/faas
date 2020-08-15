#!/bin/bash
set -eux pipefail

sudo unzip -o /srv/assets.zip -d /srv
sudo rm /srv/assets.zip
# shellcheck disable=SC2046
sudo chmod 0444 $(find /srv -type f)
# enable faas service
sudo systemctl enable httpd
