package api_test

import (
	"testing"

	"github.com/Nashmeg/Usage.git/internal/api"
)

func TestGetdiskusage(t *testing.T) {
	// init field values
	var w = api.FileWhereInput{
		MountPoint: "/tmp",
	}
	files, err := resolver.Query().Files(ctx, &w)
	if err != nil {
		t.Fatal(err)
	}
	if files == nil {
		t.Fatal("Expected files, Got nil")
	}
	if len(files) == 0 {
		t.Fatalf("Expected at least 1 File, Got %d", len(files))
	}

}
