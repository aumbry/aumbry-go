package aumbry

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/franela/goblin"
)

type SampleConfig struct {
	Database struct {
		Host string
		Port int
	}
}

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Yaml", func() {
		var tmpFile *os.File
		var options map[string]string

		g.BeforeEach(func() {
			tmpFile, _ = ioutil.TempFile(os.TempDir(), "aumbry")

			dn, fn := filepath.Split(tmpFile.Name())
			options = make(map[string]string)
			options["CONFIG_FILENAME"] = fn
			options["CONFIG_SEARCH_PATH"] = dn

			sample := "---\n" +
				"top: tester\n" +
				"database:\n" +
				"  host: 0.0.0.0\n" +
				"  port: 1234"

			ioutil.WriteFile(tmpFile.Name(), []byte(sample), os.ModeExclusive)
		})

		g.AfterEach(func() {
			os.Remove(tmpFile.Name())
		})

		g.It("Should load from a file", func() {
			var cfg SampleConfig
			a := New(YamlFile, &cfg, options)
			a.Load()

			g.Assert(cfg.Database.Host).Equal("0.0.0.0")
			g.Assert(cfg.Database.Port).Equal(1234)
		})
	})
}
