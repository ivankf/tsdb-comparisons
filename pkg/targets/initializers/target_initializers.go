package initializers

import (
	"fmt"
	"strings"

	"github.com/ivankf/tsdb-comparisons/pkg/targets"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/constants"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/influx"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/iotdb"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/tdengine"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/timescaledb"
)

func GetTarget(format string) targets.ImplementedTarget {
	switch format {
	case constants.FormatTimescaleDB:
		return timescaledb.NewTarget()
	case constants.FormatInflux:
		return influx.NewTarget()
	case constants.FormatTDEngine:
		return tdengine.NewTarget()

	case constants.FormatIOTDB:
		return iotdb.NewTarget()
	}
	supportedFormatsStr := strings.Join(constants.SupportedFormats(), ",")
	panic(fmt.Sprintf("Unrecognized format %s, supported: %s", format, supportedFormatsStr))
}
