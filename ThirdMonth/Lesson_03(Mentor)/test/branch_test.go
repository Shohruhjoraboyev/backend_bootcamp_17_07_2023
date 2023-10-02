package test

import (
	"fmt"
	"net/http"
	"projectWithGinAndSwagger/models"
	"strconv"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestCreateBranch(t *testing.T) {
	response := models.Branch{}
	yearInt, _ := strconv.Atoi(faker.Date())

	request := &models.CreateBranch{
		Name:      faker.FirstName(),
		Address:   faker.Paragraph(),
		FoundedAt: yearInt,
	}

	resp, err := makeRequest(http.MethodPost, "/branch", request, &response)
	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, 201, resp.StatusCode)
	}

	fmt.Println("Person updated successfully")

}
