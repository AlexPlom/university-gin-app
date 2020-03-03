# University Gin App

University Gin App is a simple application to display how we can leverage the minimalistic [Gin framework](https://github.com/gin-gonic/gin) with Go in order to achieve web routing.

## Installation

To install University Gin App package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.11+ is required**), then you can use the below Go command to install University Gin App.

```sh
$ go get -u github.com/alexplom/university-gin-app
```

## Usage

- Navigate to the directory where you have `university-gin-app`
- run `go run main.go`

That's it, you should be up and running on `localhost:8080`. If you need to listen to a custom port you can change `r.Run()` to `r.Run(":1337")`.

## **Oooor You can try it online**

- **`GET`**
  `https://university-gin-app.herokuapp.com/`

- **`GET`**
  `https://university-gin-app.herokuapp.com/query?name=YourName&age=30`

- **`GET`**
  `https://university-gin-app.herokuapp.com/path/YourName/30`

- **`POST`**
  `https://university-gin-app.herokuapp.com/`
  ```json
  {
    "key": "value"
  }
  ```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
