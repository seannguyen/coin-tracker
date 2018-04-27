package main

func main()  {
	var bar a
	bar = &b{}
}

type a interface {
	foo() int
}

type b struct{}

func (*b) foo() int {
	return 1
}