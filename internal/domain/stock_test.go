package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistsTicker_Exists(t *testing.T) {
	// Arrange
	ticker := Ticker("CRM")

	// Act
	exists := ExistsTicker(ticker)

	// Assert
	assert.True(t, exists)
}

func TestExistsTicker_NotExists(t *testing.T) {
	// Arrange
	ticker := Ticker("UNEXPECTED")

	// Act
	exists := ExistsTicker(ticker)

	// Assert
	assert.False(t, exists)
}
