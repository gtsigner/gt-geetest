package geetest

func _h6(d6 int, x6 int) int {
    p6 := 0
    j6 := 23
    for {
        if j6 < 0 {
            break
        }

        if _ae(x6, j6) == 1 {
            p6 = p6<<1 + _ae(d6, j6)
        }
        j6--
    }
    return p6
}
func _ae(H6 int, T6 int) int {
    return int(uint(H6) >> uint(T6) & 1)
}
func _ddd(q6 int) string {
    C6 := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "(", ")"}
    //q6 = q6 + 1
    if q6 >= 0 && q6 < 64 {
        return C6[q6]
    }
    return "."
}

func base64Encode(s6 []byte) string {
    //fmt.Println((s6))

    g7s := 1
    f6 := len(s6)
    var V6, r6 string
    var D6 = 0
    for {
        if D6 >= f6 || g7s*(g7s+1)*g7s%2 != 0 {
            break
        }
        K6 := 0
        if D6+2 < f6 {
            //println(D6)
            K6 = int(s6[D6])<<16 + int(s6[D6+1])<<8 + int(s6[D6+2])
            V6 = V6 + _ddd(_h6(K6, 7274496)) + _ddd(_h6(K6, 9483264)) + _ddd(_h6(K6, 19220)) + _ddd(_h6(K6, 235))

        } else {
            Z6 := f6 % 3
            if Z6 == 2 {
                K6 = int(s6[D6])<<16 + int(s6[D6+1])<<8
                V6 = V6 + _ddd(_h6(K6, 7274496)) + _ddd(_h6(K6, 9483264)) + _ddd(_h6(K6, 19220))
                r6 = "."
            }
            if Z6 == 1 {
                K6 = int(s6[D6]) << 16
                V6 = V6 + _ddd(_h6(K6, 7274496)) + _ddd(_h6(K6, 9483264))
                r6 = ".."
            }
        }
        if g7s > 21221 {
            g7s = g7s - 6
        } else {
            g7s = g7s + 6
        }
        D6 = D6 + 3
    }

    return V6 + r6
}
