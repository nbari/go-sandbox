package main

import "fmt"

type me struct{}

func (m *me) Do(s string) string {
	fmt.Println(s)
	out := "end: " + s
	return out
}

type Task interface {
	Do(string) string
}

type myTask func(string) string

func (f myTask) Do(s string) string {
	return f(s)
}

type Decorator func(Task) Task

func Decorate(t Task, ds ...Decorator) Task {
	decorated := t
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func First(d string) Decorator {
	return func(t Task) Task {
		return myTask(func(s string) string {
			fmt.Println(d, "first")
			return t.Do(s)
		})
	}
}

func Second(d string, i int) Decorator {
	return func(t Task) Task {
		return myTask(func(s string) string {
			fmt.Println(d, "an int: ", i)
			return t.Do(s)
		})
	}
}

func main() {
	m := &me{}
	s := "my string"
	//	fmt.Println(m.Do(s))
	//	x := First("test decorator")
	//	decorated := x(m)
	//	decorated.Do(s)
	//	fmt.Printf("decorated.Do(s) = %+v\n", decorated.Do(s))

	t := Decorate(m, First("test decorator"), Second("second", 2))
	t.Do(s)
}
