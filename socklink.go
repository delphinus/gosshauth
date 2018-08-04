package main

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"text/template"
	"time"
)

// SockLink is info for sock symlinks of SSH.
type SockLink struct {
	Path    PathString
	ModTime time.Time
}

// SockLinks is a set of *SockLink
type SockLinks []*SockLink

// Newest returns the newest SockLink
func (sls SockLinks) Newest() (sl *SockLink) {
	if len(sls) > 0 {
		sl = sls[0]
	}
	return
}

// PathString is a path string
type PathString string

// IsEvalEqual returns true if the evaled path string is equal to target.
func (ps *PathString) IsEvalEqual(target *string) (bool, error) {
	resolved, err := filepath.EvalSymlinks(string(*ps))
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	resolvedTarget, err := filepath.EvalSymlinks(*target)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return resolved == resolvedTarget, nil
}

var (
	filename      = []byte("Filename")
	modTime       = []byte("Modified Time")
	spaceByte     = []byte{' '}
	lineBreakByte = []byte{'\n'}
)

// Print prints out the sock links.
func (sls SockLinks) Print(w io.Writer) (err error) {
	var maxPathLen int
	var paths, modTimes [][]byte
	for _, s := range sls {
		paths = append(paths, []byte(s.Path))
		modTimes = append(modTimes, []byte(s.ModTime.Local().Format(time.RFC3339)))
		if maxPathLen < len(s.Path) {
			maxPathLen = len(s.Path)
		}
	}
	if err = writeHeader(w, maxPathLen); err != nil {
		return err
	}
	for i := 0; i < len(paths); i++ {
		if err = writeRow(w, maxPathLen, paths[i], modTimes[i]); err != nil {
			return err
		}
	}
	return nil
}

// PrintWithTemplate prints out by building with the supplied template.
func (sls SockLinks) PrintWithTemplate(w io.Writer, text string) error {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		return err
	}
	for i := 0; i < len(sls); i++ {
		if err := tmpl.Execute(w, sls[i]); err != nil {
			return err
		}
		_, _ = w.Write([]byte{'\n'})
	}
	return nil
}

func writeRow(out io.Writer, maxLen int, path, modTime []byte) (err error) {
	if _, err = out.Write(path); err != nil {
		return
	}
	if err = columnSpace(out, len(path), maxLen); err != nil {
		return
	}
	if _, err = out.Write(modTime); err != nil {
		return
	}
	return lineBreak(out)
}

// writeHeader writes `Filename     Modified Time`
func writeHeader(out io.Writer, maxLen int) (err error) {
	if _, err = out.Write(filename); err != nil {
		return
	}
	if err = columnSpace(out, len(filename), maxLen); err != nil {
		return
	}
	if _, err = out.Write(modTime); err != nil {
		return
	}
	return lineBreak(out)
}

func columnSpace(out io.Writer, colLen, maxLen int) error {
	for i := colLen; i < maxLen+3; i++ {
		if _, err := out.Write(spaceByte); err != nil {
			return err
		}
	}
	return nil
}

// TODO: use platform-independent way
func lineBreak(out io.Writer) error {
	_, err := out.Write(lineBreakByte)
	return err
}

// SearchSockLinks returns info of all sock links. The order is ModTime
// descendant.
func SearchSockLinks() (socks SockLinks, err error) {
	var paths []string
	paths, err = filepath.Glob(sockLinksGlob)
	if err != nil {
		return
	}
	for _, p := range paths {
		var st os.FileInfo
		st, err = os.Stat(p)
		if err != nil {
			return
		}
		if st.Mode()&os.ModeSocket != 0 {
			socks = append(socks, &SockLink{PathString(p), st.ModTime()})
		}
	}
	sort.Slice(socks, func(i, j int) bool {
		return socks[i].ModTime.After(socks[j].ModTime)
	})
	return
}
