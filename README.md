# govue-razors

A simple example of a full-stack app in Go using embedded Vue.js.

The Go server runs a rudimentary REST API (using gorilla/mux) which returns a list of philosophical razors, and the Vue app calls the API and displays the razors.

Inspired by https://hackandsla.sh/posts/2021-06-18-embed-vuejs-in-go/

## Development

Start the Vue service in watch mode:
```
cd frontend
npm run watch
```

In a separate terminal window, start the Go app in `dev` mode (so it can reload the frontend changes):
```
go run . -dev
```

## General Embedded UI Setup in Go

  - Create your Go project as usual
  - Generate a new Vue (2.x) project using the [Vue CLI](https://cli.vuejs.org)
```
vue create frontend
```
  - Modify Vue app as needed, then generate assets
```
cd frontend
npm run build
```
  - In your Go code, use the _embed_ package to embed and serve the assets (see [server.go](server.go))
```
//go:embed frontend/dist
var frontend embed.FS
```
  - In your Vue app, use axios to call the Go API routes as needed

## Author

Tenyo Grozev (tenyo.grozev@gmail.com)
