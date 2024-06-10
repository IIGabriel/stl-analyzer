package stlanalyzer

import (
	"io"
	"math"
)

type Model3D interface {
	Facets() ([]Facet, error)
	Triangles() ([]Triangle, error)
}

type Facet struct {
	Vector   Vector3D
	Triangle Triangle
}

type Triangle struct {
	V1, V2, V3 Vector3D
}

// Vector3D represents a 3D point.
type Vector3D struct {
	X, Y, Z float64
}

func CalculateTriangleArea(t Triangle) float64 {
	a := math.Sqrt(math.Pow(t.V2.X-t.V1.X, 2) + math.Pow(t.V2.Y-t.V1.Y, 2) + math.Pow(t.V2.Z-t.V1.Z, 2))
	b := math.Sqrt(math.Pow(t.V3.X-t.V2.X, 2) + math.Pow(t.V3.Y-t.V2.Y, 2) + math.Pow(t.V3.Z-t.V2.Z, 2))
	c := math.Sqrt(math.Pow(t.V1.X-t.V3.X, 2) + math.Pow(t.V1.Y-t.V3.Y, 2) + math.Pow(t.V1.Z-t.V3.Z, 2))
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func IsSTLAscii(reader io.ReadSeeker) (bool, error) {
	temp := make([]byte, 5)
	if _, err := reader.Read(temp); err != nil {
		return false, err
	}

	if _, err := reader.Seek(0, 0); err != nil {
		return false, err
	}

	if string(temp) == "solid" {
		return true, nil
	}

	return false, nil
}
