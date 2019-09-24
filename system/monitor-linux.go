package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	lastTime time.Time

	logPath   string
	cpuSample int
	memSample int

	logcpu *log.Logger
	logmem *log.Logger
)

func main() {
	flag.StringVar(&logPath, "p", "/zxtech/", "log file path")
	flag.IntVar(&cpuSample, "c", 300, "sampling interval seconds")
	flag.IntVar(&memSample, "m", 300, "sampling interval seconds")
	flag.Parse()

	fmt.Printf("logPath: %v\n", logPath)
	fmt.Printf("cpuSample: %v\n", cpuSample)
	fmt.Printf("memSample: %v\n", memSample)

	if !initLogger() {
		return
	}

	go getCPUInfo()

	getMemInfo()
}

func initLogger() bool {
	fileCPU, err := os.OpenFile(logPath+"cpu"+getNowDate()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	logcpu = log.New(fileCPU, "", log.LstdFlags)

	fileMem, err := os.OpenFile(logPath+"mem"+getNowDate()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	logmem = log.New(fileMem, "", log.LstdFlags)

	return true
}

func getCPUInfo() {
	for {
		res, _ := cpu.Percent(time.Duration(cpuSample)*time.Second, true)
		total := 0.0
		line := []string{}
		for i := 0; i < len(res); i++ {
			line = append(line, fmt.Sprintf("%4.2f", res[i]))
			total += res[i]
		}
		line = append(line, fmt.Sprintf("%6.2f", total))
		logcpu.Println(strings.Join(line, "\t"))
	}
}

func getMemInfo() {
	for {
		v, _ := mem.VirtualMemory()
		memRate := float64(v.Used) * 100.0 / float64(v.Total)
		logmem.Println(fmt.Sprintf("%4.2f", memRate))

		time.Sleep(time.Duration(memSample) * time.Second)
	}
}

func getNowDate() string {
	t := time.Now()
	res := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
	return res
}
