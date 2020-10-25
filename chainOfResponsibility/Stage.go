package main

type stage interface {
	execute(*player)
	setNext(stage)
}
