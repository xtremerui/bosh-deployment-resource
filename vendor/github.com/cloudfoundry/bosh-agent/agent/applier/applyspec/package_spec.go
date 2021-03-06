package applyspec

import (
	models "github.com/cloudfoundry/bosh-agent/agent/applier/models"
	"github.com/cloudfoundry/bosh-utils/crypto"
)

type PackageSpec struct {
	Name        string                `json:"name"`
	Version     string                `json:"version"`
	Sha1        crypto.MultipleDigest `json:"sha1"`
	BlobstoreID string                `json:"blobstore_id"`
}

func (s *PackageSpec) AsPackage() models.Package {
	return models.Package{
		Name:    s.Name,
		Version: s.Version,
		Source: models.Source{
			Sha1:        s.Sha1,
			BlobstoreID: s.BlobstoreID,
		},
	}
}
