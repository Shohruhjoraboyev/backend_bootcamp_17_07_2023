package main

import "fmt"

var blacklist = []string{"John", "Adam"}

type Message struct {
	Text     string
	Receiver string
}

func main() {
	ch := make(chan string)
	go check(ch)
	message1 := Message{Text: "Hello", Receiver: "John"}
	checkMessage(ch, message1)

	message2 := Message{Text: "Hi", Receiver: "Adam"}
	checkMessage(ch, message2)

	message3 := Message{Text: "Hi", Receiver: "Omadbek"}
	checkMessage(ch, message3)

	close(ch)
}

func check(ch chan string) {
	for checked := range ch {
		fmt.Println(checked)
	}
}

func checkMessage(ch chan string, message Message) {
	for _, name := range blacklist {
		if message.Receiver == name {
			ch <- fmt.Sprintf("Message ignored to send %s", message.Receiver)
			return
		}
	}
	ch <- fmt.Sprintf("Message Send To %s", message.Receiver)
}
