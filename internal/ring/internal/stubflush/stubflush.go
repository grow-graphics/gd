//go:build cgo

// Package stubflush provides an empty definition of the C ring-flush symbol
// so that the ring package's unit tests can link without pulling in the
// engine (gd.c and the startup callbacks). It must only ever be imported by
// tests that substitute the ring's dispatch function: importing it alongside
// the real engine would define the symbol twice.
package stubflush

// #include <stdint.h>
// #include <stddef.h>
// void gd_ring_flush(void *entries, uint32_t tail, uint32_t head, uint32_t *crash_index) {}
// void *gd_ring_flush_g0_addr(void) { return NULL; }
// void gd_ring_flush_main(void *entries, uint32_t tail, uint32_t head, uint32_t *crash_index) {}
// void *gd_ring_flush_main_g0_addr(void) { return NULL; }
// void gd_ring_adopt(void *ring, uint32_t *crash_index, void *threads_shared, void *threads_entries) {}
import "C"
