package main

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewModel(arg string) Car {
	return &Lambo(arg)
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop() {

}
