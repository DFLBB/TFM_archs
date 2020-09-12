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
echo "Estableciendo las variables de entorno"
echo "********************"

export CHANNEL_NAME=netcanchannel

echo ""
echo "********************"
echo "Creando el canal"
echo "********************"

peer channel create -o orderer.netcan.com:7050 -t 15s -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem

echo ""
echo "********************"
echo "Adhiriendo el peer FCI de Federaciones al canal"
echo "Puede tardar un poco..."
echo "********************"
sleep 60

peer channel join -b netcanchannel.block

while [ $? -eq 1 ]; 
do 
     echo ""
     echo "********************"
     echo "Reintentando en 5 segundos"
     echo "********************"
     sleep 5
     peer channel join -b netcanchannel.block;
done

echo ""
echo "********************"
echo "Adhiriendo el resto de peers de ambas organizaciones al canal"
echo "********************"

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVMadrid.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVMadrid.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/users/Admin@federaciones.netcan.com/msp/ CORE_PEER_ADDRESS=RSCE.federaciones.netcan.com:7051 CORE_PEER_LOCALMSPID="FederacionesMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/peers/RSCE.federaciones.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/users/Admin@federaciones.netcan.com/msp/ CORE_PEER_ADDRESS=TKC.federaciones.netcan.com:7051 CORE_PEER_LOCALMSPID="FederacionesMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/peers/TKC.federaciones.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/users/Admin@federaciones.netcan.com/msp/ CORE_PEER_ADDRESS=ACW.federaciones.netcan.com:7051 CORE_PEER_LOCALMSPID="FederacionesMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/federaciones.netcan.com/peers/ACW.federaciones.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block

CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVAndalucia.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVAndalucia.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVAragon.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVAragon.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVAsturias.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVAsturias.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVIllesBalears.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVIllesBalears.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVCanarias.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVCanarias.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVCantabria.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVCantabria.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVCastillayLeon.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVCastillayLeon.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVCastillalaMancha.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVCastillalaMancha.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVCataluna.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVCataluna.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVComunitatValenciana.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVComunitatValenciana.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVExtremadura.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVExtremadura.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVGalicia.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVGalicia.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVMurcia.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVMurcia.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVNavarra.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVNavarra.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVPaisVasco.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVPaisVasco.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp/ CORE_PEER_ADDRESS=CVLaRioja.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVLaRioja.colegiosveterinarios.netcan.com/tls/ca.crt peer channel join -b netcanchannel.block

echo ""
echo "********************"
echo "Declarando los pares de anclaje"
echo "********************"

peer channel update -o orderer.netcan.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/FederacionesMSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/users/Admin@colegiosveterinarios.netcan.com/msp CORE_PEER_ADDRESS=CVMadrid.colegiosveterinarios.netcan.com:7051 CORE_PEER_LOCALMSPID="ColegiosVeterinariosMSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/colegiosveterinarios.netcan.com/peers/CVMadrid.colegiosveterinarios.netcan.com/tls/ca.crt peer channel update -o orderer.netcan.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/ColegiosVeterinariosMSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem