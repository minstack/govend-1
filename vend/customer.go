// Package vend handles interactions with the Vend API.
package vend

import (
	"encoding/json"
	"log"
)

// Customer is a customer object.
type Customer struct {
	ID               *string  `json:"id,omitempty"`
	Code             *string  `json:"customer_code,omitempty"`
	FirstName        *string  `json:"first_name,omitempty"`
	LastName         *string  `json:"last_name,omitempty"`
	Email            *string  `json:"email,omitempty"`
	YearToDate       *float64 `json:"year_to_date,omitempty"`
	Balance          *float64 `json:"balance,omitempty"`
	LoyaltyBalance   *float64 `json:"loyalty_balance,omitempty"`
	Note             *string  `json:"note,omitempty"`
	Gender           *string  `json:"gender,omitempty"`
	DateOfBirth      *string  `json:"date_of_birth,omitempty"`
	CompanyName      *string  `json:"company_name,omitempty"`
	Phone            *string  `json:"phone,omitempty"`
	Mobile           *string  `json:"mobile,omitempty"`
	Fax              *string  `json:"fax,omitempty"`
	Twitter          *string  `json:"twitter,omitempty"`
	Website          *string  `json:"website,omitempty"`
	PhysicalSuburb   *string  `json:"physical_suburb,omitempty"`
	PhysicalCity     *string  `json:"physical_city,omitempty"`
	PhysicalPostcode *string  `json:"physical_postcode,omitempty"`
	PhysicalState    *string  `json:"physical_state,omitempty"`
	PostalSuburb     *string  `json:"postal_suburb,omitempty"`
	PostalCity       *string  `json:"postal_city,omitempty"`
	PostalState      *string  `json:"postal_state,omitempty"`
	CreatedAt        *string  `json:"created_at,omitempty"`
	PostalPostcode   *string  `json:"postal_postcode,omitempty"`
	PhysicalAddress1 *string  `json:"physical_address1,omitempty"`
	PhysicalAddress2 *string  `json:"physical_address2,omitempty"`
	PostalAddress1   *string  `json:"postal_address1,omitempty"`
	PostalAddress2   *string  `json:"postal_address2,omitempty"`
	PostalCountryID  *string  `json:"postal_country_id,omitempty"`
	CustomField1     *string  `json:"custom_field_1,omitempty"`
	CustomField2     *string  `json:"custom_field_2,omitempty"`
	CustomField3     *string  `json:"custom_field_3,omitempty"`
	CustomField4     *string  `json:"custom_field_4,omitempty"`
}

// Customers grabs and collates all customers in pages of 10,000.
func (c Client) Customers() ([]Customer, error) {

	customers := []Customer{}
	page := []Customer{}

	// v is a version that is used to get customers by page.
	// Here we get the first page.
	data, v, err := ResourcePage(0, c.DomainPrefix, c.Token, "customers")

	// Unmarshal payload into sales object.
	err = json.Unmarshal(data, &page)
	if err != nil {
		log.Printf("error while unmarshalling: %s", err)
	}

	customers = append(customers, page...)

	for len(page) > 0 {
		page = []Customer{}

		// Continue grabbing pages until we receive an empty one.
		data, v, err = ResourcePage(v, c.DomainPrefix, c.Token, "customers")
		if err != nil {
			return nil, err
		}

		// Unmarshal payload into customer object.
		err = json.Unmarshal(data, &page)

		// Append customer page to list of customers.
		customers = append(customers, page...)
	}

	return customers, err
}