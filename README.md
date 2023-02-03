# Sam

A Command Line Interface to Hobbit service in Go. The app name comes (of course)
from [Samwise "Sam" Gamgee](https://en.wikipedia.org/wiki/Samwise_Gamgee), was Bag End's gardener, at Hobbiton.
[Hobbit](https://github.com/pjover/hobbit) is a Kotlin Spring boot application for managing a Kindergarten business,
initially developed for [Hobbiton](http://www.hobbiton.es/) Kindergarten.

Sam will:

1. Be an easy way to interact with Hobbit service, mapping the main operations to simple commands
1. All the business logic remains at Hobbit
1. Control the monthly workflow with Hobbit
1. Generate the files locally, invoices (PDFs) and reports (Excel)

## Install

### Database

To install and run a MongoDb server and a mongo-express web interface run one of these scripts depending on your
architecture:

- For AMD-64 architecture: `docker compose -f docker-compose-amd64.yaml up --detach`
- For ARM-64 architecture: `docker compose -f docker-compose-arm64.yaml up --detach`

mongo-express will be accessible from http://localhost:8081

### Config

Copy and adapt these files to `~/.sam/`:

- ./configs/sam.yaml
- ./configs/new_customer.json
- ./configs/new_product.json

## Sam

Install

```
cd cmd/sam
go install
```

Test

```
go test ./...
```

Format

```
go fmt ./...
```


## Generate mocks with mockery

1. cd to the ports directory: `cd internal/domain/ports`
2. Run mockery for the interface you want to mock `mockery --name=DbService`