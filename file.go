package gonzo

import (
	"io"
	"os"
	"sync"
	"time"
)

// File is the object that is passed between Stages of Pipes
type file struct {
	io.ReadCloser
	fileinfo FileInfo
}

func (f file) Stat() (os.FileInfo, error) {
	return f.fileinfo, nil
}

func (f file) FileInfo() FileInfo {
	return f.fileinfo
}

//FileInfoFrom createa a FileInfo from an os.FileInfo.
//Useful when working with os.File or other APIs that
//mimics http.FileSystem.
func FileInfoFrom(fi os.FileInfo) FileInfo {
	return &fileinfo{
		&sync.RWMutex{},
		fi.Name(),
		fi.Size(),
		fi.Mode(),
		fi.ModTime(),
		fi.IsDir(),
		fi.Sys(),
		"",
	}

}

//NewFileInfo create a new empty FileInfo.
func NewFileInfo() FileInfo {
	return &fileinfo{
		lock: &sync.RWMutex{},
	}
}

// Fileinfo implements os.fileinfo.
type fileinfo struct {
	lock    *sync.RWMutex
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}

	base string //base, usually glob.Base
	//XXX: For consideration.
	//Cwd  string //Where are we?
	//Path string //Full path.

}

func (f fileinfo) Name() string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.name
}
func (f fileinfo) Size() int64 {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.size
}

func (f fileinfo) Mode() os.FileMode {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.mode
}

func (f fileinfo) ModTime() time.Time {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.modTime
}

func (f fileinfo) IsDir() bool {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.isDir
}

func (f fileinfo) Sys() interface{} {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.sys
}

func (f *fileinfo) SetName(name string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.name = name
}

func (f *fileinfo) SetSize(size int64) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.size = size
}

func (f *fileinfo) SetMode(mod os.FileMode) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.mode = mod
}

func (f *fileinfo) SetModTime(time time.Time) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.modTime = time
}
func (f *fileinfo) SetIsDir(isdir bool) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.isDir = isdir
}
func (f *fileinfo) SetSys(sys interface{}) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.sys = sys
}

//Extension

func (f fileinfo) Base() string {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.base
}

func (f *fileinfo) SetBase(base string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.base = base
}
