# GOLANG Todo Application
This a todo application in which we can add items in a list using button. This application is created using purely Test Driven Development. This application stores the added data in postgresql
in the backend. 

## Getting Started
To get started using this package follow the instructions below.

## Dependencies
- go lang
- postgresql

## Installation
- git clone https://github.com/syedfaisalsaleeem/optimum-backend
- cd optimum-backend
- go mod download

## Available Scripts

In the project directory, you can run:

### `go run index.go`

Runs the app in the development mode at [http://localhost:8010]


### Testing

1) Run the go application using `go run index.go`
2) In the second terminal test the application using `go test index.go index_test.go -v`

Launches the test runner in the interactive watch mode.\
All the test cases are described below:
RUN   TestNew
PASS: TestNew (0.00s)     
RUN   Test_testAPI        
PASS: Test_testAPI (0.00s)
RUN   TestDBConnection    
PASS: TestDBConnection (0.09s)
RUN   TestEmptyTable
PASS: TestEmptyTable (0.01s) 
RUN   TestAddTodoList        
PASS: TestAddTodoList (0.00s)
RUN   TestGetTodoListitems   
PASS: TestGetTodoListitems (0.00s)


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Authors
- Syed Faisal Saleem 

