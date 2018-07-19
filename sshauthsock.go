package gosshauth

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ErrLinkIsValid means it needs not to be fixed up.
var ErrLinkIsValid = errors.New("Link is valid")

type sshAuthPath string

var sshAuthSockPath = sshAuthPath(filepath.Join("~", ".ssh", "auth_sock"))

// FullPath returns the full path replaced ~ with the supplied dir.
func (p *sshAuthPath) FullPath() (*string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	fullPath := strings.Replace(string(*p), "~", u.HomeDir, 1)
	return &fullPath, nil
}

type envPath struct {
	path string
}

var sshAuthSockEnv = &envPath{os.Getenv("SSH_AUTH_SOCK")}

func (env *envPath) isSymlink() (bool, error) {
	st, err := os.Lstat(env.path)
	if err != nil {
		return false, err
	}
	return st.Mode()&os.ModeSymlink != 0, nil
}

func (env *envPath) fixNonSymlink(newest *PathString) error {
	if isEqual, err := newest.IsEvalEqual(&env.path); err != nil {
		return err
	} else if isEqual {
		return ErrLinkIsValid
	}
	fullPath, err := sshAuthSockPath.FullPath()
	if err != nil {
		return err
	}
	return os.Symlink(string(*newest), *fullPath)
}

func (env *envPath) fixSymlink(newest *PathString) error {
	if target, err := os.Readlink(env.path); err != nil {
		return err
	} else if isEqual, err := newest.IsEvalEqual(&target); err != nil {
		return err
	} else if isEqual {
		return ErrLinkIsValid
	} else if err := os.Remove(env.path); err != nil {
		return err
	}
	return os.Symlink(string(*newest), env.path)
}

// FixSSHAuthSock returns info for the link in $SSH_AUTH_SOCK
func FixSSHAuthSock(sls SockLinks) error {
	socks, err := SearchSockLinks()
	if err != nil {
		return err
	}
	newest := &socks.Newest().Path
	if isSymlink, err := sshAuthSockEnv.isSymlink(); err != nil {
		return err
	} else if !isSymlink {
		return sshAuthSockEnv.fixNonSymlink(newest)
	}
	return sshAuthSockEnv.fixSymlink(newest)
}
