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


func fibHandler(w http.ResponseWriter, r *http.Request) {
	
	var err error
	 	 
	infibNum := r.URL.Query().Get("fib")	
	fibNum, err := strconv.ParseUint(infibNum, 10, 64)
	//---------------------------------------------------------------------
	// lets do some input validation here
	//---------------------------------------------------------------------	
	if (err != nil) {
		w.WriteHeader(400)			
		_,_ = w.Write([]byte("Invalid Input")) 	// int,error			 
	    return
	}
	//---------------------------------------------------------------------
	// lets build the return here
	//---------------------------------------------------------------------	
	fibRet,err := buildFib(fibNum)
	 
    if (err == nil){	   		 
	     _,_ = w.Write(fibRet)                  // int,error		
	     w.WriteHeader(200)	     
    }else{
    	 w.WriteHeader(400)	      	
    }
}