#!/bin/sh -e

docker pull seannguyen/coin-tracker;
container_name=coin-tracker;
container_temp_name=coin-tracker-temp;

# Create new container
docker run -d \
  --name ${container_temp_name} \
  -v /home/seannguyen/personal-infra/coin-tracker/configs/config.yml:/go/src/github.com/seannguyen/coin-tracker/configs/config.yml \
  seannguyen/coin-tracker

# Remove old container
if [ "docker ps -q -f name=${container_name}" ]; then
    docker stop ${container_name};
fi
if [ "docker ps -aq -f status=exited name=${container_name}" ]; then
    docker rm ${container_name};
fi

# Switch over
docker rename ${container_temp_name} ${container_name}