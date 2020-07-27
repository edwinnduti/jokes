/*
JokesAPI sv443

Author : Edwin Nduti
Twitter: @Iamthe_edd

Description:
	A jokeAPI that renders a randomly choosen  joke request in json format.
	Categories of Jokes available are:
		Any, Programming, Miscellaneous,Dark
	
	All begin with a capital letter


ENJOY! AND HAVE FUN!

CIAO!
*/

package jokes

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//Categories
type Category struct{
	Any              string
	Miscellaneous    string
	Programming      string
	Dark             string
}

//response result
type Result struct{
	Category  string         `json:"category"`
	Type      string         `json:"type"`
	Setup     string         `json:"setup"`
	Delivery  string         `json:"delivery"`
	Flags     *BlackListFlag `json:"flags"`
	Id        uint           `json:"id"`
	Error     bool           `json:"error"`
}

//Blacklisted flasgs
type BlackListFlag struct{
	Nsfw      bool   `json:"nsfw"`
	Religious bool   `json:"religious"`
	Political bool   `json:"political"`
	Racist    bool   `json:"racist"`
	Sexist    bool   `json:"sexist"`
}


//GET endpoint 
func Get(category string) {

	//Categories available
	c := &Category{
                "Any",
                "Miscellaneous",                                  "Programming",                                    "Dark",
        }
	fmt.Println("CATEGORIES : ",*(c))
	fmt.Println("")

	//create a Client
	Client := &http.Client{}

	//make request
	req,err := http.NewRequest("GET",BASEURL+category,nil)
	Check(err)

	//Get response
	resp,err := Client.Do(req)
	Check(err)

	//close stream
	defer resp.Body.Close()

	//read bytes
	body,err := ioutil.ReadAll(resp.Body)
	Check(err)

	//unmarshal bytes to result struct
	result := &Result{}
	json.Unmarshal(body,&result)

	//Marshal result struct into json fomart
	resultInJson,err := json.MarshalIndent(result,"","  ")
	Check(err)


	//Display results
	fmt.Println(string(resultInJson))

}

//error handler
func Check(err error)  {
	if err != nil{
		fmt.Println("Error :",err)
	}

}

//Base url
const(
	BASEURL = "https://sv443.net/jokeapi/v2/joke/"
)
