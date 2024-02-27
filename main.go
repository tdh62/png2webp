package main

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: png2webp <input.png>|<dir> [quality] [-lossless] [-dir_walker]")
		return
	}
	// try to get quality and lossless
	quality := 80
	lossless := false
	dirWalker := false
	overwrite := false
	for _, arg := range os.Args[2:] {
		if strings.HasPrefix(arg, "-") {
			if strings.HasPrefix(arg, "-lossless") || strings.HasPrefix(arg, "-l") {
				lossless = true
			}
			if strings.HasPrefix(arg, "-dir_walker") || strings.HasPrefix(arg, "-d") {
				dirWalker = true
			}
			if strings.HasPrefix(arg, "-overwrite") || strings.HasPrefix(arg, "-o") {
				overwrite = true
			}
		} else {
			_, _ = fmt.Sscanf(arg, "%d", &quality)
		}
	}

	if dirWalker {
		// 遍历目录
		err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// check png jpg jpeg gif
			if !strings.HasSuffix(strings.ToLower(path), ".png") &&
				!strings.HasSuffix(strings.ToLower(path), ".jpg") &&
				!strings.HasSuffix(strings.ToLower(path), ".jpeg") &&
				!strings.HasSuffix(strings.ToLower(path), ".gif") {
				return nil
			}
			// convert
			png2webp(path, lossless, quality, false)
			return nil
		})
		if err != nil {
			fmt.Printf("执行过程中出错：%s\n", err)
			return
		}
	} else {
		png2webp(os.Args[1], lossless, quality, overwrite)
	}
}

func png2webp(path string, lossless bool, quality int, overwrite bool) {

	// 输入和输出文件名
	//inputFileName := os.Args[1]
	outputFileName := strings.TrimSuffix(path, filepath.Ext(path)) + ".webp"

	// check if output file exists
	if _, err := os.Stat(outputFileName); err == nil {
		if !overwrite {
			fmt.Printf("输出文件已存在：%s，跳过\n", outputFileName)
			return
		}
	}

	// 打开输入文件
	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("无法打开输入文件：%s\n", err)
		return
	}
	defer inputFile.Close()

	// 读取图片，不再仅限于 PNG 格式
	pngImage, f, err := image.Decode(inputFile)
	fmt.Printf("读取图片：%s\n", f)
	if err != nil {
		fmt.Printf("无法解码 PNG 图片：%s\n", err)
		return
	}

	// 创建一个内存缓冲区用于存储转换后的 WebP 图片
	var webpBuffer bytes.Buffer

	// 将 PNG 图片编码为 WebP 格式
	if lossless {
		err = webp.Encode(&webpBuffer, pngImage, &webp.Options{Lossless: true})
	} else {
		err = webp.Encode(&webpBuffer, pngImage, &webp.Options{Quality: float32(quality)})
	}
	if err != nil {
		fmt.Printf("无法编码为 WebP 格式：%s\n", err)
		return
	}

	// 将转换后的 WebP 数据写入到输出文件
	err = os.WriteFile(outputFileName, webpBuffer.Bytes(), 0666)
	if err != nil {
		fmt.Printf("无法写入输出文件：%s\n", err)
		return
	}

	fmt.Println("图片转换成功：" + outputFileName)
}
