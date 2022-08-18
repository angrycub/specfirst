# SpecFirst Design Playground



## Build

## Run

- **Mockoon** - Simple mock API generator with passthrough proxy capabilities
- **MrinDoc** - OpenAPI viewer served as a single page application
- **NGINX** - Used as a reverse proxy since Mockoon can only perform a single
  passthrough
- **Apicurito** - OpenAPI editor running in Docker
- **Nomad Dev Agent** - Used to catch-all other requests

```mermaid
graph TD
    Mockoon -- /openapi --> B[fa:fa-file openapi.yaml]
    Mockoon -- /openapi/view --> MrinDoc["fa:fa-file MrinDoc (static HTML)"]
    MrinDoc -- reads --> B

    Mockoon --> NGINX
    
    NGINX -- /openapi/edit --> Apicurito
    Apicurito -- reads --> B

    NGINX -- other requests --> Nomad[Nomad Dev Agent]
```

