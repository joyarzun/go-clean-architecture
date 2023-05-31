# Go clean architecture Holidays

Project to show clean architecture in Go using a Holidays backend example

## Usage

**Create a SQLlite table on _gorm.db_ file**

```sql
CREATE TABLE holidays (
	id INTEGER PRIMARY KEY,
	year INTEGER NOT NULL,
	name TEXT NOT NULL,
	date TEXT NOT NULL
);
```

**Install go and run the example**

```bash
  brew install go
  go run main.go
```

**Create a new holiday**

```bash
  curl --request POST \
  --url http://localhost:3000/holiday \
  --header 'Content-Type: application/json' \
  --data '{
	"year": 2023,
	"name": "new year",
	"date": "2023-01-01T00:00:00Z"
}'
```

**Read holidays by year**

```bash
curl --request GET \
  --url http://localhost:3000/holiday/2023
```

## Main tasks

Principal tasks are included in the makefile:

- Test: `make test`
- Test with coverage: `make testcov`
- Open coverage html report: `make opencov`
- Show coverage result: `make showcov`
- Go vet, the golang analizer for suspicious constructs: `make govet`
- Staticcheck, state of the art linter for the Go: `make staticcheck`

**Most of these tasks will be running on the CI pipeline**

For use the _staticcheck_ you will need:

```bash
  go install honnef.co/go/tools/cmd/staticcheck@latest
```

## Based on

- [Clean Architecture with go](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31)
- [Clean architecture by domains](https://www.youtube.com/watch?v=y3MWfPDmVqo&t=905s)

## FAQ

### ginkgo doesn't found

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
