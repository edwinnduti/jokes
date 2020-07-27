# JOKES Corner

A consumed [jokeAPI](https://sv443.net/jokeapi/v2/)  that render randomly choosen jokes in json format.


install:
```
go get github.com/edwinnduti/jokes
```

Example:
```go
package main
import(
	"fmt"
	"github.com/edwinnduti/jokes"
)

func main(){
	fmt.Println("IT'S JOKES TIME GUYS!")

	jokes.Get("Programming")

}

```

Enjoy!

