package main


import (
	"fmt"

	"errors"
)

const backoffAttempts = 10
type retryable func()error

func main(){

	a, b := 0, 0

	res:=0

	f := func ()(error){
		var err error
		res, err = myFunction(a, b)
		if err != nil{
			return err
		}
		return nil
	}

	err:= withSimpleRetry(f)
	if err != nil{
		fmt.Println(err)
	}






}

func withSimpleRetry(r retryable)error{

	for i:=1; i<= backoffAttempts; i++{
		if err := r(); err != nil{
			fmt.Println("received error:", err, "\n retrying... attempt", i)
			continue
		}
		// TODO: sleep for rand interval between 0 and 5 sec
	}

}

func myFunction (a, b int)(int, error){
	return 0, errors.New("simulated error")
}
