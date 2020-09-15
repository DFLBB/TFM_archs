#!/bin/bash

echo "**************************************************************************"
echo ""
echo "        _/      _/  _/_/_/_/  _/_/_/_/_/    _/_/_/    _/_/    _/      _/"
echo "       _/_/    _/  _/            _/      _/        _/    _/  _/_/    _/"
echo "      _/  _/  _/  _/_/_/        _/      _/        _/_/_/_/  _/  _/  _/"
echo "     _/    _/_/  _/            _/      _/        _/    _/  _/    _/_/"
echo "    _/      _/  _/_/_/_/      _/        _/_/_/  _/    _/  _/      _/"
echo ""
echo "**************************************************************************"

echo ""

echo "********************"
echo "Levantando la red swarm"
echo "********************"

docker network create --attachable --driver overlay netcan

echo "********************"
echo "Levantando la red Hyperledger Fabric"
echo "********************"

docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d

clear

echo "********************"
echo "Copiando scripts al CLI"
echo "********************"

docker cp /home/hyperledger/work/src/fabric-samples/TFM/scripts cli:./netcan_scripts

clear