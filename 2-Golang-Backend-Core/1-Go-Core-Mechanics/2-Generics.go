package main

import (
	"fmt"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Queue[T any] struct {
	items[]T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue()(T, bool) {
	var zeroValue T
	if len(q.items) == 0 {
		return zeroValue, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func GetMapKeyValue[K comparable, V any] (m map[K]V) ([]K, []V) {
	keys, values := make([]K, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func Sum[T Number](number[]T) T {
	var total T
	for _, num := range number {
		total += num
	}
	return total
}

func main() {
	fmt.Println("--- 1. GENERIC FUNCTION ---")
	scores := map[string]int{"Alice": 90, "Bob": 85, "Charlie": 95}
	keys, values := GetMapKeyValue(scores)
	fmt.Printf("Keys of scores: %v\n", keys)
	fmt.Printf("Values of scores: %v\n", values)

	users := map[int]string{1: "Admin", 2: "Editor"}
	keys2, values2 := GetMapKeyValue(users)
	fmt.Printf("Keys of users: %v\n", keys2)
	fmt.Printf("Values of users: %v\n", values2)

	fmt.Println("\n--- 2. CUSTOM CONSTRAINTS ---")
	intSlice :=[]int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of ints: %d\n", Sum(intSlice))

	floatSlice :=[]float64{1.5, 2.5, 3.5}
	fmt.Printf("Sum of floats: %.2f\n", Sum(floatSlice))

	fmt.Println("\n--- 3. GENERIC STRUCT (QUEUE) ---")
	jobQueue := Queue[string]{}
	jobQueue.Enqueue("JOB-001")
	jobQueue.Enqueue("JOB-002")
	job1, ok := jobQueue.Dequeue()
	if ok { fmt.Printf("Dequeued Job: %s\n", job1) }

	numberQueue := Queue[int]{}
	numberQueue.Enqueue(100)
	num1, _ := numberQueue.Dequeue()
	fmt.Printf("Dequeued Number: %d\n", num1)
}