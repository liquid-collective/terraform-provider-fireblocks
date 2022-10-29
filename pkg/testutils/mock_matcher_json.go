package testutils

import (
	"encoding/json"
	"fmt"

	"github.com/golang/mock/gomock"
)

type jsonMatcher struct {
	b string
}

func (m jsonMatcher) Matches(x interface{}) bool {
	b, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	if string(b) == m.b {
		return true
	}

	return false
}

func (m jsonMatcher) Got(got interface{}) string {
	b, err := json.Marshal(got)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (m jsonMatcher) String() string {
	return fmt.Sprintf("is equal to %v", m.b)
}

func JSONEq(x interface{}) gomock.Matcher {
	b, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	return jsonMatcher{string(b)}
}
