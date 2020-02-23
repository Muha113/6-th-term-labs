package common

import (
	"crypto/md5"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func GenKey(length int) string {
	key := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		key += string(rune(rand.Intn(26) + 97))
	}
	return Hash(key)
}

func Hash(text string) string {
	hash := md5.Sum([]byte(text))
	hashStr := ""
	for _, v := range hash {
		hashStr += string(v)
	}
	return hashStr
}

func HandleError(err error, adds string) {
	if err != nil {
		logrus.Fatal(adds, err)
	}
}