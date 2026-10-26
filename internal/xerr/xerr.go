package xerr

// 错误码  10000 0000 ~~99999 9999
// 模块id  455907
// 错误码 = 模块id + 业务错误码
var (
	ModuleId = int64(455907)
)

var (
	// 系统错误 000-099
	ErrCodeServerInternal = int64(455907001)
	ErrCodeCallRpc        = int64(455907002)

	// 业务错误码 100-199
	ErrCodeParam = int64(455907100)
)
