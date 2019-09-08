/*
@Time : 2019/8/26 14:52
@Author : louis
@File : testuint32
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/imroc/biu"
	"os"
)

//A FileMode represents a file's mode and permission bits.
//The bits have the same definition on all systems, so that
//information about files can be moved from one system
//to another portably. Not all bits apply to all systems.
//The only required bit is ModeDir for directories.
//type mode uint32
//
//const (
//	// The single letters are the abbreviations
//	// used by the String method's formatting.
//	ModeDir        mode = 1 << (32 - 1 - iota) // d: is a directory
//	ModeAppend                                     // a: append-only
//	ModeExclusive                                  // l: exclusive use
//	ModeTemporary                                  // T: temporary file; Plan 9 only
//	ModeSymlink                                    // L: symbolic link
//	ModeDevice                                     // D: device file
//	ModeNamedPipe                                  // p: named pipe (FIFO)
//	ModeSocket                                     // S: Unix domain socket
//	ModeSetuid                                     // u: setuid
//	ModeSetgid                                     // g: setgid
//	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
//	ModeSticky                                     // t: sticky
//	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file
//
//	// Mask for the type bits. For regular files, none will be set.
//	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
//
//	ModePerm mode = 0777 // Unix permission bits
//)
//
//func (m mode) String() string {
//	const str = "dalTLDpSugct?"
//	var buf [32]byte // Mode is uint32.
//	w := 0
//	for i, c := range str {
//		if m&(1<<uint(32-1-i)) != 0 {
//			fmt.Println(string(c))
//			buf[w] = byte(c)
//			w++
//		}
//	}
//	if w == 0 {
//		buf[w] = '-'
//		w++
//	}
//	//    rwx = 1 1111 1111 ==> 将字节对应到相应的字符.
//	const rwx = "rwxrwxrwx"
//	for i, c := range rwx {
//		if m&(1<<uint(9-1-i)) != 0 {
// 如果 & 结果不为零; 说明为真 ; 将该位置为相应的字符;否则置空
// 第一位即最高位为1,可读权限 1 0000 0000  = 400
// 第二位为1               0 1000 0000  = 200
// 第三位为1               0 0100 0000  = 100
// rwx ==> 700
// 第四位为1               0 0010 0000  = 040
// 第五位为1               0 0001 0000  = 020
// 第六位为1               0 0000 1000  = 010
// rwx ==> 070
// 第七位为1               0 0000 0100  = 004
// 第八位为1               0 0000 0010  = 002
// 第九位为1               0 0000 0001  = 001
//			// rwx ==> 007
//			buf[w] = byte(c)
//		} else {
//			buf[w] = '-'
//		}
//		w++
//	}
//	return string(buf[:w])
//}

func main() {
	var m1 os.FileMode
	var m2 uint16
	m2 = 0775
	fmt.Printf("%b\n", m2)
	m1 = 0775 //每一位都是8进制, 775 -> 111111101
	fmt.Printf("%b\n", m1)
	fmt.Println(m1.String())
	fmt.Println(os.ModeDir.String())

	fmt.Println(biu.ToBinaryString(os.ModePerm))
	fmt.Println(biu.ToBinaryString(os.ModeDir))
	fmt.Println(biu.ToBinaryString(os.ModeAppend))
	fmt.Println(biu.ToBinaryString(os.ModeExclusive))
	fmt.Println(biu.ToBinaryString(os.ModeTemporary))
	fmt.Println(biu.ToBinaryString(os.ModeSymlink))
	fmt.Println(biu.ToBinaryString(os.ModeDevice))
	fmt.Println(biu.ToBinaryString(os.ModeNamedPipe))
	fmt.Println(biu.ToBinaryString(os.ModeSocket))
	fmt.Println(biu.ToBinaryString(os.ModeSetuid))
	fmt.Println(biu.ToBinaryString(os.ModeSetgid))
	fmt.Println(biu.ToBinaryString(os.ModeCharDevice))
	fmt.Println(biu.ToBinaryString(os.ModeSticky))
	fmt.Println(biu.ToBinaryString(os.ModeIrregular))
	fmt.Println(biu.ToBinaryString(os.ModeType))

}

/*
[00000000 00000000 00000001 11111111]
[10000000 00000000 00000000 00000000]
[01000000 00000000 00000000 00000000]
[00100000 00000000 00000000 00000000]
[00010000 00000000 00000000 00000000]
[00001000 00000000 00000000 00000000]
[00000100 00000000 00000000 00000000]
[00000010 00000000 00000000 00000000]
[00000001 00000000 00000000 00000000]
[00000000 10000000 00000000 00000000]
[00000000 01000000 00000000 00000000]
[00000000 00100000 00000000 00000000]
[00000000 00010000 00000000 00000000]
[00000000 00001000 00000000 00000000]
[10001111 00101000 00000000 00000000]
*/
