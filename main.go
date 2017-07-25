package three

import (
	"fmt"

	"github.com/bep/gr/support"
	"github.com/gopherjs/gopherjs/js"
)

var (
	three = js.Global.Get("THREE")
)

func init() {
	if three == js.Undefined {
		// Require as a fallback
		if _, err := support.Require("three"); err != nil {
			panic(fmt.Sprintf("Couldn't find three.js"))
		}
	}
}
