package shape_test

import (
	"shapes/shape/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestShapeOperations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockshape := mocks.NewMockIShape(ctrl)
	r := float32(20.0)
	e := float32(20.0)
	mockshape.EXPECT().Area().Return(r)
	area := mockshape.Area()
	assert.Equal(t, e, area, "Area should be 20.0")
}
