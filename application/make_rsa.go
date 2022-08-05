package application

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Gcd 最大公約数を求める
func Gcd(a, b *big.Int) *big.Int {
	var tmp big.Int
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	} else {
		return Gcd(b, tmp.Mod(b, a))
	}
}

// Lcm 最小公倍数を求める
func Lcm(a, b *big.Int) *big.Int {
	var tmp big.Int
	ab := tmp.Mul(a, b)
	Result := tmp.Div(ab, Gcd(a, b))
	return Result
}

// IsPrime 素数かどうか判定する
func IsPrime(n *big.Int) bool {
	nSub := n
	if nSub.Cmp(big.NewInt(2)) == 0 {
		return true
	}

	// 40 回繰り返す
	count := 0
	for {
		if count >= 40 {
			return true
		}

		prime := big.NewInt(0)
		randomNumber, err := rand.Int(rand.Reader, prime.Sub(n, big.NewInt(-2)))
		if err != nil {
			fmt.Println(err)
			return false
		}
		// 2 ~ n のランダムな数を生成する
		tmp := big.NewInt(1)
		tmp.Add(randomNumber, big.NewInt(2))

		// number is tmp ^ (n-1) mod n
		// →         tmp ^ tmp2 mod nSub
		tmp2, tmp3 := big.NewInt(1), big.NewInt(1)
		// tmp2 is n-1
		tmp2 = tmp2.Sub(nSub, big.NewInt(1))
		// tmp3 is tmp ^ n-1 mod n
		tmp3 = tmp3.Exp(tmp, tmp2, nSub)

		number := tmp3
		//fmt.Println("randomNumber:", randomNumber.String())

		if number.Cmp(big.NewInt(1)) == 0 {
			count++
		} else {
			return false
		}
	}
}

// MakePrime n bit のランダムな素数を生成する
func MakePrime(n int64) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, int(n))
	if err != nil {
		return nil, err
	}
	return prime, nil
	//for {
	//
	//	primeRange := new(big.Int)
	//	primeRange.Exp(big.NewInt(2), big.NewInt(n), nil).Sub(big.NewInt(n), big.NewInt(1))
	//
	//	// 0 ~ 2^n のランダムな数を生成する
	//	seed, err := rand.Int(rand.Reader, primeRange)
	//	mathrand.Seed(seed.Int64())
	//	randomPrime := big.NewInt(seed.Int64())
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// 2^n ~ 2^(n+1) のランダムな数を生成する
	//	tmp := big.NewInt(1)
	//	randomPrime.Add(randomPrime, tmp.Exp(big.NewInt(2), big.NewInt(n), nil))
	//
	//	fmt.Println("randomPrime:", randomPrime.String())
	//
	//	if IsPrime(randomPrime) {
	//		return randomPrime, nil
	//	}
	//}
}

// MakeRSAKeys
// 秘密鍵 (d), 公開鍵 (n, e) を作成する x bit で作成する
func MakeRSAKeys(x int64) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {

	fmt.Println("作成中...")

	// e = 65537 とする
	e := big.NewInt(65537)

	p := big.NewInt(0)
	q := big.NewInt(0)
	l := big.NewInt(0)

	for {
		// p, q を、x bit で gcd(65537, l) = 1 となるように選ぶ
		p, _ = MakePrime(x)

		q, _ = MakePrime(x)

		// l = lcm(p-1, q-1) とする
		tmpP := big.NewInt(1)
		tmpP.Sub(p, big.NewInt(1))
		tmpQ := big.NewInt(1)
		tmpQ.Sub(q, big.NewInt(1))
		l = Lcm(tmpP, tmpQ)

		if Gcd(e, l).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	// p と q が 0 だったらエラー
	compP := p
	if compP.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("p is 0")
	}
	compQ := q
	if compQ.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("q is 0")
	}

	fmt.Println("p:", p.String())
	fmt.Println("q:", q.String())
	fmt.Println("l:", l.String())

	// n = p * q
	n := new(big.Int)
	n.Mul(p, q)
	fmt.Println("n:", n.String())

	fmt.Println("e:", e.String())

	// d = e^-1 mod l
	d := new(big.Int)
	d.ModInverse(e, l)
	fmt.Println("d:", d.String())
	return p, q, n, l, e, d, nil
}

// RSA 暗号で数字を暗号化
func RSAEncrypt(message *big.Int, n *big.Int, e *big.Int) *big.Int {
	// message^d mod n
	c := big.NewInt(0)
	message.Mod(message, n)
	return c.Exp(message, e, n)
}

// RSA 暗号で数字を復号
func RSADecrypt(cipher *big.Int, n *big.Int, d *big.Int) *big.Int {
	// cipher^d mod n
	m := big.NewInt(0)
	tmp := big.NewInt(0)
	tmp.Mod(cipher, n)
	return m.Exp(tmp, d, n)
}
