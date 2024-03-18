package enum

//go:generate enumer -type=DatabaseSystem -trimprefix DatabaseSystem -output=database_system_enumer.go
type DatabaseSystem int8

const (
	DatabaseSystemPostgres DatabaseSystem = iota
	DatabaseSystemMongo
	DatabaseSystemMySQL
)
