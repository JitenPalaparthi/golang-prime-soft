
type Rect struct {
	l, b float32
}
type RectError struct {
	ErrCode int // l is zero or less than zero 1
	Message string // b is zero or less than zero 2
	// if both of them are zeros then error code should be 3
}

func (re *RectError) Error() string {

	return ""
}

func NewRect(l, b float32) (*Rect, error)

//Area and Perimeter
// if l or b is zero, to return an error
