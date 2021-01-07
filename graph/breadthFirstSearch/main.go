package main

import (
	"container/list"
	"fmt"
)

func main() {
	graph := generateGraph()
	fmt.Println(graph)

	res := bfs(graph, "you")
	fmt.Println(res)
}

type void struct{}

type person struct {
	name     string
	isSeller bool
}

// Breadth-First-Search (BFS) implementation
func bfs(graph map[string][]string, name string) bool {
	queue := list.New() // Linked List implementation which will be used as Queue
	addElementsToQueue(queue, graph[name])

	searched := make(map[string]void)

	for queue.Len() > 0 {
		person := popLeft(queue)
		if _, ok := searched[person]; !ok {
			if check(person) {
				fmt.Printf("%v is seller\n", person)
				return true
			} else {
				addElementsToQueue(queue, graph[person])
				searched[person] = void{}
			}
		}
	}
	return false
}

func check(element string) bool {
	return element == "anuj"
}

func popLeft(queue *list.List) string {
	e := queue.Front()
	queue.Remove(e)
	return fmt.Sprintf("%v", e.Value)
}

func addElementsToQueue(queue *list.List, elems []string) {
	for _, v := range elems {
		queue.PushBack(v)
	}
}

// TODO: replace with person struct instead of dummy string
func generateGraph() map[string][]string {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	return graph
}
