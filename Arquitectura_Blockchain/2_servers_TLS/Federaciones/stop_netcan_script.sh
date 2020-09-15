#!/bin/bash

clear

echo "********************"
echo "Limpiando la instalación"
echo "********************"

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker volume prune --force
docker system prune --force
docker network prune --force

echo "********************"
echo "Limpiando chaincodes"
echo "********************"

docker rmi $(docker images dev-fci.federaciones.netcan.com-* -q)

ssh root@82.223.101.251 <<'ENDSSH'
echo "********************"
echo "Limpiando la instalación en el servidor ColegiosVeterianrios"
echo "********************"

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker volume prune --force
docker system prune --force
docker network prune --force

ENDSSH

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

echo "                   Apagada y servidor limpio"

echo ""