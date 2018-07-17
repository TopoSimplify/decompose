package decompose

import (
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/pln"
	"github.com/intdxdt/geom"
)

type scoreRelationFn func(float64) bool

//Douglas-Peucker decomposition at a given threshold
func DouglasPeucker(
	pln *pln.Polyline, scoreFn lnr.ScoreFn,
	scoreRelation scoreRelationFn, geomFn geom.GeometryFn,
) []*node.Node {
	var k, n int
	var val float64
	var coordinates []geom.Point
	var hque []*node.Node

	if pln == nil {
		return hque
	}

	var r = pln.Range()
	var stack = make([]rng.Rng, 0, (r.J-r.I)+1)
	stack = append(stack, r)

	for !(len(stack) == 0) {
		n = len(stack) - 1
		r = stack[n]
		stack = stack[:n]

		coordinates = pln.SubCoordinates(r)
		k, val = scoreFn(coordinates)
		k = r.I + k //offset

		if scoreRelation(val) {
			hque = append(hque, node.New(coordinates, r, geomFn))
		} else {
			stack = append(stack, rng.Range(k, r.J)) // right
			stack = append(stack, rng.Range(r.I, k)) // left
		}
	}
	return hque
}
