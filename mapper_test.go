package main

import (
	"path"
	"testing"
	"testing/fstest"
)

func TestTargetHasNoEntryAtSourceFilePath(t *testing.T) {
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

func TestTargetHasExistingFileAtSourceFilePath(t *testing.T) {
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

    want := FileExistsError{Path: targetPath}
    if err != want {
		t.Errorf("got error %#v, want %#v", err, want)
	}

    if action != nil {
		t.Errorf("unexpected action %#v", action)
	}
}
