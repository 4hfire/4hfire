package errors

// IMPORTANT: Code Must Be Append Only! NO REMOVE OR REPLACE ORDER
// 重要: 状态码仅可新增, 禁止删除或重排序
const (
	// 	ErrUnkonwn 未知错误
	ErrCodeUnkonwn = 999999 - iota

	//  自定义错误
	ErrCodeDIY

	// 	ErrSystem 系统错误
	ErrCodeSystem

	// 	ErrPermission 权限错误
	ErrCodePermission

	// 	ErrUnauth 未授权登录
	ErrCodeUnauth

	// 	ErrDB 数据库操作错误
	ErrCodeDB

	// 	ErrCodeReqParam 请求参数错误
	ErrCodeReqParam

	//	ErrCodeBusy 系统繁忙错误
	ErrCodeBusy

	//	ErrRpcSystem rpc通信错误
	ErrCodeRpcSystem

	// 	ErrCodeNotFound 未查找到记录错误
	ErrCodeNotFound

	// 	ErrCode 缓存错误
	ErrCodeCache

	//	数据过期错误
	ErrCodeExpired

	//	验证错误
	ErrCodeUnchecked

	//	操作禁止或用户禁止
	ErrCodeForbid

	//  存在记录无法操作
	ErrCodeExistCanNotOperate

	//	登录失败
	ErrCodeLogin

	//	数据重复
	ErrCodeRepeated

	//  上传文件格式错误
	ErrCodeFormat

	//	数据锁定 无法操作
	ErrCodeDataLocked

	//	更新失败
	ErrCodeUpdate

	//	角色关联账号 无法删除
	ErrCodeRoleRelatedCannotDelete

	// 	账号源绑定组织, 无法删除组织
	ErrCodeSourceBindOrgCannotDelete

	//  网关已被绑定到网关组 无法删除
	ErrCodeGatewayBindCannotDelete

	// 网关启用状态 无法删除
	ErrCodeGatewayCheckedDelete

	//	账号包含中文字符 无法新增
	ErrCodeAccountContainsZhChar

	//	账户重复 无法新增
	ErrCodeAccountRepeated

	// 新增规则失败,存在同名属性
	ErrCodeRuleAddRepeated

	// 新增凭证失败,存在同名属性
	ErrCodeCertAddRepeated

	// sdp控制器同步错误
	ErrCodeSdpController

	// 信用分超出范围 0-100
	ErrCodeScoreOutOfRange

	//	存在相同应用分
	ErrCodeScoreExist

	//	非系统管理员无法创建管理员账号 仅可创建普通用户账号
	ErrCodeNotAdminCreateAdmin

	//	手机号检验失败
	ErrCodeValidatePhone
	//	邮箱检验失败
	ErrCodeValidateEmail

	//	ip域名校验失败
	ErrCodeValidateIpDomain

	//	导入用户数据错误 请检查数据源
	ErrCodeImport

	//	改信用等级已经被应用使用,无法删除
	ErrCodeCreditUsed

	//  账号源配置错误 连接失败
	ErrCodeSourceConfig

	//	数据已存在, 请勿重复添加
	ErrCodeExistCanNotAdd

	//	连接无效,无法连接
	ErrCodeConnCannotConnect

	//	当前源存在正在进行的同步记录
	ErrCodeExsistSyncLog

	//	sp entity id repeated error
	ErrCodeSpEntityIDRepeated

	//	sp 信息不存在
	ErrCodeSpMissing

	//	内建用户不可加入到同步源的组织中
	ErrCodeAddBuiltInUserToSourceOrg

	//  组织不存在
	ErrCodeOrgMissing

	//	当前源已存在内建用户
	// ErrCodeLDAPADSourceAlreadyHasUser

	//	已存在相同code用户
	ErrCodeUserHaveSameCode

	//	缓存已过期
	ErrCodeCaptchaExpired

	//	账号源已禁用
	ErrCodeSourceConfigDisabled

	//	网关节点不存在
	ErrCodeGatewayNotExist

	//	网关节点异常
	ErrCodeGateway

	//	用户不存在
	ErrCodeUserNotExist

	//	设备和设备历史记录删除失败
	ErrCodeDeviceDelete

	//	平台SN异常
	ErrCodePlatformSN

	//	非法的平台激活密钥
	ErrCodeInvalidPlatformKey

	//	密钥已过期
	ErrCodeExpiredKey

	//	文件切片不存在
	ErrCodeFileSliceNotExist

	//	文件名重复
	ErrCodeFileNameRepeated

	//	不可操作自己的账号
	ErrCodeSuicide

	//	授权的设备数量不足 请先注册设备
	ErrCodeAuthorizedDeviceLack

	//  不是同一天
	ErrCodeNotSameDay

	//  获取waf日志出错
	ErrCodeEdgeApiAccessLog

	// 编辑授权组织参数错误, 仅能包含一个授权对象
	ErrCodeEditGrantOrg

	// 获取OTP信息失败
	ErrGetOTPInfo

	// OTP信息已失效
	ErrOTPInfoExpired

	// OTP二次验证码错误
	ErrOTPCode

	// 连接配置格式错误
	ErrConnConfig

	// 应用标签重复
	ErrAppTagsRepeated
)

var ErrMap = map[string]map[int32]string{
	"zh": {
		ErrCodeUnkonwn:                   "未知错误",
		ErrCodeSystem:                    "系统错误",
		ErrCodePermission:                "权限错误",
		ErrCodeUnauth:                    "未授权登录",
		ErrCodeDB:                        "数据库错误",
		ErrCodeReqParam:                  "请求参数错误",
		ErrCodeBusy:                      "系统繁忙",
		ErrCodeRpcSystem:                 "RPC系统错误",
		ErrCodeNotFound:                  "未找到记录",
		ErrCodeCache:                     "缓存错误",
		ErrCodeExpired:                   "数据过期",
		ErrCodeUnchecked:                 "验证错误",
		ErrCodeForbid:                    "用户禁止或操作禁止",
		ErrCodeExistCanNotOperate:        "已存在记录无法操作或删除",
		ErrCodeLogin:                     "用户名或密码错误",
		ErrCodeRepeated:                  "数据参数重复,不可重复添加",
		ErrCodeFormat:                    "上传文件格式错误",
		ErrCodeDataLocked:                "当前数据已锁定,无法编辑",
		ErrCodeUpdate:                    "更新失败",
		ErrCodeRoleRelatedCannotDelete:   "当前角色已关联账号,无法删除",
		ErrCodeSourceBindOrgCannotDelete: "账号源绑定组织, 无法删除组织",
		ErrCodeGatewayBindCannotDelete:   "网关已被绑定到网关组 无法删除",
		ErrCodeAccountContainsZhChar:     "账号包含中文字符 无法新增",
		ErrCodeRuleAddRepeated:           "新增规则失败,存在同名属性",
		ErrCodeCertAddRepeated:           "新增凭证失败,存在同名属性",
		ErrCodeSdpController:             "sdp控制器同步错误",
		ErrCodeScoreOutOfRange:           "信用分超出范围 0-100以内",
		ErrCodeScoreExist:                "已存在相同的信用分",
		ErrCodeNotAdminCreateAdmin:       "非系统管理员无法创建管理员账号 仅可创建普通用户账号",
		ErrCodeAccountRepeated:           "账户重复 无法新增",
		ErrCodeValidatePhone:             "手机号校验失败",
		ErrCodeValidateEmail:             "邮箱校验失败",
		ErrCodeValidateIpDomain:          "ip或域名校验失败",
		ErrCodeImport:                    "导入用户数据错误 请检查数据源",
		ErrCodeCreditUsed:                "该信用等级已经被应用使用,无法删除",
		ErrCodeSourceConfig:              "账号源配置错误,连接失败",
		ErrCodeExistCanNotAdd:            "数据已存在, 请勿重复添加",
		ErrCodeConnCannotConnect:         "连接无效,无法连接",
		ErrCodeExsistSyncLog:             "当前源存在正在进行的同步记录",
		ErrCodeSpEntityIDRepeated:        "服务提供方的实体ID不能重复",
		ErrCodeSpMissing:                 "sp info 不存在, 请在零信任管理后台注册该应用",
		ErrCodeAddBuiltInUserToSourceOrg: "内建用户不可加入到同步源的组织架构中",
		ErrCodeOrgMissing:                "组织不存在",
		ErrCodeUserHaveSameCode:          "已存在相同用户编码的用户",
		ErrCodeCaptchaExpired:            "验证图已过期, 请重新获取",
		ErrCodeSourceConfigDisabled:      "账号源已禁用, 无法同步",
		ErrCodeGatewayNotExist:           "网关节点不存在",
		ErrCodeGateway:                   "网关节点异常",
		ErrCodeUserNotExist:              "用户不存在",
		ErrCodeDeviceDelete:              "设备和设备历史记录删除失败",
		ErrCodePlatformSN:                "平台sn获取失败",
		ErrCodeInvalidPlatformKey:        "非法的平台激活密钥",
		ErrCodeExpiredKey:                "密钥已过期",
		ErrCodeFileSliceNotExist:         "文件切片不存在",
		ErrCodeFileNameRepeated:          "文件名重复",
		ErrCodeSuicide:                   "不可操作自己的账号",
		ErrCodeAuthorizedDeviceLack:      "授权的设备数量不足 请先注册设备",
		ErrCodeNotSameDay:                "开始和结束日期不是同一天",
		ErrCodeEdgeApiAccessLog:          "获取waf日志出错",
		ErrCodeEditGrantOrg:              "编辑授权组织参数错误, 仅能包含一个授权对象",
		ErrGetOTPInfo:                    "获取otp信息失败",
		ErrOTPInfoExpired:                "OTP信息已失效, 请重新获取",
		ErrOTPCode:                       "二次验证码错误",
		ErrConnConfig:                    "连接配置格式错误",
		ErrAppTagsRepeated:               "应用标签重复",
	},
	"en": {
		ErrCodeUnkonwn:                   "unknown error",
		ErrCodeSystem:                    "system error",
		ErrCodePermission:                "permission error",
		ErrCodeUnauth:                    "unauthentificated error",
		ErrCodeDB:                        "database error",
		ErrCodeReqParam:                  "request params error",
		ErrCodeBusy:                      "system busy error",
		ErrCodeRpcSystem:                 "rpc system error",
		ErrCodeNotFound:                  "record not found error",
		ErrCodeCache:                     "cache error",
		ErrCodeExpired:                   "data expired error",
		ErrCodeUnchecked:                 "data unchecked error",
		ErrCodeForbid:                    "user or operation forbidden",
		ErrCodeExistCanNotOperate:        "record exist can not be operated or deleted",
		ErrCodeLogin:                     "username or password error",
		ErrCodeRepeated:                  "data repeated, cannot insert repeatedly",
		ErrCodeFormat:                    "uoload file format error",
		ErrCodeDataLocked:                "data locked, cannot edit",
		ErrCodeUpdate:                    "update error",
		ErrCodeRoleRelatedCannotDelete:   "current role has related to a role, cannot be deleted",
		ErrCodeSourceBindOrgCannotDelete: "current org has related to source, cannot be deleted",
		ErrCodeGatewayBindCannotDelete:   "current gateway has been binded to a gateway set, cannot be delete",
		ErrCodeAccountContainsZhChar:     "account contains zh-cn character, cannot insert",
		ErrCodeRuleAddRepeated:           "rule add error, same property in record",
		ErrCodeCertAddRepeated:           "cert add error, same property in record",
		ErrCodeSdpController:             "sdp controller sync error",
		ErrCodeScoreOutOfRange:           "score out of range, need 0-100",
		ErrCodeScoreExist:                "record exists with same score",
		ErrCodeNotAdminCreateAdmin:       "current user is not amdin account, cannot create account except common staff",
		ErrCodeAccountRepeated:           "record exists with same account",
		ErrCodeValidatePhone:             "phone format error",
		ErrCodeValidateEmail:             "email format error",
		ErrCodeValidateIpDomain:          "ip or domain format error",
		ErrCodeImport:                    "import error, please check your user source",
		ErrCodeCreditUsed:                "this credit level is been used by some applications, cannot delete",
		ErrCodeSourceConfig:              "source config params errors, cannot connect",
		ErrCodeExistCanNotAdd:            "data exists, cannot add repeatedly",
		ErrCodeConnCannotConnect:         "conn invalid, connot connect",
		ErrCodeExsistSyncLog:             "current source aleady have existing sync record, wait util it's finished",
		ErrCodeSpEntityIDRepeated:        "sp entity id cannot repeate",
		ErrCodeSpMissing:                 "sp info is missing, please regist in zero-trust console",
		ErrCodeAddBuiltInUserToSourceOrg: "built-in user cannot add to organization which is from ldap/ad source",
		ErrCodeOrgMissing:                "organization is not exist",
		ErrCodeUserHaveSameCode:          "already exist user with same code",
		ErrCodeCaptchaExpired:            "captcha was expired, please re-get captcha",
		ErrCodeSourceConfigDisabled:      "source config is diabled, cannot sync",
		ErrCodeGatewayNotExist:           "gateway is not exist",
		ErrCodeGateway:                   "gateway node error",
		ErrCodeUserNotExist:              "user is not exist",
		ErrCodeDeviceDelete:              "delete device with device_history logs error",
		ErrCodePlatformSN:                "platform sn get error or is empty",
		ErrCodeInvalidPlatformKey:        "invalid platform key provided",
		ErrCodeExpiredKey:                "expired key provided",
		ErrCodeFileSliceNotExist:         "file slice is not exist",
		ErrCodeFileNameRepeated:          "file name is repeated",
		ErrCodeSuicide:                   "cannot operate yourself",
		ErrCodeAuthorizedDeviceLack:      "lack of authorized device count, please regist first",
		ErrCodeNotSameDay:                "start date and end date is not the same day",
		ErrCodeEdgeApiAccessLog:          "get waf log error",
		ErrCodeEditGrantOrg:              "edit org grant info error, only one target can be supported",
		ErrGetOTPInfo:                    "get user otp info error",
		ErrOTPInfoExpired:                "otp info was expired, please reacqure again",
		ErrOTPCode:                       "otp code error",
		ErrConnConfig:                    "connection config format error",
		ErrAppTagsRepeated:               "application tags is repeated",
	},
}
