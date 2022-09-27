package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lcrypto_aead_ascon80pqv12_ref -lascongobench
#include <stdlib.h>
#include "../ascon80pq.h"
*/
import "C"
import (
	"fmt"	
)

func main() {
	fmt.Println("Ascon-80pq binding benchmark test --------------------------------- ")

	//Bench
	ret := C.bench(0)
	
	fmt.Printf("Return code: %d\n", ret)
	fmt.Println("End of test ------------------------------------------------------- ")
}

