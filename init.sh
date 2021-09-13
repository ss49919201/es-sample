#!/bin/bash
# https://hub.docker.com/_/elasticsearch
# https://www.elastic.co/guide/en/elasticsearch/reference/7.13/docker.html
docker network create somenetwork
docker pull docker.elastic.co/elasticsearch/elasticsearch:latest:7.13.4
docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.13.4
