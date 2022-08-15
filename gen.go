package main

//go:generate go install github.com/goccmack/gogll/v3@latest
//go:generate go run ./tools/clean script
//go:generate go run ./tools/clean src
//go:generate gogll -o script -go confscript.md
//go:generate gogll -rust confscript.md
