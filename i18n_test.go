package i18n_test

import (
	"embed"
	"testing"

	"github.com/fitv/go-i18n"
	"github.com/stretchr/testify/assert"
)

//go:embed locales/*.yml
var fs embed.FS

func Test(t *testing.T) {
	assert := assert.New(t)

	i18n.Init(fs, "locales")
	i18n.SetDefaultLocale("en")
	assert.Equal(i18n.Trans("user.name"), "Name")
	assert.Equal(i18n.Trans("user.age"), "Age")
	assert.Equal(i18n.Trans("hello.world"), "World")
	assert.Equal(i18n.Trans("hello.foo", "bar"), "param bar")

	i18n.Init(fs, "locales")
	i18n.SetDefaultLocale("zh")
	assert.Equal(i18n.Trans("user.name"), "姓名")
	assert.Equal(i18n.Trans("user.age"), "年龄")
	assert.Equal(i18n.Trans("hello.world"), "世界")
	assert.Equal(i18n.Trans("hello.foo", "bar"), "参数 bar")
}