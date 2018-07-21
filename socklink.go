package gosshauth

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	if err != nil {
		return false, err
	}
	resolvedTarget, err := filepath.EvalSymlinks(*target)
	if err != nil {
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

func (sls SockLinks) String() string {
	var maxPathLen int
	var paths, modTimes [][]byte
	for _, s := range sls {
		paths = append(paths, []byte(s.Path))
		modTimes = append(modTimes, []byte(s.ModTime.Local().Format(time.RFC3339)))
		if maxPathLen < len(s.Path) {
			maxPathLen = len(s.Path)
		}
	}
	var out strings.Builder
	var err error
	if err = writeHeader(&out, maxPathLen); err != nil {
		return fmt.Sprintf("error on String(): %v", err)
	}
	for i := 0; i < len(paths); i++ {
		if err = writeRow(&out, maxPathLen, paths[i], modTimes[i]); err != nil {
			return fmt.Sprintf("error on String(): %v", err)
		}
	}
	return out.String()
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
