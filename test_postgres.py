#!/usr/local/bin/python3.5

import psycopg2

try:
    # connect to docker machine IP
    connect_str = "dbname='realtime' user='postgres' host='192.168.99.100' " + \
                  "password='secret'"
    # use our connection values to establish a connection
    conn = psycopg2.connect(connect_str)
    # create a psycopg2 cursor that can execute queries
    cursor = conn.cursor()

    # create a new table with a single column called "name"
    # cursor.execute("""CREATE TABLE tutorials (name char(40));""")
    cursor.execute("""CREATE EXTENSION "uuid-ossp" """)
    cursor.execute("""CREATE TABLE events(id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_time TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    title varchar(512) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
    );""")

    # run a SELECT statement - no data in there, but we can try it
    cursor.execute("""SELECT datname,usename,client_addr,client_port FROM pg_stat_activity""")
    rows = cursor.fetchall()
    print(rows)
    # close communication with the PostgreSQL database server
    cursor.close()
    # commit the changes
    conn.commit()
except Exception as e:
    print("Uh oh, can't connect. Invalid dbname, user or password?")
    print(e)
