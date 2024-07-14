package main 

import ( 
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    var url string 
    fmt.Println("nhập url cần bypass: ")
    _, err := fmt.Scan(&url)
    var re = "https://dhphuoc207.xyz/bypassclf.php?host="+url
    
	resp, err := http.Get(re)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))
}
