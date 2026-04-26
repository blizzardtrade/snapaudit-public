# SnapAudit API Overview

This document describes the public-facing API surface of
[SnapAudit](https://getsnapaudit.com). Internal admin/platform endpoints
are not covered here.

## Base URL

- Production: `https://getsnapaudit.com`

## Authentication

SnapAudit uses session cookies set after email+password login or
Google OAuth. There is no public REST API key today — for programmatic
integrations, contact [Enterprise sales](mailto:hello@getsnapaudit.com).

## Routes (public, indexable)

### Marketing pages

| Path | Purpose |
|------|---------|
| `/{lang}/` | Landing page (per-locale: en/es/fr/de) |
| `/{lang}/login.html` | Sign-in / sign-up form |
| `/{lang}/signup` | Redirects to `/login.html#signup` |
| `/{lang}/offer` | Terms of Service |
| `/{lang}/privacy` | Privacy Policy |

### SEO infrastructure

| Path | Purpose |
|------|---------|
| `/sitemap.xml` | XML sitemap with hreflang xhtml:link alternates and image:image entries |
| `/robots.txt` | Crawl rules + AI bot blocklist |
| `/{key}.txt` | IndexNow verification file (Bing/Yandex/Naver/Seznam) |

### Webhook receivers

| Path | Purpose |
|------|---------|
| `/webhooks/stripe` | Stripe subscription events (HMAC-verified) |

## Locale resolution

For routes under `/{lang}/`, the lang prefix is stripped and the
request's locale is set to that lang. Without a prefix, public pages
302-redirect to `/{detected_lang}<path>` based on (in order):

1. `?lang=xx` query param
2. `locale` cookie
3. `CF-IPCountry` header → mapped locale (ES→es, FR→fr, DE→de, fallback en)

System routes (`/api/*`, `/webhooks/*`, `/sitemap.xml`, `/robots.txt`,
`/css/*`, `/js/*`, `/app`) bypass the redirect logic.

## Rate limiting

Per-IP rate limits apply on:
- `/api/auth/login` — 1 req/10s, burst 5
- `/api/auth/signup` — 1 req/20s, burst 3
- `/api/billing/webhook`, `/webhooks/stripe` — 2 req/s

## Schema.org markup

The landing page emits four JSON-LD blocks in `<head>`:

1. `SoftwareApplication` (general product info)
2. `Organization` (brand info, Knowledge Graph)
3. `Product` with array of `Offer` (per-tier pricing for SERP snippets)
4. `FAQPage` (rich snippet eligible)

Plus on each gallery tile in the markup, `data-detections` JSON for the
modal's coloured-block rendering.

## Useful URLs

- Pricing: https://getsnapaudit.com/en/#pricing
- FAQ: https://getsnapaudit.com/en/#faq
- Industries demo gallery: https://getsnapaudit.com/en/#gallery
- Sales: [hello@getsnapaudit.com](mailto:hello@getsnapaudit.com)
