package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// channel <- "marwan belajar goroutine"
		fmt.Println("selesai kirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	fmt.Println("selesai")
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "marwan belajar golang"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	fmt.Println("done")
	time.Sleep(3 * time.Second)
}

func TestRanegChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("done function")
}
