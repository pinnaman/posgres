package main

import (  
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (  
  host     = "192.168.99.100"
  port     = 5432
  user     = "postgres"
  password = "secret"
  dbname   = "realtime"
)

func main() {  
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  var (
	id string
        created string
        start string
        end sql.NullString
	title string
        completed string
)
rows, err := db.Query("select * from events ")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &created,&start,&end,&title,&completed)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id,created,start,end,title,completed)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}


}

