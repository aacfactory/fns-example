package repository

import (
	"fmt"
	"math/rand"
	"unsafe"
)

const (
	m1 = 0xa0761d6478bd642f
	m2 = 0xe7037ed1a0b428db
	m3 = 0x8ebc6af09c88c6e3
	m4 = 0x589965cc75374cc3
	m5 = 0x1d8e4e27c47d124f
)

var hashkey [4]uintptr

const PtrSize = 4 << (^uintptr(0) >> 63)

func init() {
	for i := 0; i < 4; i++ {
		hashkey[i] = uintptr(rand.Int63())
	}
	hashkey[0] |= 1 // make sure these numbers are odd
	hashkey[1] |= 1
	hashkey[2] |= 1
	hashkey[3] |= 1
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func Hash(v interface{}) (s string) {
	p := memhashFallback(unsafe.Pointer(&v), 1, 36)
	s = fmt.Sprintf("%v", p)
	return
}

func memhashFallback(p unsafe.Pointer, seed, s uintptr) uintptr {
	var a, b uintptr
	seed ^= hashkey[0] ^ m1
	switch {
	case s == 0:
		return seed
	case s < 4:
		a = uintptr(*(*byte)(p))
		a |= uintptr(*(*byte)(add(p, s>>1))) << 8
		a |= uintptr(*(*byte)(add(p, s-1))) << 16
	case s == 4:
		a = r4(p)
		b = a
	case s < 8:
		a = r4(p)
		b = r4(add(p, s-4))
	case s == 8:
		a = r8(p)
		b = a
	case s <= 16:
		a = r8(p)
		b = r8(add(p, s-8))
	default:
		l := s
		if l > 48 {
			seed1 := seed
			seed2 := seed
			for ; l > 48; l -= 48 {
				seed = mix(r8(p)^m2, r8(add(p, 8))^seed)
				seed1 = mix(r8(add(p, 16))^m3, r8(add(p, 24))^seed1)
				seed2 = mix(r8(add(p, 32))^m4, r8(add(p, 40))^seed2)
				p = add(p, 48)
			}
			seed ^= seed1 ^ seed2
		}
		for ; l > 16; l -= 16 {
			seed = mix(r8(p)^m2, r8(add(p, 8))^seed)
			p = add(p, 16)
		}
		a = r8(add(p, l-16))
		b = r8(add(p, l-8))
	}

	return mix(m5^s, mix(a^m2, b^seed))
}

func memhash32Fallback(p unsafe.Pointer, seed uintptr) uintptr {
	a := r4(p)
	return mix(m5^4, mix(a^m2, a^seed^hashkey[0]^m1))
}

func memhash64Fallback(p unsafe.Pointer, seed uintptr) uintptr {
	a := r8(p)
	return mix(m5^8, mix(a^m2, a^seed^hashkey[0]^m1))
}

func mathMul64(x, y uint64) (hi, lo uint64) {
	const mask32 = 1<<32 - 1
	x0 := x & mask32
	x1 := x >> 32
	y0 := y & mask32
	y1 := y >> 32
	w0 := x0 * y0
	t := x1*y0 + w0>>32
	w1 := t & mask32
	w2 := t >> 32
	w1 += x0 * y1
	hi = x1*y1 + w2 + w1>>32
	lo = x * y
	return
}

func mix(a, b uintptr) uintptr {
	hi, lo := mathMul64(uint64(a), uint64(b))
	return uintptr(hi ^ lo)
}

func r4(p unsafe.Pointer) uintptr {
	return uintptr(readUnaligned32(p))
}

func r8(p unsafe.Pointer) uintptr {
	return uintptr(readUnaligned64(p))
}

func readUnaligned32(p unsafe.Pointer) uint32 {
	q := (*[4]byte)(p)
	if BigEndian {
		return uint32(q[3]) | uint32(q[2])<<8 | uint32(q[1])<<16 | uint32(q[0])<<24
	}
	return uint32(q[0]) | uint32(q[1])<<8 | uint32(q[2])<<16 | uint32(q[3])<<24
}

const BigEndian = false

func readUnaligned64(p unsafe.Pointer) uint64 {
	q := (*[8]byte)(p)
	if BigEndian {
		return uint64(q[7]) | uint64(q[6])<<8 | uint64(q[5])<<16 | uint64(q[4])<<24 |
			uint64(q[3])<<32 | uint64(q[2])<<40 | uint64(q[1])<<48 | uint64(q[0])<<56
	}
	return uint64(q[0]) | uint64(q[1])<<8 | uint64(q[2])<<16 | uint64(q[3])<<24 | uint64(q[4])<<32 | uint64(q[5])<<40 | uint64(q[6])<<48 | uint64(q[7])<<56
}
