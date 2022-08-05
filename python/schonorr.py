import random
import hashlib

MAX_BIT = 1024


# テスト
def test(key):
    try:
        if not ("p" in key):
            raise ValueError("key の指定が不正: p が存在しません。")
        if not ("q" in key):
            raise ValueError("key の指定が不正: q が存在しません。")
        if not ("a" in key):
            raise ValueError("key の指定が不正: a が存在しません。")
        if not ("g" in key):
            raise ValueError("key の指定が不正: g が存在しません。")

        p = key["p"]
        q = key["q"]
        a = key["a"]
        g = key["g"]

        if not isPrime(p):
            raise ValueError("p が素数ではない")
        if not isPrime(q):
            raise ValueError("q が素数ではない")
        if p != a * q + 1:
            raise ValueError("a の値が不正")
        if pow(g, q, p) != 1:
            raise ValueError("g の値が不正")
        # for i in range(2, q):
        #     if pow(g, i, p) == 1:
        #         raise ValueError("g^" + str(i) + " mod p の値が1になります。")
        print("ALL TEST PASSED!")

    except ValueError as e:
        print(key)
        print(e)
        exit(1)


# a~b の範囲で素数を生成する関数
def make_prime(a, b):
    p = random.randint(a, b)
    while not (isPrime(p)):
        p = random.randint(a, b)
        print(".", end="")
    print("!")
    return p


# p が素数かどうか判定する関数
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


# n bit であるかどうかチェックする関数
def isNbit(p, n):
    if pow(2, n) <= p:
        return True
    else:
        return False


# 秘密鍵(p, q)を n bit で作成する関数
# p = 2^n 3^m q + 1 を計算
# p が素数かどうか判定
# p が素数じゃなかった場合、n,m の値を変えて再実行
def make_p_q(n):
    key = {"n": n}
    k = int(n / 4)
    while True:
        q = make_prime(pow(2, k), pow(2, k + 1) - 1)
        a = 2
        while True:
            print("_", end="")
            p = a * q + 1

            # p が 1024bit を超えるまで 2 をかける
            if isNbit(p, MAX_BIT):
                # 素数かどうか判定をする
                if isPrime(p):
                    print("!")
                    key["a"] = a
                    key["p"] = p
                    key["q"] = q
                    return key

                # すべて 3 に置き換わったら生成不可能
                if a % 3 == 0:
                    break

                # 2 を 3 に置き換える
                a = int((a / 2) * 3)

            else:
                a *= 2


# g'をランダムに選ぶ (1 <= g' < p)
# g^((p-1)/2) mod p ,g^((p-1)/3) mod p が 1 にならなかったら g'^a が g
def make_g(a, p, q):
    while True:
        g = random.randint(1, p)

        div2 = (p - 1) // 2
        div3 = (p - 1) // 3
        if (pow(g, div2, p) != 1) and (pow(g, div3, p) != 1):

            if pow(g, q, p) == 1:
                return g
            else:
                return pow(g, a, p)


# 文章をハッシュ化する関数
def encoding_message(message, grp):
    h = hashlib.sha256()
    h.update(message.encode())
    h.update(grp.to_bytes(128, "big"))
    return int.from_bytes(h.digest(), byteorder="big")


# 署名をする関数
def to_schnorr(message, key, grp):
    c = encoding_message(message, grp)
    s = c * key["x"] + key["r"]
    return c, s


# n bit で鍵を作成する関数
def make_key(n):
    key = make_p_q(n)
    key["g"] = make_g(key["a"], key["p"], key["q"])
    key["x"] = random.randint(1, key["p"])
    key["y"] = pow(key["g"], key["x"], key["p"])
    return key


def main():
    key = make_key(MAX_BIT)
    print(key)
    test(key)

    # 署名
    message = "こんにちは"
    key["r"] = random.randint(1, key["p"])
    grp = pow(key["g"], key["r"], key["p"])
    c, s = to_schnorr(message, key, grp)
    print("署名完了: ", c, s)

    # 復号
    gsyc = pow(key["g"], s, key["p"]) * pow(key["y"], -c, key["p"])
    cp = encoding_message(message, gsyc)
    print("署名完了: ", cp)

    # if key["c"] == key["cp"]:
    #     print("署名完了")
    # else:
    #     print("署名失敗")


if __name__ == "__main__":
    print("実行開始")
    main()
    print("実行終了")