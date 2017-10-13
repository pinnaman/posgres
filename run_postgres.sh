docker stop postgres-db
docker rm postgres-db
docker run --name postgres-db -e POSTGRES_PASSWORD=secret -d pinnaman/postgres:latest
sleep 10
#docker run -it --rm --link postgres-db:postgres postgres:alpine psql -h postgres -U postgres
docker run -it --rm --link postgres-db:postgres pinnaman/postgres:latest psql -h postgres -U postgres

#docker run -v $(PWD):/data -it --link local-postgres9.6.2:postgres --rm postgres:alpine sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'
