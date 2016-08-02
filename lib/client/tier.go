// Copyright (c) 2016 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/tigera/libcalico-go/lib/api"
	"github.com/tigera/libcalico-go/lib/api/unversioned"
	"github.com/tigera/libcalico-go/lib/backend/model"
)

// TierInterface has methods to work with Tier resources.
type TierInterface interface {
	List(api.TierMetadata) (*api.TierList, error)
	Get(api.TierMetadata) (*api.Tier, error)
	Create(*api.Tier) (*api.Tier, error)
	Update(*api.Tier) (*api.Tier, error)
	Apply(*api.Tier) (*api.Tier, error)
	Delete(api.TierMetadata) error
}

// tiers implements TierInterface
type tiers struct {
	c *Client
}

// newTiers returns a tiers
func newTiers(c *Client) *tiers {
	return &tiers{c}
}

// Create creates a new tier.
func (h *tiers) Create(a *api.Tier) (*api.Tier, error) {
	return a, h.c.create(*a, h)
}

// Create creates a new tier.
func (h *tiers) Update(a *api.Tier) (*api.Tier, error) {
	return a, h.c.update(*a, h)
}

// Create creates a new tier.
func (h *tiers) Apply(a *api.Tier) (*api.Tier, error) {
	return a, h.c.apply(*a, h)
}

// Delete deletes an existing tier.
func (h *tiers) Delete(metadata api.TierMetadata) error {
	return h.c.delete(metadata, h)
}

// Get returns information about a particular tier.
func (h *tiers) Get(metadata api.TierMetadata) (*api.Tier, error) {
	if a, err := h.c.get(metadata, h); err != nil {
		return nil, err
	} else {
		return a.(*api.Tier), nil
	}
}

// List takes a Metadata, and returns the list of tiers that match that Metadata
// (wildcarding missing fields)
func (h *tiers) List(metadata api.TierMetadata) (*api.TierList, error) {
	l := api.NewTierList()
	err := h.c.list(metadata, h, l)
	return l, err
}

// Convert a TierMetadata to a TierListInterface
func (h *tiers) convertMetadataToListInterface(m unversioned.ResourceMetadata) (model.ListInterface, error) {
	hm := m.(api.TierMetadata)
	l := model.TierListOptions{
		Name: hm.Name,
	}
	return l, nil
}

// Convert a TierMetadata to a TierKeyInterface
func (h *tiers) convertMetadataToKey(m unversioned.ResourceMetadata) (model.Key, error) {
	hm := m.(api.TierMetadata)
	k := model.TierKey{
		Name: hm.Name,
	}
	return k, nil
}

// Convert an API Tier structure to a Backend Tier structure
func (h *tiers) convertAPIToKVPair(a unversioned.Resource) (*model.KVPair, error) {
	at := a.(api.Tier)
	k, err := h.convertMetadataToKey(at.Metadata)
	if err != nil {
		return nil, err
	}

	d := model.KVPair{
		Key: k,
		Value: model.Tier{
			Order: at.Spec.Order,
		},
	}

	return &d, nil
}

// Convert a Backend Tier structure to an API Tier structure
func (h *tiers) convertKVPairToAPI(d *model.KVPair) (unversioned.Resource, error) {
	bt := d.Value.(model.Tier)
	bk := d.Key.(model.TierKey)

	at := api.NewTier()
	at.Metadata.Name = bk.Name
	at.Spec.Order = bt.Order

	return at, nil
}

const (
	DefaultTierName = "default"
)
