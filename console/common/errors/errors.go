/*
 * LastEditors: lihy lihy@zhiannet.com
 * @Date: 2022-10-24 20:50:44
 * LastEditTime: 2023-10-16 11:43:24
 * @FilePath: /zero-trust/console/IAM/common/errors/errors.go
 */

package errors

import (
	"encoding/json"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RawError(s string) error {
	return errors.New(s)
}

type Error struct {
	code int32  `json:"errCode"`
	msg  string `json:"errMsg"`
}

func (e *Error) Code() int32 {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Error() string {
	return e.msg
	// val, err := json.Marshal(e)
	// _ = err
	// return string(val)
}

// Parse 根据json格式的数据解释为错误对象
// 数据格式：{"errCode":0,"errMsg":"message"}
func Parse(err string) *Error {
	var e Error
	_ = json.Unmarshal([]byte(err), &e)
	return &e
}

func FromRpcError(errRpc error) *Error {
	s, _ := status.FromError(errRpc)
	if s == nil {
		return NewErrCodeMsg(ErrCodeUnkonwn, "未知错误")
	}
	/*
		e := &Error{}
		err := json.Unmarshal([]byte(s.Message()), e)
		if err != nil {
			return NewErrCodeMsg(ErrUnkonwn, "未知错误")
		}
	*/
	return NewErrCodeMsg(int32(s.Code()), s.Message())
}

func NewErrCodeMsg(code int32, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func NewErrCode(code int32) *Error {
	if v, ok := ErrMap["zh"][code]; ok {
		return &Error{code: code, msg: v}
	}
	return &Error{code: code, msg: "unkonwn error message"}
}

func SystemError() error {
	return NewErrCodeMsg(ErrCodeSystem, "系统错误，请稍后重试")
}

func SystemBusy() error {
	return NewErrCodeMsg(ErrCodeBusy, "服务器繁忙，请稍后重试")
}

func RpcSystemError() error {
	return NewErrCodeMsg(ErrCodeRpcSystem, "系统错误，请稍后重试").RpcError()
}

func ErrorNotFound() error {
	return NewErrCodeMsg(ErrCodeNotFound, "未找到记录")
}

func PermissionError() error {
	return NewErrCodeMsg(ErrCodePermission, "权限不足")
}

func ParamsError(msg ...string) error {
	m := "参数错误"
	if len(msg) > 0 {
		m = msg[0]
	}
	return NewErrCodeMsg(ErrCodeReqParam, m)
}

func UnauthError() error {
	return NewErrCodeMsg(ErrCodeUnauth, "未授权登录")
}

func DbError() *Error {
	return NewErrCodeMsg(ErrCodeDB, "数据库错误")
}

func (e *Error) RpcError() error {
	return status.Error(codes.Aborted, e.Msg())
}

func CacheError() error {
	return NewErrCodeMsg(ErrCodeCache, "缓存错误")
}

func ExpiredError() error {
	return NewErrCodeMsg(ErrCodeExpired, "数据过期")
}

func UncheckedError() error {
	return NewErrCodeMsg(ErrCodeUnchecked, "验证错误")
}

func ForbidError() error {
	return NewErrCode(ErrCodeForbid)
}

func ExistCanNotOperate() error {
	return NewErrCode(ErrCodeExistCanNotOperate)
}

func LoginError() error {
	return NewErrCode(ErrCodeLogin)
}

func RepeatedError() error {
	return NewErrCode(ErrCodeRepeated)
}

func DiyError(s string) error {
	return NewErrCodeMsg(ErrCodeDIY, s)
}

func FormatError() error {
	return NewErrCode(ErrCodeFormat)
}

func CannotEdit() error {
	return NewErrCode(ErrCodeDataLocked)
}

func UpdateError() error {
	return NewErrCode(ErrCodeUpdate)
}

func RoleRelatedCannotDelete() error {
	return NewErrCode(ErrCodeRoleRelatedCannotDelete)
}

func SourceBindOrgCannotDelete() error {
	return NewErrCode(ErrCodeSourceBindOrgCannotDelete)
}

func GatewayUsedCannotDelete() error {
	return NewErrCode(ErrCodeGatewayBindCannotDelete)
}

func AccountContainsZhError() error {
	return NewErrCode(ErrCodeAccountContainsZhChar)
}

func AccountRepeatedError() error {
	return NewErrCode(ErrCodeAccountRepeated)
}

func RuleAddRepatedError() error {
	return NewErrCode(ErrCodeRuleAddRepeated)
}

func CertAddRepatedError() error {
	return NewErrCode(ErrCodeRuleAddRepeated)
}

func SdpControllerError() error {
	return NewErrCode(ErrCodeSdpController)
}

func ScoreOutOfRangeError() error {
	return NewErrCode(ErrCodeScoreOutOfRange)
}

func ScoreExistError() error {
	return NewErrCode(ErrCodeScoreExist)
}

func NotAdminCreateAdminError() error {
	return NewErrCode(ErrCodeNotAdminCreateAdmin)
}

func ValidatePhoneError() error {
	return NewErrCode(ErrCodeValidatePhone)
}
func ValidateEmailError() error {
	return NewErrCode(ErrCodeValidateEmail)
}

func ValidateIpDomainError() error {
	return NewErrCode(ErrCodeValidateIpDomain)
}

func ImportError() error {
	return NewErrCode(ErrCodeImport)
}

func CreditUsedError() error {
	return NewErrCode(ErrCodeCreditUsed)
}

func SourceConfigError() error {
	return NewErrCode(ErrCodeSourceConfig)
}

func ExistCanNotAdd() error {
	return NewErrCode(ErrCodeExistCanNotAdd)
}

func ConnCannotConnect() error {
	return NewErrCode(ErrCodeConnCannotConnect)
}

func ExsistSyncLog() error {
	return NewErrCode(ErrCodeExsistSyncLog)
}

func SpEntityIDRepeated() error {
	return NewErrCode(ErrCodeSpEntityIDRepeated)
}

func SpMissing() error {
	return NewErrCode(ErrCodeSpMissing)
}

func AddBuiltInUserToSourceOrg() error {
	return NewErrCode(ErrCodeAddBuiltInUserToSourceOrg)
}

func OrgMissing() error {
	return NewErrCode(ErrCodeOrgMissing)
}

// func LDAPADSourceAlreadyHasUser() error {
// 	return NewErrCode(ErrCodeLDAPADSourceAlreadyHasUser)
// }

func UserHaveSameCode() error {
	return NewErrCode(ErrCodeUserHaveSameCode)
}

func CaptchaExpired() error {
	return NewErrCode(ErrCodeCaptchaExpired)
}

func SourceConfigDisabled() error {
	return NewErrCode(ErrCodeSourceConfigDisabled)
}

func GatewayNotExistError() error {
	return NewErrCode(ErrCodeGatewayNotExist)
}

func GatewayError() error {
	return NewErrCode(ErrCodeGateway)
}

func UserNotExist() error {
	return NewErrCode(ErrCodeUserNotExist)
}

func DeviceDelete() error {
	return NewErrCode(ErrCodeDeviceDelete)
}

func PlatformSNError() error {
	return NewErrCode(ErrCodePlatformSN)
}

func InvalidPlatformKey() error {
	return NewErrCode(ErrCodeInvalidPlatformKey)
}

func ExpiredKeyProvided() error {
	return NewErrCode(ErrCodeExpiredKey)
}

func FileSliceNotExist() error {
	return NewErrCode(ErrCodeFileSliceNotExist)
}

func FileNameRepeated() error {
	return NewErrCode(ErrCodeFileNameRepeated)
}

func CommitSuicide() error {
	return NewErrCode(ErrCodeSuicide)
}

func AuthorizedDeviceLack() error {
	return NewErrCode(ErrCodeAuthorizedDeviceLack)
}

func NotSameDayError() error {
	return NewErrCode(ErrCodeNotSameDay)
}

func EdgeApiAccessLogError() error {
	return NewErrCode(ErrCodeEdgeApiAccessLog)
}

func EditGrantOrgError() error {
	return NewErrCode(ErrCodeEditGrantOrg)
}

func GetOTPInfoError() error {
	return NewErrCode(ErrGetOTPInfo)
}

func OTPInfoExpired() error {
	return NewErrCode(ErrOTPInfoExpired)
}

func OTPCodeError() error {
	return NewErrCode(ErrOTPCode)
}

func ConnConfigError() error {
	return NewErrCode(ErrConnConfig)
}

func AppTagsRepeated() error {
	return NewErrCode(ErrAppTagsRepeated)
}
