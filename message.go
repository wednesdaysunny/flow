package flow

import (
	"context"
	"time"
)

type Message struct {
	ctx     context.Context
	headers map[string]string
	data    interface{}
	ts      time.Time
}

func NewMessage(ctx context.Context) *Message {
	return &Message{
		ctx:     ctx,
		headers: make(map[string]string),
		ts:      time.Now(),
	}
}

func (m *Message) AddHeader(key string, val string) *Message {
	m.headers[key] = val
	return m
}

func (m *Message) GetHeader(key string) string {
	res, ok := m.headers[key]
	if ok {
		return res
	}
	return ""
}

func (m *Message) SetData(data interface{}) *Message {
	m.data = data
	return m
}

func (m *Message) Data() interface{} {
	return m.data
}

func (m *Message) Headers() map[string]string {
	return m.headers
}

func (m *Message) SetContext(ctx context.Context) *Message {
	m.ctx = ctx
	return m
}

func (m *Message) Context() context.Context {
	return m.ctx
}

func (m *Message) TimeStamp() int64 {
	return m.ts.UnixMilli()
}
