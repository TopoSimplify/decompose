package decompose

import (
	"simplex/rng"
	"simplex/node"
	"simplex/lnr"
	"github.com/intdxdt/deque"
	"github.com/intdxdt/stack"
	"github.com/intdxdt/geom"
	"simplex/pln"
)

type scoreRelationFn func(float64) bool

//Douglas Peucker decomposition at a given threshold
func DouglasPeucker(
	pln *pln.Polyline,
	scoreFn lnr.ScoreFn,
	scoreRelation scoreRelationFn,
	gfn geom.GeometryFn,
) *deque.Deque {
	var k int
	var val float64
	var coordinates []*geom.Point
	var hque = deque.NewDeque()

	if pln == nil {
		return hque
	}

	var rg = pln.Range()
	var s = stack.NewStack().Push(rg)

	for !s.IsEmpty() {
		rg = s.Pop().(*rng.Range)
		coordinates = pln.SubCoordinates(rg)
		k, val = scoreFn(coordinates)
		k = rg.Index(k)

		if scoreRelation(val) {
			hque.Append(node.New(coordinates, rg, gfn))
		} else {
			s.Push(
				rng.NewRange(k, rg.J()), // right
				rng.NewRange(rg.I(), k), // left
			)
		}
	}
	return hque
}
