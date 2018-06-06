package gendat

import (
	"strings"
	"strconv"
	"os"
	"io"
	"time"
	"math"
	"encoding/binary"
	"math/rand"
	"bufio"
	"flag"
)

var (
	rootFolder     string
	datFolder      string
	machineNum     int
	metricNum      int
	machineFolders []string
)

func checkFileIsExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		println("Open file error: " + dstName)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func createMachineFolders(machineNum int) []string {
	machineFolders = make([]string, machineNum)
	for i := 0; i < machineNum; i++ {
		farm := strconv.Itoa(i / 255)
		machine := strconv.Itoa(i % 255)
		fileName := strings.Join([]string{"192.168", farm, machine}, ".")
		path := strings.Join([]string{rootFolder, "farms", farm, fileName}, "/")

		if !checkFileIsExist(path) {
			os.MkdirAll(path, os.ModePerm)
		}
		machineFolders[i] = path
	}
	return machineFolders
}

func generateDATTimer() {
	logTime := time.Now().UnixNano() / 1e9
	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
			logTime += 10
			filename := createDatFile(logTime)
			filePath := strings.Join([]string{datFolder, filename}, "/")

			start := time.Now()
			for _, folder := range machineFolders {
				CopyFile(strings.Join([]string{folder, filename}, "/"), filePath)
			}

			os.Remove(filePath)

			elapsed := time.Since(start) / 1e6
			println("Copy dat file to", machineNum, "folders, cost", elapsed, "ms!")
		}
	}
}

func createDatFile(logTime int64) string {
	start := time.Now()
	t := time.Unix(0, logTime*1e9)
	filename := t.Format("20060102150405") + ".dat"
	p := strings.Join([]string{datFolder, filename}, "/")

	f, err := os.Create(p) //创建文件
	check(err)
	defer f.Close()

	writer := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		writer.Write(Int64ToBytes(logTime + int64(i)))
		for j := 0; j < metricNum; j++ {
			metric := rand.Float32()
			//println(metric)
			writer.Write(Float32ToBytes(metric))
		}
	}
	writer.Flush()
	f.Close()

	elapsed := time.Since(start) / 1e6
	println(time.Now().Format("[2006-01-02 15:04:05]"), "create file", p, ", cost", elapsed, "ms!")

	return filename
}

func Float32ToBytes(metric float32) []byte {
	bits := math.Float32bits(metric)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Int64ToBytes(i int64) []byte {
	var bytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(i))
	return bytes
}

func main() {
	flag.StringVar(&rootFolder, "root", "", "根路径,用来保存模拟数据")
	flag.IntVar(&machineNum, "machine", 10000, "机器总台数")
	flag.IntVar(&metricNum, "metric", 100, "每条记录中指标总数")

	flag.Parse()

	if len(rootFolder) == 0 {
		println("root folder path can't be blank!")
		return
	}

	if !checkFileIsExist(rootFolder) {
		println("root folder doesn't exist!")
		return
	}

	datFolder = strings.Join([]string{rootFolder, "dat"}, "/")

	if !checkFileIsExist(datFolder) {
		os.MkdirAll(datFolder, os.ModePerm)
	}

	createMachineFolders(machineNum)

	if len(machineFolders) > 0 {
		generateDATTimer()
	} else {
		println("Failed to initialize wind machine folders!")
	}
}
