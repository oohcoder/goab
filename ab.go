package main

import "fmt"
import "flag"
import "sync"
import "time"
import "strings"
import "net/http"
import "io/ioutil"

func main() {
	c := flag.Int("c", 1, "concurrency num")
	n := flag.Int("n", 1, "total request num")
	d := flag.String("d", "", "total request num")
	flag.Parse()

    var total int
    var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(*c)
	for i := 0; i < *c; i++ {
		go func() {
            for ;total <*n; {
                downFlag := false
                mutex.Lock()
                if total < *n {
                    downFlag = true
                    total++
                }
                mutex.Unlock()
                if downFlag {
                    downWeb(*d)
                }
            }
			defer wg.Done()
		}()

	}
	wg.Wait()

	//fmt.Println(*c)
	//fmt.Println(*n)
	//fmt.Println(*d)
}
func downWeb(url string) {
     resp, err := http.Get(url)
     if err != nil {
         fmt.Printf("%s\t%d\t%s\n", time.Now().Format("20060102150405"), 800, err)
         return
     }
     defer resp.Body.Close()
     body, err := ioutil.ReadAll(resp.Body)
     if err != nil {
         fmt.Printf("%s\t%d\t%s\n", time.Now().Format("20060102150405"), resp.StatusCode, err)
         return
     } 
     res := strings.Replace(string(body), "\n", "\\n", -1)
     fmt.Printf("%s\t%d\t%s\n", time.Now().Format("20060102150405"), resp.StatusCode, res)
}
