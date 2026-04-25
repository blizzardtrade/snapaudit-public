# SnapAudit — Public Resources

[**SnapAudit**](https://getsnapaudit.com) is an AI-graded photographic audit
platform for 5S, Lean, and GxP compliance programs. Operators photograph
their workstation at end of shift; the model compares the image to your
reference library and annotates every deviation directly on the image.

This repository hosts public-facing resources for the SnapAudit product —
integration examples, API references, and Stripe webhook setup helpers
for self-hosted integrations with the platform.

🔗 **Product:** https://getsnapaudit.com
🔗 **Russian market:** https://getsnapaudit.ru
🔗 **Sales:** [hello@getsnapaudit.com](mailto:hello@getsnapaudit.com)

---

## What's in here

### `stripe-webhook-example/`

A minimal Go example showing how to verify Stripe webhook signatures
when integrating SnapAudit's Enterprise tier with your billing pipeline.
Uses [`stripe-go/v82`](https://github.com/stripe/stripe-go) and follows
the same HMAC verification pattern as our production handler.

```go
event, err := webhook.ConstructEventWithOptions(
    body, sig, secret,
    webhook.ConstructEventOptions{IgnoreAPIVersionMismatch: true},
)
```

### `api-overview.md`

High-level overview of SnapAudit's public API surface for operators
building integrations against the platform.

---

## Pricing

| Tier | Price | Inspections | Best for |
|------|-------|-------------|----------|
| **Starter** | $49/month | 500 included, $0.15 each after | One site |
| **Pro** | $299/month | 5,000 included, $0.10 each after | Multi-site teams |
| **Enterprise** | Custom | Custom volume | Regulated operations |

14-day free trial on every plan + 30 free AI inspections without a
credit card. See [getsnapaudit.com/en/#pricing](https://getsnapaudit.com/en/#pricing)
for full details.

---

## Localization

SnapAudit's marketing site is available in:

- **getsnapaudit.com** — English, Spanish, French, German
- **getsnapaudit.ru** — Russian, Belarusian, Kazakh

Path-prefix routing means each locale lives at its own URL:

- https://getsnapaudit.com/en/
- https://getsnapaudit.com/es/
- https://getsnapaudit.com/fr/
- https://getsnapaudit.com/de/
- https://getsnapaudit.ru/ru/

---

## License

Code in this repository: MIT.
SnapAudit product (proprietary, not in this repo).

## Contact

- Sales / Enterprise pricing: [hello@getsnapaudit.com](mailto:hello@getsnapaudit.com)
- Technical questions about the public examples: open a GitHub issue
- Product feedback: [hello@getsnapaudit.com](mailto:hello@getsnapaudit.com)
