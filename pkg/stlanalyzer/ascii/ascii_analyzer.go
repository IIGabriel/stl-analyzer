package ascii

import (
	"bufio"
	"fmt"
	"io"
	"stl-file-analysis/pkg/stlanalyzer"
	"strings"
)

type analyzer struct {
	r      io.Reader
	name   string
	facets []stlanalyzer.Facet
}

func NewAsciiAnalyzer(reader io.ReadSeeker) (stlanalyzer.Model3D, error) {

	isAscii, err := stlanalyzer.IsSTLAscii(reader)
	if err != nil {
		return nil, err
	}
	if !isAscii {
		return nil, fmt.Errorf("not an ASCII STL file")
	}

	return &analyzer{r: reader}, nil
}

func (s *analyzer) Facets() ([]stlanalyzer.Facet, error) {
	if err := s.checkAndAnalyze(); err != nil {
		return nil, err
	}
	return s.facets, nil
}

func (s *analyzer) checkAndAnalyze() error {
	if s.facets == nil && s.name == "" {
		if err := s.analyze(); err != nil {
			return err
		}
	}
	return nil
}

func (s *analyzer) Triangles() ([]stlanalyzer.Triangle, error) {
	facets, err := s.Facets()
	if err != nil {
		return nil, err
	}
	var triangles []stlanalyzer.Triangle
	for _, facet := range facets {
		triangles = append(triangles, facet.Triangle)
	}
	return triangles, nil
}

func (s *analyzer) analyze() error {

	var (
		scanner = bufio.NewScanner(s.r)
		facets  []stlanalyzer.Facet
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "facet normal") {
			// Read the facet
			var currentFacet stlanalyzer.Facet
			_, err := fmt.Sscanf(line, "facet normal %f %f %f", &currentFacet.Vector.X, &currentFacet.Vector.Y, &currentFacet.Vector.Z)
			if err != nil {
				continue
			}

			var vertexCount int

			// Read the facet loop
			for scanner.Scan() {
				loopLine := strings.TrimSpace(scanner.Text())
				if strings.HasPrefix(loopLine, "outer loop") {
					continue
				}
				if strings.HasPrefix(loopLine, "endloop") {
					break
				}

				if strings.HasPrefix(loopLine, "vertex") {
					var x, y, z float64
					_, _ = fmt.Sscanf(loopLine, "vertex %f %f %f", &x, &y, &z)
					switch vertexCount {
					case 0:
						currentFacet.Triangle.V1 = stlanalyzer.Vector3D{X: x, Y: y, Z: z}
					case 1:
						currentFacet.Triangle.V2 = stlanalyzer.Vector3D{X: x, Y: y, Z: z}
					case 2:
						currentFacet.Triangle.V3 = stlanalyzer.Vector3D{X: x, Y: y, Z: z}
					}
					vertexCount++
					if vertexCount == 3 {
						vertexCount = 0
						break
					}
				}
			}
			facets = append(facets, currentFacet)

		} else if strings.HasPrefix(line, "solid") {
			_, _ = fmt.Sscanf(line, "solid %s", &s.name)
		}
	}
	s.facets = facets

	return scanner.Err()

}
