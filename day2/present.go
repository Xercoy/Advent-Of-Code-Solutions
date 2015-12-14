package main

type Present struct {
	dimensions
}

func NewPresent(l float64, w float64, h float64) *Present {
	return &Present{
		dimensions{
			Length: l,
			Width:  w,
			Height: h,
		},
	}
}

// Surface Area = 2*l*w + 2*w*h + 2*h*l
func (p *Present) SurfaceArea() float64 {
	return (2 * p.Length * p.Width) +
		(2 * p.Width * p.Height) +
		(2 * p.Height * p.Length)
}

func (p *Present) SmallestSideArea() float64 {
	side1 := p.Length * p.Width
	side2 := p.Width * p.Height
	side3 := p.Height * p.Length

	smallestSide := side1

	if side2 < smallestSide {
		smallestSide = side2
	}

	if side3 < smallestSide {
		smallestSide = side3
	}

	return smallestSide
}

func (p *Present) SmallestPerimeter() float64 {
	perim1 := (2 * p.Width) + (2 * p.Height)
	perim2 := (2 * p.Width) + (2 * p.Length)
	perim3 := (2 * p.Length) + (2 * p.Height)

	smallestPerim := perim1

	if perim2 < perim1 {
		smallestPerim = perim2
	}

	if perim3 < smallestPerim {
		smallestPerim = perim3
	}

	return smallestPerim
}

func (p *Present) Volume() float64 {
	return (p.Length * p.Width * p.Height)
}
