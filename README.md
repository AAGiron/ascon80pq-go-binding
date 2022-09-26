# Ascon-80pq go-binding
This repository offers a very simple Go binding for the `ascon-80pq` C implementation (aead_encrypt/decrypt operations). Ascon is one finalist of the Lightweight Cryptography standardization process conducted by NIST ([LWC](https://csrc.nist.gov/Projects/lightweight-cryptography/lwc-publications)). 

## Clone instructions

In a single line: `git clone --recursive https://github.com/AAGiron/ascon80pq-go-binding`

Or use `cd ascon80pq-go-binding && git submodule update --init --recursive` after cloning.

## Ascon-c building process

In order to compile the required implementations for the binding, use:

```
cd ascon-c
mkdir build && cd build
cmake .. -DALG_LIST="ascon80pq" -DIMPL_LIST="ref" -DTEST_LIST="genkat" -DBUILD_SHARED_LIBS=ON
cmake --build .
```

Install it (and export) to `LD_LIBRARY_PATH` (e.g., `sudo cp *.so /usr/local/lib/` if that is your path).

## Usage

After cloning and building Ascon (above), go back `cd ..` to main directory and then test the binding using `go run test.go`. 


## Benchmarking `ascon-80pq`

We also provide a benchmark in an attempt to evaluate the cost of such a binding (C to Go). It is almost identical to [getcycles.c](https://github.com/ascon/ascon-c/blob/main/tests/getcycles.c), with small modifications for building the benchmark. (We might provide results in the future).

The benchmark requires the above build process, a `cd ..`  and the following steps:
```		
	cd bench/ 
	cmake  -DBUILD_SHARED_LIBS=ON .
	cmake --build .
	sudo cp libascongobench.so /usr/local/lib/
```

Run the benchmark using `go run bench.go`, it calls the benchmarking c-functions for `ascon-80pq` and displays the results.

## Disclaimer

This is a very simple and third-party binding implementation: it does not offer warranty of any kind (nor any security guarantees).
