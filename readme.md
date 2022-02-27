# Casserole memes

Silly site to generate casserole memes from existing memes.


# required dependencies
1. [go](https://go.dev/dl/) (at least 1.17)
2. [air](https://github.com/cosmtrek/air) (optional for hot reloading during dev)

# development
just run `air` and start changing main.go

if air doesn't work (sometimes it dosen't on windows)
then just make changes and run the following script
<br>
`GOARCH=wasm GOOS=js go build -o ./web/app.wasm && go run main.go`
