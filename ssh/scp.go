/*
 * Author: fasion
 * Created time: 2019-09-17 17:26:20
 * Last Modified by: fasion
 * Last Modified time: 2019-11-06 10:21:38
 */

package ssh

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	shellquote "github.com/kballard/go-shellquote"
	ssh "golang.org/x/crypto/ssh"
)

func CopyTo(client *ssh.Client, content io.Reader, size int64, path, name string, mode os.FileMode) error {
	if name == "" {
		name = filepath.Base(path)
	}

	// create new session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// stdin for feed file content
	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	cmd := shellquote.Join("scp", "-t", path)
	if err := session.Start(cmd); err != nil {
		stdin.Close()
		return err
	}

	errors := make(chan error)
	go func() {
		errors <- session.Wait()
	}()

	// feed content
	fmt.Fprintf(stdin, "C%#o %d %s\n", mode, size, name)
	io.Copy(stdin, content)
	fmt.Fprint(stdin, "\x00")
	stdin.Close()

	return <-errors
}

func CopyBytesTo(client *ssh.Client, content []byte, path, name string, mode os.FileMode) error {
	return CopyTo(client, bytes.NewReader(content), int64(len(content)), path, name, mode)
}

func CopyFileTo(client *ssh.Client, content *os.File, path, name string, mode os.FileMode) error {
	info, err := content.Stat()
	if err != nil {
		return err
	}

	return CopyTo(client, content, info.Size(), path, info.Name(), mode)
}
