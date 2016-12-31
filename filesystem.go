package main

import "io/ioutil"
import "os"
import "io"

// IFileSystem .
type IFileSystem interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
	Copy(io.Writer, io.Reader) (int64, error)
	MkdirAll(string, os.FileMode) error
	RemoveAll(string) error
	Create(string) (*os.File, error)
}

// FileSystem .
type FileSystem struct {
}

// WriteFile .
func (fs FileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

// ReadFile .
func (fs FileSystem) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// Copy .
func (fs FileSystem) Copy(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}

// MkdirAll .
func (fs FileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// RemoveAll .
func (fs FileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// Create .
func (fs FileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}
