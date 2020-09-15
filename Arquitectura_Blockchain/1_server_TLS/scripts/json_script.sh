#!/bin/bash

clear

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
echo "Copiando los archivos json de carga inicial de datos"
echo "********************"

#git clone https://github.com/DFLBB/json/json /home/hyperledger/work/src/fabric-samples/TFM/json

for DOCKER in $(docker ps --filter "since=cli" -q); do

    echo ""
    echo "********************"
    echo "Copiando los archivos json de carga inicial de datos al docker $DOCKER"
    echo "********************"

    docker cp /home/hyperledger/work/src/fabric-samples/TFM/json $DOCKER:./json
done