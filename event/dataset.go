package event

import (
	"github.com/qri-io/qri/dsref"
)

const (
	// ETDatasetNameInit is when a dataset is initialized
	// payload is a DsChange
	ETDatasetNameInit = Type("dataset:Init")
	// ETDatasetCommitChange is when a dataset changes its newest commit
	// payload is a DsChange
	ETDatasetCommitChange = Type("dataset:CommitChange")
	// ETDatasetDeleteAll is when a dataset is entirely deleted
	// payload is a DsChange
	ETDatasetDeleteAll = Type("dataset:DeleteAll")
	// ETDatasetRename is when a dataset is renamed
	// payload is a DsChange
	ETDatasetRename = Type("dataset:Rename")
	// ETDatasetCreateLink is when a dataset is linked to a working directory
	// payload is a DsChange
	ETDatasetCreateLink = Type("dataset:CreateLink")
)

// DsChange represents the result of a change to a dataset
type DsChange struct {
	InitID     string
	TopIndex   int
	ProfileID  string
	Username   string
	PrettyName string
	HeadRef    string
	Info       *dsref.VersionInfo
	Dir        string
}
