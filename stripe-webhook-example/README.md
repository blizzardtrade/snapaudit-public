# Stripe Webhook Example

Minimal Go example verifying Stripe webhook signatures the same way
[SnapAudit](https://getsnapaudit.com)'s production handler does.

## Quick start

```bash
go mod init stripe-webhook-example
go get github.com/stripe/stripe-go/v82
STRIPE_WEBHOOK_SECRET=whsec_xxx go run main.go
```

In another terminal:

```bash
stripe listen --forward-to localhost:8080/webhook
stripe trigger customer.subscription.created
```

## Key points

1. **Read raw body before parsing.** Any middleware that buffers,
   modifies, or re-serialises the body will break the HMAC.
2. **`IgnoreAPIVersionMismatch: true`** lets your `stripe-go` SDK be
   newer or older than the version your webhook endpoint was created
   under. The HMAC still verifies — only field shapes can drift.
3. **Always 200 on unknown event types.** Non-200 triggers retry up
   to 72 hours; logging once and returning 200 is the safe default
   for unhandled types.

## License

MIT. Use freely. Reference: https://getsnapaudit.com
