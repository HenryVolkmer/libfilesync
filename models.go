package libfilesync

/**
 * Syncable Interface describes a syncable Entity
 */
type Syncable interface {
	/**
	 * get the Name of the File
	 */
	GetName() string
	/**
	 * get the absolute file path to the File to sync
	 */
	GetPath() string
	/**
	 * get the sha1 message digest
	 */
	GetSha1() string
	/**
	 * get the last modified time as string
	 */
	GetModTime() string
}

/**
 * SyncableFile implements the Syncable Interface
 * Intention is to use it as Superclass 
 */
type syncableFile struct {
	Name string ""
	Path string
	Sha1 string
	ModTime string
}

func (sf syncableFile) GetName() string {
	return sf.Name
}

func (sf syncableFile) GetPath() string {
	return sf.Path
}

func (sf syncableFile) GetSha1() string {
	return sf.Sha1
}

func (sf syncableFile) GetModTime() string {
	return sf.ModTime
}