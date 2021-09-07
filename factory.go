package libfilesync

import (
	"os"
	"io"
	"errors"
	"encoding/hex"
	"crypto/sha1"
)

/**
 * Generate a File Struct by os.DirEntry
 */
func NewSyncableFile(entry os.DirEntry,absDir string) (Syncable,error) {

	finfo,err := entry.Info()
	if err != nil {
		panic("Could not get Fileinfo")
	}

	file, err := os.Open(absDir)
	defer file.Close()

	hash := sha1.New()

	if _, err := io.Copy(hash, file); err != nil {
		return syncableFile{}, errors.New("Could not generate sha1 hash for file " + finfo.Name())
	}	
	sf := syncableFile{}
	//Get the 20 bytes hash
	hashInBytes := hash.Sum(nil)[:20]
	sf.Sha1 = hex.EncodeToString(hashInBytes)
	sf.Name = finfo.Name()
	sf.Path = absDir

	return &sf,nil
}