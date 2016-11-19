package rhythmtool_test

import (
	"fmt"

	"github.com/peter-mueller/rhythmtool"
)

func Example_name() {
	fmt.Print(`"Auto": `)
	r := rhythmtool.UseString("Auto")
	fmt.Println(r)

	// Output:
	// "Auto": [x x - x]
}
