package migrations

import "embed"

//go:embed "tables"
var MigrationFiles embed.FS
