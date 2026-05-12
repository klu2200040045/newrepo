package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Task struct {
    ID   int
    Name string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("Worker %d started task %d (%s)\n", id, task.ID, task.Name)
        time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
        fmt.Printf("Worker %d finished task %d\n", id, task.ID)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    tasks := make(chan Task)
    var wg sync.WaitGroup

    // Start workers
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, tasks, &wg)
    }

    // Send tasks
    for i := 1; i <= 10; i++ {
        tasks <- Task{
            ID:   i,
            Name: fmt.Sprintf("Task-%d", i),
        }
    }

    close(tasks)
    wg.Wait()

    fmt.Println("All tasks completed!")
}