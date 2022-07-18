package main

import "github.com/broothie/jog"

func main() {
	jog.Mount("root", jog.BuilderFunc(func(state *jog.State) *jog.Node {
		state.Set("query", "")

		return jog.Div(nil,
			jog.Input(jog.Attr{"onkeyup": jog.Listener(func(event jog.Event) any {
				state.Set("query", event.JSValue.Get("target").Get("value").String())
				return nil
			})}),
			jog.Div(nil, state.Watch("query", func(value any) *jog.Node {
				query := value.(string)
				return jog.P(nil, jog.Text(query))
			})),
		)
	}))

	jog.Wait()
}
