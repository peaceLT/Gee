module example

go 1.23.2

require (
	gee v1.0.0
)

replace (
	gee v1.0.0 => ./gee
)