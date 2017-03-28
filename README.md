# Twall - MIDDLE

Live wall of tweets for the #devfestlille 2017
To combine with [Twall - FRONT](https://github.com/fgruchala/twall-front-vuejs)

## Made with ...
* Love
* [GoLang](https://golang.org/) 
* [GoDep](https://github.com/tools/godep)
* [Go-Twitter](https://github.com/dghubble/go-twitter)
* [Search API](https://dev.twitter.com/rest/public/search) from [Twitter](https://twitter.com/) 

## Requirement for Twitter
You have to generate token to access the Twitter API, in particular consumer key and secret. 
Go to https://apps.twitter.com/ and generate them.

## To develop
1. Initialize by `<WORKSPACE_GO_SRC_DIR>/github.com/fgruchala/twall-middle-go/$ godep go install`
2. Launch by `<WORKSPACE_GO_BIN_DIR>$ twall-middle-go -consumerKey=XXX -consumerSecret=XXX`

If you want a complete list of options, play `<WORKSPACE_GO_BIN_DIR>$ twall-middle-go -h`

## Contact
### Via GitHub > [Issues](https://github.com/fgruchala/twall-middle-go/issues)
Helpful for **question**, **bug** and **contribution request**.

### Via Twitter > François GRUCHALA [@FGruchala](https://twitter.com/FGruchala)
