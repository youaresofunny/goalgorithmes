package main

// Weighted quick-union algorithm with pass reduction
// M union-find operations on set of N array, N + M*lgN

import (
	"fmt"
	math "math/rand"
)

type Model struct {
	data []int
	size []int
}

func main() {
	m := &Model{}
	m.init(50)
	m.randomConnections(100)
	m.out()
}

//Init model data and sizes arrays
func (m *Model) init(size int) {
	m.data = make([]int, size)
	m.size = make([]int, size)
	for i := 0; i < size; i++ {
		m.data[i] = i
		m.size[i] = 1
	}
	m.out()
}

//Searching for root of the element and reducing the tree
func (m *Model) root(elem int) int {
	i := elem
	for m.data[i] != i {
		/m.data[i] = m.data[m.data[i]]
		i = m.data[i]
	}
	return i
}

//Union two elements, this is reduced version because of compare
func (m *Model) union(a int, b int) {
	aRoot := m.root(a)
	bRoot := m.root(b)
	if aRoot == bRoot {
		return
	}
	if m.size[aRoot] > m.size[bRoot] {
		m.data[aRoot] = bRoot
		m.size[bRoot] += m.size[aRoot]
	} else {
		m.data[bRoot] = aRoot
		m.size[aRoot] += m.size[bRoot]
	}
}

func (m *Model) connected(a int, b int) bool {
	return m.root(a) == m.root(b)
}

func (m *Model) randomConnections(count int) {
	for i := 0; i < count; i++ {
		m.union(math.Intn(len(m.data)), math.Intn(len(m.data)))
	}
}

func (m *Model) out() {
	fmt.Printf("data: %v\n", m.data)
	fmt.Printf("sizes: %v\n", m.size)
}
