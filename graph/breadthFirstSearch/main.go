package main

import (
	"container/list"
	"fmt"
)

func main() {
	graph, start := generateGraph()
	fmt.Println(graph)

	res := bfs(graph, start)
	fmt.Println(res)
}

type void struct{}

type person struct {
	name     string
	isSeller bool
}

// Breadth-First-Search (BFS) implementation
func bfs(graph map[person][]person, start person) bool {
	queue := Queue{
		list.New(),
	}
	queue.addElementsToQueue(graph[start])

	searched := make(map[person]void)

	for queue.inner.Len() > 0 {
		person := queue.popLeft()
		if _, ok := searched[person]; !ok {
			if person.isSeller {
				fmt.Printf("%v is seller\n", person.name)
				return true
			} else {
				queue.addElementsToQueue(graph[person])
				searched[person] = void{}
			}
		}
	}
	return false
}

// Queue is a wrapper for the linked list
type Queue struct {
	inner *list.List
}

func (q Queue) popLeft() person {
	e := q.inner.Front()
	q.inner.Remove(e)
	return e.Value.(person)
}

func (q Queue) addElementsToQueue(elems []person) {
	for _, v := range elems {
		q.inner.PushBack(v)
	}
}

// return the generated graph and the root node
func generateGraph() (map[person][]person, person) {
	anuj := person{"anuj", true}
	peggy := person{"peggy", false}
	thom := person{"thom", false}
	jonny := person{"jonny", false}
	claire := person{"claire", false}
	alice := person{"alice", false}
	bob := person{"bob", false}
	you := person{"you", false}

	graph := make(map[person][]person)
	graph[you] = []person{alice, bob, claire}
	graph[bob] = []person{anuj, peggy}
	graph[alice] = []person{peggy}
	graph[claire] = []person{thom, jonny}
	graph[anuj] = []person{}
	graph[peggy] = []person{}
	graph[thom] = []person{}
	graph[jonny] = []person{}
	return graph, you
}
