package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/child6yo/rago/services/storage/internal/pkg/database/mock"
	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/schema"
)

func TestHandleDocMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expError bool
		mockFunk func(*mock.MockVectorDB)
	}{
		{
			name:     "OK",
			input:    `{"content":"test content","metadata":{"url":"test.com"}}`,
			expError: false,
			mockFunk: func(mvd *mock.MockVectorDB) {
				mvd.PutFunc = func(ctx context.Context, docs []schema.Document) error {
					return nil
				}
			},
		},
		{
			name:     "empty message",
			input:    `{}`,
			expError: true,
			mockFunk: func(mvd *mock.MockVectorDB) {
				mvd.PutFunc = func(ctx context.Context, docs []schema.Document) error {
					return nil
				}
			},
		},
		{
			name:     "invalid data",
			input:    `{"c":"123","mt":"123.com"}`,
			expError: true,
			mockFunk: func(mvd *mock.MockVectorDB) {
				mvd.PutFunc = func(ctx context.Context, docs []schema.Document) error {
					return nil
				}
			},
		},
		{
			name:     "database error",
			input:    `{"content":"test content","metadata":{"url":"test.com"}}`,
			expError: true,
			mockFunk: func(mvd *mock.MockVectorDB) {
				mvd.PutFunc = func(ctx context.Context, docs []schema.Document) error {
					return errors.New("error")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &mock.MockVectorDB{}
			tt.mockFunk(mockDB)
			handler := NewDocHandlerService(mockDB)

			msg := []byte(tt.input)

			err := handler.HandleDocMessage(msg)
			if tt.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
