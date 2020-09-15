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
echo "Copiando los chaincode al servidor"
echo "********************"

#git clone https://github.com/DFLBB/json/chaincode /home/hyperledger/work/src/fabric-samples/TFM/chaincode

echo ""
echo "********************"
echo "Estableciendo las variables de entorno de chaincodes"
echo "********************"

export CHANNEL_NAME=netcanchannel
export CA_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem
export ORDENER_URL=orderer.netcan.com:7050
export CC_VERSION=1.0.0

LISTACHAINCODES=`ls /opt/gopath/src/github.com/chaincode/netcan`
for CHAINCODE in $LISTACHAINCODES; do
    if [ "$CHAINCODE" != "netcan" ]; then

        echo ""
        echo "********************"
        echo "Instalando e instanciando el chaincode $CHAINCODE"
        echo "********************"

        export CC_NOMBRE=$CHAINCODE
        export CC_FILE=github.com/chaincode/netcan/$CHAINCODE/cc

        peer chaincode install     -n $CC_NOMBRE -v $CC_VERSION -p $CC_FILE
        peer chaincode instantiate -n $CC_NOMBRE -v $CC_VERSION -c '{"Args":["init"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME -P "OR ('FederacionesMSP.peer','ColegiosVeterinariosMSP.peer')"
    fi
done