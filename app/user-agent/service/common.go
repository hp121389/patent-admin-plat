package service

import "fmt"

var ErrConflictBindPatent = fmt.Errorf("can not bind patent repeatly")
var ErrCanNotUpdate = fmt.Errorf("can not update record")
