package _menu

import (
	"bufio"
	"fmt"
	"os"
)

type Menu interface {
	AddItem(name, description string, command Command)
	GetItemCommand(name string) *Command
	Run()
}

func NewMenu() Menu {
	return &menu{}
}

type menu struct {
	items []Item
}

func (m *menu) AddItem(name, description string, command Command) {
	item := NewItem(name, description, command)
	m.items = append(m.items, item)
}

func (m *menu) GetItemCommand(name string) *Command {
	for _, item := range m.items {
		if item.Name() == name {
			command := item.Command()
			return &command
		}
	}
	return nil
}

func (m *menu) Run() {

//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//
//	for {
//		var str string
//		fmt.Fscan(in, &str)
//
//
//		fmt.Fprintln(out, output)
//	}
//}
