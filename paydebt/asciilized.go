/*
Copyright 2019 louis.
@Time : 2019/10/24 1:01
@Author : louis
@File : asciilized
@Software: GoLand

*/

package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"image/jpeg"
	"math"
	"os"
)

// max block size = 4096 X 4096

// gray to ascii map
var (
	grayMap                = [...]string{"@", "w", "#", "$", "k", "d", "t", "j", "i", ".", " "}
	fileName string
	blx      int
	bly      int
	sampling int
	h        bool
)

func usages() {
	_, _ = fmt.Fprintf(os.Stderr, `picascii version:1.0.0
Example: 
	picascii -f hello.jpg -x 4 -y 6 -j 2 
Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.StringVarP(&fileName,"file","f","a.jpg","Define the source jpeg file")
	flag.IntVarP(&blx, "block-x", "x", 5, "Define the block size in X axis")
	flag.IntVarP(&bly, "block-y", "y", 10, "Define the block size in Y axis")
	flag.IntVarP(&sampling, "sampling", "s", 25, "sampling step")
	flag.BoolVarP(&h, "help", "h", false, "this help")
	flag.Usage=usages
}

func main() {
	flag.Parse()
	if h || blx < 1 || bly < 1 {
		flag.Usage()
		return
	}

	_, asciiData, err := AsciiSeized(fileName, blx, bly, sampling)
	if nil != err {
		flag.Usage()
		panic(err)
	}
	OutputAsciiLiSedData(asciiData)

}

func AsciiSeized(jpegFile string, blockSizeX int, blockSizeY int, jump int) (grayData [][]uint8, asciiData [][]string, err error) {
	blockX := blockSizeX
	blockY := blockSizeY
	// open file
	mFile, err := os.Open(jpegFile)
	if err != nil {
		return nil, nil, err
	}
	defer mFile.Close()
	// decode jpeg
	mImg, err := jpeg.Decode(mFile)
	if err != nil {
		return nil, nil, err
	}

	// pre
	mBounds := mImg.Bounds()

	// split blocks
	nBlockX := (int)(math.Ceil((float64)(mBounds.Dx()) / float64(blockX)))
	nBlockY := (int)(math.Ceil((float64)(mBounds.Dy()) / float64(blockY)))

	// gen 2d array
	result := make([][]uint8, nBlockY)
	for row := 0; row < nBlockY; row++ {
		subResult := make([]uint8, nBlockX, nBlockX)
		for col := 0; col < nBlockX; col++ {
			subResult[col] = 0
		}
		result[row] = subResult
	}

	// calc grey value
	var pickColorCount uint = 0
	for x := 0; x < nBlockX; x++ {
		for y := 0; y < nBlockY; y++ {
			// every block
			var graySum uint32 = 0
			pixelX := x * blockX
			pixelY := y * blockY
			breakPixel := 0
			stepCount := 0
			for i := 0; i < blockX; i++ {
				for j := 0; j < blockY; j++ {
					if 0 != jump {
						stepCount %= jump
						if 0 != stepCount {
							breakPixel++
							stepCount++
							continue
						} else {
							stepCount++
						}
					}
					targetX := pixelX + i
					targetY := pixelY + j
					if targetX >= mBounds.Dx() || targetY >= mBounds.Dy() {
						breakPixel++
						continue
					}
					// m_color := m_img.At(targetX, targetY)
					// gray := image.NewGray(image.Rect(0, 0, m_bounds.Dx(), m_bounds.Dy())).GrayAt(targetX, targetY)
					pickColorCount++
					r, g, b, a := mImg.At(targetX, targetY).RGBA()
					r /= 256
					g /= 256
					b /= 256
					a /= 256
					graySum += (uint32)((float64(r)*float64(0.3) + float64(g)*float64(0.59) + float64(b)*float64(0.11)) * (float64(a) / float64(256.0)))
				}
			}
			graySum /= uint32(blockX*blockY - breakPixel)
			result[y][x] = uint8(graySum)
		}
	}
	// output pick time
	fmt.Printf("Pick Color %d times\n", pickColorCount)

	// to ascii 2d array
	step := float64(256) / float64(len(grayMap))
	// var ascii [][]string
	ascii := make([][]string, nBlockY)
	for row := 0; row < nBlockY; row++ {
		subAscii := make([]string, nBlockX)
		for col := 0; col < nBlockX; col++ {
			index := int(float64(result[row][col]) / step)
			subAscii[col] = grayMap[index]
		}
		ascii[row] = subAscii
	}
	grayData = result
	asciiData = ascii
	return
}

func OutputAsciiLiSedData(asciiData [][]string) {
	// format output
	nBlockY := len(asciiData)
	if 0 == nBlockY {
		return
	}
	nBlockX := len(asciiData[0])
	for row := 0; row < nBlockY; row++ {
		for col := 0; col < nBlockX; col++ {
			fmt.Print(asciiData[row][col])
		}
		fmt.Println()
	}
}
