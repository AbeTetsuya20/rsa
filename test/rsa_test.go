package test

import (
	"fmt"
	"math/big"
	"rsa/application"
	"testing"
)

// RSA 暗号が正しく動作するかテストする
func TestRSA(t *testing.T) {

	// 1. RSA暗号の各関数が正しい動作をしていることのテスト
	fmt.Println("<----- RSA Key Parts テスト開始 ----->")

	fmt.Println("12 と 18 の最大公約数が 6 になること")
	if application.Gcd(big.NewInt(12), big.NewInt(18)).Cmp(big.NewInt(6)) != 0 {
		panic(fmt.Sprintf("Error: %s", "テスト失敗"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("12 と 18 の最小公倍数が 36 になること")
	if application.Lcm(big.NewInt(12), big.NewInt(18)).Cmp(big.NewInt(36)) != 0 {
		panic(fmt.Sprintf("Error: %s", "テスト失敗"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("5224287 が素数でないこと")
	if application.IsPrime(big.NewInt(5224287)) {
		panic(fmt.Sprintf("Error: %s", "5224287 が素数として判定されている"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("2147483647 が素数であること")
	if !application.IsPrime(big.NewInt(2147483647)) {
		panic(fmt.Sprintf("Error: %s", "2147483647 が素数として判定されていない"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("素数が生成できること")
	prime, err := application.MakePrime(10)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
	}

	if !application.IsPrime(prime) {
		panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
	} else {
		fmt.Println("--> PASS 素数:", prime)
	}

	fmt.Println("巨大な数でも素数が生成できること")
	prime, err = application.MakePrime(1024)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
	}

	if !application.IsPrime(prime) {
		panic(fmt.Sprintf("Error: %s", "素数を生成できない"))
	} else {
		fmt.Println("--> PASS 素数:", prime)
	}

	fmt.Println("<-------- RSA Key Parts テスト終了 -------->")
}

// RSA 暗号のキーが正しいかテストする
func TestRSAKey(t *testing.T) {
	fmt.Println("<----- RSA Key テスト開始 ----->")

	fmt.Println("RSA 鍵が正しく生成できること")
	p, q, n, l, e, d, err := application.MakeRSAKeys(100)
	if err != nil {
		panic(fmt.Sprintf("Error: %s", "RSA 鍵が生成できない"))
	}
	fmt.Println("--> PASS")

	fmt.Println("p と q が等しくないこと")
	comp := p
	if comp.Cmp(q) == 0 {
		panic(fmt.Sprintf("Error: %s", "p と q が等しい"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("p * q = n になっていること")
	tmp := big.NewInt(0)
	if tmp.Mul(p, q).Cmp(n) != 0 {
		panic(fmt.Sprintf("Error: %s", "p * q != n"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("l が (p-1) と (q-1) の最小公倍数になっていること")
	tmpP := big.NewInt(0)
	tmpQ := big.NewInt(0)
	if application.Lcm(tmpP.Sub(p, big.NewInt(1)), tmpQ.Sub(q, big.NewInt(1))).Cmp(l) != 0 {
		panic(fmt.Sprintf("Error: %s", "l != (p-1) * (q-1)"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("e * d = 1 mod l になっていること")
	tmp = big.NewInt(1)
	if tmp.Mul(e, d).Mod(tmp, l).Cmp(big.NewInt(1)) != 0 {
		panic(fmt.Sprintf("Error: %s", "e * d != 1 mod l"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("e と l が互いに素であること")
	if application.Gcd(e, l).Cmp(big.NewInt(1)) != 0 {
		panic(fmt.Sprintf("Error: %s", "e と l が互いに素でない"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("e と d が同じ値ではないこと")
	if e.Cmp(d) == 0 {
		panic(fmt.Sprintf("Error: %s", "e と d が同じ値"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("平文を暗号化して復号化した時に平文と復号結果が同じになること")

	// 平文を 100 とする
	plainOriginal := big.NewInt(100)

	// 暗号化
	cipher := application.RSAEncrypt(plainOriginal, n, e)
	fmt.Println("plainOriginal:", plainOriginal)

	// 復号化
	plainAfter := application.RSADecrypt(cipher, n, d)
	fmt.Println("plainAfter:", plainAfter)

	if plainOriginal.Cmp(plainAfter) != 0 {
		panic(fmt.Sprintf("Error: %s", "平文が復号結果と一致しない"))
	} else {
		fmt.Println("--> PASS")
	}

	fmt.Println("<-------- RSA Key テスト終了 -------->")
}
