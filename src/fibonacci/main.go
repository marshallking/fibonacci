package main 

import (
    "net/http"  
    "github.com/gorilla/mux"   
    "fmt" 
    "strconv"
    "errors"
)
const  FIBITEMS int = 94
const  FIBITEMSDISPLAY = 95
var    gFibNumbers[FIBITEMS]uint64  
const  MAXNUMBER    uint64 =  18446744073709551615  
const  MAXFIBNUMBER uint64 =  12200160415121876738

var mFib map[int]string

type CompareFunc func(int)CompareResult

type CompareResult int

const (
    Smaller CompareResult = -1
    Equal   CompareResult   = 0
    Larger  CompareResult  = 1
)

                               
func main() {
	 
	r := mux.NewRouter()
	 
    r.HandleFunc("/fibonacci", fibHandler).Methods("GET")  
    r.HandleFunc("/", defaultHandler)   
    http.Handle("/", r)	   
    
    // load all fibonacci number and store in global memory     
    setupFibNumbers()  
    setupFibMap()    
      
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
	  		if(gFibNumbers[i] == (MAXFIBNUMBER)){			
	  			break
	  		} 
	  	}
	  }
	 
  return nil
}
//-----------------------------------------------------------------
// This method will load all the fibonacci number into a map
// that will later be used for retuning the fibonacci numbers
// per user request through the API.
//-----------------------------------------------------------------
func setupFibMap() error {
	
	  var i int 
	  // There are 94 fibonacci numbers that will fit into an uint64. 
	  // This routine will load them into memory.
	  mFib = make(map[int]string, 100)
	  
	  for i = 0;i < FIBITEMS ;i++{	  
	  	byteReturn,err := buildFib(gFibNumbers[i])
	  	if(err!=nil){
	  		fmt.Printf("Error writing fibonacci value - %v",err)      		  		
	  	}else{
	  		mFib[i]=string(byteReturn[:])	  		
	  	}
	  }
	  //--------------------------------------------
	  // add the last item. Note: I add one to the 
	  // last fibonacci number so that it will be
	  // added to the last set of numbers. 
	  //--------------------------------------------
	  var uLastNum = gFibNumbers[FIBITEMS-1] + 1      
	  byteReturn,err := buildFib(uLastNum)	
	  if(err!=nil){
	  		fmt.Printf("Error writing fibonacci value - %v",err)      		  		
	  	}else{
	  		mFib[FIBITEMS]=string(byteReturn[:])	  
	  } 
  return nil
}


//-----------------------------------------------------------------
// This method will find the index into the map of fibonacci numbers.
//-----------------------------------------------------------------
func findIndex(inputNum uint64) int {
		  
	  fibIndex := FindFibIndex(FIBITEMS, func(index int) CompareResult {
	    if inputNum > gFibNumbers[index] {
	        return Smaller
	    } else if inputNum == gFibNumbers[index] {
	        return Equal
	    } else {
	        return Larger
	    }
	  })    
  return fibIndex
}
//-----------------------------------------------------------------
// This method is the binary compare for the fibonacci numbers.
// O(log(n))
//-----------------------------------------------------------------
func FindFibIndex(numArrayItems int, compare CompareFunc) int {
	startIndex, endIndex := 0, numArrayItems
	for startIndex < endIndex {
		h := startIndex + (endIndex-startIndex)/2
		result := int(compare(h))
		if result < 0 {
			startIndex = h + 1
		} else if result == 0 {
			return h
		} else {
			endIndex = h
		}
	}
	return startIndex
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
	for i:=1;i< int(FIBITEMS);i++{
		tmpNum := strconv.FormatUint(gFibNumbers[i], 10)
		if(gFibNumbers[i] >= inputNum ){
			if(gFibNumbers[i] < MAXFIBNUMBER){
			    break
			}
		}else{
		  retStr += "," + tmpNum
		}
	}
	   	
	retStr += "]"	 
    return []byte(retStr),nil
}

func getFib(inputNum uint64) ([]byte,error) {
    	 
    return []byte(mFib[findIndex(inputNum)]),nil
}

 

