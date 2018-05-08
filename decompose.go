package decompose

import (
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/pln"
	"github.com/intdxdt/stack"
	"github.com/intdxdt/geom"
)

type scoreRelationFn func(float64) bool

//Douglas-Peucker decomposition at a given threshold
func DouglasPeucker(
	pln *pln.Polyline,
	scoreFn lnr.ScoreFn,
	scoreRelation scoreRelationFn,
	gfn geom.GeometryFn,
) []*node.Node {
	var k int
	var val float64
	var coordinates []*geom.Point
	var hque []*node.Node

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
			hque = append(hque, node.New(coordinates, rg, gfn))
		} else {
			s.Push(
				rng.NewRange(k, rg.J), // right
				rng.NewRange(rg.I, k), // left
			)
		}
	}
	return hque
}
