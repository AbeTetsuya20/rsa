import random
import math
import hashlib
import time

# n は鍵の bit 長
Max_bit = 100

# 100: 0.007398843765258789
# 500: 0.33533811569213867
# 1024: 2.70386004447937
# 1500: 6.355698108673096
# 2048: 15.969524145126343

# テスト
def key_test(key):
    print("-----テストを開始します-----")

    n = key["n"]
    e = key["e"]
    p = key["p"]
    q = key["q"]
    l = key["l"]
    d = key["d"]

    try:
        if not isPrime(p) or not isPrime(q):
            raise ValueError("p,q error: p の値が素数ではない")
        if n != p * q:
            raise ValueError("n error: n の値が p * q と不一致")
        if l != lcm(p - 1, q - 1):
            raise ValueError("l error: l の値が p-1, q-1 の最小公倍数になっていない")
        if math.gcd(e, l) != 1:
            raise ValueError("e error: e が l と互いに素でない")
        if ((e * d) % l) != 1:
            raise ValueError("d error: e * d mod l の値が 1 にならない")

        message = 100
        enc_result = rsa_enc(key, message)
        dec_result = rsa_dec(key, enc_result)
        if not (message == dec_result):
            print("平文: ", message)
            print("暗号結果: ", enc_result)
            print("復号結果: ", dec_result)
            raise ValueError("復号結果が平文と不一致")

        print("All Test Passed!")
        print("")
    except ValueError as e:
        print(e)
        exit(1)


# 最大公約数を求める関数
def lcm(p, q):
    return p // math.gcd(p, q) * q


# n bit の素数を作成する関数
def isPrime(p):
    count = 0

    while count <= 40:
        a = random.randint(1, p)

        num = pow(a, p - 1, p)
        if num == 1:
            count = count + 1
        else:
            return False

    return True

# a~b の範囲で素数を生成する関数
def make_prime(a, b):
    p = random.randint(a, b)
    while not (isPrime(p)):
        p = random.randint(a, b)
        print(".", end="")
    print("!")
    return p


# n bit の素数を二つ作成する関数
def make_p_q(n):
    global Max_bit
    key = {"p": make_prime(pow(2, n - 1), pow(2, n) - 1), "q": make_prime(pow(2, n - 1), pow(2, n) - 1)}
    return key


# 秘密鍵 e を作成
# e は l と互いに素
def make_e(l):
    return make_prime(1, l)


# 情報逆元を求める関数1: pow関数で逆元を求める
def make_d_1(e, l):
    return pow(e, -1, l)


# 情報逆元を求める関数2: 自作の関数で逆元を求める
def make_d_2(e, l):
    (A, B) = (l, e)
    (a, b) = (0, 1)
    while True:
        if A % B == 0:
            return l - b
        q = A // B
        (A, B) = (B, A % B)
        (a, b) = (b, a - q * b)


# n bit の秘密鍵・公開鍵を作成する関数
def make_key(n):
    key = make_p_q(n)
    l = lcm(key["p"] - 1, key["q"] - 1)
    key["n"] = key["p"] * key["q"]
    key["l"] = l
    e = make_e(l)
    key["e"] = e
    d = make_d_1(e, l)
    key["d"] = d
    return key


# 鍵を表示する print 関数
def printer(key):
    print("------------公開鍵------------")
    print("n(=p*q)=", key["n"])
    print("e=", key["e"])
    print("")
    print("------------秘密鍵------------")
    print("d=", key["d"])
    print("")
    print("------------その他------------")
    print("p=", key["p"])
    print("q=", key["q"])
    print("l=", key["l"])
    print("")


# 暗号化
def rsa_enc(pub_key, message):
    e = pub_key["e"]
    n = pub_key["n"]
    result = pow(message, e, n)
    return result


# 復号化
def rsa_dec(priv_key, message):
    d = priv_key["d"]
    n = priv_key["n"]
    result = pow(message, d, n)
    return result


def encoding_message(message):
    h = hashlib.sha256()
    h.update(message.encode())
    h_byte = h.digest()
    h_byte = b'\x30\x0B\x06\x09\x60\x86\x48\x01\x65\x03\x04\x02\x01' + h_byte
    while len(h_byte) < (2048 - 16) // 8:
        h_byte = b'\xff' + h_byte
    return int.from_bytes(h_byte, byteorder="big")


# main 関数
def main():
    global Max_bit

    # 素数生成の時間を3回計測して、その平均を求める
    start = time.time()
    for i in range(3):
        prime = make_prime(pow(2, Max_bit - 1), pow(2, Max_bit) - 1)
    end = time.time()
    print("素数生成時間: ", (end - start) / 3)


#     # 鍵の生成とテスト
#     key = make_key(Max_bit)
#     printer(key)
#     key_test(key)
#
#     # 暗号化・復号化
#     message = encoding_message("hello")
#     enc_result = rsa_enc(key, message)
#     dec_result = rsa_dec(key, enc_result)
#
#     print("平文: " , message)
#     print("エンコード後: " , enc_result)
#     print("デコード後: " , dec_result)


if __name__ == "__main__":
    print("実行開始")
    main()
