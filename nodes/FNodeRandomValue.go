package nodes

import (
	"fnode2/core"
	"math/rand"
	"time"
)

var randomValueOutput = core.NewNodeOutput(core.FTypeFloat, "Result",
	func(node *core.Node) any {
		min := node.GetInputFloat(0)
		max := node.GetInputFloat(1)
		//TODO bias := inputs[2].(float64)
		r := rand.New(rand.NewSource(time.Now().Unix() + int64(core.RunState.CurrentIteration)))
		random0To1 := r.Float64()
		return min + random0To1*(max-min)
	},
	false)

func newRandomValueNode() *core.Node {
	return core.NewNodeCreator(
		"RandomValue",
		"Math",
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
