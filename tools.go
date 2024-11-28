//go:build tools

package main

// import dev tools here to lock its version.

import (
	_ "github.com/99designs/gqlgen"
)

//go:generate go run github.com/99designs/gqlgen generate
//go:generate gqlgenc
