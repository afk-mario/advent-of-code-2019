package vect

import (
	"advent/utils"
	"math"
	"sort"
)

// Vector3 that hoolds coord x,y,z
type Vector3 struct {
	X int
	Y int
	Z int
}

// Vector2 holds coord x and y
type Vector2 struct {
	X int
	Y int
}

// Equals Check if two vectors are equal
func (a *Vector3) Equals(b Vector3) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

// ToVec2 converts vector 3 to vector2
func (a *Vector3) ToVec2() Vector2 {
	return Vector2{a.X, a.Y}
}

// Add two vectors
func (a *Vector2) Add(b Vector2) Vector2 {
	return Vector2{a.X + b.X, a.Y + b.Y}
}

// Equals Check if two vectors are equal
func (a *Vector2) Equals(b Vector2) bool {
	return a.X == b.X && a.Y == b.Y
}

// MDistance calculates the distance between center and the vector
func (a *Vector2) MDistance() int {
	c := Vector2{0, 0}
	return utils.Abs(a.X-c.X) + utils.Abs(a.Y-c.Y)
}

func (a *Vector2) cross(b Vector2) int {
	return a.X*b.Y - b.X*a.Y
}

// Angle returns the angle between two vectors
func (a *Vector2) Angle(b Vector2) int {
	angle := math.Atan2(float64(a.Y-b.Y), float64(a.X-b.X))
	angle += 1.5707963268
	result := int(angle * (180 / math.Pi))
	if result < 0 {
		result = 360 - (result * -1)
	}
	return result
}

// SortByX ...
func SortByX(s []Vector2) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].X < s[j].X
	})
}

// SortByY ...
func SortByY(s []Vector2) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Y < s[j].Y
	})
}

// SortByZ ...
func SortByZ(s []Vector3) {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Z < s[j].Z
	})
}

// Line segment that holds point A and B
type Line struct {
	A Vector2
	B Vector2
}

// IsPointOn check if a point is in line
func (l *Line) IsPointOn(p Vector2) bool {
	aTmp := Line{Vector2{0, 0}, Vector2{l.B.X - l.A.X, l.B.Y - l.A.Y}}
	bTmp := Vector2{p.X - l.A.X, p.Y - l.A.Y}
	r := aTmp.B.cross(bTmp)
	return utils.Abs(r) == 0
}

// IsBetween checks if a point is between a line
func (l *Line) IsBetween(c Vector2) bool {
	a := l.A
	b := l.B
	crossproduct := (c.Y-a.Y)*(b.X-a.X) - (c.X-a.X)*(b.Y-a.Y)
	if utils.Abs(crossproduct) != 0 {
		return false
	}

	dotproduct := (c.X-a.X)*(b.X-a.X) + (c.Y-a.Y)*(b.Y-a.Y)
	if dotproduct < 0 {
		return false
	}

	squaredlengthba := (b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y)
	if dotproduct > squaredlengthba {
		return false
	}

	return true
}

func (l *Line) isPointRight(b Vector2) bool {
	aTmp := Line{Vector2{0, 0}, Vector2{l.B.X - l.A.X, l.B.Y - l.A.Y}}
	bTmp := Vector2{b.X - l.A.X, b.Y - l.A.Y}
	return aTmp.B.cross(bTmp) < 0
}

func (l *Line) touchOrCross(b Line) bool {
	return l.IsPointOn(b.A) || l.IsPointOn(b.B) || (l.isPointRight(b.A) != l.isPointRight(b.B))
}

// Intersects checks if two lines intersect
func (l *Line) Intersects(b Line) bool {
	return l.touchOrCross(b) && b.touchOrCross(*l)
}

// Magnitude returns vector magnitude
func (l *Line) Magnitude() int {
	a := utils.Abs(l.B.X - l.A.X)
	b := utils.Abs(l.B.Y - l.A.Y)
	// fmt.Printf("[%d,%d] -> [%d, %d] m %d \n", l.A.X, l.A.Y, l.B.X, l.B.Y, a+b)
	return a + b
}

// GetIntersection get's the coordinates where the lines intersected
func (l *Line) GetIntersection(b Line) Vector2 {
	x, y := -1, -1

	if l.A.X == l.B.X {
		x = l.A.X
	} else if b.A.X == b.B.X {
		x = b.A.X
	}

	if l.A.Y == l.B.Y {
		y = l.A.Y
	} else if b.A.Y == b.B.Y {
		y = b.A.Y
	}

	return Vector2{x, y}
}
