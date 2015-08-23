package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"
    "net/url"
               
    "github.com/mrjones/oauth"
)

func authorize() *http.Client{
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
    return client
}

func main(){
    twitterClient := authorize()
    
    response,err :=twitterClient.PostForm("https://api.twitter.com/1.1/statuses/update.json",url.Values{"status":{"Golang Example"}})

     defer response.Body.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Status Posted",response)
}
