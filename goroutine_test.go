package main

import "testing"

func BenchmarkContextSwitch(b *testing.B) {
	die := make(chan struct{})
	ch := make(chan struct{})
	v := struct{}{}
	go func() {
		for {
			select {
			case <-ch:
				ch <- v
			case <-die:
				return
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- v
		<-ch
	}
	close(die)
}
