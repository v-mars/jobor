package fsembed

import "embed"

//go:embed jobor3.sql
var SqlFs embed.FS

//go:embed dist/*
//go:embed dist/assets/_*
var DistFs embed.FS
