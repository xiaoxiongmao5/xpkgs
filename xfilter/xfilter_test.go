package xfilter

import (
	"testing"
	"time"
)

type User struct {
	Name      string
	CreatedTm string
}

func (u *User) GetName() string {
	return u.Name
}
func (u *User) GetCreatedAt() time.Time {
	tm, _ := time.Parse("2006-01-02", u.CreatedTm)
	return tm
}

func TestFilterList(t *testing.T) {
	users := make([]*User, 0)
	users = append(users, &User{Name: "a张三", CreatedTm: "2021-01-01"})
	users = append(users, &User{Name: "李四", CreatedTm: "2021-01-02"})
	users = append(users, &User{Name: "b王五", CreatedTm: "2021-01-03"})
	users = append(users, &User{Name: "1赵六", CreatedTm: "2021-01-04"})
	users = append(users, &User{Name: "c孙七", CreatedTm: "2021-01-05"})
	users = append(users, &User{Name: "周八", CreatedTm: "2021-01-06"})
	users = append(users, &User{Name: "吴九", CreatedTm: "2021-01-07"})
	users = append(users, &User{Name: "郑十", CreatedTm: "2021-01-08"})

	// 按照名字拼音首字母升序排序
	list := FilterList[*User](users,
		&SortFilter[*User]{SortType: "asc", SortField: "name"},
	)
	expectList := []string{"1赵六", "a张三", "b王五", "c孙七", "李四", "吴九", "郑十", "周八"}
	Check(t, list, expectList)

	// 按照创建时间降序排序
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "createdTm"},
	)
	expectList = []string{"郑十", "吴九", "周八", "c孙七", "1赵六", "b王五", "李四", "a张三"}
	Check(t, list, expectList)

	// 按照创建时间升序排序，分页
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "asc", SortField: "createdTm"},
		&PageFilter[*User]{PageNo: 1, PageSize: 3},
	)
	expectList = []string{"a张三", "李四", "b王五"}
	Check(t, list, expectList)

	// 按照名字拼音首字母降序排序，分页
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "name"},
		&PageFilter[*User]{PageNo: 3, PageSize: 3},
	)
	expectList = []string{"a张三", "1赵六"}
	Check(t, list, expectList)

	// 测试边界
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "name"},
		&PageFilter[*User]{PageNo: 5, PageSize: 10},
	)
	expectList = []string{}
	Check(t, list, expectList)

	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "name"},
		&PageFilter[*User]{PageNo: 0, PageSize: 10},
	)
	expectList = []string{"周八", "郑十", "吴九", "李四", "c孙七", "b王五", "a张三", "1赵六"}
	Check(t, list, expectList)
}

func Check(t *testing.T, list []*User, expectList []string) {
	for i, user := range list {
		if user.Name != expectList[i] {
			t.Errorf("Expected name to be %s, but got %s", expectList[i], user.Name)
		}
	}
}
