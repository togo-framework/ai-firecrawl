// Package firecrawl is a togo AI data-source plugin: scrape/crawl web pages to
// clean markdown via Firecrawl — both a self-hosted instance (FIRECRAWL_URL) and
// the hosted API (FIRECRAWL_API_KEY) — so ai-rag ingest and agents can pull web
// content. Registers an "ai-firecrawl" service + REST endpoints under
// /api/ai/firecrawl. Self-hosted default base http://localhost:3002.
package firecrawl

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/togo-framework/togo"
)

// Source talks to a Firecrawl instance (self-hosted or hosted).
type Source struct {
	base   string
	key    string
	client *http.Client
}

// New builds a Source. FIRECRAWL_URL selects a self-hosted base; otherwise the
// hosted API (https://api.firecrawl.dev) with FIRECRAWL_API_KEY.
func New() *Source {
	base := os.Getenv("FIRECRAWL_URL")
	if base == "" {
		if os.Getenv("FIRECRAWL_API_KEY") != "" {
			base = "https://api.firecrawl.dev"
		} else {
			base = "http://localhost:3002"
		}
	}
	return &Source{base: base, key: os.Getenv("FIRECRAWL_API_KEY"), client: &http.Client{Timeout: 90 * time.Second}}
}

func (s *Source) post(ctx context.Context, path string, in any) (map[string]any, error) {
	buf, _ := json.Marshal(in)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.base+path, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if s.key != "" {
		req.Header.Set("Authorization", "Bearer "+s.key)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// Scrape returns the markdown of a single page.
func (s *Source) Scrape(ctx context.Context, url string) (string, error) {
	out, err := s.post(ctx, "/v1/scrape", map[string]any{"url": url, "formats": []string{"markdown"}})
	if err != nil {
		return "", err
	}
	if data, ok := out["data"].(map[string]any); ok {
		if md, ok := data["markdown"].(string); ok {
			return md, nil
		}
	}
	return "", nil
}

// Crawl starts a crawl job for a site and returns the raw response (job id/status).
func (s *Source) Crawl(ctx context.Context, url string) (map[string]any, error) {
	return s.post(ctx, "/v1/crawl", map[string]any{"url": url})
}

// FromKernel returns the registered Source, or nil.
func FromKernel(k *togo.Kernel) *Source {
	if v, ok := k.Get("ai-firecrawl"); ok {
		if s, ok := v.(*Source); ok {
			return s
		}
	}
	return nil
}

func init() {
	togo.RegisterProviderFunc("ai-firecrawl", togo.PriorityService, func(k *togo.Kernel) error {
		s := New()
		k.Set("ai-firecrawl", s)
		mount(k.Router, s)
		return nil
	})
}

func mount(r chi.Router, s *Source) {
	r.Route("/api/ai/firecrawl", func(r chi.Router) {
		r.Post("/scrape", func(w http.ResponseWriter, req *http.Request) {
			var b struct {
				URL string `json:"url"`
			}
			if err := json.NewDecoder(req.Body).Decode(&b); err != nil || b.URL == "" {
				http.Error(w, `{"error":"url required"}`, http.StatusBadRequest)
				return
			}
			md, err := s.Scrape(req.Context(), b.URL)
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadGateway)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"markdown": md})
		})
		r.Post("/crawl", func(w http.ResponseWriter, req *http.Request) {
			var b struct {
				URL string `json:"url"`
			}
			if err := json.NewDecoder(req.Body).Decode(&b); err != nil || b.URL == "" {
				http.Error(w, `{"error":"url required"}`, http.StatusBadRequest)
				return
			}
			out, err := s.Crawl(req.Context(), b.URL)
			if err != nil {
				http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadGateway)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(out)
		})
	})
}
