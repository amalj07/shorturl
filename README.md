# shorturl
Url shortener written in Golang, create custom shorturl by passing custom shortids

# Install
```
go get github.com/amljs/shorturl
```

# Usage
### Default shorturls
```
url, err := shorturl.CreateShortURL(longurl, baseurl)
```

### Custom shorturls
```
url, err := shorturl.CreateShortURL(longurl, baseurl, shortid)
```