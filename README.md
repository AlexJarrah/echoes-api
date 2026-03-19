# Echoes API

API contracts & generated bindings.

## Structure

```
bindings/    # Generated bindings
contracts/   # OpenAPI & AsyncAPI schemas
docs/        # Generated documentation
generation/  # Templatess & configs
```

## Requirements

- [just](https://github.com/casey/just)
- [redocly](https://redocly.com/docs/cli/)
- [asyncapi-cli](https://github.com/asyncapi/cli)
- [Go](https://go.dev/) 1.21+
- [Node.js](https://nodejs.org/en) 18+

## Usage

```bash
just           # List all actions
just generate  # Generate all bindings and docs
just validate  # Validate API contracts
just clean     # Remove generated files
```
