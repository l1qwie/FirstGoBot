package async

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Имитация длительной работы
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
	fmt.Println("SHALOM")
}

func Main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Запускаем несколько горутин-работников
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Передаем задачи в канал jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs) // Закрываем канал jobs, чтобы горутины завершили выполнение

	// Получаем результаты из канала results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}
