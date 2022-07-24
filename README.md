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

## How to

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

## Sample config

Copy and adapt these files to `~/.ssm/`:

- docs/sam.yaml
- docs/new_customer.json
- docs/new_product.json