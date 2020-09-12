#!/bin/bash

echo "********************"
echo "Copiando los archivos json de carga inicial de datos"
echo "********************"

for DOCKER in $(docker ps --filter "since=cli" -q); do

    echo "********************"
    echo "Copiando los archivos json de carga inicial de datos al docker $DOCKER"
    echo "********************"

    docker cp /home/hyperledger/work/src/fabric-samples/TFM/json $DOCKER:./json
done