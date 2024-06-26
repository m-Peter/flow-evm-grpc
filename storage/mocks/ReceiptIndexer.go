// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/onflow/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	models "github.com/onflow/flow-evm-gateway/models"

	types "github.com/onflow/go-ethereum/core/types"
)

// ReceiptIndexer is an autogenerated mock type for the ReceiptIndexer type
type ReceiptIndexer struct {
	mock.Mock
}

// BloomsForBlockRange provides a mock function with given fields: start, end
func (_m *ReceiptIndexer) BloomsForBlockRange(start *big.Int, end *big.Int) ([]*types.Bloom, []*big.Int, error) {
	ret := _m.Called(start, end)

	if len(ret) == 0 {
		panic("no return value specified for BloomsForBlockRange")
	}

	var r0 []*types.Bloom
	var r1 []*big.Int
	var r2 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) ([]*types.Bloom, []*big.Int, error)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) []*types.Bloom); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Bloom)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int) []*big.Int); ok {
		r1 = rf(start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(*big.Int, *big.Int) error); ok {
		r2 = rf(start, end)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByBlockHeight provides a mock function with given fields: height
func (_m *ReceiptIndexer) GetByBlockHeight(height *big.Int) ([]*models.StorageReceipt, error) {
	ret := _m.Called(height)

	if len(ret) == 0 {
		panic("no return value specified for GetByBlockHeight")
	}

	var r0 []*models.StorageReceipt
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int) ([]*models.StorageReceipt, error)); ok {
		return rf(height)
	}
	if rf, ok := ret.Get(0).(func(*big.Int) []*models.StorageReceipt); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.StorageReceipt)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTransactionID provides a mock function with given fields: ID
func (_m *ReceiptIndexer) GetByTransactionID(ID common.Hash) (*models.StorageReceipt, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByTransactionID")
	}

	var r0 *models.StorageReceipt
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*models.StorageReceipt, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *models.StorageReceipt); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.StorageReceipt)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: receipt
func (_m *ReceiptIndexer) Store(receipt *models.StorageReceipt) error {
	ret := _m.Called(receipt)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.StorageReceipt) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReceiptIndexer creates a new instance of ReceiptIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReceiptIndexer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReceiptIndexer {
	mock := &ReceiptIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
