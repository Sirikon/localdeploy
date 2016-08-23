package main

import "container/list"

type Promise struct{
	chain list.List
	catchFunc CatchFunc
}

type ThenFunc func() error
type CatchFunc func(error)

func (p *Promise) Then(fn ThenFunc) (*Promise) {
	p.chain.PushBack(fn)
	return p
}

func (p *Promise) Catch(fn CatchFunc) (*Promise) {
	p.catchFunc = fn
	return p
}

func (p *Promise) Run() {
	var foundError error
	for e := p.chain.Front(); (e != nil && foundError == nil); e = e.Next() {
		foundError = e.Value.(ThenFunc)()
	}
	if foundError != nil {
		p.catchFunc(foundError)
	}
}