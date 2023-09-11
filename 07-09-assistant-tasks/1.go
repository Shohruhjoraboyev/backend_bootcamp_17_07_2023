package main

// func main() {
// 	ch := make(chan int)

// 	var wg sync.WaitGroup

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			ch <- i
// 		}(i)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for num := range ch {
// 		fmt.Println(num)
// 	}
// }
