package files

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
    dir := os.TempDir()

    l, err := NewLocal(dir, 1024*1000*5)
    if err != nil {
        t.Fatal(err)   
    }

    return l, dir, func() {
        // cleanup
        //os.RemoveAll(dir)
    }
}

func TestSavesContentsOfReader(t *testing.T) {
    savePath := "/1/test.png"
    fileContents := "Hello test file"
    l, dir, cleanup := setupLocal(t)
    defer cleanup()

    err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
    assert.NoError(t, err)

    f, err := os.Open(filepath.Join(dir, savePath))
    assert.NoError(t, err)

    d, err := io.ReadAll(f)
    assert.NoError(t, err)
    assert.Equal(t, fileContents, string(d))
}

func TestGetsContentsAndWritesToWriter(t *testing.T) {
    savePath := "/1/test.png"
    fileContents := "Hello test file"

    l, _, cleanup := setupLocal(t)
    defer cleanup()

    err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
    assert.NoError(t, err)

    r, err := l.Get(savePath)
    assert.NoError(t, err)
    defer r.Close()

    d, err := io.ReadAll(r)
    assert.Equal(t, fileContents, string(d))
}
