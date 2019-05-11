package dante

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	datasetPath = "test_dataset.txt"
)

func TestParseDataset(t *testing.T) {
	fmt.Println(os.Getwd())
	docs, err := ParseDataset(datasetPath)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, 6, len(docs))
}
