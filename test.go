//
// main.go
// Copyright (C) 2019 Tim Hughes
//
// Distributed under terms of the MIT license.
//
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
	//name := C.CString("Test")
	//defer C.free(unsafe.Pointer(name))

	//year := C.int(2022)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	/*unsigned char n[CRYPTO_NPUBBYTES] = {0, 1, 2,  3,  4,  5,  6,  7,
                                       8, 9, 10, 11, 12, 13, 14, 15};
  unsigned char k[CRYPTO_KEYBYTES] = {0, 1, 2,  3,  4,  5,  6,  7,
                                      8, 9, 10, 11, 12, 13, 14, 15};
  unsigned char a[16] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15};
  unsigned char m[16] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15};
  unsigned char c[32], h[32], t[32];
  unsigned long long alen = 8;
  unsigned long long mlen = 8;
  unsigned long long clen = CRYPTO_ABYTES;
  int result = 0;
*/
	n := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	k := [20]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	//k := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	a := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	m := [16]byte{0, 1, 2,  3,  4,  5,  6,  7, 8, 9, 10, 11, 12, 13, 14, 15}
	c := []byte{}
	//h := []byte{}
	//t := []byte{}
	var alen uint64 = 8
	var mlen uint64 = 8
	var clen uint64 = 16

	fmt.Println("Ascon-80pq binding test ---------------------- ")
/*
result |= crypto_aead_encrypt(c, &clen, m, mlen, a, alen, (void*)0, n, k);
  print('c', c, clen - CRYPTO_ABYTES);
  printf(" ");
  print('t', c + clen - CRYPTO_ABYTES, CRYPTO_ABYTES);
  printf(" -> ");
  result |= crypto_aead_decrypt(m, &mlen, (void*)0, c, clen, a, alen, n, k);
  print('a', a, alen);
  printf(" ");
  print('m', m, mlen);
  printf("\n");

  */
  	pointClen := C.ulonglong(clen)
  	pointMlen := C.ulonglong(mlen)
  	sliceM := m[:]
  	sliceA := a[:]
  	sliceN := n[:]
  	sliceK := k[:]
  	zero := C.uchar(0)
  	
  	//pointers
  	cp := (*C.uchar)(C.CBytes(c))
  	sliceMp := (*C.uchar)(C.CBytes(sliceM))
  	sliceAp := (*C.uchar)(C.CBytes(sliceA))
  	sliceNp := (*C.uchar)(C.CBytes(sliceN))
  	sliceKp := (*C.uchar)(C.CBytes(sliceK))
  	//Encrypt
	ret := C.crypto_aead_encrypt(cp, &pointClen, 
			sliceMp, C.ulonglong(mlen), 
			sliceAp, C.ulonglong(alen), &(zero), 
			sliceNp, sliceKp)

	//fmt.Println(cp)
	fmt.Print("Message: ") //+ string(sliceM))
	fmt.Printf("%v \n", m)
	fmt.Print("Encryption code return (0 is success):")
	fmt.Println(ret)

	//Decrypt
	
	ret = C.crypto_aead_decrypt( (*C.uchar)(ptr), &pointMlen, 
	 							&(zero), cp, pointClen,  
	 							sliceAp, C.ulonglong(alen),
	 							sliceNp, sliceKp);
	

	if (ret == 0) {
		//decrypted := C.GoBytes(ptr, sliceM) //Not ret
		fmt.Print("Decrypted message:")
		//fmt.Println(string(sliceA))
		//fmt.Println(string(sliceM))	
		//
		b := C.GoBytes(ptr, 16)

		fmt.Println((b))
	}else{
		fmt.Print("Decryption code error: ")
		fmt.Println(ret)
	}
	
}


//PS: cgo LDFLAGS: -Lascon/build2 -llibcrypto_aead_ascon80pqv12_ref could be -L${SRCDIR}/libs or something
//${SRCDIR}/ascon/build2
//
//cgo LDFLAGS: -L. -lcrypto_aead_ascon128v12_ref