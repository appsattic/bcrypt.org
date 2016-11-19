#!/bin/bash
## ----------------------------------------------------------------------------

set -e

## ----------------------------------------------------------------------------
# Set these to your preferred values.

THIS_USER=`id -un`
THIS_GROUP=`id -gn`
THIS_PWD=`pwd`

NAME=org-bcrypt
NAKED_DOMAIN=bcrypt.org
PORT=8791

## ----------------------------------------------------------------------------

# building the server
echo "Building the server ..."
gb build
echo

# set up the server
echo "Setting up various directories ..."
sudo mkdir -p /var/log/$NAME/
sudo chown $THIS_USER:$THIS_GROUP /var/log/$NAME/
sudo mkdir -p /var/lib/$NAME/
sudo chown $THIS_USER:$THIS_GROUP /var/lib/$NAME/
echo

# copy the Supervisor config
echo "Copying supervisor script ..."
sudo cp etc/supervisor/conf.d/org-bcrypt.conf /etc/supervisor/conf.d/org-bcrypt.conf
echo

# copy the Caddy config
echo "Copying caddy script ..."
sudo cp etc/caddy/vhosts/org.bcrypt.conf /etc/caddy/vhosts/org.bcrypt.conf
echo

# restart services
echo "Restarting services ..."
sudo systemctl restart caddy.service
sudo systemctl restart supervisor.service
echo

## --------------------------------------------------------------------------------------------------------------------
