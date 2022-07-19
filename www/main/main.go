package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/broothie/gh"
	"github.com/samber/lo"
)

func main() {
	gh.Mount("root", func(state *gh.State) gh.Generator {
		state.Set("results", []string{})

		return gh.Div(nil,
			gh.Input(gh.Attr{"onkeyup": gh.Listener(func(event gh.Event) any {
				query := event.Target().Value().String()

				go func() {
					response, err := http.Get(fmt.Sprintf("/api/users?query=%s", query))
					if err != nil {
						gh.Console.Error(err)
						return
					}

					var results []string
					if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
						gh.Console.Error(err)
						return
					}

					state.Set("results", results)
				}()

				return nil
			})}),
			gh.Div(nil, state.Watch("results", func(value any) gh.Generator {
				results := value.([]string)
				return gh.Div(nil,
					lo.Map(results, func(s string, i int) gh.Generator {
						return gh.P(nil, gh.Text(s))
					})...,
				)
			})),
		)
	})

	gh.Wait()
}
