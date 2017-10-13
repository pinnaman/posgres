docker rm -v $(docker ps -a -q -f status=exited)
docker rm -v $(docker ps -a -q -f status=Created)
docker rmi -f $(docker images | grep "<none>" | awk '{print $3}')
