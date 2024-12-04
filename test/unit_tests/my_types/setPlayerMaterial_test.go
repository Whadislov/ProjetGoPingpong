package my_types_test

import (
	"fmt"
	"github.com/Whadislov/ProjetGoPingPong/internal/my_types"
	"testing"
)

func TestSetPlayerMaterial(t *testing.T) {

	var julien my_types.Player
	julien.Name = "Julien"
	expectedMaterial := []string{"Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood"}

	t.Run(fmt.Sprintf("Set player material of %s", julien.Name), func(t *testing.T) {
		julien.SetPlayerMaterial("Victas V20 double Extra", "Victas V20 double Extra", "Koki Niwa Wood")
		for i := range expectedMaterial {
			if expectedMaterial[i] != julien.Material[i] {
				t.Errorf("Material of %s is currently %s and is expected to be %s", julien.Name, julien.Material[i], expectedMaterial[i])
			} else {
				fmt.Printf("Material of %s is currently %s and is expected to be %s", julien.Name, julien.Material[i], expectedMaterial[i])
			}
		}
	})
}
