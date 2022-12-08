package cache

import (
	"calculator/pkg/calc/command"
	"calculator/pkg/calc/operation"
)

func NewCacheCmdBuilder(builder command.CmcBuilder, cache CacheStorage) command.CmcBuilder {
	return &cacheCmdBuilder{
		builder: builder,
		cache:   cache,
	}
}

type cacheCmdBuilder struct {
	builder command.CmcBuilder
	cache   CacheStorage
}

func (c cacheCmdBuilder) CreateFnIdCommand(id, id2 string) *command.Command {
	cmd := c.builder.CreateFnIdCommand(id, id2)
	ccmd := NewCacheCommandFn(id, []string{id2}, *cmd, c.cache)
	return &ccmd
}

func (c cacheCmdBuilder) CreateFnOpCommand(id, id2 string, op operation.Operation, id3 string) *command.Command {
	cmd := c.builder.CreateFnOpCommand(id, id2, op, id3)
	ccmd := NewCacheCommandFn(id, []string{id2, id3}, *cmd, c.cache)
	return &ccmd
}
