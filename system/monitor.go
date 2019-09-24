package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const (
	postURL = "http://10.98.202.84:9876/util/sendsms.do"
	logPath = "d:/"
	// threshold number of cpu core
	highNum = 15
)

var (
	// how many times of continuous exceed threshold (6 times)
	count       = 0
	lastTime    time.Time
	smsSendFlag = false

	phoneNums string
	cpuSample int
	memSample int
	cpuRate   float64
	sendFlag  bool

	logcpu *log.Logger
	logmem *log.Logger
)

// WarnInfo sms info content
type WarnInfo struct {
	ip      string
	overNum int
	cpuSum  float64
	memRate float64
}

func main() {
	flag.StringVar(&phoneNums, "phone", "18040236152", "Phone numbers")
	flag.IntVar(&cpuSample, "c", 60, "sampling interval seconds")
	flag.IntVar(&memSample, "m", 300, "sampling interval seconds")
	flag.Float64Var(&cpuRate, "cp", 70.0, "CPU useage rate")
	flag.BoolVar(&sendFlag, "ts", false, "Test send sms message")
	flag.Parse()

	smsSendFlag = false

	fmt.Printf("phoneNums: %v\n", phoneNums)
	fmt.Printf("cpuSample: %v\n", cpuSample)
	fmt.Printf("memSample: %v\n", memSample)
	fmt.Printf("cpuRate: %v\n", cpuRate)
	fmt.Printf("sendFlag: %v\n", sendFlag)

	if !initLogger() {
		return
	}

	warnBean := WarnInfo{getIntranetIP(), 0, 0.0, 0.0}

	if sendFlag {
		sendSms(&warnBean)
		return

	}

	go getCPUInfo(&warnBean)

	getMemInfo(&warnBean)
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

func getCPUInfo(warnBean *WarnInfo) {
	for {
		res, _ := cpu.Percent(time.Duration(cpuSample)*time.Second, true)
		total := 0.0
		innerNum := 0
		line := []string{}
		for i := 0; i < len(res); i++ {
			line = append(line, fmt.Sprintf("%4.2f", res[i]))
			if res[i] > cpuRate {
				innerNum++
			}
			total += res[i]
		}
		line = append(line, fmt.Sprintf("%6.2f", total))
		logcpu.Println(strings.Join(line, "\t"))

		if innerNum >= highNum {
			count++
		} else {
			count = 0
		}

		if count > 6 {
			warnBean.overNum = innerNum
			warnBean.cpuSum = total
			count = 0
			sendSms(warnBean)
		}
	}
}

func getMemInfo(warnBean *WarnInfo) {
	for {
		v, _ := mem.VirtualMemory()
		warnBean.memRate = float64(v.Used) * 100.0 / float64(v.Total)
		logmem.Println(fmt.Sprintf("%4.2f", warnBean.memRate))

		time.Sleep(time.Duration(memSample) * time.Second)
	}
}

func getNowDate() string {
	t := time.Now()
	res := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
	return res
}

func sendSms(warnBean *WarnInfo) {
	if !smsSendFlag || (smsSendFlag && time.Now().Sub(lastTime) > 3*time.Hour) {
		content := fmt.Sprintf("[Alarm] Server %s Alarm, cpu_total_rate %6.2f%%, mem_usage_rate %6.2f%%, core_over_num %d", warnBean.ip, warnBean.cpuSum, warnBean.memRate, warnBean.overNum)

		//把post表单发送给目标服务器
		res, err := http.PostForm(postURL, url.Values{"phoneNums": {phoneNums}, "content": {content}})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer res.Body.Close()

		lastTime = time.Now()
		smsSendFlag = true
	}
}

func getIntranetIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tmpArr := make([]string, 0)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				tmpArr = append(tmpArr, ipnet.IP.String())
			}
		}
	}

	return strings.Join(tmpArr, " | ")
}
