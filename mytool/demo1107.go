package mytool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Main1107 entry
func Main1107() {
	// randomtest()
	// md5test()
	sha256testfile()
}

func randomtest() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int31())
	fmt.Println(rand.Int63())
	fmt.Println(rand.Uint32())
	fmt.Println(rand.Uint64())
	fmt.Println(rand.Intn(19))
	fmt.Println(rand.Intn(19))

	fmt.Println(rand.Float64())

	//正态分布标准差，期望
	stddev, exp := 10.0, 0.0
	fmt.Println(rand.NormFloat64())
	fmt.Println(rand.NormFloat64()*stddev + exp)

	// 指数分布，率参数是期望的倒数
	rate := 0.5
	fmt.Println(rand.ExpFloat64())
	fmt.Println(rand.ExpFloat64() / rate)

	fmt.Println(rand.Perm(10))
}

func md5test() {
	fmt.Println(md5.Sum([]byte("123")))

	// 需要从数组转化成切片，md5.Sum这个函数返回值比较变态
	mdkarr := make([]byte, md5.Size)
	mdarr := md5.Sum([]byte("123"))
	mdkarr = mdarr[:]
	fmt.Println(mdkarr)
	fmt.Println(hex.EncodeToString(mdkarr))

	hashfun := md5.New()
	hashfun.Write([]byte("123"))
	resArr := hashfun.Sum(nil)
	fmt.Println(resArr)
	fmt.Println(hex.EncodeToString(resArr))
}

func sha256testfile() {
	var hashStr string

	file, err := os.Open("/Users/allen/Desktop/ttt.txt")
	if err != nil {
		fmt.Println("Open file Error!!")
		return
	}
	defer file.Close()

	hashfun := sha256.New()
	if _, err := io.Copy(hashfun, file); err != nil {
		fmt.Println("cal error!!")
		return
	}

	hashInBytes := hashfun.Sum(nil)
	hashStr = hex.EncodeToString(hashInBytes)

	fmt.Println(hashStr)
}
