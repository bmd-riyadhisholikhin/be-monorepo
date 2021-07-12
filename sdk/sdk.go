// Code generated by candi v1.5.32. DO NOT EDIT.

package sdk

import (
	"sync"

	// @candi:serviceImport
	"monorepo/sdk/shark"
)

// Option func type
type Option func(*sdkInstance)

var (
	sdk  SDK
	once sync.Once
)

// SetGlobalSDK constructor with each sdk service option.
func SetGlobalSDK(opts ...Option) {
	s := new(sdkInstance)
	for _, o := range opts {
		o(s)
	}
	once.Do(func() {
		sdk = s
	})
}

// GetSDK get global sdk instance
func GetSDK() SDK {
	return sdk
}

// @candi:construct

// SetShark option func
func SetShark(shark shark.Shark) Option {
	return func(s *sdkInstance) {
		s.shark = shark
	}
}

// SDK instance abstraction
type SDK interface {
	// @candi:serviceMethod
	Shark() shark.Shark
}

// sdkInstance implementation
type sdkInstance struct {
	// @candi:serviceField
	shark	shark.Shark
}

// @candi:instanceMethod
func (s *sdkInstance) Shark() shark.Shark {
	return s.shark
}
