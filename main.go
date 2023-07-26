package main

import (
	"log"
	"math/rand"
	"time"
)

// ping-pong sederhana
// buat sebuah aplikasi sederhana dengan kriteria :
// 2 pemain ping-pong saling mengembalikan bola
// permainan berhenti setelah 1 detik
// gunakan konkurensi dengan 1 chanel

func main() {
	table := make(chan *ball)
	done := make(chan *ball)

	go player("marwan", table, done)
	go player("elsa", table, done)
	// go player("sean", table, done)
	// go player("danish", table, done)

	referree(table, done)
}

type ball struct {
	hits       int
	lastPLayer string
}

func referree(table chan *ball, done chan *ball) {
	table <- new(ball)
	for {
		select {
		case ball := <-done:
			log.Println("winner is", ball.lastPLayer)
			return
		}
	}
}

func player(name string, table chan *ball, done chan *ball) {
	for {

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		select {
		case ball := <-table:
			v := r.Intn(1000)
			log.Println(name, "drop the ball")
			if v%11 == 0 {
				done <- ball
				return
			}
			ball.hits++
			ball.lastPLayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Microsecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return
		}

	}
}
