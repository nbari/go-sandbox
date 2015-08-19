package main

import (
	"fmt"
	"github.com/nbari/violetear"
)

func main() {
	d := violetear.NewDynamic()
	d.Set(":ip", `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
	d.Set(":uuid", "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

	uuid := "2E9C64A5-FF13-4DC5-A957-F39E39ABDC48"
	for k, v := range d {
		if v.MatchString(uuid) {
			fmt.Printf("Match ---> %v", k)
		}
		fmt.Println(k)

	}

}
