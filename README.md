# URL Shortening Service

Proyect based on [Url Shortening Service of roadmap.sh](https://roadmap.sh/projects/url-shortening-service)

## Before you start
You should have MySQL or MariaDB installed and change the dsn variable of main.go according to your configuration.

### Other Requeriments
- GORM
- GIN
- [ShortID](github.com/teris-io/shortid)
## How to use

- Run the proyect
- Go to localhost:8080

Currently the user interface is minimal with a small form to shorten a URL.

According to the proyect requirement here are the functionalities
### Create a Short URL
Just fill the form field and press the button (Must be an complete url like "http://example-url.com")

### Retrieve original URL
Write the short url like "localhost:8080/abc123". Currently you should get the JSON with the data and not being redirected, I haven't implemented the actual front end for the user yet.

### Update Short URL
Via terminal

> curl -X PUT -H "Content-Type: application/json" -d '{"url": "https://github.com/long-url"}' localhost:8080/shorten/abc123

### Delete Short URL
Via terminal

> curl -X DELETE http://localhost:8080/shorten/abc123

### Get URL Statistics
Write at the direcction bar "localhost:8080/abc123/stats"

## Additional notes

On the project page from roadmap.sh, the example short codes have a length of 6, but I used the short IDs from the [teris-io](https://github.com/teris-io) package, which generate short IDs with a length of 9. The solution worked so well at the beginning that I decided to continue using this length. Additionally, the requirements are not strict about this aspect, so I left it as it is.

I am aware that there are still many details to refine, but at this stage, the project requirements are fulfilled.