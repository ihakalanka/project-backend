package main

import (
    "log"
    "net/http"
    "github.com/stripe/stripe-go/v72"
    "github.com/stripe/stripe-go/v72/checkout/session"
)

func main() {
  // This is your test secret API key.
  stripe.Key = "sk_test_51LD7toGQ8mDQYrPf5JpNTNcTIDuBKbsYxiBu4kwFPEmEMxeNuyFhoNdup5gy3BgzydFwwdCjKLwBpyQwYh2hHnFD00ZVMXkEXo"

  http.Handle("/", http.FileServer(http.Dir("public")))
  http.HandleFunc("/create-checkout-session", createCheckoutSession)
  addr := "localhost:3000"
  log.Printf("Listening on %s", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
  domain := "http://localhost:3000"
  params := &stripe.CheckoutSessionParams{
    LineItems: []*stripe.CheckoutSessionLineItemParams{
      &stripe.CheckoutSessionLineItemParams{
        // Provide the exact Price ID (for example, pr_1234) of the product you want to sell
        Price: stripe.String("{{PRICE_ID}}"),
        Quantity: stripe.Int64(1),
      },
    },
    Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
    SuccessURL: stripe.String(domain + "?success=true"),
    CancelURL: stripe.String(domain + "?canceled=true"),
  }

  s, err := session.New(params)

  if err != nil {
    log.Printf("session.New: %v", err)
  }

  http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
