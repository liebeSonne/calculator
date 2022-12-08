package cache

import (
	"calculator/pkg/calc/command"
)

func NewCacheCommandFn(id string, deps []string, cmd command.Command, cache CacheStorage) command.Command {
	cache.SetDependencies(id, deps)
	return &cacheCommandFn{
		deps:  deps,
		cmd:   cmd,
		cache: cache,
	}
}

type cacheCommandFn struct {
	cmd   command.Command
	id    string
	deps  []string
	cache CacheStorage
}

func (c cacheCommandFn) Execute() {
	c.cmd.Execute()
}
