package template

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"os/user"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-sprout/sprout"
	"github.com/sethvargo/go-password/password"
)

// custom symbols: removed \ and "
var customSymbols = "~!@#$%^&*()_+`-={}|[]:<>?,./"

type boilrRegistry struct {
	handler sprout.Handler
}

func newBoilrRegistry() *boilrRegistry { return &boilrRegistry{} }

func (r *boilrRegistry) UID() string { return "ilyes512/boilr" }

func (r *boilrRegistry) LinkHandler(fh sprout.Handler) error {
	r.handler = fh
	return nil
}

func (r *boilrRegistry) RegisterFunctions(fnMap sprout.FunctionMap) error {
	sprout.AddFunction(fnMap, "hostname", r.hostname)
	sprout.AddFunction(fnMap, "username", r.username)
	sprout.AddFunction(fnMap, "toBinary", r.toBinary)
	sprout.AddFunction(fnMap, "formatFilesize", r.formatFilesize)
	sprout.AddFunction(fnMap, "toTitle", strings.ToTitle)
	sprout.AddFunction(fnMap, "password", r.generatePassword)
	sprout.AddFunction(fnMap, "randomBase64", r.randomBase64)
	return nil
}

func (r *boilrRegistry) hostname() string {
	return os.Getenv("HOSTNAME")
}

func (r *boilrRegistry) username() string {
	t, err := user.Current()
	if err != nil {
		return "Unknown"
	}
	return t.Name
}

func (r *boilrRegistry) toBinary(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return s
	}
	return fmt.Sprintf("%b", n)
}

func (r *boilrRegistry) formatFilesize(value interface{}) string {
	var size float64

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		size = float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		size = float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		size = v.Float()
	default:
		return ""
	}

	var KB float64 = 1 << 10
	var MB float64 = 1 << 20
	var GB float64 = 1 << 30
	var TB float64 = 1 << 40
	var PB float64 = 1 << 50

	filesizeFormat := func(filesize float64, suffix string) string {
		return strings.Replace(fmt.Sprintf("%.1f %s", filesize, suffix), ".0", "", -1)
	}

	var result string
	if size < KB {
		result = filesizeFormat(size, "bytes")
	} else if size < MB {
		result = filesizeFormat(size/KB, "KB")
	} else if size < GB {
		result = filesizeFormat(size/MB, "MB")
	} else if size < TB {
		result = filesizeFormat(size/GB, "GB")
	} else if size < PB {
		result = filesizeFormat(size/TB, "TB")
	} else {
		result = filesizeFormat(size/PB, "PB")
	}

	return result
}

// generatePassword generates a random password.
// password.Generate(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error)
func (r *boilrRegistry) generatePassword(length, numDigits, numSymbols int, noUpper, allowRepeat bool) string {
	generator, err := password.NewGenerator(&password.GeneratorInput{Symbols: customSymbols})
	if err != nil {
		return fmt.Sprintf("failed to generate password generator: %s", err)
	}

	res, err := generator.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
	if err != nil {
		return fmt.Sprintf("failed to generate password: %s", err)
	}

	return res
}

// randomBase64 generates a random base64 string based on random bytes of length n.
func (r *boilrRegistry) randomBase64(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("failed to generate randomBase64: %s", err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
