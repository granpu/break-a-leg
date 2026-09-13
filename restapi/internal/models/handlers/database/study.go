package main
import (
	"fmt"
	"net/http"
	"sync"
)

func main(){

urls := []string{
//список ссылок
}

var wg sync.WaitGroup

for _, url := range urls{
wg.Add(1)

go func(url string){
	defer wg.Done()

	res, err := http.Get(url){
		if err != nil{
			fmt.Printf("%s is not Ok: %v\n", url, err)
			return
		}
		defer res.Body.Close()

		if res.StatusCode == 200{
			fmt.Printf("200 OK")
		} else {
			fmt.Printf("200 not OK")
		}
	}
}(url)
}
wg.Wait()
}