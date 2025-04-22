package main

import (
	"fmt"
	"time"
	"slices"
	"strings"
)

type Task struct {
	deadline time.Time
	done bool
}

type Tasks map[string]Task

func (tasks Tasks) print_all_tasks() {
	keys := make([]string, 0, len(tasks))

	for k := range tasks {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(i, j string) int {
		return strings.Compare(i, j)
	})

	for _, key := range keys {
		fmt.Printf("Taske name: %s, Deadline : %s, Finished: %t\n", key, tasks[key].deadline.Format(time.DateTime), tasks[key].done)
	}
}

func (tasks Tasks) create_fake_tasks() {
	for i := 0; i < 10; i++  {
		tasks[fmt.Sprintf("%.2d", i)] = Task{deadline : time.Now(), done : i%2 == 0}
	}
}

func main() {
	tasks := make(Tasks)
	tasks.create_fake_tasks()
	tasks.print_all_tasks()
}
