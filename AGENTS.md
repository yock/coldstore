# AGENTS.md

This file provides guidelines for agentic coding agents working on the coldstore repository.

## Build/Lint/Test Commands

- Build: go build
- Run: go run main.go
- Lint: go vet ./... && go fmt ./...
- Test: go test ./... (Note: Currently no test files)
- Run single test: go test -run ^TestFunctionName$ [path/to/package]

## Code Style Guidelines

- Formatting: Use go fmt for standard Go formatting. Indent with 2 spaces per .editorconfig.
- Imports: Group imports - standard library first, then third-party, then local packages.
- Naming Conventions: Use CamelCase for exported identifiers; lowercase camelCase for private.
- Types: Use structs for models (e.g., GORM models). Prefer interfaces for abstractions.
- Error Handling: Always check if err != nil; use log.Fatal for critical errors or return errors.
- Templates: Use html/template with embed.FS for embedding templates.
- Handlers: Use gorilla/mux for routing, subrouters for modules.
- Database: Use GORM with PostgreSQL; connect via data.Connect().
- General: Keep code concise, no unnecessary comments. Follow Go best practices.
