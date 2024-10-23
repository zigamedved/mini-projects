package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
)

type node struct {
	mu sync.Mutex
	kv map[string]string
}

func newNode() *node {
	return &node{
		kv: make(map[string]string),
	}
}

func (n *node) Apply(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.kv[key] = value
}

func (n *node) Snapshot(ctx context.Context) (io.ReadCloser, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	buf, err := json.Marshal(&n.kv)
	if err != nil {
		fmt.Printf("error marshaling node kv, ctx: %v", ctx)
		return nil, err
	}

	return io.NopCloser(strings.NewReader(string(buf))), nil
}

func (n *node) Restore(ctx context.Context, r io.ReadCloser) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	buf, err1 := io.ReadAll(r)
	if err1 != nil {
		fmt.Printf("error reading node kv, ctx: %v", ctx)
		return err1
	}

	err2 := json.Unmarshal(buf, &n.kv)
	if err2 != nil {
		fmt.Printf("error unmarshalling node kv, ctx: %v", ctx)
		return err2
	}

	return r.Close()
}

func (n *node) Read(key string) string {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.kv[key]
}
