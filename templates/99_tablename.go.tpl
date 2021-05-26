{{- $alias := .Aliases.Table .Table.Name -}}

// GetTableName returns the real table name.
func (_ {{$alias.UpSingular}}) GetTableName() string { return "{{.Table.Name}}" }
