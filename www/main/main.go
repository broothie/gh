package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/broothie/gh"
	"github.com/broothie/gh/model"
	"github.com/broothie/gh/store"
	"github.com/samber/lo"
)

func main() {
	userStore := store.New(make(map[string]model.User))

	gh.Mount("root", gh.Div(nil,
		gh.Input(gh.Attr{"onkeyup": gh.Listener(func(event gh.Event) any {
			query := event.Target().Value().String()

			go func() {
				response, err := http.Get(fmt.Sprintf("/api/users?query=%s", query))
				if err != nil {
					fmt.Println(err)
					return
				}

				var users []model.User
				if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
					fmt.Println(err)
					return
				}

				userStore.Update(lo.Reduce(users, func(users map[string]model.User, user model.User, _ int) map[string]model.User {
					users[user.ID] = user
					return users
				}, make(map[string]model.User)))
			}()

			return nil
		})}),
		userStore.Watch(
			store.WatchMapLengthChanged[string, model.User],
			func(users map[string]model.User) gh.Generator {
				return gh.Div(nil,
					lo.Map(lo.Keys(users), func(id string, _ int) gh.Generator {
						return store.WatchSelection(userStore,
							func(users map[string]model.User) model.User { return users[id] },
							userEntry,
						)
					})...,
				)
			},
		),
	))

	gh.Wait()
}

func userEntry(user model.User) gh.Generator {
	return gh.Div(gh.Attr{
		"class": gh.Class{
			"d-flex flex-row": true,
		},
	},
		gh.P(gh.Attr{"class": "pe-1"}, gh.Text(user.ID)),
		gh.P(gh.Attr{"class": "pe-1"}, gh.Text(user.Username)),
		gh.P(nil, gh.Text(user.Age)),
	)
}
