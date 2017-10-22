package decompose

import (
	"simplex/rng"
	"simplex/node"
	"simplex/lnr"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/stack"
	"github.com/intdxdt/geom"
)

type scoreRelationFn func(float64) bool

//Douglas Peucker decomposition at a given threshold
func DouglasPeucker(self lnr.Linear, scoreRelation scoreRelationFn, gfn geom.GeometryFn) *deque.Deque {
	var pln = self.Polyline()
	var hque = deque.NewDeque()
	var k int
	var val float64
	var rg = pln.Range()

	s := stack.NewStack().Push(rg)

	for !s.IsEmpty() {
		rg = s.Pop().(*rng.Range)
		k, val = self.Score(self, rg)
		if scoreRelation(val) {
			hque.Append(node.New(pln, rg, gfn))
		} else {
			s.Push(
				rng.NewRange(k, rg.J()), // right
				rng.NewRange(rg.I(), k), // left
			)
		}
	}
	return hque
}
