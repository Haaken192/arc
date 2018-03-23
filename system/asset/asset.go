package asset

import (
	"github.com/haakenlabs/arc/core"
)

// RegisterHandler registers an asset handler.
func RegisterHandler(h core.AssetHandler) error {
	return core.GetAssetSystem().RegisterHandler(h)
}

// GetHandler gets an asset handler by name.
func GetHandler(name string) (core.AssetHandler, error) {
	return core.GetAssetSystem().GetHandler(name)
}

// Get gets an asset by name from a handler by kind.
func Get(kind, name string) (core.Object, error) {
	return core.GetAssetSystem().Get(kind, name)
}

// // MustGet is like Get, but panics if an error is encountered.
func MustGet(kind, name string) core.Object {
	return core.GetAssetSystem().MustGet(kind, name)
}

func MountPackage(name string) error {
	return core.GetAssetSystem().MountPackage(name)
}

// UnmountPackage unmounts a mounted package given by name.
func UnmountPackage(name string) error {
	return core.GetAssetSystem().UnmountPackage(name)
}

// UnmountAllPackages unmounts all mounted packages.
func UnmountAllPackages() {
	core.GetAssetSystem().UnmountAllPackages()
}

// LoadManifest loads a manifest of assets.
func LoadManifest(files ...string) error {
	return core.GetAssetSystem().LoadManifest(files...)
}

func ReadResource(r *core.Resource) error {
	return core.GetAssetSystem().ReadResource(r)
}
