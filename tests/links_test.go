package iiq_test

import "testing"

func TestLinkAssets(t *testing.T) {
	parent := "f27780c8-9495-495f-baa2-8cd267b9eaca"
	children := []string{"eddd51cc-986f-440f-86c1-5bd029438a6d", "e2c3255d-01b2-45a5-a8ef-97633c648335"}
	result, err := client.LinkAssets(parent, children...)
	if err != nil {
		t.Error(err)
	}

	if result != "Assets successfully linked" {
		t.Errorf("Link message = '%s'; expected 'Assets successfully linked'", result)
	}
}
