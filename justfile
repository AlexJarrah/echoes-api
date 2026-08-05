set shell := ["bash", "-euo", "pipefail", "-c"]

contracts  := absolute_path("contracts")
generation := absolute_path("generation")
bindings   := absolute_path("bindings")
docs       := absolute_path("docs")

openapi_schema  := absolute_path(contracts / "openapi/openapi.yaml")
asyncapi_schema := absolute_path(contracts / "asyncapi/asyncapi.yaml")

bundles         := `mktemp -d`
openapi_bundle  := absolute_path(bundles / "openapi.yaml")
asyncapi_bundle := absolute_path(bundles / "asyncapi.yaml")

_default:
    @just --list --unsorted

# Generate all bindings
generate: clean schemas bundle docs go ts

schemas:
  python3 {{generation / "schemas/generate_openapi.py"}} \
    "$(dirname {{openapi_schema}})/openapi_template.yaml" \
    {{openapi_schema}}

# Validate OpenAPI & AsyncAPI contracts
validate:
    redocly lint {{openapi_schema}}
    asyncapi validate {{asyncapi_schema}}

# Bundle split contracts
bundle:
    mkdir -p {{bundles}}
    redocly bundle {{openapi_schema}} -o {{openapi_bundle}}
    asyncapi bundle {{asyncapi_schema}} -o {{asyncapi_bundle}}

# Remove generated bindings
clean:
    rm -f {{openapi_schema}} 
    rm -rf {{docs}}/*
    rm -rf {{bindings}}/go/{openapi,asyncapi,schemas}
    rm -rf {{bindings}}/ts/{openapi,asyncapi}

# Generate documentation
docs: bundle
    cp {{generation}}/docs/scalar-template.html {{docs}}/openapi.html
    sed -i "s|<doc-url>|{{openapi_bundle}}|g" {{docs}}/openapi.html
    cp {{generation}}/docs/scalar-template.html {{docs}}/asyncapi.html
    sed -i "s|<doc-url>|{{asyncapi_bundle}}|g" {{docs}}/asyncapi.html

# Generate Go bindings
go: go-openapi go-asyncapi go-schemas

go-openapi: bundle
    mkdir -p {{bindings}}/go/openapi
    (cd {{bindings}}/go && go install tool)
    ogen --target {{bindings}}/go/openapi \
         --package openapi \
         --clean \
         --config {{generation}}/go/ogen.yml \
         {{openapi_bundle}}

go-asyncapi: bundle
    mkdir -p {{bindings}}/go/asyncapi
    (cd {{bindings}}/go && go install tool)
    go-asyncapi code {{asyncapi_bundle}} \
      --target-dir {{bindings}}/go/asyncapi/ \
      -M gitlab.com/AlexJarrah/echoes-api/bindings/go/asyncapi

go-schemas: bundle
    mkdir -p {{bindings}}/go/schemas
    printf 'package schemas\n\nconst OpenAPI = `%s`\n' \
        "$(sed 's/`/`+"`"+`/g' {{openapi_bundle}})" \
        > {{bindings}}/go/schemas/openapi.go
    printf 'package schemas\n\nconst AsyncAPI = `%s`\n' \
        "$(sed 's/`/`+"`"+`/g' {{asyncapi_bundle}})" \
        > {{bindings}}/go/schemas/asyncapi.go
    gofmt -w {{bindings}}/go/schemas/*

# Generate TypeScript bindings
ts: ts-openapi ts-asyncapi

ts-openapi: ts-openapi-types ts-openapi-client

ts-asyncapi: ts-asyncapi-types

ts-openapi-types: bundle
    mkdir -p {{bindings}}/ts/openapi
    npx openapi-typescript {{openapi_bundle}} \
        -o {{bindings}}/ts/openapi/openapi-typescript.d.ts

ts-openapi-client: bundle
    mkdir -p {{bindings}}/ts/openapi/client
    npx @hey-api/openapi-ts -i {{openapi_bundle}} \
                            -o {{bindings}}/ts/openapi/client \
                            --plugins "@hey-api/client-fetch"

ts-asyncapi-types: bundle
    mkdir -p {{bindings}}/ts/asyncapi/types
    npx @asyncapi/cli generate models typescript \
        {{asyncapi_bundle}} \
        -o {{bindings}}/ts/asyncapi/types
