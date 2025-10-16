package ast

import "fmt"

type Env map[string]int

type Node interface {
	Execute(Env)
}

func ExecuteNodes(nodes []Node, env Env) {
	for _, node := range nodes {
		node.Execute(env)
	}
}

// nodes
type (
	Incr struct {
		Var string
	}

	Decr struct {
		Var string
	}

	Clear struct {
		Var string
	}

	While struct {
		Var  string
		Not  int
		Body []Node
	}
)

func (i *Incr) Execute(env Env) {
	if _, ok := env[i.Var]; !ok {
		env[i.Var] = 0
	}

	env[i.Var]++

	fmt.Printf("INCR %s -> %s", i.Var, state(env))
}

func (d *Decr) Execute(env Env) {
	if _, ok := env[d.Var]; !ok {
		env[d.Var] = 0
	}

	env[d.Var]--
	
	fmt.Printf("DECR %s -> %s", d.Var, state(env))
}

func (c *Clear) Execute(env Env) {
	env[c.Var] = 0
	fmt.Printf("CLEAR %s -> %s", c.Var, state(env))
}

func (w *While) Execute(env Env) {
	if _, ok := env[w.Var]; !ok {
		env[w.Var] = 0
	}

	for env[w.Var] != w.Not {
		ExecuteNodes(w.Body, env)
	}

	fmt.Printf("WHILE %s NOT %d -> %s", w.Var, w.Not, state(env))
}

// get a string representing global variable state
func state(env Env) string {
	s := ""

	for name, value := range env {
		s += fmt.Sprintf("%s=%d ", name, value)
	}

	s = s[:len(s) - 1]
	s += "\n"

	return s
}
