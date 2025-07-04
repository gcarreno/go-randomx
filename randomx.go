package randomx

/*
#cgo CFLAGS: -IRandomX/src
#cgo LDFLAGS: -LRandomX/build -lrandomx
#include "randomx.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type RandomX struct {
	cache *C.randomx_cache
	vm    *C.randomx_vm
}

// Initialize RandomX with a seed key
func NewRandomX(seed []byte) (*RandomX, error) {
	// Create flags for RandomX
	flags := C.randomx_get_flags()

	// Create cache for RandomX
	cache := C.randomx_alloc_cache(flags)
	if cache == nil {
		return nil, errors.New("failed to allocate RandomX cache")
	}

	// Initialize cache with seed
	C.randomx_init_cache(cache, unsafe.Pointer(&seed[0]), C.size_t(len(seed)))

	// Allocate VM
	vm := C.randomx_create_vm(flags, cache, nil)
	if vm == nil {
		C.randomx_release_cache(cache)
		return nil, errors.New("failed to create RandomX VM")
	}

	return &RandomX{
		cache: cache,
		vm:    vm,
	}, nil
}

// Compute a RandomX hash
func (r *RandomX) ComputeHash(input []byte) ([]byte, error) {
	output := make([]byte, 32)
	C.randomx_calculate_hash(r.vm, unsafe.Pointer(&input[0]), C.size_t(len(input)), unsafe.Pointer(&output[0]))
	return output, nil
}

// Clean up resources
func (r *RandomX) Close() {
	if r.vm != nil {
		C.randomx_destroy_vm(r.vm)
	}
	if r.cache != nil {
		C.randomx_release_cache(r.cache)
	}
}
