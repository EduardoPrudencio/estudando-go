package tests

import (
	"strings"
	"testing"
)

const msgs = "%s (parte: %s) - valores esperados (%d) <> encontrado: (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		text     string
		part     string
		expected int
	}{
		{"Go é muito legal", "Go", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Eduardo", "u", 2},
	}

	for _, test := range testes {
		t.Logf("Massa: %v", testes)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(msgs, test.text, test.part, test.expected, actual)
		}

	}
}
