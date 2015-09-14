// The 'recovery' part of a panic has to happen on the same goroutine as the
// panic itself, otherwise it bubbles to the top without being captured. so it
// needed to be "attached" to the same thing that called foo, which was your
// anonymous func the recovery itself is a deferred anonymous func and if it
// feels difficult to use, then that's excellent. you shouldn't be using this,
// and should avoid code that does 'neat' things in this direction

package main

import "fmt"
import "time"

func foo() {
	panic("no recovery")
}

func main() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered")
			}
		}()
		foo()
	}()
	<-time.After(10 * time.Second)

}
