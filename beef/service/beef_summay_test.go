package service

import (
	mockapi "lottery-plus/beef/repository/api/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBeefSummary_WhenCallSuccess_ShouldReturnData(t *testing.T) {
	mockCtl := gomock.NewController(t)
	repo := mockapi.NewMockAPI(mockCtl)

	dataExpected := make(map[string]int)
	dataExpected["bresaola"] = 1
	dataExpected["enim"] = 1
	dataExpected["fatback"] = 1
	dataExpected["jowl"] = 1
	dataExpected["meatloaf"] = 1
	dataExpected["pastrami"] = 1
	dataExpected["pork"] = 1
	dataExpected["t-bone"] = 1
	expected := BeefSummaryResponse{
		Beef: dataExpected,
	}

	repo.EXPECT().GetText().Return("bresaola enim fatback jowl meatloaf pastrami pork t-bone", nil)
	s := New(repo)
	actual, err := s.BeefSummary()

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
