package main

type Data struct {
	i int
}

func func1(c chan *Data) {
	for {
		var t *Data
		t = <-c   //receive
		t.i += 10 //increment
		c <- t    //send it back
	}
}

func main() {
	c := make(chan *Data)
	t := Data{10}
	go func1(c)
	println(t.i)
	c <- &t  //send a pointer to our t
	i := <-c //receive the result
	println(i.i)
	println(t.i)
}
