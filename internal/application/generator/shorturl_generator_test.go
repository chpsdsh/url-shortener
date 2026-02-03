package generator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLink(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2"
	shortLink := GenerateShortLink(initialLink, UserId)
	fmt.Println(shortLink)
	assert.Equal(t, "SUbyVBdbmF7", shortLink)
}
