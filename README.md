# bookstore
A case study project to demonstrate an online bookstore based on Golang

## Generate API Documents

- Download [swaggo](https://github.com/swaggo/swag) by using: 
   ```bash
   $ go install github.com/swaggo/swag/cmd/swag@latest
   ```

## Make commands
- `make first`: makes project ready to run
- `make tidy`: installs project dependencies
- `make vet`: examines the source code and reports suspicious constructs that could be bugs or syntactical errors
- `make build`: builds project and generate binary in root dir
- `make run`: runs project
- `make doc`: generates api docs
- `make mock`: generates mock data (user and book)
- `make clean`: cleans project