package service

type CheckItem struct {
	Desc   string
	Result uint8
}

type CheckItems []*CheckItem
