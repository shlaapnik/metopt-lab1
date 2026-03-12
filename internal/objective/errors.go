package objective

import "errors"

var ErrInvalidDimension = errors.New("размерность должна быть не меньше 2")
var ErrPointDimensionMismatch = errors.New("размерность точки не совпадает с размерностью целевой функции")
