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
echo "Configurando el servidor ColegiosVeterianrios"
echo "********************"

ssh root@82.223.101.251 bash -c "'
cd /home/hyperledger/work/src/fabric-samples/TFM/scripts/
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
echo "Limpiando la instalación"
echo "********************"

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker volume prune --force
docker system prune --force
docker network prune --force

echo "********************"
echo "Levantando la red en el servidor ColegiosVeterinarios"
echo "********************"
docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml up -d
 
'"