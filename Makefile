dev:
	gin --appPort 8090 -i run main.go


check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models	

serve-swagger: check-swagger
	swagger serve -F=swagger --port=35081 --no-open swagger.yaml	