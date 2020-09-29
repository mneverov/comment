package comment

// https://github.com/census-instrumentation/opencensus-go/blob/d7677d6af5953e0506ac4c08f349c62b917a443a/stats/view/worker.go#L34

/*
	This is invalid go code because of this line: it is not a valid comment (no //).
	Hence this comment will be treated as useful info.

	func main() {
		if 1 == 1 {
			_ = 2 + 2
		}
	}
*/

// return something

// someFunc does something
func someFunc() {
	// this is a comment inside the function
	_ = 2 + 2
}

// this is an ordinary comment

// 1

// ms

// label:

// empty comments are indistinguishable from processor directives

//

/*

 */

//nolint:unparam

// what // now
