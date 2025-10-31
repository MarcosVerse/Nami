# Nome do arquivo principal
MAIN=cmd/main.go

# Roda o servidor Go
run:
	go run $(MAIN)
	
# Gera a documentação Swagger
swagger:
	swag init -g $(MAIN)

# Gera o Swagger e roda o servidor
swaginit: swagger run
