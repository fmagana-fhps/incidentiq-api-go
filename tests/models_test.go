package iiq_test

import (
	"testing"
)

func TestModelById(t *testing.T) {
	id := "fdf1109c-7b3f-287b-7b56-4a0d46c2e046"
	model, err := client.ModelById(id)
	if err != nil {
		t.Error(err)
	}

	if model.Body.Item.ModelID != id {
		t.Errorf("ModelId = %s; expected %s", model.Body.Item.ModelID, id)
	}
}
