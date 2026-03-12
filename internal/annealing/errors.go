package annealing

import "errors"

var ErrNilObjective = errors.New("целевая функция не должна быть nil")
var ErrNonPositiveDimension = errors.New("размерность должна быть положительной")
var ErrInvalidIterations = errors.New("число итераций должно быть положительным")
var ErrInvalidInitialTemperature = errors.New("начальная температура должна быть положительной")
var ErrInvalidMinTemperature = errors.New("минимальная температура должна быть положительной")
var ErrInvalidCoolingRate = errors.New("коэффициент охлаждения должен быть в интервале (0, 1)")
var ErrInvalidStepSize = errors.New("размер шага должен быть положительным")
var ErrInvalidEpsilon = errors.New("epsilon должен быть положительным")
var ErrInvalidBounds = errors.New("границы должны соответствовать размерности и нижняя граница должна быть меньше верхней")
var ErrStartDimensionMismatch = errors.New("размерность начальной точки не совпадает с размерностью функции")
var ErrStartOutOfBounds = errors.New("начальная точка должна лежать внутри заданных границ")
