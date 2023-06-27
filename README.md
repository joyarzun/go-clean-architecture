# Go clean architecture Holidays

Project to show clean architecture in Go using a Holidays backend example

## Requirements

```bash
  brew install go
  brew install sqlite
```

## Usage

**Create a SQLlite table on _gorm.db_ file**

```bash
sqlite3 gorm.db 'CREATE TABLE holidays ( id INTEGER PRIMARY KEY, year INTEGER NOT NULL, name TEXT NOT NULL, date TEXT NOT NULL );'
```

**(Optional) Create a holiday record**

```bash
sqlite3 gorm.db "INSERT INTO holidays VALUES (1, 2024, 'Example', '2024-01-01 00:00:00+00:00')"
```

**Run the example**

```bash
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

## Clean architecture

[Clean Architecture was presented by Uncle Bob (Robert C. Martin) on 2012](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) as a compilation of different ideas, like Hexagonal Architecture, Onion Architecture, etc, that produce similar advantages.

![](img/clean-architecture.jpeg)

We will not explore the details of the architecture here, but we will explain how this example follow the clean diagram.

Inside `src` folder we are separating folder following [domain groups idea](https://www.youtube.com/watch?v=y3MWfPDmVqo&t=905s). This tiny example only have one domain (holiday). Each domain folder have a Clean Architecture structure like:

- `entities`: Contains the business rules and entities. Holiday entity is compound of a Name and the holiday Date
- `usecases`: Contains the exposed features or use cases that the business need. For example, it is exposed the Holiday Creation and Retrieve
- `interfaceadapter`: Contains the http controllers and its own details that use the features exposed by `usecases`
- `registry`: Part of _frameworks and drivers_. Contains the details to generate the web controller
- `infra`: Contains the details of DB and Router implementation

## TODO

- `interfaceadapter` controller have dependency on `echo.Context` that is part of the web framework

## Based on

- [Clean Architecture by Uncle Bob (Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Clean Architecture with go](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31)
- [Clean architecture by domains](https://www.youtube.com/watch?v=y3MWfPDmVqo&t=905s)

## FAQ

### ginkgo doesn't found

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
