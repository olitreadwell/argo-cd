package tgzstream_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/argoproj/argo-cd/v3/util/tgzstream"
)

func TestCloseAndDeleteRemovesTempDirectory(t *testing.T) {
	// Given a file compressed into a fresh temp directory
	appPath := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(appPath, "app.yaml"), []byte("kind: ConfigMap\n"), 0600))

	tgz, _, _, err := tgzstream.CompressFiles(appPath, nil, nil)
	require.NoError(t, err)
	tempDir := filepath.Dir(tgz.Name())
	require.DirExists(t, tempDir)

	// When the stream is closed and deleted
	tgzstream.CloseAndDelete(tgz)

	// Then the temp file and its parent directory are both removed
	_, err = os.Stat(tgz.Name())
	assert.True(t, os.IsNotExist(err), "temp file should be removed")
	_, err = os.Stat(tempDir)
	assert.True(t, os.IsNotExist(err), "temp directory should be removed")
}
