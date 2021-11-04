#!/bin/bash

echo "Info: starting infra"
docker-compose up -d kong kafka zookeeper
echo "Info: wait for 10 seconds to start infra"
sleep 10

echo "Info: starting menu service postgres"
docker-compose up -d menu_postgres
echo "Info: wait for 10 seconds to start postgres"
sleep 10

echo "Info: starting serach service elasticsearch"
docker-compose up -d search_elasticsearch
echo "Info: wait for 10 seconds to start elasticsearch"
sleep 10

if ! docker exec -it kafka kafka-topics --list --bootstrap-server kafka:9092 | grep -q 'menuitems'; then
    echo "Info: creating menu servie kafka topics"
    docker exec -it kafka kafka-topics --create --bootstrap-server kafka:9092 --topic menuitems --partitions 1 --replication-factor 1
fi

echo "Info: starting service"
docker-compose build menu_service search_service
docker-compose up -d menu_service search_service

echo "Info: setup done"
