package dbs

import "math/cmplx"

func signum(x float64) int8 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return +1
}

func signumPair(c complex128) (pair [2]int8) {
	pair[0] = signum(real(c))
	pair[1] = signum(imag(c))
	return
}

// Bounds return bounding rectangle of a Point
func (me *Point) Bounds() Rect {
	return Rect{Min: *me, Max: *me}
}

// Bounds return bounding rectangle of a Span
func (me *Span) Bounds() Rect {
	result := me.A.Bounds()
	result.AddPoint(&me.Z)

	var outs int8
	straight := me.Vector().C()
	curve := complex(1, -me.Bulge)
	curve *= curve
	for i := 0; i < 2; i++ {
		tangent := straight * curve
		t2 := signumPair(tangent)
		s2 := signumPair(straight)
		for j, sign := range t2 {
			if sign != 0 && sign != s2[j] {
				outs |= 1 << (uint(j) + uint(sign+1))
			}
		}
		if 0 == i {
			// Other end of Span
			straight = -straight
			curve = cmplx.Conj(curve)
		}
	}
	if outs != 0 {
		Radius := me.Radius()
		Center := me.Center()

		if outs&1 != 0 {
			result.Min.X = Center.X - Radius
		}
		if outs&2 != 0 {
			result.Min.Y = Center.Y - Radius
		}
		if outs&4 != 0 {
			result.Max.X = Center.X + Radius
		}
		if outs&8 != 0 {
			result.Max.Y = Center.Y + Radius
		}
	}
	return result
}
