# Echoes API

Echoes API contracts, documentation, and generated bindings.

## Structure

```
bindings/    # Language bindings
contracts/   # API contracts
docs/        # ocumentation
generation/  # Generation utilities
```

## Requirements

- [just](https://github.com/casey/just)
- [redocly](https://redocly.com/docs/cli/)
- [asyncapi-cli](https://github.com/asyncapi/cli)
- [Go](https://go.dev/) 1.21+
- [Node.js](https://nodejs.org/en) 18+

## Usage

```bash
$ just
Available recipes:
    generate          # Generate all bindings & documentation
    schemas           # Generate OpenAPI schema
    validate          # Validate API contracts
    bundle            # Bundle split contracts
    clean             # Remove generated bindings
    docs              # Generate documentation
    go                # Generate Go bindings
    go-openapi        # Generate Go OpenAPI bindings
    go-asyncapi       # Generate Go AsyncAPI bindings
    go-schemas        # Generate Go schema constants
    ts                # Generate TypeScript bindings
    ts-openapi        # Generate TypeScript OpenAPI bindings
    ts-asyncapi       # Generate TypeScript AsyncAPI bindings
    ts-openapi-types  # Generate TypeScript OpenAPI types
    ts-openapi-client # Generate TypeScript OpenAPI client
    ts-asyncapi-types # Generate TypeScript AsyncAPI types
```
