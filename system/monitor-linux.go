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

const (
	logPath = "/Users/allen/"
	// threshold number of cpu core
	highNum = 15
)

var (
	// how many times of continuous exceed threshold (6 times)
	count    = 0
	lastTime time.Time

	cpuSample int
	memSample int
	cpuRate   float64
	sendFlag  bool

	logcpu *log.Logger
	logmem *log.Logger
)

func main() {
	flag.IntVar(&cpuSample, "c", 60, "sampling interval seconds")
	flag.IntVar(&memSample, "m", 300, "sampling interval seconds")
	flag.Float64Var(&cpuRate, "cp", 70.0, "CPU useage rate")
	flag.Parse()

	fmt.Printf("cpuSample: %v\n", cpuSample)
	fmt.Printf("memSample: %v\n", memSample)
	fmt.Printf("cpuRate: %v\n", cpuRate)

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
	logcpu = log.New(fileCPU, "", log.LstdFlags|log.Llongfile)
	logcpu.SetFlags(log.LstdFlags)

	fileMem, err := os.OpenFile(logPath+"mem"+getNowDate()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		return false
	}
	logmem = log.New(fileMem, "", log.LstdFlags|log.Llongfile)
	logmem.SetFlags(log.LstdFlags)

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
