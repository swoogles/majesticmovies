build:
	hugo -v
	mkdir -p functions
	# go get ./...
	go build -o functions/movies-lambda  ./lambdas
	go build -o functions/submitStripeCharge  ./lambdas/stripe
