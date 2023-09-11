package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Service struct {
// 	blackList map[string]bool
// 	mu        sync.Mutex
// }

// type Message struct {
// 	Text     string
// 	Receiver string
// }

// func NewService(blackList []string) *Service {
// 	service := &Service{
// 		blackList: make(map[string]bool),
// 	}

// 	for _, name := range blackList {
// 		service.blackList[name] = true
// 	}

// 	return service
// }

// func (s *Service) SendMessage(message Message, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	if !s.blackList[message.Receiver] {
// 		fmt.Printf("Message to %s sent: %s\n", message.Receiver, message.Text)
// 	} else {
// 		fmt.Printf("Message to %s ignored: %s\n", message.Receiver, message.Text)
// 	}
// }

// func main() {
// 	blackList := []string{"John"}
// 	service := NewService(blackList)

// 	messages := []Message{
// 		{Text: "hello", Receiver: "John"},
// 		{Text: "hi", Receiver: "Adam"},
// 	}

// 	var wg sync.WaitGroup
// 	for _, message := range messages {
// 		wg.Add(1)
// 		go service.SendMessage(message, &wg)
// 	}

// 	wg.Wait()
// }
