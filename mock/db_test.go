package mock

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	v := GetFromDB(m, "Tom")
	if v != -1 {
		t.Fatal("expected -1, but got", v)
	}

	fmt.Println("获取的值v", v)
}
