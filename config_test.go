package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TmpDirBlock(f func(string) error) error {
	cur, err := os.Getwd()
	if err != nil {
		return err
	}

	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	if err := os.Chdir(tmpDir); err != nil {
		return err
	}
	defer os.Chdir(cur)

	return f(tmpDir)
}

func TestSearchConfig(t *testing.T) {
	err := TmpDirBlock(func(root string) error {
		RootDir = root

		// path=/, file=(no)
		c, e := SearchConfig(root)
		if e != os.ErrNotExist {
			t.Errorf("unexpected error: %s\n", e.Error())
		}
		if c != "" {
			t.Errorf("unexpected result: %v\n", c)
		}

		// path=/, file=/
		if err := ioutil.WriteFile(filepath.Join(root, ConfigFileName), []byte(""), 0644); err != nil {
			return err
		}
		c, e = SearchConfig(root)
		if e != nil {
			t.Errorf("unexpected error: %s\n", e.Error())
		}
		if c != filepath.Join(root, ConfigFileName) {
			t.Errorf("unexpected result: %v\n", c)
		}

		// path=/foo, file=/
		c, e = SearchConfig(filepath.Join(root, "foo"))
		if e != nil {
			t.Errorf("unexpected error: %s\n", e.Error())
		}
		if c != filepath.Join(root, ConfigFileName) {
			t.Errorf("unexpected result: %v\n", c)
		}

		// path=/foo, file=/foo
		if e := os.MkdirAll(filepath.Join(root, "foo"), 0755); e != nil {
			return e
		}
		if e := ioutil.WriteFile(filepath.Join(root, "foo", ConfigFileName), []byte(""), 0644); e != nil {
			return e
		}

		c, e = SearchConfig(filepath.Join(root, "foo"))
		if e != nil {
			t.Errorf("unexpected error: %s\n", e.Error())
		}
		if c != filepath.Join(root, "foo", ConfigFileName) {
			t.Errorf("unexpected result: %v\n", c)
		}

		return nil
	})

	if err != nil {
		t.Errorf("test block failed: %s\n", err)
	}
}

func TestParseConfig(t *testing.T) {
	err := TmpDirBlock(func(root string) error {
		yaml := `---
- name: foo
  exec: exec foo
`

		configs, err := ParseConfig([]byte(yaml))
		if err != nil {
			t.Errorf("unexpected error: %s\n", err.Error())
		}

		if len(configs) != 1 {
			t.Errorf("unexpected config length: %d\n", len(configs))
		}

		if configs[0].Name != "foo" {
			t.Errorf("unexpected namd: %s\n", configs[0].Name)
		}
		if configs[0].Exec != "exec foo" {
			t.Errorf("unexpected namd: %s\n", configs[0].Exec)
		}

		return nil
	})

	if err != nil {
		t.Errorf("test block failed: %s\n", err)
	}

	err = TmpDirBlock(func(root string) error {
		yaml := `---
- name: foo
  exec: exec foo
- name: bar
  exec: exec bar
`

		configs, err := ParseConfig([]byte(yaml))
		if err != nil {
			t.Errorf("unexpected error: %s\n", err.Error())
		}

		if len(configs) != 2 {
			t.Errorf("unexpected config length: %d\n", len(configs))
		}

		if configs[0].Name != "foo" {
			t.Errorf("unexpected namd: %s\n", configs[0].Name)
		}
		if configs[0].Exec != "exec foo" {
			t.Errorf("unexpected namd: %s\n", configs[0].Exec)
		}
		if configs[1].Name != "bar" {
			t.Errorf("unexpected namd: %s\n", configs[0].Name)
		}
		if configs[1].Exec != "exec bar" {
			t.Errorf("unexpected namd: %s\n", configs[0].Exec)
		}

		return nil
	})

	if err != nil {
		t.Errorf("test block failed: %s\n", err)
	}

}
