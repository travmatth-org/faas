#!/bin/bash -xe

# send script output to /tmp/ApplicationStart.log for debugging
exec >> /tmp/ApplicationStart.log 2>&1

sudo systemctl enable httpd
sudo systemctl start httpd
