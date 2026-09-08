package main

// Methods are useful when the behaviors are associated with the type
import (
	"fmt"
)

type SMS struct {
	Message string `json:"message"`
}

// Expense interface represents anything that has a cost
type expense interface {
	Cost() float64
}

func (s SMS) Cost() float64 {
	const costPerChar = 0.002
	return costPerChar * float64(len(s.Message))
}

// Method to send sms after validating its length
func (s SMS) Send() (float64, error) {

	const maxMessageLength = 50

	if len(s.Message) > maxMessageLength {
		return 0.0, fmt.Errorf("message too long, can't send text over %v characters", maxMessageLength)
	}

	return s.Cost(), nil
}

func SendSMSToCouple(msgToCustomer string, msgToSpouse string) (customerCost float64, spouseCost float64, err error) {

	customerSMS := SMS{
		Message: msgToCustomer,
	}

	spouseSMS := SMS{
		Message: msgToSpouse,
	}

	customerCost, err = customerSMS.Send()
	if err != nil {
		return 0.0, 0.0, err
	}

	spouseCost, err = spouseSMS.Send()
	if err != nil {
		return 0.0, 0.0, err
	}

	return customerCost, spouseCost, nil

}

func main() {

	customerMessage := "Your package has arrived"
	spouseMessage := "Hey baby I am coming home"

	// customerMessage = "Hey Stacy, I have a client I would like to refer to you, she wants a 1kg vanilla cake"

	customerCost, spouseCost, err := SendSMSToCouple(customerMessage, spouseMessage)

	if err != nil {
		fmt.Println(err)
		return
	}

	totalCost := customerCost + spouseCost

	fmt.Println("=============================================")
	fmt.Printf("Customer SMS Cost : %.4f\n\n ", customerCost)
	fmt.Printf("Spouse SMS Cost : %.4f\n\n ", spouseCost)
	fmt.Printf("Total SMS Cost : %.4f\n\n ", totalCost)

}
