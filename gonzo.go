package gonzo

import (
	"flag"
	"io"
	"os"
	"time"

	"github.com/omeid/gonzo/context"
)

var help = flag.Bool("help", false, "show help")

//File is the virtual-file API used by Gonzo.
// This method `State` is provided for interoperability with os.File,
// so consider using the following interface for a Gonzo and os.File
// compatiable API
//
//  type File interface {
//    io.ReadCloser
//    Stat() (os.FileInfo, error)
//  }
type File interface {
	io.ReadCloser
	Stat() (os.FileInfo, error)
	FileInfo() FileInfo
}

//NewFile returns a file using the provided io.ReadCloser and FileInfo.
func NewFile(rc io.ReadCloser, fi FileInfo) File {
	if fi == nil {
		fi = &fileinfo{}
	}
	return file{rc, fi}
}

// FileInfo is a mutuable superset of os.FileInfo
// with the additon of "Base"
type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}

	Base() string

	SetName(string)
	SetSize(int64)
	SetMode(os.FileMode)
	SetModTime(time.Time)
	SetIsDir(bool)

	SetBase(string)
}

// Stage is a function that takes a context, a channel of files to read and
// an output chanel.
// There is no correlation between a stages input and output, a stage may
// decided to pass the same files after transofrmation or generate new files
// based on the input or drop files.
//
// A stage must not close the output channel based on the simple
// "Don't close it if you don't own it" principle.
// A stage must either pass on a file or call the `Close` method on it.
type Stage func(context.Context, <-chan File, chan<- File) error

//Pipe handles stages and talks to other pipes.
type Pipe interface {
	Context() context.Context
	Files() <-chan File
	Pipe(stages ...Stage) Pipe
	Then(stages ...Stage) error
	Wait() error
}

//NewPipe returns a pipe using the context provided and
//channel of files. If you don't need a context, use context.Background()
func NewPipe(ctx context.Context, files <-chan File) Pipe {
	return pipe{
		context: ctx,
		files:   files,
	}
}
