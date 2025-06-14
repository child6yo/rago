package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/child6yo/rago/services/storage/internal"
	"github.com/child6yo/rago/services/storage/internal/app/repository/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleDocMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expError bool
		mockFunk func(*mock.VectorDB)
	}{
		{
			name:     "OK",
			input:    `{"content":"test content","metadata":{"url":"test.com"}}`,
			expError: false,
			mockFunk: func(mvd *mock.VectorDB) {
				mvd.PutDocumentFunc = func(ctx context.Context, docs internal.Document) error {
					return nil
				}
			},
		},
		{
			name:     "empty message",
			input:    `{}`,
			expError: true,
			mockFunk: func(mvd *mock.VectorDB) {
				mvd.PutDocumentFunc = func(ctx context.Context, docs internal.Document) error {
					return nil
				}
			},
		},
		{
			name:     "invalid data",
			input:    `{"c":"123","mt":"123.com"}`,
			expError: true,
			mockFunk: func(mvd *mock.VectorDB) {
				mvd.PutDocumentFunc = func(ctx context.Context, docs internal.Document) error {
					return nil
				}
			},
		},
		{
			name:     "database error",
			input:    `{"content":"test content","metadata":{"url":"test.com"}}`,
			expError: true,
			mockFunk: func(mvd *mock.VectorDB) {
				mvd.PutDocumentFunc = func(ctx context.Context, docs internal.Document) error {
					return errors.New("error")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &mock.VectorDB{}
			tt.mockFunk(mockDB)
			loader := NewLoader(mockDB)

			msg := []byte(tt.input)

			err := loader.LoadDocument(msg)
			if tt.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
