/**
 * FileName:   kvfile.go
 * Author:     Fasion Chan
 * @contact:   fasionchan@gmail.com
 * @version:   $Id$
 *
 * Description:
 *
 * Changelog:
 *
 **/

package storage

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "path"
)

var _ = fmt.Println

const (
    FILE_SIZE_LIMIT = 100 * 1024 * 1024
)

type FileBasedKeyValue struct {
    workdir string
}

func (self *FileBasedKeyValue) Fetch(key string, data interface{}) (error) {
    filePath := path.Join(self.workdir, key)
    file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
    if err != nil {
        return err
    }

    info, err := file.Stat()
    if err != nil {
        return err
    }

    fileSize := info.Size()
    if fileSize > FILE_SIZE_LIMIT {
        return errors.New("file too big")
    }

    content := make([]byte, fileSize)
    _, err = file.Read(content)
    if err != nil {
        return err
    }

    fmt.Println(string(content))

    err = json.Unmarshal(content, data)
    if err != nil {
        return err
    }

    return nil
}

func (self *FileBasedKeyValue) Save(key string, data interface{}) (error) {
    content, err := json.Marshal(data)
    if err != nil {
        return err
    }

    filePath := path.Join(self.workdir, key)
    file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
    if err != nil {
        return err
    }

    for {
        n, err := file.Write(content)
        if err != nil {
            return err
        }

        if len(content) > n {
            content = content[n:]
        } else {
            break
        }
    }

    return nil
}

func NewFileBasedKeyValue(workdir string) (*FileBasedKeyValue, error) {
    return &FileBasedKeyValue{
        workdir: workdir,
    }, nil
}
