package main

import (
	//"fmt"
	"socketyi"
)

func main() {
	so := socketyi.NewSocketYi(nil)
	so.ListenAndServe(":8000")
}

