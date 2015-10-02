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

# figure out if we need to install nginx-generator
echo "Installing nginx-generator if needed ..."
which nginx-generator || npm install -g nginx-generator
echo

# setting $GOPATH
echo "Setting GOPATH ..."
GOPATH=`pwd`
echo

# building the server
echo "Building the server ..."
go build server.go
echo

# set up Nginx
echo "Setting up Nginx ..."
FILE=/tmp/$NAME
cat /dev/null > $FILE
nginx-generator \
    --name $NAME \
    --domain $NAKED_DOMAIN \
    --type proxy \
    --var host=localhost \
    --var port=$PORT \
    - >> $FILE
nginx-generator \
    --name $NAME-www \
    --domain www.$NAKED_DOMAIN \
    --type redirect \
    --var to=$NAKED_DOMAIN \
    - >> $FILE
nginx-generator \
    --name $NAME-ww \
    --domain ww.$NAKED_DOMAIN \
    --type redirect \
    --var to=$NAKED_DOMAIN \
    - >> $FILE
nginx-generator \
    --name $NAME-w \
    --domain w.$NAKED_DOMAIN \
    --type redirect \
    --var to=$NAKED_DOMAIN \
    - >> $FILE
sudo cp $FILE /etc/nginx/sites-enabled/
rm -f $FILE
echo

# set up the server
echo "Setting up various directories ..."
sudo mkdir -p /var/log/$NAME/
sudo chown $THIS_USER:$THIS_GROUP /var/log/$NAME/
sudo mkdir -p /var/lib/$NAME/
sudo chown $THIS_USER:$THIS_GROUP /var/lib/$NAME/
echo

# add the supervisor scripts
echo "Copying supervisor script ..."
m4 \
    -D __USER__=$THIS_USER \
    -D  __PWD__=$THIS_PWD  \
    -D __NAME__=$NAME      \
    -D __PORT__=$PORT      \
    etc/supervisor/conf.d/$NAME.conf.m4 | sudo tee /etc/supervisor/conf.d/$NAME.conf
echo

# restart services
echo "Restarting services ..."
sudo supervisorctl reload
sudo service nginx restart
echo

## --------------------------------------------------------------------------------------------------------------------
