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
	queue := Queue{
		list.New(),
	}
	queue.addElementsToQueue(graph[name])

	searched := make(map[string]void)

	for queue.inner.Len() > 0 {
		person := queue.popLeft()
		if _, ok := searched[person]; !ok {
			if check(person) {
				fmt.Printf("%v is seller\n", person)
				return true
			} else {
				queue.addElementsToQueue(graph[person])
				searched[person] = void{}
			}
		}
	}
	return false
}

func check(element string) bool {
	return element == "anuj"
}

// Queue is a wrapper for the linked list
type Queue struct {
	inner *list.List
}

func (q Queue) popLeft() string {
	e := q.inner.Front()
	q.inner.Remove(e)
	return fmt.Sprintf("%v", e.Value)
}

func (q Queue) addElementsToQueue(elems []string) {
	for _, v := range elems {
		q.inner.PushBack(v)
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
