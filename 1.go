//问题1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//答：可以直接抛给上层，但如果想给上层提供更多的信息并且不在乎细节泄漏给上层，也可以wrap这个error抛给上层。


package GeekHomework

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	id uint64
	name string
}

type Dao interface {
	GetById(id uint64) interface{}
}

type UserDao struct {}

func (user *UserDao) GetById(id uint64) (*User, error) {
	user := User{}
	err := db.QueryRow("select name from users where id = ?", id).Scan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, fmt.Sprintf("find user null, user id: %v", id))
	} else {
		return nil, err
	}
	return &user, nil
}