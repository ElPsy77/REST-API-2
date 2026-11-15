package todo

import "errors"

var ErrTaskNotFound = errors.New("задача не найдена")
var ErrTaskAlreadyExist = errors.New("задача уже создана")
