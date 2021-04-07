package models

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) Equal(o *Coordinates) bool {
	if o == nil {
		return false
	}

	return c.X == o.X && c.Y == o.Y
}