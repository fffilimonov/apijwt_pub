package layer

import (
	"io"
	"sync"

	"github.com/hyperhq/hypercli/pkg/archive"
)

type mountedLayer struct {
	name       string
	mountID    string
	initID     string
	parent     *roLayer
	layerStore *layerStore

	references map[RWLayer]*referencedRWLayer
}

func (ml *mountedLayer) cacheParent() string {
	if ml.initID != "" {
		return ml.initID
	}
	if ml.parent != nil {
		return ml.parent.cacheID
	}
	return ""
}

func (ml *mountedLayer) TarStream() (io.ReadCloser, error) {
	archiver, err := ml.layerStore.driver.Diff(ml.mountID, ml.cacheParent())
	if err != nil {
		return nil, err
	}
	return archiver, nil
}

func (ml *mountedLayer) Name() string {
	return ml.name
}

func (ml *mountedLayer) Parent() Layer {
	if ml.parent != nil {
		return ml.parent
	}

	// Return a nil interface instead of an interface wrapping a nil
	// pointer.
	return nil
}

func (ml *mountedLayer) Mount(mountLabel string) (string, error) {
	return ml.layerStore.driver.Get(ml.mountID, mountLabel)
}

func (ml *mountedLayer) Unmount() error {
	return ml.layerStore.driver.Put(ml.mountID)
}

func (ml *mountedLayer) Size() (int64, error) {
	return ml.layerStore.driver.DiffSize(ml.mountID, ml.cacheParent())
}

func (ml *mountedLayer) Changes() ([]archive.Change, error) {
	return ml.layerStore.driver.Changes(ml.mountID, ml.cacheParent())
}

func (ml *mountedLayer) Metadata() (map[string]string, error) {
	return ml.layerStore.driver.GetMetadata(ml.mountID)
}

func (ml *mountedLayer) getReference() RWLayer {
	ref := &referencedRWLayer{
		mountedLayer: ml,
	}
	ml.references[ref] = ref

	return ref
}

func (ml *mountedLayer) hasReferences() bool {
	return len(ml.references) > 0
}

func (ml *mountedLayer) deleteReference(ref RWLayer) error {
	rl, ok := ml.references[ref]
	if !ok {
		return ErrLayerNotRetained
	}

	if err := rl.release(); err != nil {
		return err
	}
	delete(ml.references, ref)

	return nil
}

type referencedRWLayer struct {
	*mountedLayer

	activityL     sync.Mutex
	activityCount int
}

func (rl *referencedRWLayer) release() error {
	rl.activityL.Lock()
	defer rl.activityL.Unlock()

	if rl.activityCount > 0 {
		return ErrActiveMount
	}

	rl.activityCount = -1

	return nil
}

func (rl *referencedRWLayer) Mount(mountLabel string) (string, error) {
	rl.activityL.Lock()
	defer rl.activityL.Unlock()

	if rl.activityCount == -1 {
		return "", ErrLayerNotRetained
	}

	rl.activityCount++
	return rl.mountedLayer.Mount(mountLabel)
}

func (rl *referencedRWLayer) Unmount() error {
	rl.activityL.Lock()
	defer rl.activityL.Unlock()

	if rl.activityCount == 0 {
		return ErrNotMounted
	}
	if rl.activityCount == -1 {
		return ErrLayerNotRetained
	}
	rl.activityCount--

	return rl.mountedLayer.Unmount()
}
