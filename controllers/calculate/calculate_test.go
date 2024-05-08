package calculate

import "testing"

func checkValidToken(i, j Token) bool {
	if i.Type == j.Type && i.Value == j.Value {
		return true
	}
	return false
}

func checkValidTokenWithPosition(i, j Token) bool {
	if i.Type == j.Type && i.Value == j.Value {
		return true
	}
	return false
}
func TestTokenizer(t *testing.T) {

	type exprToToken struct {
		expression string
		tokens     []Token
		fun        func(Token, Token) bool
	}

	exprToTokens := map[string]exprToToken{
		"simpleNumber": {
			expression: "1",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "1",
				},
			},
			fun: checkValidToken,
		},
	}

	for name, exptt := range exprToTokens {
		outputTokens := tokenize(exptt.expression)
		if len(outputTokens) != len(exptt.tokens) {
			t.Errorf("Test for %v failed:\n different length of tokens.", name)
		}
		for i := 0; i < len(outputTokens); i++ {
			if !exptt.fun(outputTokens[i], exptt.tokens[i]) {
				t.Errorf("Test for %v failed:\n Expected %v\n Got %v", name, exptt.tokens[i], outputTokens[i])
			}
		}

	}

}
