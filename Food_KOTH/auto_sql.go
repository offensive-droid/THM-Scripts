package main

import (
        "database/sql"
        "fmt"
        "log"
        "os"

        _ "github.com/go-sql-driver/mysql"
)

type Auth interface {
        Connect()
}

type MySQL struct {
        ip string
}

type User struct {
        username string
        password string
}

func (m MySQL) Connect() {
        db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s:3306)/users", m.ip))
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        err = db.Ping()
        if err != nil {
                log.Fatal(err)
        }

        rows, err := db.Query("SELECT username, password FROM User")
        if err != nil {
                log.Fatal(err)
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
                var user User
                err := rows.Scan(&user.username, &user.password)
                if err != nil {
                        log.Fatal(err)
                }
                users = append(users, user)
        }

        if err = rows.Err(); err != nil {
                log.Fatal(err)
        }

        fmt.Println("Users:")
        for _, user := range users {
                fmt.Println(user.username, user.password)
        }
}

func callAuth(a Auth) {
        a.Connect()
}

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: go run main.go <ip>")
                return
        }

        c := &MySQL{
                ip: os.Args[1],
        }

        callAuth(c)
}

                                                                                                                                             

