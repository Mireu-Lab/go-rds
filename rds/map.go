package rds

import (
	"fmt"
	"sync"
)

type Rds struct {
	Data map[string]map[string]string
	mu   sync.RWMutex
}

// 전역 RDS 인스턴스
var GlobalRds = &Rds{
	Data: make(map[string]map[string]string),
}

func (r *Rds) Set(channel, key, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Data == nil {
		r.Data = make(map[string]map[string]string)
	}

	if r.Data[channel] == nil {
		r.Data[channel] = make(map[string]string)
	}

	r.Data[channel][key] = value
	return nil
}

func (r *Rds) Remove(channel, key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Data == nil || r.Data[channel] == nil {
		return fmt.Errorf("channel '%s' does not exist", channel)
	}

	if _, exists := r.Data[channel][key]; !exists {
		return fmt.Errorf("key '%s' does not exist in channel '%s'", key, channel)
	}

	delete(r.Data[channel], key)

	// Remove channel if it becomes empty
	if len(r.Data[channel]) == 0 {
		delete(r.Data, channel)
	}

	return nil
}

func (r *Rds) List(channel string) (map[string]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.Data == nil || r.Data[channel] == nil {
		return nil, fmt.Errorf("channel '%s' does not exist", channel)
	}

	result := make(map[string]string)
	for k, v := range r.Data[channel] {
		result[k] = v
	}

	return result, nil
}

func (r *Rds) Get(channel, key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.Data == nil || r.Data[channel] == nil {
		return "", fmt.Errorf("channel '%s' does not exist", channel)
	}

	value, exists := r.Data[channel][key]
	if !exists {
		return "", fmt.Errorf("key '%s' does not exist in channel '%s'", key, channel)
	}

	return value, nil
}

func (r *Rds) ListChannels() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.Data == nil {
		return []string{}
	}

	channels := make([]string, 0, len(r.Data))
	for channel := range r.Data {
		channels = append(channels, channel)
	}

	return channels
}

// 편의 함수들 - 전역 인스턴스를 사용
func Set(channel, key, value string) error {
	return GlobalRds.Set(channel, key, value)
}

func Remove(channel, key string) error {
	return GlobalRds.Remove(channel, key)
}

func List(channel string) (map[string]string, error) {
	return GlobalRds.List(channel)
}

func Get(channel, key string) (string, error) {
	return GlobalRds.Get(channel, key)
}

func ListChannels() []string {
	return GlobalRds.ListChannels()
}
