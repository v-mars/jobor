package fs

import "embed"

//go:embed jobor3.sql
var SqlFs embed.FS

//go:embed dist/*
var DistFs embed.FS
