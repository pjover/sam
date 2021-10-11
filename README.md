# Sam

A Command Line Interface to Hobbit app in Go. The app name comes (of course) from [Samwise "Sam" Gamgee](https://en.wikipedia.org/wiki/Samwise_Gamgee), was Bag End's gardener, at Hobbiton.
[Hobbit](https://github.com/pjover/hobbit) is a Kotlin Spring boot application for managing a Kindergarten business, initially developed for [Hobbiton](http://www.hobbiton.es/) Kindergarten.

Sam will:

1. Be an easy way to interact with Hobbit service, mapping the main operations to simple commands
2. Control the monthly workflow with Hobbit
3. Generate the files locally, invoices (PDFs) and reports (Excel)

## Commands

### Working directory

`sam dir [--month=YYYY-MM]`

Creates the working directory for a month:

- if month is not specified, takes the current month
- changes the current directory to the newly created directory
- creates the configuration file `sam.yaml` inside the created directory


### Add child consumption

1. The child's account can be modified with `add` and `del`
2. Once the consumptions are ok, verified with `read`
3. Then all children's pending consumptions are moved to child's invoices account to generate a single invoice, with `move`
4. And finally generate the invoice PDF as many times you want, with `invoice`

Adds a consumption to the child's account.

`sam add "QME, 0.5 MME, MHX, 1.5 HEX" -c 1520 -n "This is a note" [--bdd | --cash | --voucher | transfer]`


### Delete child consumption

Delete a consumption from the child's consumption account

`sam del "0.5 MME" -c 1520`


### Read child

Reads all consumptions from the child's consumption account

`sam read -c 1520`


### Move child consumption to invoice account

Moves all consumptions from the child's consumption account to the child's invoices account
Generate a single invoice for all pending consumptions in the child's invoices account

`sam move -c 1520`


### Generate the invoice PDF file

If no client is selected, it will create all invoices for the current month

`sam invoice [-c 1520]`


### Generate the Bank Direct Debit file

Generates the Bank Direct Debit file `bdd.q1x` for all BDD invoices for the current month

`sam bdd`


### Generate the month report

Generates the spreadsheet `Month report.xlsx` with all invoices for the current month

`sam report`