package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/broothie/jog"
	"github.com/broothie/jog/promise"
	"github.com/samber/lo"
)

func main() {
	jog.Mount("root", func(state *jog.State) jog.Generator {
		state.Set("results", []string{})

		return jog.Div(nil,
			jog.Input(jog.Attr{"onkeyup": jog.Listener(func(event jog.Event) any {
				query := event.JSValue.Get("target").Get("value").String()

				promise.From(func() (*http.Response, error) {
					return http.Get(fmt.Sprintf("/api/users?query=%s", query))
				}).
					Then(func(response *http.Response) {
						var results []string
						if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
							jog.Console.Log(err)
						}

						state.Set("results", results)
					})

				return nil
			})}),
			jog.Div(nil, state.Watch("results", func(value any) jog.Generator {
				results := value.([]string)
				return jog.Div(nil,
					lo.Map(results, func(s string, i int) jog.Generator {
						return jog.P(nil, jog.Text(s))
					})...,
				)
			})),
		)
	})

	jog.Wait()
}
