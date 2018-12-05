#!/usr/bin/env bash
./build.sh
sudo docker build -t connectus/api-server .
sudo docker push connectus/api-server
