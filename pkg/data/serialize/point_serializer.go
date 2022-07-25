package serialize

import (
	"github.com/ivankf/tsdb-comparisons/pkg/data"
	"io"
)

// PointSerializer serializes a Point for writing
type PointSerializer interface {
	Serialize(p *data.Point, w io.Writer) error
}
