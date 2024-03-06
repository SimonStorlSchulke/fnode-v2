package nodes

import (
	"fnode2/core"
	"math/rand"
	"time"
)

var randomValueOutput = core.NewNodeOutput(core.FTypeFloat, "Result",
	func(inputs []any, _ map[string]*core.NodeOption) any {
		min := inputs[0].(float64)
		max := inputs[1].(float64)
		//TODO bias := inputs[2].(float64)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		random0To1 := r.Float64()
		return min + random0To1*(max-min)
	},
	false)

func newRandomValueNode() *core.Node {
	return core.NewNodeCreator(
		"Random Value",
		[]core.NodeInput{
			core.NewNodeInput(core.FTypeFloat, "Min", 0.0),
			core.NewNodeInput(core.FTypeFloat, "Max", 1.0),
			core.NewNodeInput(core.FTypeFloat, "Bias", 1.0),
		},
		[]*core.NodeOutput{
			randomValueOutput,
		},
	)
}
