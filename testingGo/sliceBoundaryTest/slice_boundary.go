//You will build a mock system component that tracks an array of Error IDs. 
// If the tracking buffer fills up past its allocated capacity boundary, your code must intercept the failure, 
// track the underlying memory address shift, and return a custom error struct that implements the native Go error interface

package main

import "fmt"

type Overflow struct{
	CurrentLen int
	MaxCap int
	Syscall string
}

func (e Overflow) Error() string{
	return fmt.Sprintf("System Error: [%s] failed. Buffer length %d reached max capacity %d", e.Syscall, e.CurrentLen, e.MaxCap)
}

func AppendErrorMetrics(buffer[] int, errorID int)([]int, error){
    fmt.Printf("func address is:%p\n",buffer)

	if(len(buffer)== cap(buffer)){
            return buffer, Overflow{
                CurrentLen: len(buffer),
				MaxCap:     cap(buffer),
				Syscall:    "write",				
			}
	}
	buffer=append(buffer,errorID)
	return buffer,nil
}

func main(){

	slice1 := make([]int,0,3)
	var err error
	fmt.Printf("main adress : %p\n",slice1)
	for i:=1;i<=4;i++{
		slice1, err =AppendErrorMetrics(slice1,i*100)
		if(err!=nil){
			fmt.Println("-----------------------------------------------")
			fmt.Println("OVERFLOW DETECTED")
			fmt.Println("eror details:",err)
			if overflowErr, ok := err.(Overflow); ok {
					fmt.Printf("Trapped Field  : Found failed syscall target -> %s\n", overflowErr.Syscall)
			}
	    }

	}

}