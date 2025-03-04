package store

import "github.com/guneyin/yarbay/utils"

type Value[T any] struct {
	val T
}

func (v *Value[T]) Parse(dest T) error {
	_, err := utils.Convert(v.val, dest)
	if err != nil {
		return err
	}
	return nil
}
