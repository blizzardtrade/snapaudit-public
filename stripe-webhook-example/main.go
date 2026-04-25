// Package main is a minimal Go example showing how to verify Stripe
// webhook signatures the same way SnapAudit's production handler does.
//
// Build & run:
//   go mod init stripe-webhook-example
//   go get github.com/stripe/stripe-go/v82
//   STRIPE_WEBHOOK_SECRET=whsec_xxx go run main.go
//
// Test with the Stripe CLI:
//   stripe listen --forward-to localhost:8080/webhook
//   stripe trigger customer.subscription.created
//
// SnapAudit reference: https://getsnapaudit.com
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v82/webhook"
)

func main() {
	secret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if secret == "" {
		log.Fatal("STRIPE_WEBHOOK_SECRET env var required")
	}
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "read body", http.StatusBadRequest)
			return
		}
		// IgnoreAPIVersionMismatch=true mirrors SnapAudit's production
		// pattern: HMAC verifies regardless of which API version the
		// endpoint was registered under, so SDK upgrades don't break
		// signature validation. The subscription/invoice fields you
		// read on the event ARE version-sensitive — read them
		// defensively.
		event, err := webhook.ConstructEventWithOptions(
			body,
			r.Header.Get("Stripe-Signature"),
			secret,
			webhook.ConstructEventOptions{IgnoreAPIVersionMismatch: true},
		)
		if err != nil {
			http.Error(w, "signature: "+err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("received event: %s id=%s\n", event.Type, event.ID)
		// Stripe expects 200 even on unhandled types — non-200 triggers
		// retry up to 72h. Handle each event.Type explicitly and 200
		// for the rest (including unknown types).
		w.WriteHeader(http.StatusOK)
	})
	log.Println("listening on :8080/webhook")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
