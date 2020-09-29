package comment

// package main // want "commented code"
/*
// this is a valid go doc comment.
func main() {
	if 1 == 1 {
		_ = 2 + 2
	}
}
*/

// package main // want "commented code"
/*
	func main() {
		https://github.com/census-instrumentation/opencensus-go/blob/d7677d6af5953e0506ac4c08f349c62b917a443a/stats/view/worker.go#L34
	}
*/

func someOtherFunc() {
	// _ = 2 + 2 // want "commented code"
	_ = 2 + 2
}

// 1 // want "commented code"
// 2

// func foo(){} // want "commented code"

// // this double comment doesn't make any sense // want "empty double comment"

// this = that // want "commented code"
