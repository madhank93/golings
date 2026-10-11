// modules10
// go.mod has drifted from the code. It still requires oldfmt, which nothing
// imports any more, and it never recorded pad, which table needs. Don't fix it
// by hand: from this directory, let the toolchain rewrite go.mod to match the
// imports. The test checks the result.

// I AM NOT DONE

module example.com/report

go 1.26.0

require (
	example.com/oldfmt v0.0.0
	example.com/table v0.0.0
)

replace (
	example.com/oldfmt => ./oldfmt
	example.com/pad => ./pad
	example.com/table => ./table
)
