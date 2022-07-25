package constants

// Formats supported for generation
const (
	FormatInflux      = "influx"
	FormatTimescaleDB = "timescaledb"
	FormatTDEngine    = "tdengine"
	FormatIOTDB       = "iotdb"
)

func SupportedFormats() []string {
	return []string{
		FormatInflux,
		FormatTimescaleDB,
		FormatTDEngine,
		FormatIOTDB,
	}
}
