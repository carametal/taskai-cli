# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test Commands
- Build: `go build -o taskai-cli main.go`
- Run: `./taskai-cli`
- Test: `go test ./...`
- Test single package: `go test ./path/to/package`
- Test single function: `go test -run TestFunctionName ./path/to/package`
- Format code: `go fmt ./...`
- Lint: `go vet ./...`

## Code Style Guidelines
- **Formatting**: Follow Go standard formatting with `go fmt`
- **Imports**: Group imports (standard lib first, then external, then internal)
- **Types**: Use descriptive names for types, prefer strong typing with enums via const blocks
- **Naming**: Use camelCase for variables, PascalCase for exported types/functions
- **Error Handling**: Always check errors, return them to caller when appropriate
- **Documentation**: Add comments for exported functions using Go doc conventions
- **File Structure**: Follow standard Go project layout (cmd/, internal/, pkg/)
- **Testing**: Write tests in *_test.go files with table-driven testing approach
- **Dependencies**: Minimize external dependencies, use go modules for dependency management
