package main

import "container/list"

// Promise struct is an implementation of
// Javascript's promises for flow control
type Promise struct {
	chain     list.List
	catchFunc catchFunc
}

type thenFunc func() error
type catchFunc func(error)

// Then stablishes a new step
func (p *Promise) Then(fn thenFunc) *Promise {
	p.chain.PushBack(fn)
	return p
}

// Catch stablishes what to do to handle
func (p *Promise) Catch(fn catchFunc) *Promise {
	p.catchFunc = fn
	return p
}

// Run the chain
func (p *Promise) Run() {
	var foundError error
	for e := p.chain.Front(); e != nil && foundError == nil; e = e.Next() {
		foundError = e.Value.(thenFunc)()
	}
	if foundError != nil {
		p.catchFunc(foundError)
	}
}
