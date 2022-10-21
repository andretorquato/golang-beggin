package main

func main() {
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go worker(tasks, results)
	}

	for i := 0; i < 100; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i < 100; i++ {
		<-results
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-1) + fibonacci(position-2)
}
