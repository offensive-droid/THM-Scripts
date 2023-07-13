//This is not yet finished.
package main

import (
        "fmt"
        "os"

        "github.com/gocolly/colly"
)

type methods interface {
        get()
}

type RSA struct {
        ip        string
        directory string
}

func call(m methods) {
        m.get()
}

func (r *RSA) get() {
        c := colly.NewCollector(
                colly.AllowedDomains(r.ip, "www."+r.ip, "https://"+r.ip, "http://"+r.ip, "https://"+r.ip+"directory", "http://"+r.ip+"directory"),
        )
        // Allow URLs with the specified directory in the path

        // Set a custom User-Agent header
        c.OnRequest(func(r *colly.Request) {
                r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
                fmt.Println("Visiting", r.URL)
        })

        c.OnResponse(func(r *colly.Response) {
                fmt.Println("Visited", r.Request.URL)
                fmt.Println(r.StatusCode)
        })

        c.OnHTML("p", func(e *colly.HTMLElement) {
                fmt.Println(e.Text)
        })

        err := c.Visit("http://" + r.ip + r.directory)
        if err != nil {
                fmt.Println("Error:", err)
        }
}

func main() {
        if len(os.Args) < 2 {
                fmt.Println("Please provide a URL as a command-line argument.")
                return
        }

        rsa := &RSA{ip: os.Args[1]}
        if len(os.Args) > 2 {
                rsa.directory = os.Args[2]
        }

        call(rsa)
}
