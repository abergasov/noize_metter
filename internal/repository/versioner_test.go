package repository_test

import (
	testhelpers "noize_metter/internal/test_helpers"
	"noize_metter/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo_FetchLatestVersion(t *testing.T) {
	container := testhelpers.GetClean(t)
	version, err := container.Repo.FetchLatestVersion(container.Ctx)
	require.NoError(t, err)
	t.Logf("latest version: %s, current %s", version, utils.GetLastCommitHash())
}
