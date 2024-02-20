package shapes

import "fmt"

const (
	SmallMultiplier = 0.25
	LargeMultiplier = 10
)

type Drawer interface {
	DrawCircle(x int, y int, radius float64)
}

type SmallCircleDrawer struct {
	radiusMultiplier float64
}

func NewSmallCircleDrawer() *SmallCircleDrawer {
	return &SmallCircleDrawer{
		radiusMultiplier: SmallMultiplier,
	}
}

func (d *SmallCircleDrawer) DrawCircle(x int, y int, radius float64) {
	fmt.Printf("Small circle center %d, %d radius = %0.2f\n", x, y, radius*d.radiusMultiplier)
}

type LargeCircleDrawer struct {
	radiusMultiplier float64
}

func NewLargeCircleDrawer() *LargeCircleDrawer {
	return &LargeCircleDrawer{
		radiusMultiplier: LargeMultiplier,
	}
}

func (d *LargeCircleDrawer) DrawCircle(x int, y int, radius float64) {
	fmt.Printf("Large circle center %d, %d radius = %0.2f\n", x, y, radius*d.radiusMultiplier)
}

type Shape interface {
	Draw()
	EnlargeRadius(multiplier float64)
}

type Circle struct {
	x      int
	y      int
	radius float64
	drawer Drawer
}

func NewCircle(x int, y int, radius float64, drw Drawer) *Circle {
	return &Circle{
		drawer: drw,
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (c *Circle) Draw() {
	c.drawer.DrawCircle(c.x, c.y, c.radius)
}

func (c *Circle) EnlargeRadius(multiplier float64) {
	c.radius *= multiplier
}

func RenderShapes() {
	var shapes = []Shape{
		&Circle{
			x:      5,
			y:      10,
			radius: 10,
			drawer: NewLargeCircleDrawer(),
		},
		&Circle{
			x:      20,
			y:      30,
			radius: 100,
			drawer: NewSmallCircleDrawer(),
		},
	}

	for _, shape := range shapes {
		shape.Draw()
	}

	for _, shape := range shapes {
		shape.EnlargeRadius(float64(5))
		shape.Draw()
	}

}
