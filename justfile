set shell := ["bash", "-euo", "pipefail", "-c"]

contracts  := "./contracts"
generation := "./generation"
bindings   := "./bindings"
docs       := "./docs"

_default:
    @just --list --unsorted

# Generate all bindings
generate: docs go ts

# Validate OpenAPI & AsyncAPI contracts
validate:
    redocly lint {{contracts}}/openapi.yaml
    asyncapi validate {{contracts}}/asyncapi.yaml

# Remove generated bindings
clean:
    rm -rf {{docs}}/*
    rm -rf {{bindings}}/go/{openapi,asyncapi}
    rm -rf {{bindings}}/ts/{openapi,asyncapi}

# Generate documentation
docs:
    cp {{generation}}/docs/scalar-template.html {{docs}}/openapi.html
    sed -i "s|<doc-url>|./contracts/openapi.yaml|g" {{docs}}/openapi.html
    cp {{generation}}/docs/scalar-template.html {{docs}}/asyncapi.html
    sed -i "s|<doc-url>|./contracts/asyncapi.yaml|g" {{docs}}/asyncapi.html

# Generate Go bindings
go: go-openapi go-asyncapi

go-openapi:
    mkdir -p {{bindings}}/go/openapi
    cd {{bindings}}/go && go install tool
    ogen --target {{bindings}}/go/openapi \
         --package openapi \
         --clean \
         --config {{generation}}/go/ogen.yml \
         {{contracts}}/openapi.yaml

go-asyncapi:
    mkdir -p {{bindings}}/go/asyncapi
    cd {{bindings}}/go && go install tool
    go-asyncapi code {{contracts}}/asyncapi.yaml \
      --target-dir {{bindings}}/go/asyncapi/ \
      -M gitlab.com/AlexJarrah/echoes-api/bindings/go/asyncapi

# Generate TypeScript bindings
ts: ts-openapi ts-asyncapi

ts-openapi: ts-openapi-types ts-openapi-client

ts-asyncapi: ts-asyncapi-types

ts-openapi-types:
    mkdir -p {{bindings}}/ts/openapi
    npx openapi-typescript {{contracts}}/openapi.yaml \
        -o {{bindings}}/ts/openapi/openapi-typescript.d.ts

ts-openapi-client:
    mkdir -p {{bindings}}/ts/openapi/client
    npx @hey-api/openapi-ts -i {{contracts}}/openapi.yaml \
                            -o {{bindings}}/ts/openapi/client \
                            --plugins "@hey-api/client-fetch"

ts-asyncapi-types:
    mkdir -p {{bindings}}/ts/asyncapi/types
    npx @asyncapi/cli generate models typescript \
        {{contracts}}/asyncapi.yaml \
        -o {{bindings}}/ts/asyncapi/types
