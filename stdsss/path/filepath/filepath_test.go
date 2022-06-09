package tests

import (
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
