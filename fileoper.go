package ssui

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// file modified time
func FileModTime(path string) int64 {
	f, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return f.ModTime().Unix()
}

// file size bytes
func FileSize(path string) int64 {
	f, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return f.Size()
}

//file lines
func FileLineNum(path string) (num int64) {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		_, isPrefix, e := buf.ReadLine()
		for isPrefix {
			_, isPrefix, e = buf.ReadLine()
		}
		if e != nil {
			return
		}
		num++
	}
	return
}

// delete file
func FileDelete(path string) error {
	return os.Remove(path)
}

// rename file
func FileRename(path string, to string) error {
	return os.Rename(path, to)
}

// is file
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// is exist
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//make dir
func FileMakeDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

//create file
func FileCreate(file string) (*os.File, error) {
	d := FileDir(file)
	if d != "" {
		err := os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(file)
}

//open appending file
func FileOpenAppend(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND, 0666)
}

//create a new file and write content
func FileCreateAndWrite(path, content string) error {
	f, err := FileCreate(path)
	if err != nil {
		return err
	}
	if _, err = f.WriteString(content); err != nil {
		return err
	}
	f.Close()

	return nil
}

func FileCreateAndWriteWithBom(path, content string) error {
	f, err := FileCreate(path)
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
		return err
	}
	if _, err = f.WriteString(content); err != nil {
		return err
	}
	f.Close()

	return nil
}

//append string to file
func FileWriteAndAppend(path, content string) error {
	if FileIsExist(path) {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		if _, err = f.WriteString(content); err != nil {
			return err
		}
		f.Close()
		return nil
	} else {
		return FileCreateAndWrite(path, content)
	}
}

//reand file all content with bom
func FileReadAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	c, err := ioutil.ReadAll(f)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return c, nil
}

//read line whithout bom
func FileIterateLine(path string, callback func(num int, line string) bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	n := 0
	for {
		n++
		line, isPrefix, err := buf.ReadLine()
		for isPrefix {
			var next []byte
			next, isPrefix, err = buf.ReadLine()
			line = append(line, next...)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		if n == 1 && len(line) > 2 { //has bom?
			if line[0] == 239 && line[1] == 187 && line[2] == 191 {
				line = line[3:]
			}
		}
		if !callback(n, string(line)) {
			break
		}
	}
	return nil
}

//walk files
func FileIterateDir(path, filter string, callback func(file string) bool) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			//return filepath.SkipDir
			return nil
		}
		if filter != "" {
			fs := strings.Split(filter, "|")
			isMatch := false
			for _, v := range fs {
				if strings.HasSuffix(path, v) {
					isMatch = true
					break
				}
			}
			if !isMatch {
				return nil
			}
		}
		if !callback(path) {
			return fmt.Errorf("walk file over")
		}
		return nil
	})

	return err
}

func FileStdinHasData() bool {
	stat, statErr := os.Stdin.Stat()

	if statErr != nil {
		return false
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}
}

type STFile struct {
	path string
	file *os.File
	buf  *bufio.Reader
	line int
}

func NewSTFile(path string) (*STFile, error) {
	var fp *os.File
	var err error

	if path == "stdin" {
		fp = os.Stdin
		err = nil
	} else {
		fp, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	}

	return &STFile{path, fp, bufio.NewReader(fp), 0}, nil
}

func (f *STFile) ReadLine() (string, int) {
	if f.file == nil {
		return "", -1
	}

	f.line++
	line, isPrefix, err := f.buf.ReadLine()
	for isPrefix {
		var next []byte
		next, isPrefix, err = f.buf.ReadLine()
		line = append(line, next...)
	}
	if err != nil {
		return "", -1
	}
	return string(line), f.line
}

func (f *STFile) Close() {
	if f.file == nil {
		return
	}
	f.file.Close()
	f.file = nil
}

//filepath
func FileFullPath(path string) (s string) {
	if path != "" {
		s, _ = filepath.Abs(path)
	} else {
		s, _ = filepath.Abs(os.Args[0])
	}
	return s
}
func FileDir(path string) (s string) {
	if path != "" {
		s, _ = filepath.Abs(filepath.Dir(path))
	} else {
		s, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	return s
}
func FileBase(path string) string {
	if path == "" {
		path = os.Args[0]
	}
	return filepath.Base(path)
}
func FileOnlyName(path string) string {
	if path == "" {
		path = os.Args[0]
	}
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

func FileIsHidden(path string) bool {
	ps := strings.Split(path, string(filepath.Separator))
	for _, s := range ps {
		if strings.HasPrefix(s, ".") {
			return true
		}
	}
	return false
}

//get all files and dirs
func FileReadDir(path string, noHidden bool, fileList map[string]os.FileInfo) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, info := range infos {
		f := path + string(filepath.Separator) + info.Name()
		if _, ok := fileList[f]; ok {
			continue
		}
		if noHidden && FileIsHidden(f) {
			continue
		}
		fileList[f] = info
		if info.IsDir() {
			FileReadDir(f, noHidden, fileList)
		}
	}
}

func FileMD5(path string) (string, error) {
	f, e := os.Open(path)
	if e != nil {
		return "", e
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5hash.Sum(nil)), nil
}
