package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lcrypto_aead_ascon80pqv12_ref
#include <stdlib.h>
#include "ascon80pq.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	//parameters from https://github.com/ascon/ascon-c/blob/main/tests/demo.c
	n := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	k := [20]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	a := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	m := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}		
	var len uint64 = 16	

	fmt.Println("Ascon-80pq binding test --------------------------------- ")
  	clen := C.ulonglong(16)
  	mlen := C.ulonglong(len)
  	alen := C.ulonglong(8)  	  
  	zerobyte := C.uchar(0)
  	
  	//pointer conversions
  	cbytes := (*C.uchar)(C.CBytes([]byte{}))
  	mbytes := (*C.uchar)(C.CBytes(m[:]))
  	abytes := (*C.uchar)(C.CBytes(a[:]))
  	nbytes := (*C.uchar)(C.CBytes(n[:]))
  	kbytes := (*C.uchar)(C.CBytes(k[:]))
  	
  	//Encrypt
	ret := C.crypto_aead_encrypt(cbytes, &clen, mbytes, mlen, abytes, 
								 alen, &(zerobyte), nbytes, kbytes)

	fmt.Printf("\tMessage:  %v \n", m)
	fmt.Printf("Encryption return code (0 is success): %d\n", ret)

	//Decrypt in ptr	
	ptr := C.malloc(C.sizeof_char * C.ulong(len))
	defer C.free(unsafe.Pointer(ptr))
	ret = C.crypto_aead_decrypt((*C.uchar)(ptr), &mlen, &(zerobyte), cbytes, 
								clen, abytes, alen, nbytes, kbytes);
	
	if (ret == 0) {		
		fmt.Print("\tDecrypted:")
		b := C.GoBytes(ptr, C.int(len))
		fmt.Println((b))
	}else{
		fmt.Printf("Decryption code error: %d\n", ret)
	}	
}
