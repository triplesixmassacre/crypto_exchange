package eth

import "testing"

func TestImportWallet(t *testing.T) {
	tests := []struct {
		name        string
		inputKey    string
		errExpected bool
	}{
		{"Пустая строка", "", false},
		{"Специальные символы", "\t		\n", false},
		{"Специальные символы и запрещенная буква", "\tx		\n", true},
		{"Специальные символы и разрешенная буква", "\ta		\n", true},
		{"Нормальный ключ", "d34157f1b67603f397b43b7aa4142e63aaac8ca9cfe27b31ab1f4ae1e50dda17", false},
		{"Выглядит как нормальный, содержит запрещенную букву", "d34157f1b67603f397b43b7aa4142e63aGac8ca9cfe27b31ab1f4ae1e50dda17", true},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := preparePrivateKey(tt.inputKey)
				if (err != nil) != tt.errExpected {
					t.Errorf("Произошла ошибка: %v", err)
				}
			},
		)
	}
}
