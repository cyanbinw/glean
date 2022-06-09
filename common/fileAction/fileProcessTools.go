package fileAction

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"sync"
)

type FilesData struct {
	FileName string
	FileType string
	Tag      []string
	Describe string
	Address  string
	Child    []FilesData
}

type WriteCounter struct {
	Total int64
	Item  int
}

var FileTotal int64
var FileSize int64
var lock sync.Mutex
var wg sync.WaitGroup

// GetAllFilesData is Get all files in the folder
func GetAllFilesData(dirPth string) (*FilesData, error) {
	filesData := new(FilesData)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, fi := range dir {
		if strings.Contains(fi.Name(), "json") {
			break
		} else if fi.IsDir() { // 目录, 递归遍历
			data := FilesData{
				FileName: fi.Name(),
				FileType: "Files",
				Address:  dirPth + "/" + fi.Name(),
			}
			child, err := GetAllFilesData(dirPth + "/" + fi.Name())
			if err != nil {
				fmt.Println(err)
			}
			data.Child = child.Child
			filesData.Child = append(filesData.Child, data)
		} else {
			data := FilesData{
				FileName: fi.Name(),
				FileType: "File",
				Address:  dirPth + "/" + fi.Name(),
			}
			filesData.Child = append(filesData.Child, data)
		}
	}

	return filesData, nil
}

func copyFile(selectAddress string, toAddress string) (written int64, err error) {
	dstFileName := selectAddress

	files := strings.Split(toAddress, "/")

	verifyFolder := dstFileName + "/" + files[len(files)-2] + "/" + files[len(files)-1]

	//pathExists(verifyFolder)

	srcFile, err := os.Open(toAddress)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过src fileAction ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(verifyFolder, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	counter := &WriteCounter{}
	data, err := io.Copy(writer, io.TeeReader(reader, counter))

	defer wg.Done()

	return data, nil
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total = int64(n)

	wc.PrintProgress()
	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	lock.Lock()
	FileSize += wc.Total
	num := float64(FileSize) / float64(FileTotal)
	f := int(math.Floor((num * 100) + 0.5))

	fmt.Printf("\r %d %%", f)
	lock.Unlock()
}
