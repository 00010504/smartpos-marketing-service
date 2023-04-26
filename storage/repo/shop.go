package repo

import "genproto/common"

type ShopI interface {
	Upsert(*common.ShopCreatedModel) error
}
