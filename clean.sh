#!/bin/bash
docker rm $(docker stop elasticsearch) 
docker network rm somenetwork
