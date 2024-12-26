package rectangle

func (r *Rect) Perimeter() float32 {
	return 2 * (r.l + r.b)
}
