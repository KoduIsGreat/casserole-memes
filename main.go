package main

import (
	"embed"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

//go:embed assets
var assets embed.FS

type casserole struct {
	app.Compo
	baking bool
	name   string
}
type s map[string]string

func (c *casserole) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Img().Styles(s{
				"width":  "100%",
				"height": "auto",
			}).
				Src("/assets/banner.png"),
		),
		app.Div().Body(
			app.Button().Type("button").Text("Bake").OnClick(func(ctx app.Context, e app.Event) {
				c.baking = !c.baking
			}),
			app.If(c.baking,
				app.Img().Styles(s{
					"display":      "block",
					"margin-left":  "auto",
					"margin-right": "auto",
					"width":        "40%",
				}).Src("/assets/oven.gif"),
			)))
}

func main() {
	// Components routing:
	app.Route("/", &casserole{})
	app.Route("/casserole", &casserole{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})
	http.Handle("/assets/", http.FileServer(http.FS(assets)))

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
