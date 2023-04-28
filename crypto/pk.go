package crypto

import (
	"fmt"
	"leetcode.com/leetcode/utils"
	"math/big"
)

var (
	zero = big.NewInt(0)
)

func NaiveModExp(base, exp, mod int) int {
	ans := 1
	for i := 0; i < exp; i++ {
		ans = (base * ans) % mod
	}
	return ans
}

func FasterModExp(base, exp, mod int) int {
	shift := 1
	ans := 1
	for shift <= exp {
		if shift & exp == shift {
			ans = ans * modExpPow2(base, utils.Log2(shift), mod) % mod
		}
		// shift shift left by 1
		shift = shift << 1
	}
	return ans
}

// use when the exponent of base is a power of two, can perform log2(exp) operations, i.e 3^16 (4) = 
func modExpPow2(base, log2exp, mod int) int {
	ans := base % mod
	for i := 0; i < log2exp; i++ {
		ans = (ans * ans) % mod
	}
	return ans
}

func Exp(base, exp int) int {
	return exp_(base, exp, 1)
}

func exp_(base, exp, accum int) int {
	if exp == 0 {
		return accum
	}
	if exp % 2 == 0 {
		return exp_(base, exp >> 1, accum * accum)
	}
	return exp_(base, exp - 1, accum * base)
}

func GCD(a, b int) int {
	// reverse if necessary
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	} else {
		return GCD(a - b, b)
	}
}

func ExtendedEuclidean(a, b *big.Int) (u, v, d *big.Int) {
	fmt.Println(a, b)
	// switch if necessary
	if a.Cmp(b) < 0{
		a, b = b, a
	}
	// initialize aux variables
	var (
		v_1 = big.NewInt(0)
		v_3 = big.NewInt(0).Set(b)
		t_1 = big.NewInt(0)
		t_3 = big.NewInt(0)
	)
	u = big.NewInt(1)
	d = big.NewInt(0).Set(a)
	// short circuit if possible
	if b.Cmp(zero) == 0 {
		return u, zero, a
	}
	for v_3.Cmp(zero) != 0 {
		t_3.Mod(d,v_3)
		// t_1 = u - (d / v_3) * v_1
		q := d.Div(d,v_3)
		t_1.Sub(u, q.Mul(q, v_1))
		u.Set(v_1)
		d.Set(v_3)
		v_1.Set(t_1)
		v_3.Set(t_3)
	}
	v = big.NewInt(0).Sub(d, a.Mul(a, u))
	v.Div(v, b)
	return u, v, d
}

func FasterModExpBig(base, exp, mod *big.Int) *big.Int {	
	fmt.Println(base, exp, mod)
	shift := 0
	ans := big.NewInt(1)
	for shift <= exp.BitLen(){
		if exp.Bit(shift) == 1 {
			ans.Mod(ans.Mul(ans, bigModExpPow2(base, mod, shift)), mod)
		}
		// shift shift left by 1
		shift++
	}
	return ans
}

func bigModExpPow2(base, mod *big.Int, exp int) *big.Int {
	accum := base.Mod(base, mod)
	for i := 0; i < exp; i++ {
		accum.Mod(accum.Mul(accum, accum), mod)
	}
	return accum
}
