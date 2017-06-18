package main 

import (
    "net/http"  
    "github.com/gorilla/mux"   
    "fmt" 
    "strconv"
    "errors"
)
const  FIBITEMS int = 94
var    gFibNumbers[FIBITEMS]uint64  
const  MAXNUMBER    uint64 =  18446744073709551615  
const  MAXFIBNUMBER uint64 =  12200160415121876738
                                   
func main() {
	 
	r := mux.NewRouter()
	 
    r.HandleFunc("/fibonacci", fibHandler).Methods("GET")  
    r.HandleFunc("/", defaultHandler)   
    http.Handle("/", r)	   
    
    // load all fibonacci number and store in global memory     
    setupFibNumbers() 
      
    fmt.Printf("Listening on port 8080") 
    http.ListenAndServe(":8080", nil)
}
//-----------------------------------------------------
// load all fibonacci number and store in global memory 
//-----------------------------------------------------    
func setupFibNumbers() error {
	  
	  // There are 94 fibonacci numbers that will fit into an uint64. 
	  // This routine will load them into memory.
	  for i := 0;i< FIBITEMS ;i++{
	  	if(i==0){
	  		gFibNumbers[i]=0
	  	}else if(i==1){
	  		gFibNumbers[i]=1  			
	  		
	  	}else{		 
	  		gFibNumbers[i] = gFibNumbers[i-1] + gFibNumbers[i-2]
	  		if(gFibNumbers[i] >= (MAXFIBNUMBER)){	
	  			break
	  		} 
	  	}
	  }
  return nil
}
//---------------------------------------------------------------
// This method is responsible for setting up the return of
// fibonacci numbers based from the user input.
//---------------------------------------------------------------
func buildFib(inputNum uint64) ([]byte,error) {
      
	var retStr string	
	retStr = "["
	 
	if ( inputNum < 0 || inputNum > MAXNUMBER){
       return []byte(""),errors.New("Invalid Input")
	}	
	if (inputNum > 0) {
		retStr += "0"
	}
	for i:=1;i<FIBITEMS;i++{
		tmpNum := strconv.FormatUint(gFibNumbers[i], 10)
		if(gFibNumbers[i] >= inputNum ){
			break
		}else{
		  retStr += "," + tmpNum
		}
	}
	retStr += "]"	 
    return []byte(retStr),nil
}

 
