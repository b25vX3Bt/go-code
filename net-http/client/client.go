package main

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "io/ioutil"
)

func get_2(i int, q chan int) {
    url := "http://127.0.0.1:8080/"

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("X-Hoge", "hogehoge")
    req.Header.Set("X-Number", fmt.Sprintf("%d",i) )

    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Printf("%s", dump)


    client := new(http.Client)
    resp, _ := client.Do(req)
    defer resp.Body.Close()

    dumpResp, _ := httputil.DumpResponse(resp, true)
    fmt.Printf("%s\n", dumpResp)

    q <- i
    //byteArray, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(byteArray))
}

func get_1() {
    url := "http://127.0.0.1:8080/"

    resp, _ := http.Get(url)

    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byteArray)) // htmlをstringで取得
}

func main() {

    n := 5
    queue := make(chan int, 100)

    for i := 0; i < n; i++ {
        go get_2(i,queue)
    }
    for i := 0; i < n; i++ {
        val := <-queue
        fmt.Printf("recv [%d]\n",val)
    }
}

