<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/ai-firecrawl</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/ai-firecrawl"><img src="https://pkg.go.dev/badge/github.com/togo-framework/ai-firecrawl.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/ai-firecrawl
```

<!-- /togo-header -->

# ai-firecrawl

A togo **AI data-source** plugin — scrape/crawl web pages to clean **markdown** with [Firecrawl](https://firecrawl.dev), for `ai-rag` ingest and agents. Supports **self-hosted** and the **hosted API**.

```
togo install togo-framework/ai-firecrawl
```
- Self-hosted: set `FIRECRAWL_URL` (default `http://localhost:3002`).
- Hosted: set `FIRECRAWL_API_KEY` (base defaults to `https://api.firecrawl.dev`).

## Use
- Go: `firecrawl.FromKernel(k).Scrape(ctx, "https://example.com")` → markdown · `.Crawl(ctx, url)`
- REST: `POST /api/ai/firecrawl/scrape` `{"url":"…"}` · `POST /api/ai/firecrawl/crawl` `{"url":"…"}`

Part of the [togo AI kit](https://to-go.dev/ai). MIT.

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
