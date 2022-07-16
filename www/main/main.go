package main

import "github.com/broothie/jog"

func main() {
	jog.Mount("root", nil)

	jog.Wait()
}
