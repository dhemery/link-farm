package main

import (
	"io/fs"
	"path"
	"testing"
	"testing/fstest"
)


func addDir(fsys fstest.MapFS, path string) {
    fsys[path] = &fstest.MapFile{Mode: fs.ModeDir}
}

func addFile(fsys fstest.MapFS, path string) {
    fsys[path] = &fstest.MapFile{}
}

func addLink(fsys fstest.MapFS, oldname, newname string) {
    fsys[newname] = &fstest.MapFile{Mode: fs.ModeSymlink, Data: []byte(oldname)}
}

func TestFilePathInSourceRootDoesNotExistInTarget(t *testing.T) {
    targetDir := "target-dir"
    sourceDir := "source-dir"
    fsys := fstest.MapFS{}

    addDir(fsys, targetDir)

    baseFileName := "file-in-root"
    sourcePath := path.Join(sourceDir, baseFileName)

    mapper := Mapper{SourceDir: sourceDir, TargetDir: targetDir}

    action, err := mapper.Map(sourcePath)

    if err != nil {
        t.Errorf("unexpected error %s", err)
    }

    targetPath := path.Join(targetDir, baseFileName)
    want := CreateLink{From: targetPath, To: sourcePath}

    if action != want {
        t.Errorf("got action %#v, want %#v", action, want)
    }
}

func TestTargetHasExistingFileAtFilePathInSourceRoot(t *testing.T) {
}
