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
echo "Levantando la red"
echo "********************"

docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d

echo "********************"
echo "Copiando scripts al CLI"
echo "********************"

docker cp /home/hyperledger/work/src/fabric-samples/TFM2/scripts cli:./netcan_scripts
export TERM=xterm
clear