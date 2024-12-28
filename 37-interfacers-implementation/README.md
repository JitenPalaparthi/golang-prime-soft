mockgen -source=shape/shape.go -destination=mocks/mock_shape.go -package=mocks

- to test everything
go test ./...

- to test a package

go test shapes/shape/square

- to test only one test
go test -timeout 30s -run ^TestWhat$ shapes/shape/square

- go test with cover profile

go test shapes/shape/square --cover
profile=coverage.out  

go tool cover -html=coverage.out

to run go generate

go generate ./...