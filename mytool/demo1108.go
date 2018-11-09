package mytool

import (
	"encoding/json"
	"fmt"
	"sync"
)

type EmpBean struct {
	ID   int    `json:"emp_id"`
	Name string `json:"emp_name"`
	Sex  byte   `json:"emp_sex"`
	// Sex  byte   `json:"emp_sex,string"`
	IsWorker bool `json:"isWorker"`
	// IsWorker bool `json:"-"`
}

// Main1108 entry
func Main1108() {
	// jsontest()
	gotest()
}

func appendtest() {
	arr := [...]byte{1, 2, 3, 4, 5}
	fmt.Println(arr)

	b := append(make([]byte, 0), arr[:]...)
	fmt.Println(b)

	b = append(b, 5, 6, 7, 8, 9)
	fmt.Println(b)
}

func jsontest() {
	bean := EmpBean{12, "Allen", 1, false}

	// jsonarr, _ := json.Marshal(bean)
	jsonarr, _ := json.MarshalIndent(bean, "", "\t")
	jsonStr := string(jsonarr)

	fmt.Println(bean)
	fmt.Println(jsonStr)

	var outBean EmpBean
	json.Unmarshal([]byte(`{"emp_id": 111,   "emp_name": "张三",
	   "emp_sex":0,
	   "isWorker":true}`), &outBean)
	fmt.Println(outBean)
}

func gotest() {
	var ws sync.WaitGroup

	chArr := []byte{'A', 'a'}

	for _, ch := range chArr {
		ws.Add(1)
		go func(c byte) {
			for count := 0; count < 3; count++ {
				for ch := c; ch < c+26; ch++ {
					fmt.Printf("%c ", ch)
				}
			}
			ws.Done()
		}(ch)
	}

	// ws.Add(1)
	// go func() {
	// 	for count := 0; count < 3; count++ {
	// 		for ch := 'A'; ch < 'A'+26; ch++ {
	// 			fmt.Printf("%c ", ch)
	// 		}
	// 	}
	// 	ws.Done()
	// }()

	// ws.Add(1)
	// go func() {
	// 	for count := 0; count < 3; count++ {
	// 		for ch := 'a'; ch < 'a'+26; ch++ {
	// 			fmt.Printf("%c ", ch)
	// 		}
	// 	}
	// 	ws.Done()
	// }()

	ws.Wait()

	fmt.Println("=======FIN========")
}
