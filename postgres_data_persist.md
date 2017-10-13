####
    docker create -v /var/lib/postgresql/data --name postgres9.6.2-data busybox
    docker run --name local-postgres9.6.2 -e POSTGRES_PASSWORD=secret -d --volumes-from postgres9.6.2-data postgres:alpine
    docker ps
    docker run -it --link local-postgres9.6.2:postgres --rm postgres:alpine sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres'
    docker stop local-postgres9.6.2
    docker rm -v local-postgres9.6.2
    docker run --name local-postgres9.6.2 -e POSTGRES_PASSWORD=secret -d --volumes-from postgres9.6.2-data postgres:alpine
    docker run -it --link local-postgres9.6.2:postgres --rm postgres:alpine sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U postgres -l'
