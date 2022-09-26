cmake_minimum_required(VERSION 3.10)
project(ascon-go-bindings LANGUAGES C ASM)
set(DEFAULT_VERSIONS 1.0)
message("PS: Please make sure that ascon-c is already built.")

get_filename_component(PARENT_DIR ${CMAKE_SOURCE_DIR} DIRECTORY)
file(GLOB_RECURSE ${PARENT_DIR} "ascon-c/*.[chS]")

#creates benchmark lib
add_library(ascongobench bench.c)
target_include_directories(ascongobench PUBLIC ${PARENT_DIR}/ascon-c/crypto_aead/ascon80pqv12/ref/)
target_include_directories(ascongobench PUBLIC ${PARENT_DIR}/ascon-c/src/)
target_include_directories(ascongobench PUBLIC ${PARENT_DIR}/ascon-c/tests/)

#link libraries
find_library(LIBCRYPTO_AEAD_ASCON80PQV12_REF crypto_aead_ascon80pqv12_ref)
target_link_libraries(ascongobench PUBLIC ${LIBCRYPTO_AEAD_ASCON80PQV12_REF})
