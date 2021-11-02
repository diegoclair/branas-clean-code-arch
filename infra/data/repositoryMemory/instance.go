package repositorymemory

import "github.com/diegoclair/branas-clean-code-arch/domain/contract"

type repoMemory struct {
}

func New() contract.RepoManager {
	return &repoMemory{}
}

func (c *repoMemory) Coupon() contract.CouponRepository {
	return newCouponRepositoryMemory()
}

func (c *repoMemory) Item() contract.ItemRepository {
	return newItemRepositoryMemory()
}

func (c *repoMemory) Order() contract.OrderRepository {
	return newOrderRepositoryMemory()
}
