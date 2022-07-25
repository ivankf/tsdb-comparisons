package factories

import (
	"github.com/ivankf/tsdb-comparisons/cmd/generate_queries/databases/influx"
	"github.com/ivankf/tsdb-comparisons/cmd/generate_queries/databases/iotdb"
	"github.com/ivankf/tsdb-comparisons/cmd/generate_queries/databases/tdengine"
	"github.com/ivankf/tsdb-comparisons/cmd/generate_queries/databases/timescaledb"

	"github.com/ivankf/tsdb-comparisons/pkg/query/config"
	"github.com/ivankf/tsdb-comparisons/pkg/targets/constants"
)

func InitQueryFactories(config *config.QueryGeneratorConfig) map[string]interface{} {
	factories := make(map[string]interface{})
	factories[constants.FormatInflux] = &influx.BaseGenerator{}
	factories[constants.FormatTimescaleDB] = &timescaledb.BaseGenerator{
		UseJSON:       config.TimescaleUseJSON,
		UseTags:       config.TimescaleUseTags,
		UseTimeBucket: config.TimescaleUseTimeBucket,
	}
	factories[constants.FormatTDEngine] = &tdengine.BaseGenerator{}
	factories[constants.FormatIOTDB] = &iotdb.BaseGenerator{}

	return factories
}
