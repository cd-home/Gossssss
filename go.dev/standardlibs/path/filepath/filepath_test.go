package tests

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func TestAbsPath(t *testing.T) {
	// returns an absolute representation of path
	abspath, _ := filepath.Abs("filepath_test.go")
	t.Log(abspath)
}

func TestPathBase(t *testing.T) {
	// returns the last element of path
	tests := "/a/b/c"
	t.Log(filepath.Base(tests))
}

func TestCleanPath(t *testing.T) {
	// 返回正确的路径格式
	// 去除多余斜杠
	// Eliminate each . path name element (the current directory) 去除 current
	tests := []string{"./a//b/", "../a//b//", ".//a/", "/../a/b/c"}
	for _, test := range tests {
		t.Run("", func(tt *testing.T) {
			tt.Log(filepath.Clean(test))
		})
	}
}

func TestPathOfDir(t *testing.T) {
	// Dir returns all but the last element of path
	abspath, _ := filepath.Abs("filepath_test.go")
	dir := filepath.Dir(abspath)
	t.Log(dir)
}

func TestEvalSymlinks(t *testing.T) {
	// EvalSymlinks returns the path name after the evaluation of any symbolic links
	sympath, _ := filepath.EvalSymlinks("")
	t.Log(sympath)
}

func TestIsAbsPath(t *testing.T) {
	// IsAbs reports whether the path is absolute.
	tests := []string{"/a/b/c", "./a/b/", "../a/b/c"}
	for _, test := range tests {
		t.Run(test, func(tt *testing.T) {
			tt.Log(filepath.IsAbs(test))
		})
	}
}

func TestJoinpath(t *testing.T) {
	// Join joins any number of path elements into a single path, separating them with an OS specific Separator
	p1 := "/a/b"
	p2 := "c"
	p3 := "/d"
	t.Log(filepath.Join(p1, p2))
	t.Log(filepath.Join(p2, p3))
	t.Log(filepath.Join(p3, p1))
}

func TestPathMatch(t *testing.T) {
	// Match reports whether name matches the shell file name pattern
	// ? singel char
	t.Log(filepath.Match("/a/b/?", "/a/b/c"))
	t.Log(filepath.Match("/a/b/?", "/a/b/cc"))

	t.Log(filepath.Match("/a/?olang", "/a/golang"))
	t.Log(filepath.Match("/a/?olang", "/a/polang"))

	// * multi chars
	t.Log(filepath.Match("/a/b/*", "/a/b/c"))
	t.Log(filepath.Match("/a/b/*", "/a/b/ccccc"))

	// range
	// default single
	t.Log(filepath.Match("a/b/[a-d]", "a/b/c"))
	t.Log(filepath.Match("a/b/[a-d]", "a/b/g"))

	// single and multi
	t.Log(filepath.Match("a/b/[1-9]?", "a/b/1"))
	t.Log(filepath.Match("a/b/[1-9]*", "a/b/12"))

	// Not
	t.Log(filepath.Match("a/b/[^a]", "a/b/a"))
	t.Log(filepath.Match("a/b/[^a]", "a/b/c"))
}

func TestRelativePath(t *testing.T) {
	// 返回基于 Base 的 相对路径
	tests := []string{"/a/b/c", "/b/c", "./b/c", "../b/c"}
	base := "/a"
	for _, path := range tests {
		t.Run(path, func(tt *testing.T) {
			rel, err := filepath.Rel(base, path)
			tt.Log(rel, err)
		})
	}
}

func TestSplitPath(t *testing.T) {
	// Split Dir and File == > in fact split last slash
	// Split splits path immediately following the final Separator,
	// separating it into a directory and file name component.
	// If there is no Separator in path, Split returns an empty dir
	// and file set to path.
	// The returned values have the property that path = dir+file.
	tests := []string{"/a/b/c.go", "a/b/c/d", "a/b/c//d"}
	for _, test := range tests {
		t.Run(test, func(tt *testing.T) {
			tt.Log(filepath.Split(test))
		})
	}
}

func TestSplitListPath(t *testing.T) {
	// SplitList splits a list of paths joined by the OS-specific ListSeparator,
	// usually found in PATH or GOPATH environment variables.
	// 通常是来分割 环境变量
	test := "a/b/c:/d/e/f"
	t.Log(filepath.SplitList(test))
}

func TestToSlashPath(t *testing.T) {
	// 不知道咋用
	// ToSlash returns the result of replacing each separator character
	// in path with a slash ('/') character. Multiple separators are
	// replaced by multiple slashes.
	test := "a/b/c"
	t.Log(filepath.ToSlash(test))
}

func TestVolumeNamePath(t *testing.T) {
	// 不知道咋用
	// VolumeName returns leading volume name.
	// Given "C:\foo\bar" it returns "C:" on Windows.
	// Given "\\host\share\foo" it returns "\\host\share".
	// On other platforms it returns "".
	tests := []string{"/a/b/c", "C:/a/b/c", `\\home\share`}
	for _, test := range tests {
		t.Run(test, func(tt *testing.T) {
			tt.Log(filepath.VolumeName(test))
		})
	}
}

func TestWalkPath(t *testing.T) {
	// 遍历目录
	test := "/Users/admin/Documents/otherSpace/Gossssss/stdsss"
	filepath.Walk(test, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			t.Log(info.Name())
		}
		return nil
	})
}

func TestWalkDirPath(t *testing.T) {
	// Better than Walk
	// Walk is less efficient than WalkDir, introduced in Go 1.16,
	// which avoids calling os.Lstat on every visited file or directory.
	test := "/Users/admin/Documents/otherSpace/Gossssss/stdsss"
	filepath.WalkDir(test, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		t.Log(d.Name())
		return nil
	})
}
