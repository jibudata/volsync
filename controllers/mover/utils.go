package mover

import "encoding/json"

const (
	// VolsyncNodeSelectorAnnoKey is the key to set volsync node selector
	// volsync daemon will be deployed on this node
	VolsyncNodeSelectorAnnoKey = "volsync.backube/volsync-node-selector"
)

func AffinityFromAnnotations(annotations map[string]string) (map[string]string, error) {
	if annotations == nil {
		return nil, nil
	}
	v := annotations[VolsyncNodeSelectorAnnoKey]
	if v == "" {
		return nil, nil
	}
	lbs := make(map[string]string)
	err := json.Unmarshal([]byte(v), &lbs)
	if err != nil {
		return nil, err
	}
	return lbs, nil
}
