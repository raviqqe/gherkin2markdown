package convert

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	gherkin "github.com/cucumber/gherkin/go/v27"
	"github.com/raviqqe/gherkin2markdown/renderer"
)

const featureFileExtension = ".feature"

func ConvertFile(s string, w io.Writer) error {
	f, err := os.Open(s)

	if err != nil {
		return err
	}

	d, err := gherkin.ParseGherkinDocument(f, func() string { return s })

	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, renderer.NewRenderer().Render(d))
	return err
}

func ConvertFiles(s, d string) error {
	ps := []string{}

	err := filepath.Walk(s, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && filepath.Ext(p) == featureFileExtension {
			ps = append(ps, p)
		}

		return nil
	})

	if err != nil {
		return err
	}

	w := sync.WaitGroup{}
	es := make(chan error, len(ps))

	for _, p := range ps {
		w.Add(1)

		go func(p string) {
			defer w.Done()

			f, err := openDestFile(p, s, d)

			if err != nil {
				es <- err
				return
			}

			err = ConvertFile(p, f)

			if err != nil {
				es <- err
				return
			}
		}(p)
	}

	w.Wait()

	if len(es) != 0 {
		return <-es
	}

	return nil
}

func openDestFile(p, s, d string) (*os.File, error) {
	p, err := filepath.Rel(s, p)

	if err != nil {
		return nil, err
	}

	p = strings.TrimSuffix(filepath.Join(d, p), featureFileExtension) + ".md"

	if err := os.MkdirAll(filepath.Dir(p), 0700); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return nil, err
	}

	err = f.Truncate(0)

	if err != nil {
		return nil, err
	}

	return f, nil
}
