package translator

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

// Translator is a viper wrapper
type Translator struct {
	m map[string]interface{}
}

// NewTranslator creates a new Translator
func New(m map[string]interface{}) *Translator {
	return &Translator{m: m}
}

// Trans returns language translation by the given key
func (t *Translator) Trans(key string, args ...interface{}) string {
	value, ok := t.get(key)
	if !ok {
		return key
	}

	if len(args) > 0 {
		return fmt.Sprintf(value, args...)
	}
	return value
}

// get returns language translation by the given key
func (t *Translator) get(key string) (str string, exists bool) {
	source := t.m
	keys := strings.Split(key, ".")
	last := len(keys) - 1

	for i, k := range keys {
		val, ok := source[k]
		if !ok {
			return
		}

		switch v := val.(type) {
		case string:
			if i == last {
				return v, true
			}
			return
		case map[interface{}]interface{}:
			source = cast.ToStringMap(val)
		case map[string]interface{}:
			source = v
		default:
			return
		}
	}
	return
}
