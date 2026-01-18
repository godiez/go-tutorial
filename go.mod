module test-package

go 1.25.6

replace goodbye-module => ./goodbye

require goodbye-module v0.0.0-00010101000000-000000000000 // indirect
