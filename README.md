# Ascon-80pq go-binding
This repository offers Go bindings for the `ascon-80pq` C implementation. Ascon is one finalist of the Lightweight Cryptography standardization process conducted by NIST ([LWC](https://csrc.nist.gov/Projects/lightweight-cryptography/lwc-publications)). 

## Usage

To do.

## Clone instructions

In a single line: `git clone --recursive https://github.com/AAGiron/ascon80pq-go-binding`

Or use `cd ascon80pq-go-binding && git submodule update --init --recursive` after cloning.

## Ascon-c compiling

In order to compile the required implementations for the binding, use:

```
mkdir build && cd build
#cmake .. -DALG_LIST="ascon80pq" -DIMPL_LIST="ref" -DTEST_LIST="genkat"
cmake .. -DALG_LIST="ascon80pq" -DIMPL_LIST="ref" -DTEST_LIST="genkat" -DBUILD_SHARED_LIBS=ON
cmake --build .
```
