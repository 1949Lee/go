package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"log"
	"os"
)

//常量ti  uint32(abs(sin(i+1))*(2pow32))
var ti = []uint32{
	0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
	0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501, 0x698098d8,
	0x8b44f7af, 0xffff5bb1, 0x895cd7be, 0x6b901122, 0xfd987193,
	0xa679438e, 0x49b40821, 0xf61e2562, 0xc040b340, 0x265e5a51,
	0xe9b6c7aa, 0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
	0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed, 0xa9e3e905,
	0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a, 0xfffa3942, 0x8771f681,
	0x6d9d6122, 0xfde5380c, 0xa4beea44, 0x4bdecfa9, 0xf6bb4b60,
	0xbebfbc70, 0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
	0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665, 0xf4292244,
	0x432aff97, 0xab9423a7, 0xfc93a039, 0x655b59c3, 0x8f0ccc92,
	0xffeff47d, 0x85845dd1, 0x6fa87e4f, 0xfe2ce6e0, 0xa3014314,
	0x4e0811a1, 0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
}

//向左位移数
var s = []uint32{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

// 32位int转化为[]byte，8位一个字节，高位到低位的顺序存储在[]byte
func uint32ToBytes(i uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.LittleEndian, i)
	return bytesBuffer.Bytes()
}

// 获取每分组的16个uint32
func getGroup(data []byte, index int) *[]uint32 {
	group := make([]uint32, 16)
	for i := range group {
		_ = binary.Read(bytes.NewBuffer([]byte{
			data[index+i*4],
			data[index+1+i*4],
			data[index+2+i*4],
			data[index+3+i*4],
		}), binary.LittleEndian, &group[i])
		//log.Printf("%x",group[i])
	}
	return &group
}

// i循环左移s位
func shift(i uint32, s uint32) uint32 {
	return (i << s) | (i >> (32 - s))
}

func main() {
	file, err := os.OpenFile("test.html", os.O_RDONLY, os.ModePerm)
	if err == os.ErrNotExist {
		log.Println("文件不存在")
		return
	}
	if err != nil {
		log.Printf("打开文件发生错误：%v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data := make([]byte, 0)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("读取文件发生错误：%v", err)
		}
		data = append(data, b)
	}
	//data := []byte("abcde")
	//log.Printf("%x",md5.Sum(data))
	length := uint64(len(data))
	l1 := uint32((length * 8) & 0x00000000ffffffff)
	l2 := uint32((length * 8) >> 16)

	// 根据原始数据的长度补齐512位的整倍数（64字节的整倍数）
	mod := length % 64
	// 补一位1（一位1的uint8的值是128）
	data = append(data, 0x80)

	// 用0补充需要的字节。go语言中数值类型的默认值是0。所以创建空的字节数组即可。
	if mod == 0 {
		data = append(data, make([]byte, 55)...)
	} else if mod < 55 {
		data = append(data, make([]byte, 55-mod)...)
	} else if mod > 55 {
		// 这种情况是你的原始数据以512位（64字节）一个分组，这样分下来之后，最后一组剩下56或更大的字节
		data = append(data, make([]byte, 64-mod-1+(64-8))...)
	}

	// 最后一个512位的分组的最后64位（8字节）需要用原始数据的长度来表示。
	data = append(data, uint32ToBytes(l1)...)
	data = append(data, uint32ToBytes(l2)...)
	var (
		A uint32 = 0x67452301
		B uint32 = 0xefcdab89
		C uint32 = 0x98badcfe
		D uint32 = 0x10325476
	)

	// 64字节为一个分组，
	groupCount := len(data) / 64

	// 遍历每个分组
	for i := 0; i < groupCount; i++ {
		startIndex := i * 64
		aTemp := A
		bTemp := B
		cTemp := C
		dTemp := D
		group := getGroup(data, startIndex)
		var (
			F     uint32
			index int
		)
		for j := 0; j < 64; j++ {
			if j < 16 {
				F = (bTemp & cTemp) | ((^bTemp) & dTemp)
				index = j
			} else if j < 32 {
				F = (dTemp & bTemp) | ((^dTemp) & cTemp)
				index = (5*j + 1) % 16
			} else if j < 48 {
				F = bTemp ^ cTemp ^ dTemp
				index = (3*j + 5) % 16
			} else {
				F = cTemp ^ (bTemp | (^dTemp))
				index = (7 * j) % 16
			}
			tmp := dTemp
			dTemp = cTemp
			cTemp = bTemp
			bTemp = bTemp + shift(aTemp+F+ti[j]+(*group)[index], s[j])
			aTemp = tmp
		}
		A += aTemp
		B += bTemp
		C += cTemp
		D += dTemp
	}
	log.Println(
		hex.EncodeToString(uint32ToBytes(A)) +
			hex.EncodeToString(uint32ToBytes(B)) +
			hex.EncodeToString(uint32ToBytes(C)) +
			hex.EncodeToString(uint32ToBytes(D)))
}
