package go_ini

import (
	"testing"

	"gopkg.in/ini.v1"
)

func TestBasic(t *testing.T) {
	t.Run("test load file", func(t *testing.T) {
		_, err := ini.Load(FileName)

		assertNoError(t, err)
	})
	t.Run("test load defalut section", func(t *testing.T) {
		cfg, err := ini.Load(FileName)
		got := cfg.Section("").Key("app_mode").String()
		want := "development"

		assertNoError(t, err)
		assertIsEqual(t, got, want)
	})
}
func TestInit(t *testing.T) {
	got := Init()
	assertNoError(t, got)
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
func assertIsEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
