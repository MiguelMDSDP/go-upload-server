#!/usr/bin/env bash

if [ ! -f ../base-stack/devstack-header.sh ]; then
  echo "No devstack-header.sh file found."
  echo "Make sure you have the base-stack repo in the same folder as this repo."
  exit 1
fi

source ../base-stack/devstack-header.sh

case "$1" in
  dev )
    export DATA_PATH=$DEV_DATA_PATH
  ;;

  prod )
    export DATA_PATH=$PROD_DATA_PATH
  ;;

  -h )
    echo "This script starts the stack."
    echo "You need to specify a valid environment [prod/dev]"
    echo "You can also use './configure.sh [mode] stop' to stop the stack. "
    exit 0
  ;;

  *)
    echo "Invalid mode specified, valid modes are:"
    echo "./configure.sh prod"
    echo "./configure.sh dev"
    exit 1
  ;;
esac

if [ "$2" = "stop" ]; then
  echo "Stopping services."
  docker stack rm fs
  exit 0
fi

setHosts
echo "Checking if docker network exists."
docker network create proxy --driver overlay

# if [ ! -f $DATA_PATH/traefik/cert.crt ]; then
#   echo "Generating self signed cert."
#   prevDir=$PWD
#   cd $DATA_PATH/traefik/
#   openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout cert.key -out cert.crt
#   chmod 644 cert.crt && chmod 600 cert.key
#   cd $prevDir
# fi

#If we need to use letsencrypt.
# if ! [ -f /mnt/traefik/acme.json ]; then
#   echo "Creating acme.json for traefik."
#   touch /mnt/traefik/acme.json
#   chmod 600 /mnt/traefik/acme.json
# fi

echo "Starting admin."
docker stack deploy -c totfs-stack.yml fs
