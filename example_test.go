package errorv_test

import (
	"fmt"

	"github.com/jjeffery/errorv"
)

func Example() {
	err := errorv.New("first error",
		errorv.KV("card", "ace"),
		errorv.KV("suite", "spades"))
	fmt.Println(err)

	err = errorv.Wrap(err, "second error",
		errorv.KV("piece", "rook"),
		errorv.KV("color", "black"),
		errorv.Caller(0))
	fmt.Println(err)

	// Output:
	// first error card=ace suite=spades
	// second error piece=rook color=black github.com/jjeffery/errorv/example_test.go:18: first error card=ace suite=spades
}

var userID, documentID string

func ExampleContext() error {
	// ... if a function has been called with userID and DocumentID ...
	errorv := errorv.NewContext(errorv.KV("userID", userID), errorv.KV("documentID", documentID))

	n, err := doOneThing()
	if err != nil {
		// will include key value pairs for userID and document ID
		return errorv.Wrap(err, "cannot do one thing")
	}

	if err := doAnotherThing(n); err != nil {
		// will include key value pairs for userID, document ID and n
		return errorv.Wrap(err, "cannot do another thing", errorv.KV("n", n))
	}

	return nil
}

func doOneThing() (int, error) {
	return 0, nil
}

func doAnotherThing(n int) error {
	return nil
}
