package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var service, _ = NewStorageService()

func TestStorageServiceSaveAndRetrieval(t *testing.T) {
	url := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortUrl := "7DasdasW"

	err := service.SaveUrlMapping(shortUrl, url)
	assert.Nil(t, err)

	res, err := service.RetrieveOriginalUrl(shortUrl)
	assert.Nil(t, err)
	assert.Equal(t, url, res)
}
