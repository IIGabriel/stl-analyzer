package models

type TrianglesHTTPResponse struct {
	NumberOfTriangles uint64  `json:"number_of_triangles"`
	SurfaceArea       float64 `json:"surface_area"`
}
