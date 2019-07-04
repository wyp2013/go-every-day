package common

import "errors"

const (
	ERR_OK             = 0
	ERR_PARAM          = 1
	ERR_SYS            = 2
	ERR_UNKNOWN        = 3
	ERR_JSON_ENCODE    = 4
	ERR_JSON_DECODE    = 5
	ERR_TOKEN          = 6
	ERR_DB             = 7
	ERR_HTTP           = 8
	ERR_LOCK           = 9
	AUTH_INVALID       = 10
	ERR_PART_FAILED    = 11
	ERR_VERSION_IDENTICAL = 12    // 线下环境，agent轮询接口时，如果version与db相同，则返回该错误
	ERR_REGISTRY_BASE  = 10000 //注册时的错误
	ERR_ROUTING_BASE   = 20000 //路由相关的错误
	ERR_GROUP_BASE     = 30000 //group相关的错误
	ERR_ILIB_BASE      = 50000 //inner lib相关错误
	ERR_DUMP_BASE      = 60000 //dump 相关的错误
	ERR_REG_BASE       = 70000 //服务注册相关错误
	ERR_PRETASK_BASE   = 80000 //灰度任务相关的错误
	ERR_ENDPOINT_BASE  = 90000 //Endpoint相关的错误
	ERR_DEPENDENT_BASE = 100000
	ERR_RELEASE_BASE   = 110000
	ERR_RAPID_BASE     = 120000 //极速接入相关的错误
	ERR_OPEN_API       = 130000 // openapi相关错误
	ERR_FAS_ERR        = 140000
)

const (
	ERR_RELEASE_GET_FAIL = ERR_RELEASE_BASE + iota
)

const (
	ERR_DEPENDENT_GET_FAIL = ERR_DEPENDENT_BASE + iota
)

const (
	ERR_REGISTRY_ADD_FAIL = ERR_REGISTRY_BASE + iota
	ERR_REGISTRY_DEL_FAIL
	ERR_REGISTRY_GET_FAIL
	ERR_REGISTRY_LIST_FAIL
	ERR_REGISTRY_UPDATE_FAIL
	ERR_REGISTRY_SET_GROUP_ROUTING_FAIL
)

const (
	ERR_ROUTING_GET_FAIL = ERR_ROUTING_BASE + iota
	ERR_ROUTING_ADD_FAIL
	ERR_ROUTING_DEL_FAIL
	ERR_ROUTING_LIST_FAIL
	ERR_ROUTING_UPDATE_FAIL
)

const (
	ERR_PRE_TASK_LIST_FAIL = ERR_PRETASK_BASE + iota
	ERR_PRE_TASK_NO_EXIST
	ERR_PRE_TASK_DO_FAIL
	ERR_PRE_TASK_FINISH_FAIL
)

const (
	ERR_ENDPOINT_LIST_FAIL = ERR_ENDPOINT_BASE + iota
	ERR_ENDPOINT_ADD_FAIL
	ERR_ENDPOINT_DEL_FAIL
	ERR_ENDPOINT_UPDATE_FAIL
)

const (
	ERR_GROUP_ADD_FAIL = ERR_GROUP_BASE + iota
	ERR_GROUP_DEL_FAIl
	ERR_GROUP_LIST_FAIL
	ERR_GROUP_UPDATE_FAIL
	ERR_GROUP_DEL_ALL_FAIL
	ERR_GROUP_DUPLICATE_FAIL
	ERR_GROUP_EXIST
	ERR_GROUP_NOT_FOUND
	ERR_GROUP_UESED
)

const (
	ERR_ILIB_EXIST = ERR_ILIB_BASE + iota
	ERR_ILIB_NOT_EXIST
	ERR_ILIB_LIST_FAIL
	ERR_ILIB_GET_FAIL
	ERR_ILIB_ADD_FAIL
	ERR_ILIB_DEL_FAIL
	ERR_ILIB_UPDATE_FAIL
	ERR_ILIB_USED
	ERR_ILIB_SERVICE_ADD_FAIL
	ERR_ILIB_SERVICE_DEL_FAIL
	ERR_ILIB_SERVICE_GET_FAIL
)

const (
	ERR_DUMP_SU_NOT_FIND = ERR_DUMP_BASE + iota
	ERR_DUMP_MERGE_DEP_FAIL
	ERR_DUMP_ALIAS_FAIL
	ERR_DUMP_PROVIDER_FAIL
	ERR_DUMP_PROVIDER_CHECK_FAIL
	ERR_DUMP_RAPID_ACCESS_USER_MODIFY
	ERR_DUMP_NOT_ORGANIZATION
	ERR_DUMP_RAPID_ACCESS_ERR
)

const (
	ERR_REG_FAIL = ERR_REG_BASE + iota
	ERR_UNREG_FAIL
	ERR_ENDPOINT_GET_FAIL
)

const (
	ERR_RAPID_PARAM = ERR_RAPID_BASE + iota
	ERR_RAPID_SYS
)

const (
	ERR_RESOURCE_EXISTS = ERR_OPEN_API + iota
	ERR_RESOURCE_NOT_EXISTS
)

var ErrnoDesc = map[int]string{
	ERR_OK:                "ok",
	ERR_PARAM:             "参数错误",
	ERR_SYS:               "系统错误",
	ERR_UNKNOWN:           "未知错误",
	ERR_JSON_DECODE:       "格式解析错误",
	ERR_JSON_ENCODE:       "json序列化错误",
	ERR_TOKEN:             "token验证失败",
	ERR_DB:                "DB错误",
	ERR_HTTP:              "Http任务转发错误",
	ERR_LOCK:              "灰度任务加锁失败",
	ERR_PART_FAILED:       "批量执行部分失败",
	AUTH_INVALID:          "权限异常",
	ERR_VERSION_IDENTICAL: "版本与中心一致，无需重新dump",
	//registry的错误
	ERR_REGISTRY_ADD_FAIL:               "服务注册失败",
	ERR_REGISTRY_DEL_FAIL:               "服务删除失败",
	ERR_REGISTRY_GET_FAIL:               "服务信息获取失败",
	ERR_REGISTRY_LIST_FAIL:              "服务列表获取失败",
	ERR_REGISTRY_UPDATE_FAIL:            "服务更新失败",
	ERR_REGISTRY_SET_GROUP_ROUTING_FAIL: "设置组路由失败",
	ERR_ROUTING_GET_FAIL:                "路由获取失败",
	ERR_ROUTING_ADD_FAIL:                "路由添加失败",
	ERR_ROUTING_DEL_FAIL:                "路由删除失败",
	ERR_ROUTING_LIST_FAIL:               "获取路由列表失败",
	ERR_ROUTING_UPDATE_FAIL:             "路由更新失败",
	ERR_GROUP_ADD_FAIL:                  "添加服务到组失败",
	ERR_GROUP_DEL_FAIl:                  "删除组中的服务失败",
	ERR_GROUP_LIST_FAIL:                 "获取组列表失败",
	ERR_GROUP_UPDATE_FAIL:               "更新组失败",
	ERR_GROUP_DEL_ALL_FAIL:              "删除组失败",
	ERR_GROUP_DUPLICATE_FAIL:            "组复制失败",
	ERR_GROUP_EXIST:                     "组已经存在",
	ERR_GROUP_NOT_FOUND:                 "没有找到对应的group",
	//inner lib
	ERR_ILIB_EXIST:            "LIB已经存在",
	ERR_ILIB_NOT_EXIST:        "LIB不存在",
	ERR_ILIB_LIST_FAIL:        "获取列表失败",
	ERR_ILIB_GET_FAIL:         "获取LIB信息失败",
	ERR_ILIB_ADD_FAIL:         "添加LIB失败",
	ERR_ILIB_DEL_FAIL:         "删除LIB失败",
	ERR_ILIB_UPDATE_FAIL:      "更新LIB失败",
	ERR_ILIB_SERVICE_ADD_FAIL: "添加LIB成员失败",
	ERR_ILIB_SERVICE_DEL_FAIL: "删除LIB成员失败",
	ERR_ILIB_SERVICE_GET_FAIL: "获取LIB成员失败",

	ERR_DUMP_SU_NOT_FIND:         "指定的DISF CUUID不存在",
	ERR_DUMP_MERGE_DEP_FAIL:      "合并libs的依赖失败",
	ERR_DUMP_ALIAS_FAIL:          "dump别名失败",
	ERR_DUMP_PROVIDER_FAIL:       "dump服务依赖失败",
	ERR_DUMP_PROVIDER_CHECK_FAIL: "dump服务依赖校验失败",
	ERR_REG_FAIL:                 "服务节点注册失败",
	ERR_UNREG_FAIL:               "服务节点移除失败",
	ERR_ENDPOINT_GET_FAIL:        "获取服务节点信息失败",
	ERR_GROUP_UESED:              "服务组被使用",
	ERR_ILIB_USED:                "lib被使用",
	ERR_PRE_TASK_LIST_FAIL:       "灰度任务查询失败",
	ERR_PRE_TASK_NO_EXIST:        "灰度任务不存在",
	ERR_PRE_TASK_DO_FAIL:         "灰度任务执行失败",
	ERR_PRE_TASK_FINISH_FAIL:     "全量执行失败",
	ERR_ENDPOINT_LIST_FAIL:       "Endpoint查询失败",
	ERR_ENDPOINT_ADD_FAIL:        "Endpoint增加失败",
	ERR_ENDPOINT_DEL_FAIL:        "Endpoint删除失败",
	ERR_ENDPOINT_UPDATE_FAIL:     "Endpoint更新失败",

	ERR_DEPENDENT_GET_FAIL: "获取depedent失败",
	ERR_RELEASE_GET_FAIL:   "获取release信息失败",

	ERR_RESOURCE_EXISTS:     "资源已存在",
	ERR_RESOURCE_NOT_EXISTS: "资源不存在",
}

var (
	ErrClusterNotFound          = errors.New("cluster not found")
	ErrClusterAlreadyExisted    = errors.New("cluster already existed")
	ErrClusterBatchGetNotEnough = errors.New("cluster batch get not enough")
	ErrGroupBatchGetNotEnough   = errors.New("group batch get not enough")
	ErrServiceBatchGetNotEnough = errors.New("service batch get not enough")
	ErrServiceNotFound          = errors.New("service not found")
	ErrInvalidParam             = errors.New("invalid param")
	ErrVersionIdentical         = errors.New("version identical")
	ErrDbWrongAffectedCount     = errors.New("db wrong affected count")
)
