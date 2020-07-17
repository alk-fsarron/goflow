// Code generated by lib-go/goflow DO NOT EDIT.

// +build !codeanalysis

package main

import (
	"github.com/alkemics/goflow/example/nodes"
)

/*

 */
type Ouputs struct{}

func NewOuputs() Ouputs {
	return Ouputs{}
}

func newOuputs(id string) Ouputs {
	return Ouputs{}
}

/*

 */
func (g *Ouputs) Run() (random_int int, random_ints []int) {
	// __output_random_int_builder outputs
	var __output_random_int_builder_random_int int

	// __output_random_ints_builder outputs
	var __output_random_ints_builder_random_ints []int

	// produce_another_random_int outputs
	var produce_another_random_int_n int

	// produce_random_int outputs
	var produce_random_int_n int

	igniteNodeID := "ignite"
	doneNodeID := "done"

	done := make(chan string)
	defer close(done)

	steps := map[string]struct {
		deps        map[string]struct{}
		run         func()
		alreadyDone bool
	}{

		"__output_random_int_builder": {
			deps: map[string]struct{}{
				"produce_random_int": {},
				igniteNodeID:         {},
			},
			run: func() {
				__output_random_int_builder_random_int = produce_random_int_n
				random_int = __output_random_int_builder_random_int
				done <- "__output_random_int_builder"
			},
			alreadyDone: false,
		},
		"__output_random_ints_builder": {
			deps: map[string]struct{}{
				"produce_random_int":         {},
				"produce_another_random_int": {},
				igniteNodeID:                 {},
			},
			run: func() {
				__output_random_ints_builder_random_ints = append(__output_random_ints_builder_random_ints, produce_random_int_n)
				__output_random_ints_builder_random_ints = append(__output_random_ints_builder_random_ints, produce_another_random_int_n)
				random_ints = __output_random_ints_builder_random_ints
				done <- "__output_random_ints_builder"
			},
			alreadyDone: false,
		},
		"produce_another_random_int": {
			deps: map[string]struct{}{
				igniteNodeID: {},
			},
			run: func() {
				produce_another_random_int_n = nodes.RandomIntProducer()
				done <- "produce_another_random_int"
			},
			alreadyDone: false,
		},
		"produce_random_int": {
			deps: map[string]struct{}{
				igniteNodeID: {},
			},
			run: func() {
				produce_random_int_n = nodes.RandomIntProducer()
				done <- "produce_random_int"
			},
			alreadyDone: false,
		},
		igniteNodeID: {
			deps: map[string]struct{}{},
			run: func() {
				done <- igniteNodeID
			},
			alreadyDone: false,
		},
		doneNodeID: {
			deps: map[string]struct{}{
				"__output_random_int_builder":  {},
				"__output_random_ints_builder": {},
				"produce_another_random_int":   {},
				"produce_random_int":           {},
			},
			run: func() {
				done <- doneNodeID
			},
			alreadyDone: false,
		},
	}

	// Ignite
	ignite := steps[igniteNodeID]
	ignite.alreadyDone = true
	steps[igniteNodeID] = ignite
	go steps[igniteNodeID].run()

	// Resolve the graph
	for resolved := range done {
		if resolved == doneNodeID {
			// If all the graph was resolved, get out of the loop
			break
		}

		for name, step := range steps {
			delete(step.deps, resolved)
			if len(step.deps) == 0 && !step.alreadyDone {
				step.alreadyDone = true
				steps[name] = step
				go step.run()
			}
		}
	}

	return random_int, random_ints
}
