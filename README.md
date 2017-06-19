# finbonacci
Web service to return fibonacci numbers

Update to Initial Requirements:

Hi Mark – Questions:
Does the number given have to be a Fibonacci number? For example, can the input b the number 4? YES
Is there a maximum input number for this project?  The maximum for this project will be the uint64 maximum
Are there any time constraints for processing? 
What should the input of the number 0 return? Will return an "[]"
Is it acceptable for me to define the appropriate errors? Generally when I design a Restful API … error codes are defined. 
So I’m assuming that the error codes I define will be OK for this project. 
   The return codes will be:
                  200 - status ok
                  400 - Invalid Request
 
 
Initial Requirements:
 
Please provide a sample project for review:
1. The project should provide a RESTful web service.
  a. The web service accepts a number, n, as input and returns the first n Fibonacci numbers, starting from 0. I.e. given n = 5, 
     appropriate output would represent the sequence [0, 1, 1, 2, 3].
  b. Given a negative number, it will respond with an appropriate error.
2. Include whatever instructions are necessary to build and deploy/run the project, where "deploy/run" 
   means the web service    is accepting requests and responding to them as appropriate.
3. Include some unit and/or functional tests
4. Use any language that you know well

While this project is admittedly trivial, approach it as representing a more complex problem that you'll
have to put into production and maintain for 5 years. Providing a link to a github/bitbucket repo with the
project would probably be the easiest way to submit. Do the best that you can and please reply back to all
with a link to your results when it is completed.

TO BUILD THIS PROJECT:

1. download the code.
2. go get github.com/gorilla/mux
3. go build *.go

TO DEPLOY/RUN:

    1. ./fibHandler

API:

    /fibonacci

PARAMS:

    fib -- Any positive number from 0 to MAX uint64 (12200160415121876738)

RETURNS - All fibonacci numbers from 0 to input number. Note: will exclude the input number.

Example:

    curl -s -v 'http://127.0.0.1:8080/fibonacci?fib=5'

    returns [0,1,1,2,3]

    curl -s -v 'http://127.0.0.1:8080/fibonacci?fib=-1'

    returns 400
    
    curl -s -v 'http://127.0.0.1:8080/fibonacci?fib=abc'

    returns 400
    
    curl -s -v 'http://127.0.0.1:8080/fibonacci?fib=12200160415121876739'

    returns 400



INCLUDED FILES IN THIS PROJECT

main.go                - MAIN program
fibHandler.go          - API Handler 
fibonacci_test.go      - Unit and Functional Tests.

TO RUN TESTS

1. go test -v







