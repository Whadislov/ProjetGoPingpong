package my_types

func (t *Team) IsEmpty() bool {
	return t.Name == "" &&
		t.PlayerList == nil
}
