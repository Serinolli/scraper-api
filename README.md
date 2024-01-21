# scraper-api
Study repo. Project integrated with [WebScraper](https://github.com/Serinolli/RedditScraper). <br>
Simple REST API designed for managing data obtained through the associated scraper tool.

## Setting up...

### Installing MongoDB with Docker
```
docker run --name reddit-scraper -p 27017:27017 -d mongo
```

## Testing
```
go run api/main.go
```
