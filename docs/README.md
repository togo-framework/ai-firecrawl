# ai-firecrawl — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package firecrawl is a togo AI data-source plugin: scrape/crawl web pages to
clean markdown via Firecrawl — both a self-hosted instance (FIRECRAWL_URL) and
the hosted API (FIRECRAWL_API_KEY) — so ai-rag ingest and agents can pull web
content. Registers an "ai-firecrawl" service + REST endpoints under
/api/ai/firecrawl. Self-hosted default base http://localhost:3002.

## Install

```bash
togo install togo-framework/ai-firecrawl
```

A capability plugin — it self-registers on boot; no driver selector needed.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `FIRECRAWL_API_KEY` | _see provider docs_ |
| `FIRECRAWL_URL` | _see provider docs_ |
| `G` | _see provider docs_ |

## Usage

```go
// A data source for ai-rag / agents: fetch/scrape/search web content.
src := firecrawl.FromKernel(k)
docs, err := src.Fetch(ctx, "https://example.com")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/ai-firecrawl
- README: ../README.md
