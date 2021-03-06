package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}
/* adding a fibHandler comment holder*/
func fibHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	infibNum := r.URL.Query().Get("fib")
	fibNum, err := strconv.ParseUint(infibNum, 10, 64)
	//---------------------------------------------------------------------
	// lets do some input validation here
	//---------------------------------------------------------------------
	if err != nil {
		w.WriteHeader(400)
		_, err = w.Write([]byte("Invalid Input"))
		if err != nil {
			fmt.Printf("Error writing - %v", err)
		}
		return
	}
	//---------------------------------------------------------------------
	// lets build the return here
	//---------------------------------------------------------------------
	fibRet, err := getFib(fibNum)

	if err == nil {
		_, err = w.Write(fibRet)
		if err == nil {
			w.WriteHeader(200)
		} else {
			fmt.Printf("Error writing - %v", err)
		}
	} else {
		w.WriteHeader(400)
	}
	return
}
