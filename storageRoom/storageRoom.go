package storageRoom

func GetCurrentStorageAmount(product string) (int, bool, error) {
	toolAmountConfig := map[string]int{
		"PA": 0,
		"PB": 0,
		"PC": 0,
		"PD": 0,
	}
	storageAmount, ok := toolAmountConfig[product]
	return storageAmount, ok, nil
}
