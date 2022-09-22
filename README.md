# Ascon-80pq go-binding
This repository offers a very simple Go binding for the `ascon-80pq` C implementation (aead_encrypt/decrypt operations). Ascon is one finalist of the Lightweight Cryptography standardization process conducted by NIST ([LWC](https://csrc.nist.gov/Projects/lightweight-cryptography/lwc-publications)). 

## Clone instructions

In a single line: `git clone --recursive https://github.com/AAGiron/ascon80pq-go-binding`

Or use `cd ascon80pq-go-binding && git submodule update --init --recursive` after cloning.

## Ascon-c building process

In order to compile the required implementations for the binding, use:

```
mkdir build && cd build
cmake .. -DALG_LIST="ascon80pq" -DIMPL_LIST="ref" -DTEST_LIST="genkat" -DBUILD_SHARED_LIBS=ON
cmake --build .
```

Install it (and export) to `LD_LIBRARY_PATH` (e.g., `sudo cp *.so /usr/local/lib/` if that is your path).

## Usage

After cloning and building Ascon (above), test the binding using `go run test.go`. 


## Disclaimer

This is a very simple and third-party binding implementation: it does not offer warranty of any kind (nor any security guarantees).