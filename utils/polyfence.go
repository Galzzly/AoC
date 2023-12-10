package utils

import "image"

type Polygon []image.Point

type Polyfence struct {
	p Polygon
}

func NewPolyfence(p Polygon) *Polyfence {
	polyfence := &Polyfence{
		p: p.Copy(),
	}

	if len(polyfence.p) > 0 &&
		polyfence.p[0] != polyfence.p[len(polyfence.p)-1] {
		polyfence.p = append(polyfence.p, polyfence.p[0])
	}

	return polyfence
}

func (p Polygon) Copy() Polygon {
	newPolygon := make(Polygon, len(p))
	copy(newPolygon, p)
	return newPolygon
}

func (pf *Polyfence) Inside(P image.Point) bool {
	if checkOutside(pf.p, P) == 0 {
		return false
	}
	return true
}

func checkOutside(p Polygon, P image.Point) (res int) {
	if len(p) < 3 {
		return 0
	}
	edges := len(p) - 2
	for i := 0; i <= edges; i++ {
		if p[i].Y <= P.Y {
			if p[i+1].Y > P.Y {
				if isLeft(p[i], p[i+1], P) > 0 {
					res++
				}
			}
		} else {
			if p[i+1].Y <= P.Y {
				if isLeft(p[i], p[i+1], P) < 0 {
					res--
				}
			}
		}
	}
	return
}

func isLeft(A, B, P image.Point) int {
	return (B.X-A.X)*(P.Y-A.Y) - (P.X-A.X)*(B.Y-A.Y)
}
