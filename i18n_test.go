package i18n_test

import (
	"embed"
	"testing"

	"github.com/fitv/go-i18n"
	"github.com/stretchr/testify/assert"
)

//go:embed locales/*.yml
var fs embed.FS

func TestI18n(t *testing.T) {
	assert := assert.New(t)

	i18n, err := i18n.New(fs, "locales")
	if err != nil {
		t.Fatal(err)
	}

	user := map[string]interface{}{"name": "Jack", "email": "jack@example.com"}

	i18n.SetDefaultLocale("en")
	assert.Equal(i18n.Trans("user.name"), "Name")
	assert.Equal(i18n.Trans("user.email"), "Email")
	assert.Equal(i18n.Trans("hello.world"), "World")
	assert.Equal(i18n.Trans("hello.foo", "bar"), "param bar")
	assert.Equal(i18n.Locale("zh").Trans("hello.world"), "世界")
	assert.Equal(i18n.Trans("user.description", user), "Name: Jack, Email: jack@example.com")

	i18n.SetDefaultLocale("zh")
	assert.Equal(i18n.Trans("user.name"), "姓名")
	assert.Equal(i18n.Trans("user.email"), "邮箱")
	assert.Equal(i18n.Trans("hello.world"), "世界")
	assert.Equal(i18n.Trans("hello.foo", "bar"), "参数 bar")
	assert.Equal(i18n.Locale("en").Trans("hello.world"), "World")
	assert.Equal(i18n.Trans("user.description", user), "姓名: Jack, 邮箱: jack@example.com")
}
