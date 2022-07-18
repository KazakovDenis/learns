COVERAGE=coverage.out

install:
	go get github.com/stretchr/testify

test:
	go test foo -cover -coverprofile=${COVERAGE}

coverage:
	go tool cover -html=${COVERAGE}
