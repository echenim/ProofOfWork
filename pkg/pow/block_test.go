package pow

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	tests := []struct {
		name     string
		block    Block
		expected string
	}{
		{
			name: "Empty block",
			block: Block{
				Index:     0,
				Timestamp: "test-time",
				Data:      "",
				PrevHash:  "",
				Nonce:     0,
			},
			expected: "d577273ff885c3f84f48e8059022aef53e6e145acb0c662983ef14fae9a8e509",
		},
		{
			name: "Non-empty block",
			block: Block{
				Index:     1,
				Timestamp: "test-time",
				Data:      "some data",
				PrevHash:  "somehash",
				Nonce:     1,
			},
			expected: "b2213295d564916f89a6a42455567c87c3f480fcd7a1c15e220f17d7169a790b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.block.calculateHash(); got != tt.expected {
				t.Errorf("Block.calculateHash() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsValidHash(t *testing.T) {
	tests := []struct {
		name       string
		hash       string
		difficulty int
		expected   bool
	}{
		{
			name:       "Valid hash low difficulty",
			hash:       "000abc",
			difficulty: 3,
			expected:   true,
		},
		{
			name:       "Invalid hash low difficulty",
			hash:       "123abc",
			difficulty: 3,
			expected:   false,
		},
		{
			name:       "Valid hash high difficulty",
			hash:       "000000abc",
			difficulty: 6,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := Block{Hash: tt.hash}
			if got := block.isValidHash(tt.difficulty); got != tt.expected {
				t.Errorf("Block.isValidHash() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockchain()
	tests := []struct {
		name       string
		data       string
		difficulty int
		wantLength int
	}{
		{
			name:       "Add one block",
			data:       "first block",
			difficulty: 2, // Reduced difficulty for testing purposes
			wantLength: 2,
		},
		{
			name:       "Add two blocks",
			data:       "second block",
			difficulty: 2,
			wantLength: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc.AddBlock(tt.data, tt.difficulty)
			if got := len(bc.blocks); got != tt.wantLength {
				t.Errorf("len(Blockchain.blocks) = %v, want %v", got, tt.wantLength)
			}
		})
	}
}
