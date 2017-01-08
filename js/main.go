package main

import "github.com/gopherjs/gopherjs/js"
import "github.com/peter-mueller/rhythmtool"

func main() {
	js.Global.Set("rhythmtool", map[string]interface{}{
		"random":    rhythmtool.Random,
		"bjorklund": rhythmtool.GenerateBjorklund,
		"useString": rhythmtool.UseString,
	})
}
