#!/usr/bin/env bash

source setenv.sh

# Rabbitmq send
echo "Subindo o go-rabbitmq-send..."
docker run -d --name go-rabbitmq-receive --network rabbitmq-net  \
-p 8181:8080 \
-e MYSQL_USER=${MYSQL_USER} \
-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
-e MYSQL_HOSTNAME=${MYSQL_HOSTNAME} \
-e MYSQL_DATABASE=${MYSQL_DATABASE} \
-e MYSQL_PORT=${MYSQL_PORT} \
-e RABBITMQ_USER=${RABBITMQ_USER} \
-e RABBITMQ_PASS=${RABBITMQ_PASS} \
-e RABBITMQ_HOSTNAME=${RABBITMQ_HOSTNAME} \
-e RABBITMQ_PORT=${RABBITMQ_PORT} \
-e RABBITMQ_VHOST=${RABBITMQ_VHOST} \
marceloagmelo/go-rabbitmq-receive

# Listando os containers
docker ps
