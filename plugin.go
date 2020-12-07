package main

import (
	"fmt"
)

var LoadedPlugins = make(map[string]Plugin)

type Plugin interface {
	Run(string) error
	Register()
}

type Hello struct {
	Name string
}

func (h Hello) Run(s string) error {
	fmt.Println("Hello,", s)
	return nil
}
func (h Hello) Register() {
	LoadedPlugins[h.Name] = h
}

func main() {
	h := Hello{Name: "hello"}
	Plugin.Register(h)
	fmt.Println("Loaded Plugins:", len(LoadedPlugins))
	if plug, ok := LoadedPlugins["hello"]; ok {
		plug.Run("GO plugins")
	}
}