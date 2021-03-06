package main

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	// ErrLinkIsValid means it needs not to be fixed up.
	ErrLinkIsValid = errors.New("Link is valid")
)

// AuthPath will be used for the path suitable for SSH_AUTH_SOCK
type AuthPath string

// SSHAuthSockPath means the path suitable for SSH_AUTH_SOCK
var SSHAuthSockPath = AuthPath(filepath.Join("~", ".ssh", "auth_sock"))

// FullPath returns the full path replaced ~ with the supplied dir.
func (p *AuthPath) FullPath() (*string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	fullPath := strings.Replace(string(*p), "~", u.HomeDir, 1)
	return &fullPath, nil
}

// EnvPath will be used for the path in SSH_AUTH_SOCK
type EnvPath struct {
	path string
}

// SSHAuthSockEnv means the path in SSH_AUTH_SOCK
var SSHAuthSockEnv = &EnvPath{os.Getenv("SSH_AUTH_SOCK")}

// FixWith fixes the path by the newest path for sock.
func (p *EnvPath) FixWith(newest *PathString) error {
	if isSymlink, err := p.isSymlink(); err != nil && !os.IsNotExist(err) {
		return err
	} else if !isSymlink {
		return p.fixNonSymlink(newest)
	}
	return p.fixSymlink(newest)
}

func (p *EnvPath) isSymlink() (bool, error) {
	st, err := os.Lstat(p.path)
	if err != nil {
		return false, err
	}
	return st.Mode()&os.ModeSymlink != 0, nil
}

func (p *EnvPath) fixNonSymlink(newest *PathString) error {
	if isEqual, err := newest.IsEvalEqual(&p.path); err != nil {
		return err
	} else if isEqual {
		return ErrLinkIsValid
	}
	fullPath, err := SSHAuthSockPath.FullPath()
	if err != nil {
		return err
	}
	return symLink(string(*newest), *fullPath)
}

func (p *EnvPath) fixSymlink(newest *PathString) error {
	if target, err := os.Readlink(p.path); err != nil {
		return err
	} else if isEqual, err := newest.IsEvalEqual(&target); err != nil {
		return err
	} else if isEqual {
		return ErrLinkIsValid
	}
	return symLink(string(*newest), p.path)
}

func symLink(oldname, newname string) (err error) {
	switch _, err = os.Lstat(newname); {
	case err == nil: // if exists
		if err = os.Remove(newname); err != nil {
			return
		}
	case os.IsNotExist(err): // do nothing
	default:
		return
	}
	return os.Symlink(oldname, newname)
}
