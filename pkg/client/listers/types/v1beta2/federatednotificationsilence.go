/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta2 "kubesphere.io/api/types/v1beta2"
)

// FederatedNotificationSilenceLister helps list FederatedNotificationSilences.
// All objects returned here must be treated as read-only.
type FederatedNotificationSilenceLister interface {
	// List lists all FederatedNotificationSilences in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta2.FederatedNotificationSilence, err error)
	// Get retrieves the FederatedNotificationSilence from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta2.FederatedNotificationSilence, error)
	FederatedNotificationSilenceListerExpansion
}

// federatedNotificationSilenceLister implements the FederatedNotificationSilenceLister interface.
type federatedNotificationSilenceLister struct {
	indexer cache.Indexer
}

// NewFederatedNotificationSilenceLister returns a new FederatedNotificationSilenceLister.
func NewFederatedNotificationSilenceLister(indexer cache.Indexer) FederatedNotificationSilenceLister {
	return &federatedNotificationSilenceLister{indexer: indexer}
}

// List lists all FederatedNotificationSilences in the indexer.
func (s *federatedNotificationSilenceLister) List(selector labels.Selector) (ret []*v1beta2.FederatedNotificationSilence, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta2.FederatedNotificationSilence))
	})
	return ret, err
}

// Get retrieves the FederatedNotificationSilence from the index for a given name.
func (s *federatedNotificationSilenceLister) Get(name string) (*v1beta2.FederatedNotificationSilence, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta2.Resource("federatednotificationsilence"), name)
	}
	return obj.(*v1beta2.FederatedNotificationSilence), nil
}
