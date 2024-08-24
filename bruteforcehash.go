package bootcamp

import (
	"crypto/md5"
)

const hextable = "0123456789abcdef"

var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

func EncodedLen(n int) int {
	return n * 2
}

func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

func solveBruteForceHash(hash string, temp string, ans *string) {
	for i := 0; i < len(charset) && len(*ans) == 0; i++ {
		if len(temp) < 5 {
			temp += string(charset[i])
			h := md5.Sum([]byte(temp))
			if EncodeToString(h[:]) == hash {
				*ans += temp
				return
			}
			solveBruteForceHash(hash, temp, ans)
			temp = temp[:len(temp)-1]
		}
	}
}

func BruteForceHash(hash string) string {
	if len(hash) != 32 {
		return ""
	}
	var ans string
	solveBruteForceHash(hash, "", &ans)
	return ans
}
