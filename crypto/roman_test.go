package crypto_test

import (
	"fmt"
	"testing"
	"leetcode.com/leetcode/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
)

var (
	p, _ = big.NewInt(0).SetString("857504083339712752489993810777", 10)
	q, _ = big.NewInt(0).SetString("1029224947942998075080348647219", 10)
)

var x *big.Int = nil

func TestRoman(t *testing.T) {
	crypto.SolveCaesar()
}

func TestModExp(t *testing.T ) {
	// test naive against 2^4(2) = 0
	assert.Equal(t, crypto.NaiveModExp(2, 4, 2), 0)
	assert.Equal(t, crypto.NaiveModExp(3, 5, 2), 1)
	// assert.Equal(t, crypto.FasterModExp(3, 5, 2), 1)
	assert.Equal(t, crypto.FasterModExp(2,10, 17), 4)
	fmt.Println(crypto.FasterModExp(101, 17, 22663))
	fmt.Println(crypto.FasterModExpBig(big.NewInt(101), big.NewInt(17), big.NewInt(22663)))
	fmt.Println(crypto.FasterModExp(12, 65537, 17 * 23))
	fmt.Println(crypto.FasterModExpBig(big.NewInt(12), big.NewInt(65537), big.NewInt(17 * 23)))
	assert.Equal(t, crypto.Exp(3,3), 27)
}

func TestGCD(t *testing.T) {
	assert.Equal(t, crypto.GCD(81, 57),3)
	u, v, d := crypto.ExtendedEuclidean(big.NewInt(5), big.NewInt(3))
	assert.Equal(t, d.Int64(), int64(1))
	assert.Equal(t, u.Int64(), int64(-1))
	assert.Equal(t, v.Int64(), int64(2))
	// modular inverse of 65537 mod \phi(pq) is priv key
	u, v, d = crypto.ExtendedEuclidean(big.NewInt(65537), big.NewInt(0).Set(p.Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))))
	// what is u?
	fmt.Println(u, v, d)
	fmt.Println(big.NewInt(0).ModInverse(big.NewInt(65537), p))
	c, _ := big.NewInt(0).SetString("77578995801157823671636298847186723593814843845525223303932", 10)
	N, _ := big.NewInt(0).SetString("882564595536224140639625987659416029426239230804614613279163", 10)
	fmt.Println("message", crypto.FasterModExpBig(c, v, N))
	fmt.Println("message-expect", big.NewInt(0).Exp(c, v, N))
	m, _ := big.NewInt(0).SetString("90069129621022947985964673588167290009238865977637281349495", 10)
	fmt.Println(crypto.FasterModExpBig(m, big.NewInt(65537), N))
}
