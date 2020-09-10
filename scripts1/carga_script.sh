#!/bin/bash

export CHANNEL_NAME=netcanchannel
export CA_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem
export ORDENER_URL=orderer.netcan.com:7050

echo "*************************************************"
echo "Carga inicial de datos"
echo "*************************************************"

echo "*************************************************"
echo "Cargando datos de PERFILES DE PERSONAS"
echo "*************************************************"

peer chaincode invoke -n perfiles -c '{"function":"cargarDatosIniciales","Args":["./json/perfiles.json", "{\"IDPersona\":0}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n perfiles -c '{"function":"asignarEstado_OnlyAdmin","Args":["PERFILES_PERSONAS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":48}", "{\"IDPersona\":0}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n perfiles -c '{"function":"cancelarPerfilPersona","Args":["{\"IDPersona\":0,\"CODPerfil\":\"ADMINISTRADOR\"}", "{\"IDPersona\":0}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME


echo "*************************************************"
echo "Cargando datos de GRUPOS"
echo "*************************************************"

peer chaincode invoke -n razas -c '{"function":"cargarDatosIniciales_Grupos","Args":["./json/grupos.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n razas -c '{"function":"asignarEstado","Args":["GRUPOS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":12}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de RAZAS"
echo "*************************************************"

peer chaincode invoke -n razas -c '{"function":"cargarDatosIniciales","Args":["./json/razas.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n razas -c '{"function":"asignarEstado","Args":["RAZAS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":346}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de PERSONAS"
echo "*************************************************"

peer chaincode invoke -n personas -c '{"function":"cargarDatosIniciales","Args":["./json/personas_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n personas -c '{"function":"asignarEstado","Args":["PERSONAS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":1000}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de AFIJOS"
echo "*************************************************"

peer chaincode invoke -n afijos -c '{"function":"cargarDatosIniciales","Args":["./json/afijos.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n afijos -c '{"function":"asignarEstado","Args":["AFIJOS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":50}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de PROPIETARIOS DE AFIJOS"
echo "*************************************************"

peer chaincode invoke -n afijos -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/afijos_propietarios.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n afijos -c '{"function":"asignarEstado","Args":["AFIJOS_PROPIETARIOS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":50}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de PERROS"
echo "*************************************************"

peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales","Args":["./json/perros_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales","Args":["./json/perros_002.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales","Args":["./json/perros_003.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales","Args":["./json/perros_004.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales","Args":["./json/perros_005.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"asignarEstado","Args":["PERROS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":5759}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de PROPIETARIOS DE PERROS"
echo "*************************************************"

peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/perros_propietarios_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/perros_propietarios_002.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/perros_propietarios_003.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/perros_propietarios_004.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"cargarDatosIniciales_Propietarios","Args":["./json/perros_propietarios_005.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n perros -c '{"function":"asignarEstado","Args":["PERROS_PROPIETARIOS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":5759}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de VETERINARIOS"
echo "*************************************************"

peer chaincode invoke -n veterinarios -c '{"function":"cargarDatosIniciales","Args":["./json/veterinarios.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n veterinarios -c '{"function":"asignarEstado","Args":["VETERINARIOS_PERSONAS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":21}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de MICROCHIPS"
echo "*************************************************"

peer chaincode invoke -n microchips -c '{"function":"cargarDatosIniciales","Args":["./json/microchips_perros_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n microchips -c '{"function":"cargarDatosIniciales","Args":["./json/microchips_perros_002.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n microchips -c '{"function":"cargarDatosIniciales","Args":["./json/microchips_perros_003.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n microchips -c '{"function":"cargarDatosIniciales","Args":["./json/microchips_perros_004.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n microchips -c '{"function":"cargarDatosIniciales","Args":["./json/microchips_perros_005.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n microchips -c '{"function":"asignarEstado","Args":["MICROCHIPS_PERROS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":5759}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "Cargando datos de VACUNAS"
echo "*************************************************"

peer chaincode invoke -n vacunas -c '{"function":"cargarDatosIniciales_VacunasProteccion","Args":["./json/vacunas_proteccion_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s


echo "*************************************************"
echo "Cargando datos de VACUNACIONES"
echo "*************************************************"

peer chaincode invoke -n vacunas -c '{"function":"asignarEstado","Args":["VACUNAS_PROTECCION", "{\"docType\":\"CONTADOR\",\"IDMaximo\":11}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n vacunas -c '{"function":"cargarDatosIniciales","Args":["./json/vacunas_perros_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n vacunas -c '{"function":"asignarEstado","Args":["VACUNAS_PERROS", "{\"docType\":\"CONTADOR\",\"IDMaximo\":89}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
peer chaincode invoke -n vacunas -c '{"function":"cargarDatosIniciales_VacunasProteccion","Args":["./json/vacunas_perros_proteccion_001.json"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME
sleep 7s
peer chaincode invoke -n vacunas -c '{"function":"asignarEstado","Args":["VACUNAS_PERROS_PROTECCION", "{\"docType\":\"CONTADOR\",\"IDMaximo\":512}"]}' -o $ORDENER_URL --tls --cafile $CA_FILE -C $CHANNEL_NAME

echo "*************************************************"
echo "DATOS INICIALES CARGADOS"
echo "*************************************************"

