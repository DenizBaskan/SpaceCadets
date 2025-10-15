package ast

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
		Var, Not Value
		Body     []Node
	}
)

func (i *Incr) Execute(env Env) {
	if _, ok := env[i.Var]; !ok {
		env[i.Var] = 0
	}

	env[i.Var]++
}

func (d *Decr) Execute(env Env) {
	if _, ok := env[d.Var]; !ok {
		env[d.Var] = 0
	}

	env[d.Var]--
}

func (c *Clear) Execute(env Env) {
	env[c.Var] = 0
}

func (w *While) Execute(env Env) {
	for w.Var.ToInt(env) != w.Not.ToInt(env) {
		ExecuteNodes(w.Body, env)
	}
}

// helpers
type Value interface {
	ToInt(Env) int
}

type IntegerLiteral struct {
	Num int
}

func (i *IntegerLiteral) ToInt(_ Env) int {
	return i.Num
}

type Identifer struct {
	Name string
}

func (i *Identifer) ToInt(env Env) int {
	value, ok := env[i.Name]
	if !ok {
		value = 0
	}

	return value
}
