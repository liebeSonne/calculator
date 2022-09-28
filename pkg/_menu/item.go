package _menu

type Item interface {
	Name() string
	Description() string
	Command() Command
}

func NewItem(name, description string, command Command) Item {
	return &item{
		name:        name,
		description: description,
		command:     command,
	}
}

type item struct {
	name        string
	description string
	command     Command
}

func (i *item) Name() string {
	return i.name
}

func (i *item) Description() string {
	return i.description
}

func (i *item) Command() Command {
	return i.command
}
