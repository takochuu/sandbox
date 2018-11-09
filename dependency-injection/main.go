package main

func main() {
	sample := factory()
	NewExecutor(sample)
}

type Executor struct {
	sample I
}

func NewExecutor(s I) *Executor {
	return &Executor{
		sample: s,
	}
}

func (e *Executor) Create() error {
	return e.sample.Create()
}

func factory() I {
	if true {
		return NewSample1()
	}
	return NewSample2()
}

type I interface {
	Create() error
}

type Sample1 struct {
}

func NewSample1() I {
	return nil
}

func (s *Sample1) Create() error {
	return nil
}

type Sample2 struct {
}

func NewSample2() I {
	return nil
}

func (s *Sample2) Create() error {
	return nil
}
