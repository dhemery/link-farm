package main

import (
	"errors"
	"io/fs"
	"path"
	"testing"
	"testing/fstest"
)

func TestLinkIfSourceFileTargetDoesNotExist(t *testing.T) {
	fsys := fstest.MapFS{}
	targetDir := "target-dir"
	sourceDir := "source-dir"

	filename := "file"
	sourcePath := path.Join(sourceDir, filename)
	targetPath := path.Join(targetDir, filename)

	fsys[sourcePath] = &fstest.MapFile{}
	delete(fsys, targetPath) // Does not exist

	mapper := Mapper{FS: fsys, SourceDir: sourceDir, TargetDir: targetDir}
	action, err := mapper.Map(sourcePath)

	want := CreateLinkAction{From: targetPath, To: sourcePath}
	if action != want {
		t.Errorf("got action %#v, want %#v", action, want)
	}

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestErrorIfSourceFileTargetIsExistingFile(t *testing.T) {
	fsys := fstest.MapFS{}
	targetDir := "target-dir"
	sourceDir := "source-dir"

	filename := "file"
	sourcePath := path.Join(sourceDir, filename)
	targetPath := path.Join(targetDir, filename)

	fsys[sourcePath] = &fstest.MapFile{}
	fsys[targetPath] = &fstest.MapFile{} // Already exists

	mapper := Mapper{FS: fsys, SourceDir: sourceDir, TargetDir: targetDir}
	action, err := mapper.Map(sourcePath)

	if !errors.Is(err, fs.ErrExist) {
		t.Errorf("got error %#v, want %s", err, fs.ErrExist.Error())
	}

	if action != nil {
		t.Errorf("unexpected action %#v", action)
	}
}
