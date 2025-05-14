package examplecommon

type TaskInput struct {
	Key   string
	Value string
}

// This function computes simulates the vault setting behavior
func VaultSet(taskIndex uint32, input TaskInput) ([32]byte, error) {
	// TODO: complete
	return [32]byte{0}, nil
}
