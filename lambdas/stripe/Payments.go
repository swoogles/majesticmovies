package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stripe/stripe-go/charge"
	//"os" use to get stripe.Key
	"github.com/stripe/stripe-go"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	for k, v := range  request.MultiValueHeaders{
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	stripe.Key = "SOME FAKE KEY HERE"

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

