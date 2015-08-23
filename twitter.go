package main
import (
    "fmt"
    "flag"
    "log"
    "io/ioutil"
   
   "github.com/mrjones/oauth"
)

func main(){
    fmt.Println("This is a twitter oauth example")
    
    var consumerKey *string = flag.String("consumerKey","","Consumer Key from Twitter")

    var consumerSecret *string = flag.String("consumerSecret","","Consumer Secret from Twitter")
    
    flag.Parse()
    fmt.Printf("Consumer Key: %v\nConsumerSecret : %v",*consumerKey,*consumerSecret)

    c := oauth.NewConsumer(
		*consumerKey,
		*consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
    c.Debug(true)

    requestToken,u,err := c.GetRequestTokenAndUrl("oob")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(requestToken)
    fmt.Println("(1) Go to: " + u)
    fmt.Println("(2) Click on Authorize app : It will provide you a verification code")
    fmt.Println("(3) Enter the verification code here")

    verificationCode:=""
    fmt.Scanln(&verificationCode);

    accessToken,err := c.AuthorizeToken(requestToken,verificationCode)
    if err != nil {
        log.Fatal(err)
    }
    client, err := c.MakeHttpClient(accessToken)
    
    if err != nil {
        log.Fatal(err)
    }

    response, err := client.Get(
            "https://api.twitter.com/1.1/statuses/home_timeline.json?count=1")

    defer response.Body.Close()

    bits, err := ioutil.ReadAll(response.Body)

    fmt.Println("The newest item in your home timeline is: " + string(bits))

}
