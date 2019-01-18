package main

import (
	"sync"
	"net"
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
	"strconv"
	"math/rand"
)

// assuming localhost:7777 is emitting messages separated by newline
// with format: Q|DELL|NYSE|10|101

type num struct {
	Exchange string
	Value int
}

type quote struct {
	BestBid *num
	BestOffer *num
}

var (
	addr = "localhost:7777"
	lock = sync.Mutex{}
	quotes = make(map[string]*quote, 0)
)


func main() {
	go runServer()
	go runParser()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func runServer() {
	// Create a server, listen on address, emit message every 2s to connected clients
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening: ", err)
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + addr)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	symbols := []string{"dell", "ibm", "google", "walmart"}
	exchanges := []string{"nyse", "nasdaq", "bats"}

	for {
		quoteString := fmt.Sprintf("Q|%s|%s|%d|%d\n",
			symbols[rand.Intn(4)], // num symbols
			exchanges[rand.Intn(3)], // num exchanges
			rand.Intn(100),
			rand.Intn(100))

		fmt.Printf(">> Server generated quote: %v", quoteString)

		if _, err := conn.Write([]byte(quoteString)); err != nil {
			fmt.Printf("unable to write to client, dropping connection: %v\n", err)
			break
		}

		time.Sleep(2 * time.Second)
	}
}

func runParser() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("unable to connect to addr: %v", err)
		os.Exit(1)
	}

	defer conn.Close()

	timeoutDuration := 5 * time.Second
	bufReader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Unable to read from socket: %v\n", err)
			os.Exit(1)
		}

		go parse(string(bytes))
	}
}

func parse(line string) {
	// parse the line
	line = strings.TrimSuffix(line, "\n")
	s := strings.Split(line, "|")

	if len(s) != 5 {
		fmt.Printf("Unexpected number of elements after split: %d -- %v\n", len(s), s)
		return
	}

	if s[0] != "Q" {
		fmt.Printf("Non-quote - skipping: %v\n", s[0])
		return
	}

	// now we have all the data
	symbol := s[1]
	exchange := s[2]
	bid, err := strconv.Atoi(s[3])
	if err != nil {
		fmt.Printf("unable to parse bid string to i: %v\n", err)
		return
	}

	offer, err := strconv.Atoi(s[4])
	if err != nil {
		fmt.Printf("unable to parse offer string to i: %v\n", err)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	dingding := "DING DING! Symbol: %s || Best bid: %d on '%s' || Best offer: %d on '%s'\n"

	// first time we see it?
	if _, ok := quotes[symbol]; !ok {
		fmt.Printf(dingding, symbol, bid, exchange, offer, exchange)

		quotes[symbol] = &quote{
			BestBid: &num{
				Value: bid,
				Exchange: exchange,
			},
			BestOffer: &num{
				Value: offer,
				Exchange: exchange,
			},
		}
		return
	}

	updated := false

	// not first time we see it
	if bid < quotes[symbol].BestBid.Value {
		quotes[symbol].BestBid.Value = bid
		quotes[symbol].BestBid.Exchange = exchange
		updated = true
	}

	if offer > quotes[symbol].BestOffer.Value {
		quotes[symbol].BestOffer.Value = offer
		quotes[symbol].BestOffer.Exchange = exchange
		updated = true
	}

	if updated {
		fmt.Printf(dingding,
			symbol, quotes[symbol].BestBid.Value, quotes[symbol].BestBid.Exchange,
			quotes[symbol].BestOffer.Value, quotes[symbol].BestOffer.Exchange)
	}
}
