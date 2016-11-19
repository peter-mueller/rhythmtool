package rhythmtool_test

import (
	"fmt"

	"github.com/peter-mueller/rhythmtool"
)

func Example_bjorklund() {
	cinquillo := rhythmtool.GenerateBjorklund(5, 8)
	fmt.Print("Cuban cinquillo: ")
	fmt.Println(cinquillo)

	tresillo := rhythmtool.GenerateBjorklund(3, 8)
	tresillo = tresillo.Rotate(-1)
	fmt.Print("Cuban tresillo, rotated one to the left: ")
	fmt.Println(tresillo)

	fmt.Print("Merge of both:")
	merged := cinquillo.MergeWith(tresillo)
	fmt.Println(merged)

	fmt.Print("Reverse of the Merge: ")
	reversed := merged.Reverse()
	fmt.Println(reversed)

	// Output:
	// Cuban cinquillo: [x - x x - x x -]
	// Cuban tresillo, rotated one to the left: [- - x - - x - x]
	// Merge of both:[x - x x - x x x]
	// Reverse of the Merge: [x x x - x x - x]
}
