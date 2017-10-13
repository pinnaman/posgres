docker-machine start
docker-machine status
docker-machine env
eval "$(docker-machine env default)"

echo "Create Postgres sql volume container"
docker create -v /var/lib/postgresql/data --name postgres9.6.2-data busybox

docker images

docker ps

docker run --name mac-postgres9.6.2 -e POSTGRES_PASSWORD=secret -d --volumes-from postgres9.6.2-data pinnaman/postgres

docker run -it --link mac-postgres9.6.2:postgres --rm pinnaman/postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'

docker stop mac-postgres9.6.2
docker rm -v mac-postgres9.6.2

echo "Attach container to persistent volumes"
docker run --name mac-postgres9.6.2 -e POSTGRES_PASSWORD=secret -d --volumes-from postgres9.6.2-data pinnaman/postgres

docker ps

echo "login to psql shell"
docker run -it --link mac-postgres9.6.2:postgres --rm pinnaman/postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'

docker ps

echo "Open postgres sql to host"
docker run --name mac-postgres9.6.2 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d --volumes-from postgres9.6.2-data pinnaman/postgres

# Install pgcli client
sudo pip3.5 install -U pgcli

# connect to server listening on port 5432..
pgcli -h 192.168.99.100 -p 5432 -U postgres realtime
