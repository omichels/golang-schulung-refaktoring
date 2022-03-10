package numbergen

import "sync"

var c chan string

var once sync.Once

type Generator struct {
	mu sync.Mutex
	n  int
}

var instance *Generator

func GeneratorInstance() (g *Generator) {
	once.Do(func() {
		instance = &Generator{
			n: 0,
		}
	})
	return instance
}

func (g *Generator) GetNumber() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	instance.n = instance.n + 1
	return instance.n
}
