package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go/charge"
	"os"
	"encoding/json"
	"github.com/stripe/stripe-go"
)

type BodyRequest struct {
	RequestStripeToken string `json:"stripeToken"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {


	bodyRequest := BodyRequest{
		RequestStripeToken: "",
	}
	fmt.Println("I don't know how to handle the following body:")
	fmt.Println(request.Body)
	err := json.Unmarshal([]byte(request.Body), &bodyRequest)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "BAD REQUEST!!",
		}, nil
	}
	fmt.Println(bodyRequest.RequestStripeToken)


	for k, v := range  request.MultiValueHeaders{
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	stripe.Key = os.Getenv("STRIPE_KEY")
	fmt.Println("Stripe Key: " + stripe.Key)

	// Token is created using Checkout or Elements!
	// Get the payment token ID submitted by the form:
	token :=
		"some junk value"
		//r.FormValue("stripeToken")

	params := &stripe.ChargeParams{
		Amount: stripe.Int64(999),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Example charge"),
	}
	fmt.Println("hi")
	params.SetSource(token)
	ch, _ := charge.New(params)
	fmt.Println(ch)


	fullLineupData := "faux payment response"

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fullLineupData,
	}, nil
}

func main() {
	lambda.Start(handler)
}

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys

