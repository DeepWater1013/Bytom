package txdb

import (
	"errors"
	"fmt"
	"sync"

	"github.com/golang/groupcache/lru"
	"github.com/golang/groupcache/singleflight"

	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/legacy"
)

const maxCachedBlocks = 30

func newBlockCache(fillFn func(hash *bc.Hash) *legacy.Block) blockCache {
	return blockCache{
		lru:    lru.New(maxCachedBlocks),
		fillFn: fillFn,
	}
}

type blockCache struct {
	mu     sync.Mutex
	lru    *lru.Cache
	fillFn func(hash *bc.Hash) *legacy.Block
	single singleflight.Group
}

func (c *blockCache) lookup(hash *bc.Hash) (*legacy.Block, error) {
	if b, ok := c.get(hash); ok {
		return b, nil
	}

	block, err := c.single.Do(hash.String(), func() (interface{}, error) {
		b := c.fillFn(hash)
		if b == nil {
			return nil, errors.New(fmt.Sprintf("There are no block with given hash %s", hash.String()))
		}

		c.add(b)
		return b, nil
	})
	if err != nil {
		return nil, err
	}
	return block.(*legacy.Block), nil
}

func (c *blockCache) get(hash *bc.Hash) (*legacy.Block, bool) {
	c.mu.Lock()
	block, ok := c.lru.Get(hash)
	c.mu.Unlock()
	if block == nil {
		return nil, ok
	}
	return block.(*legacy.Block), ok
}

func (c *blockCache) add(block *legacy.Block) {
	c.mu.Lock()
	c.lru.Add(block.Hash(), block)
	c.mu.Unlock()
}
