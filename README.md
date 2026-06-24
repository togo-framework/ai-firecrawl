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
