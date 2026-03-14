module calculator_api

go 1.26.1

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/google/uuid v1.6.0 // indirect
	shared v0.0.0
)

require (
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.49.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace shared => ../shared
