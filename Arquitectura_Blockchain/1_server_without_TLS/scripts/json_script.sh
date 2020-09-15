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

for DOCKER in $(docker ps --filter "since=cli" -q); do

    echo ""
    echo "********************"
    echo "Copiando los archivos json de carga inicial de datos al docker $DOCKER"
    echo "********************"

    docker cp /home/hyperledger/work/src/fabric-samples/TFM2/json $DOCKER:./json
done