package rhythmtool_test

import (
	"fmt"
	"math/rand"

	"github.com/peter-mueller/rhythmtool"
)

func ExampleNote_String() {
	fmt.Printf("HIT: %v\n", rhythmtool.HIT)
	fmt.Printf("PAUSE: %v", rhythmtool.PAUSE)
	// Output:
	// HIT: x
	// PAUSE: -
}

func ExampleRhythm() {
	r := rhythmtool.Rhythm{rhythmtool.HIT, rhythmtool.PAUSE, rhythmtool.PAUSE}
	fmt.Printf("Rhythm: %v", r)
	// Output:
	// Rhythm: [x - -]
}

func ExampleRandom() {
	rand.Seed(42)
	fmt.Println(rhythmtool.Random(4))
	rand.Seed(24)
	fmt.Println(rhythmtool.Random(4))
	// Output:
	// [- - x x]
	// [- x - x]
}

func ExampleRhythm_MergeWith() {
	// Get two different initial rhythms.
	people := rhythmtool.UseString("People")
	car := rhythmtool.UseString("Car")

	// Create the merge.
	merge := people.MergeWith(car)

	fmt.Printf("People: %v\n", people)
	fmt.Printf("Car:    %v\n", car)
	fmt.Println()
	fmt.Printf("Merged: %v", merge)
	// Output:
	// People: [- x x - - x]
	// Car:    [x x -]
	//
	// Merged: [x x x - - x]
}

func ExampleGenerateBjorklund() {
	pulses := 3
	length := 8
	tresillo := rhythmtool.GenerateBjorklund(pulses, length)
	fmt.Printf("Tresillo: %v", tresillo)
	// Output: Tresillo: [x - - x - - x -]
}
