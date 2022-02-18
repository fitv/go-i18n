package i18n

import (
	"embed"
	"fmt"
	"strings"

	"github.com/fitv/go-i18n/internal/translator"
	"gopkg.in/yaml.v3"
)

var (
	defaultLocale string
	emptyTrans    = translator.New(make(map[string]interface{}))
	transMap      = make(map[string]*translator.Translator)
)

// Init init the translator
func Init(fs embed.FS, path string) {
	dirEntries, _ := fs.ReadDir(path)

	for _, entry := range dirEntries {
		m := make(map[string]interface{})

		file, err := fs.ReadFile(fmt.Sprintf("%s/%s", path, entry.Name()))
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(file, &m)
		if err != nil {
			panic(err)
		}

		local := strings.Split(entry.Name(), ".")[0]
		transMap[local] = translator.New(m)
	}
}

// SetDefaultLocale set the default locale
func SetDefaultLocale(local string) {
	defaultLocale = local
}

// Locale returns the translator instance by the given locale
func Locale(locale string) *translator.Translator {
	trans, ok := transMap[locale]
	if ok {
		return trans
	}
	return emptyTrans
}

// Trans returns language translation by the given key
func Trans(key string, args ...interface{}) string {
	return Locale(defaultLocale).Trans(key, args...)
}
