package result

type Code int64

const (
	Code_Success      Code = 0
	Code_ParamInvalid Code = 1
	Code_DBErr        Code = 2
	Code_RTErr        Code = 3
)
