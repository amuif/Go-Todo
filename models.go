package main

type Option struct {
	ID    int
	Label string
}
type model struct {
	choices  []Option
	cursor   int
	selected int
}
