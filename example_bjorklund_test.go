package rhythmtool_test

import (
	"fmt"

	"github.com/peter-mueller/rhythmtool"
)

// This example shows how to use the rhythmtool package to create two initial
// rhythms using the bjorklund algorithm. They are printed to the console to
// see how they look like.
//
// Then these rhythms are further manipulated. At first they are merged (overlayed) to
// a single rhythm which is then reversed.
func Example_bjorklund() {
	// Use the Bjorklund algorithm to create a rhythm with
	// length 8 containing exactly 5 beats.
	cinquillo := rhythmtool.GenerateBjorklund(5, 8)
	fmt.Print("Cuban cinquillo: ")
	fmt.Println(cinquillo)

	// The same but with only 3 beats in the rhythm.
	tresillo := rhythmtool.GenerateBjorklund(3, 8)
	tresillo = tresillo.Rotate(-1)
	fmt.Print("Cuban tresillo, rotated one to the left: ")
	fmt.Println(tresillo)

	fmt.Print("Merge of both: ")
	merged := cinquillo.MergeWith(tresillo)
	fmt.Println(merged)

	fmt.Print("Reverse of the Merge: ")
	reversed := merged.Reverse()
	fmt.Println(reversed)

	// Output:
	// Cuban cinquillo: [x - x x - x x -]
	// Cuban tresillo, rotated one to the left: [- - x - - x - x]
	// Merge of both: [x - x x - x x x]
	// Reverse of the Merge: [x x x - x x - x]
}
