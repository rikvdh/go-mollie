package mollie

import (
	"testing"
	"time"
	"math/rand"
)

func TestCustomersCreateAndGetError(t *testing.T) {
	mollieApi := Get("test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU")

	rand.Seed(time.Now().Unix())
	id := rand.Intn(999999 - 1) + 1

	userEmail := "user_" + string(id)

	data := NewCustomerData{
		Name: "test_user",
		Email: userEmail,
		Locale: "nl_NL",
	}

	customer, err := mollieApi.Customers().New(data)

	if err == nil {
		t.Errorf("Error is nil: %v", err)
	}

	if customer != nil {
		t.Errorf("Something may went wrong unmarchalling the customer\n")
	}

	// get that user
	q, err := mollieApi.Customers().Get("user1")
	if err == nil {
		t.Errorf("Error is not nil\n")
	}
	if q != nil {
		t.Errorf("Something may went wrong unmarchalling the customer\n")
	}
}
