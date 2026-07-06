package caching

import (
	"golang/globals"
	"golang/structs"
)

func CacheFetchBlocks() map[string]*structs.Block {
	Block := make(map[string]*structs.Block, 0)
	for key, val := range globals.CacheBlocks {
		Block[key] = &structs.Block{Partial: val.Partial, Full: val.Full}
	}
	return Block
}
