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

## API Endpoints Documentation

### Get All Posts
- **URL:** `/posts`
- **Method:** GET
- **Description:** Retrieves all posts.

### Create Posts (main method used by the scraper)
- **URL:** `/posts`
- **Method:** POST
- **Description:** Creates a list of new posts. If some posts already exist in the database, updates them.

### Create Post
- **URL:** `/post`
- **Method:** POST
- **Description:** Creates a new post.

### Get Post
- **URL:** `/posts/{postId}`
- **Method:** GET
- **Description:** Retrieves a specific post by ID.

### Update Post
- **URL:** `/posts/{postId}`
- **Method:** PUT
- **Description:** Updates a specific post by ID.

### Delete Post
- **URL:** `/posts/{postId}`
- **Method:** DELETE
- **Description:** Deletes a specific post by ID.

