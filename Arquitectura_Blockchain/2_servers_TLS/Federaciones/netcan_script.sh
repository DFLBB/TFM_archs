#!/bin/bash

clear
echo "**************************************************************************"
echo "                                  Bienvenido a"
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

rm -rf ./json
rm -rf ./chaincode
rm -rf ./scripts
git clone https://github.com/DFLBB/TFM_archs /home/hyperledger/work/src/fabric-samples/TFM/TFM_archs
mv ./TFM_archs/chaincode ./chaincode
mv ./TFM_archs/json ./json
mv ./TFM_archs/scripts1 ./scripts
cd scripts/ && chmod +x *.sh && cd ..
rm -rf ./TFM_archs

echo ""
echo "********************"
echo "Limpiando chaincodes"
echo "********************"

docker rmi $(docker images dev-fci.federaciones.netcan.com-* -q)

echo ""
echo "********************"
echo "Copiando los scripts de arranque al servidor"
echo "********************"

clear

./scripts/red_script.sh
./scripts/Serv2_script.sh
docker exec cli /netcan_scripts/config_script.sh
docker exec cli /netcan_scripts/chaincode_script.sh
./scripts/json_script.sh
docker exec cli /netcan_scripts/carga_script.sh