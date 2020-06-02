// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/tomkennedy513/password-gen/pkg/apis/passwordgen/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PasswordTypeLister helps list PasswordTypes.
type PasswordTypeLister interface {
	// List lists all PasswordTypes in the indexer.
	List(selector labels.Selector) (ret []*v1.PasswordType, err error)
	// PasswordTypes returns an object that can list and get PasswordTypes.
	PasswordTypes(namespace string) PasswordTypeNamespaceLister
	PasswordTypeListerExpansion
}

// passwordTypeLister implements the PasswordTypeLister interface.
type passwordTypeLister struct {
	indexer cache.Indexer
}

// NewPasswordTypeLister returns a new PasswordTypeLister.
func NewPasswordTypeLister(indexer cache.Indexer) PasswordTypeLister {
	return &passwordTypeLister{indexer: indexer}
}

// List lists all PasswordTypes in the indexer.
func (s *passwordTypeLister) List(selector labels.Selector) (ret []*v1.PasswordType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PasswordType))
	})
	return ret, err
}

// PasswordTypes returns an object that can list and get PasswordTypes.
func (s *passwordTypeLister) PasswordTypes(namespace string) PasswordTypeNamespaceLister {
	return passwordTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PasswordTypeNamespaceLister helps list and get PasswordTypes.
type PasswordTypeNamespaceLister interface {
	// List lists all PasswordTypes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.PasswordType, err error)
	// Get retrieves the PasswordType from the indexer for a given namespace and name.
	Get(name string) (*v1.PasswordType, error)
	PasswordTypeNamespaceListerExpansion
}

// passwordTypeNamespaceLister implements the PasswordTypeNamespaceLister
// interface.
type passwordTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PasswordTypes in the indexer for a given namespace.
func (s passwordTypeNamespaceLister) List(selector labels.Selector) (ret []*v1.PasswordType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PasswordType))
	})
	return ret, err
}

// Get retrieves the PasswordType from the indexer for a given namespace and name.
func (s passwordTypeNamespaceLister) Get(name string) (*v1.PasswordType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("passwordtype"), name)
	}
	return obj.(*v1.PasswordType), nil
}