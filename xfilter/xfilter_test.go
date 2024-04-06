package xfilter

import (
	"testing"
	"time"
)

type User struct {
	Name      string
	CreatedTm string
	UpdateTm  time.Time
	Nick      string
}

func (u *User) GetName() string {
	return u.Name
}
func (u *User) GetCreatedAt() time.Time {
	tm, _ := time.Parse("2006-01-02", u.CreatedTm)
	return tm
}

func getUserlist() []*User {
	users := make([]*User, 0)
	users = append(users, &User{Name: "a张三", CreatedTm: "2021-01-01", UpdateTm: time.Now().AddDate(0, 0, -1), Nick: "a张三_nick"})
	users = append(users, &User{Name: "李四", CreatedTm: "2021-01-02", UpdateTm: time.Now().AddDate(0, 0, -2), Nick: "李四_nick"})
	users = append(users, &User{Name: "b王五", CreatedTm: "2021-01-03", UpdateTm: time.Now().AddDate(0, 0, -3), Nick: "b王五_nick"})
	users = append(users, &User{Name: "21赵六", CreatedTm: "2021-01-04", UpdateTm: time.Now().AddDate(0, 0, -4), Nick: "21赵六_nick"})
	users = append(users, &User{Name: "c孙七", CreatedTm: "2021-01-05", UpdateTm: time.Now().AddDate(0, 0, -5), Nick: "c孙七_nick"})
	users = append(users, &User{Name: "周八", CreatedTm: "2021-01-06", UpdateTm: time.Now().AddDate(0, 0, -6), Nick: "周八_nick"})
	users = append(users, &User{Name: "吴九", CreatedTm: "2021-01-07", UpdateTm: time.Now().AddDate(0, 0, -7), Nick: "吴九_nick"})
	users = append(users, &User{Name: "郑十", CreatedTm: "2021-01-08", UpdateTm: time.Now().AddDate(0, 0, -8), Nick: "郑十_nick"})
	return users
}

func getUserlist1() []User {
	users1 := make([]User, 0)
	users1 = append(users1, User{Name: "a张三", CreatedTm: "2021-01-01", UpdateTm: time.Now().AddDate(0, 0, -1), Nick: "a张三_nick"})
	users1 = append(users1, User{Name: "李四", CreatedTm: "2021-01-02", UpdateTm: time.Now().AddDate(0, 0, -2), Nick: "李四_nick"})
	users1 = append(users1, User{Name: "b王五", CreatedTm: "2021-01-03", UpdateTm: time.Now().AddDate(0, 0, -3), Nick: "b王五_nick"})
	users1 = append(users1, User{Name: "21赵六", CreatedTm: "2021-01-04", UpdateTm: time.Now().AddDate(0, 0, -4), Nick: "21赵六_nick"})
	users1 = append(users1, User{Name: "c孙七", CreatedTm: "2021-01-05", UpdateTm: time.Now().AddDate(0, 0, -5), Nick: "c孙七_nick"})
	users1 = append(users1, User{Name: "周八", CreatedTm: "2021-01-06", UpdateTm: time.Now().AddDate(0, 0, -6), Nick: "周八_nick"})
	users1 = append(users1, User{Name: "吴九", CreatedTm: "2021-01-07", UpdateTm: time.Now().AddDate(0, 0, -7), Nick: "吴九_nick"})
	users1 = append(users1, User{Name: "郑十", CreatedTm: "2021-01-08", UpdateTm: time.Now().AddDate(0, 0, -8), Nick: "郑十_nick"})
	return users1
}

func TestFilterList(t *testing.T) {
	users := getUserlist()
	// 按照名字拼音首字母升序排序
	list := FilterList[*User](users,
		&SortFilter[*User]{SortType: "asc", SortField: "name"},
	)
	Check(t, list, []string{"21赵六", "a张三", "b王五", "c孙七", "李四", "吴九", "郑十", "周八"})

	// 按照创建时间降序排序
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "createdTm"},
	)
	Check(t, list, []string{"郑十", "吴九", "周八", "c孙七", "21赵六", "b王五", "李四", "a张三"})

	// 按照创建时间升序排序，分页
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "asc", SortField: "createdTm"},
		&PageFilter[*User]{PageNo: 1, PageSize: 3},
	)
	Check(t, list, []string{"a张三", "李四", "b王五"})

	// 按照名字拼音首字母降序排序，分页
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "name"},
		&PageFilter[*User]{PageNo: 3, PageSize: 3},
	)
	Check(t, list, []string{"a张三", "21赵六"})

	// 测试边界
	list = FilterList[*User](users,
		&SortFilter[*User]{SortType: "desc", SortField: "name"},
		&PageFilter[*User]{PageNo: 5, PageSize: 10},
	)
	Check(t, list, []string{})

	// 测试反射排序
	users1 := getUserlist1()
	list1 := FilterList[User](users1,
		&SortFilterByReflect[User]{SortType: "desc", SortField: "Nick"},
		&PageFilter[User]{PageNo: 0, PageSize: 10},
	)
	Check1(t, list1, []string{"周八", "郑十", "吴九", "李四", "c孙七", "b王五", "a张三", "21赵六"})

	list1 = FilterList[User](users1,
		&SortFilterByReflect[User]{SortType: "asc", SortField: "UpdateTm"},
	)
	Check1(t, list1, []string{"郑十", "吴九", "周八", "c孙七", "21赵六", "b王五", "李四", "a张三"})
}

func Check(t *testing.T, list []*User, expectList []string) {
	for i, user := range list {
		if user.Name != expectList[i] {
			t.Errorf("Expected name to be %s, but got %s", expectList[i], user.Name)
		}
	}
}
func Check1(t *testing.T, list []User, expectList []string) {
	for i, user := range list {
		if user.Name != expectList[i] {
			t.Errorf("Expected name to be %s, but got %s", expectList[i], user.Name)
		}
	}
}
