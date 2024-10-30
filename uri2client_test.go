package qdrantURItoClient

import (
	"errors"
	"testing"

	"github.com/qdrant/go-client/qdrant"
)

func TestUriToClient(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		want    *qdrant.Client
		wantErr error
	}{
		{
			name:    "Valid URI",
			uri:     "qdrant://1234567890@localhost:6333?UseTLS=1",
			want:    &qdrant.Client{}, // You need to create a mock client for this test
			wantErr: nil,
		},
		{
			name:    "Wrong URI",
			uri:     "wrong://1234567890@localhost:6333?UseTLS=1",
			want:    nil,
			wantErr: errors.New("wrong protocol, support only 'qdrant://'"),
		},
		{
			name:    "Empty Host",
			uri:     "qdrant://1234567890@:6333?UseTLS=1",
			want:    nil,
			wantErr: errors.New("Empty host"),
		},
		{
			name:    "Invalid UseTLS",
			uri:     "qdrant://1234567890@localhost:6333?UseTLS=invalid",
			want:    &qdrant.Client{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UriToClient(tt.uri)
			if err != nil && tt.wantErr == nil {
				t.Errorf("UriToClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("UriToClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("UriToClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want == nil {
				t.Errorf("UriToClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
