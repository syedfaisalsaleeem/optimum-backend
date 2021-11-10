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

## DB Migrations
- After cloning the repository follow these steps for migration
- make a database in postgresql using pgadmin or using psql `Create DATABASE test`
- update the .env file with DB_USER, DB_PASSWORD, DB_NAME, DB_HOST
- Install all the go package in go.mod using `go get ./`
- go to migrations folder using cd migrations
- `go run .` command will run the migrations in the table

## How To Run DB On LocalHost
- Install postgresql in the operating system
- The postgresql db will start running on port 5432
- Using pgadmin the database tables can be visualized

## How to setup db on LocalHost
- Using pgadmin the db can be setup on localhost
- update .env file to connect with db on local host
- use docker for running the db on cloud

## Available Scripts

In the project directory, you can run:

### `go run .`

Runs the app in the development mode at [http://localhost:8080]


### Testing

Test the application using `go test . -v`

Launches the test runner in the interactive watch mode.\
All the test cases are described below:
- RUN   TestNew
- PASS: TestNew (0.00s)     
- RUN   Test_testAPI        
- PASS: Test_testAPI (0.00s)
- RUN   TestDBConnection    
- PASS: TestDBConnection (0.09s)
- RUN   TestEmptyTable
- PASS: TestEmptyTable (0.01s) 
- RUN   TestAddTodoList        
- PASS: TestAddTodoList (0.00s)
- RUN   TestGetTodoListitems   
- PASS: TestGetTodoListitems (0.00s)
- RUN   TestPostEmptyItem
- PASS: TestPostEmptyItem (0.00s)
- RUN   TestPostJsonFormat
- PASS: TestPostJsonFormat (0.00s)
- RUN   TestGettodolistjson
- PASS: TestGettodolistjson (0.00s)
- RUN   TestLenoftodolistafteradding
- PASS: TestLenoftodolistafteradding (0.00s)

## Running Rest Api through Docker
The rest apis are dockerized and they can be run through these commands:
- before building the docker image update the .env file
- `docker-compose up -d --build`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Authors
- Syed Faisal Saleem 

