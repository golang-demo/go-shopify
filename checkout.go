package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const checkoutsBasePath = "checkouts"

// Options for checkout list
type CheckoutListOptions struct {
	Role   string `url:"role,omitempty"`
	Fields string `url:"fields,omitempty"`
}

// CheckoutService is an interface for interfacing with the checkouts endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/checkout
type CheckoutService interface {
	Create(Checkout) (*Checkout, error)
	Get(string, interface{}) (*Checkout, error)
	Update(Checkout) (*Checkout, error)
}

// CheckoutServiceOp handles communication with the checkout related methods of
// the Shopify API.
type CheckoutServiceOp struct {
	client *Client
}

// {
//     "completed_at": null,
//     "created_at": "2012-10-12T07:05:27-04:00",
//     "currency": "USD",
//     "presentment_currency": "USD",
//     "customer_id": 1073339459,
//     "customer_locale": "en",
//     "device_id": null,
//     "discount_code": null,
//     "discount_codes": [],
//     "email": "john.smith@example.com",
//     "legal_notice_url": null,
//     "location_id": null,
//     "name": "#862052962",
//     "note": "",
//     "note_attributes": {
//       "custom engraving": "Happy Birthday",
//       "colour": "green"
//     },
//     "order_id": null,
//     "order_status_url": null,
//     "order": null,
//     "payment_due": "398.00",
//     "payment_url": "https://app.local/cardserver/sessions",
//     "payments": [],
//     "phone": null,
//     "shopify_payments_account_id": null,
//     "privacy_policy_url": null,
//     "refund_policy_url": null,
//     "requires_shipping": true,
//     "reservation_time_left": 0,
//     "reservation_time": null,
//     "source_identifier": null,
//     "source_name": "web",
//     "source_url": null,
//     "subscription_policy_url": null,
//     "subtotal_price": "398.00",
//     "shipping_policy_url": null,
//     "tax_exempt": false,
//     "taxes_included": false,
//     "terms_of_sale_url": null,
//     "terms_of_service_url": null,
//     "token": "exuw7apwoycchjuwtiqg8nytfhphr62a",
//     "total_price": "398.00",
//     "total_tax": "0.00",
//     "total_tip_received": "0.00",
//     "total_line_items_price": "398.00",
//     "updated_at": "2022-10-03T12:25:22-04:00",
//     "user_id": null,
//     "web_url": "https://checkout.local/548380009/checkouts/exuw7apwoycchjuwtiqg8nytfhphr62a",
//     "total_duties": null,
//     "total_additional_fees": null,
//     "line_items": [

//     ],
//     "gift_cards": [],
//     "tax_lines": [],
//     "tax_manipulations": [],
//     "shipping_line": null,
//     "shipping_rate": null,
//     "shipping_address": {
//       "id": 550558813,
//       "first_name": "John",
//       "last_name": "Smith",
//       "phone": "(123)456-7890",
//       "company": null,
//       "address1": "126 York St.",
//       "address2": "",
//       "city": "Los Angeles",
//       "province": "California",
//       "province_code": "CA",
//       "country": "United States",
//       "country_code": "US",
//       "zip": "90002"
//     },
//     "credit_card": null,
//     "billing_address": {
//       "id": 550558813,
//       "first_name": "Bob",
//       "last_name": "Norman",
//       "phone": "+1(502)-459-2181",
//       "company": null,
//       "address1": "Chestnut Street 92",
//       "address2": "",
//       "city": "Louisville",
//       "province": "Kentucky",
//       "province_code": "KY",
//       "country": "United States",
//       "country_code": "US",
//       "zip": "40202"
//     },
//     "applied_discount": null,
//     "applied_discounts": [],
//     "discount_violations": []
//   }
// Checkout represents a Shopify checkout
type Checkout struct {
	ID                  int64      `json:"id,omitempty"`
	CompletedAt         *time.Time `json:"completed_at,omitempty"`
	Currency            string     `json:"currency,omitempty"`
	PresentmentCurrency string     `json:"presentment_currency,omitempty"`
	CreatedAt           time.Time  `json:"created_at,omitempty"`
	UpdatedAt           time.Time  `json:"updated_at,omitempty"`
	LandingSite         string     `json:"landing_site,omitempty"`
	// CustomerId          int64      `json:"customer_id,omitempty"`
	// CustomerLocale      string     `json:"customer_locale,omitempty"`
	// // device_id            string     `json:"device_id,omitempty"`
	DiscountCode string `json:"discount_code,omitempty"`
	// DiscountCodes []DiscountCode `json:"discount_codes,omitempty"`
	Email string `json:"email,omitempty"`
	// // legal_notice_url           `json:"legal_notice_url,omitempty"`
	// // location_id           `json:"location_id,omitempty"`
	// Name string `json:"name,omitempty"`
	// Note string `json:"note,omitempty"`
	// NoteAttributes  map[string]string `json:"note_attributes,omitempty"`
	LineItems []CheckoutLineItem `json:"line_items,omitempty"`
	// TaxLines        []TaxLine  `json:"tax_lines,omitempty"`
	// ShippingAddress Address    `json:"shipping_address,omitempty"`
	Token  string `json:"token,omitempty"`
	WebURL string `json:"web_url,omitempty"`
}

type CheckoutLineItem struct {
	ID                 string           `json:"id,omitempty"`
	Key                string           `json:"key,omitempty"`
	ProductID          int64            `json:"product_id,omitempty"`
	VariantID          int64            `json:"variant_id,omitempty"`
	Quantity           int              `json:"quantity,omitempty"`
	Price              *decimal.Decimal `json:"price,omitempty"`
	CompareAtPrice     *decimal.Decimal `json:"compare_at_price,omitempty"`
	LinePrice          *decimal.Decimal `json:"line_price,omitempty"`
	Title              string           `json:"title,omitempty"`
	VariantTitle       string           `json:"variant_title,omitempty"`
	SKU                string           `json:"sku,omitempty"`
	Vendor             string           `json:"vendor,omitempty"`
	ImageURL           string           `json:"image_url,omitempty"`
	GiftCard           bool             `json:"gift_card,omitempty"`
	Taxable            bool             `json:"taxable,omitempty"`
	FulfillmentService string           `json:"fulfillment_service,omitempty"`
	RequiresShipping   bool             `json:"requires_shipping,omitempty"`
	// Properties           `json:"properties,omitempty"`
	Grams    int       `json:"grams,omitempty"`
	TaxLines []TaxLine `json:"tax_lines,omitempty"`
}

// CheckoutResource is the result from the checkouts/X.json endpoint
type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

// CheckoutsResource is the result from the checkouts.json endpoint
type CheckoutsResource struct {
	Checkouts []Checkout `json:"checkouts"`
}

// Create a checkout
func (s *CheckoutServiceOp) Create(checkout Checkout) (*Checkout, error) {
	path := fmt.Sprintf("%s.json", checkoutsBasePath)
	wrappedData := CheckoutResource{Checkout: &checkout}
	resource := new(CheckoutResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Checkout, err
}

// Get a checkout
func (s *CheckoutServiceOp) Get(Token string, options interface{}) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s.json", themesBasePath, Token)
	resource := new(CheckoutResource)
	err := s.client.Get(path, resource, options)
	return resource.Checkout, err
}

// Update a checkouts
func (s *CheckoutServiceOp) Update(checkout Checkout) (*Checkout, error) {
	path := fmt.Sprintf("%s/%s.json", themesBasePath, checkout.Token)
	wrappedData := CheckoutResource{Checkout: &checkout}
	resource := new(CheckoutResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Checkout, err
}
