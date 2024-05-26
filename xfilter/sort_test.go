package xfilter

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name   string
	No     int64
	BornTm time.Time
}

func (u *User) GetStrSortField() string {
	return u.Name
}
func (u *User) GetIntSortField() int64 {
	return u.No
}
func (u *User) GetTmSortField() time.Time {
	return u.BornTm
}

func TestSort(t *testing.T) {
	list := []*User{}
	list = append(list, &User{Name: "a2", No: 4, BornTm: time.Now().Add(time.Hour * 4)})
	list = append(list, &User{Name: "1", No: 2, BornTm: time.Now().Add(time.Hour * 2)})
	list = append(list, &User{Name: "中a", No: 8, BornTm: time.Now().Add(time.Hour * 8)})
	list = append(list, &User{Name: "啊bc", No: 7, BornTm: time.Now().Add(time.Hour * 7)})
	list = append(list, &User{Name: "阿ab", No: 6, BornTm: time.Now().Add(time.Hour * 6)})
	list = append(list, &User{Name: "b1", No: 5, BornTm: time.Now().Add(time.Hour * 5)})
	list = append(list, &User{Name: "2a", No: 3, BornTm: time.Now().Add(time.Hour * 3)})
	list = append(list, &User{Name: "~!天", No: 1, BornTm: time.Now().Add(time.Hour * 1)})

	fmt.Printf("old list: \n")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%+v \n", list[i])
	}

	list = SortByString(list, "asc")

	fmt.Printf("SortByString new list: \n")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%+v \n", list[i])
	}

	list = SortByNum(list, "desc")

	fmt.Printf("SortByNum new list: \n")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%+v \n", list[i])
	}

	list = SortByTime(list, "asc")

	fmt.Printf("SortByTime new list: \n")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%+v \n", list[i])
	}
}
