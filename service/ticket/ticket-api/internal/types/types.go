// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt,optional"`
}

// The response data of attribute information | Attribute信息
// swagger:model AttributeInfo
type AttributeInfo struct {
	BaseIDInfo
	// EntityId
	EntityId uint64 `json:"entityId,optional"`
	// AttributeCode
	AttributeCode string `json:"attributeCode,optional"`
	// BackendClass
	BackendClass string `json:"backendClass,optional"`
	// BackendType
	BackendType string `json:"backendType,optional"`
	// BackendTable
	BackendTable string `json:"backendTable,optional"`
	// FrontendClass
	FrontendClass string `json:"frontendClass,optional"`
	// FrontendType
	FrontendType string `json:"frontendType,optional"`
	// FrontendLabel
	FrontendLabel string `json:"frontendLabel,optional"`
	// SourceClass
	SourceClass string `json:"sourceClass,optional"`
	// DefaultValue
	DefaultValue string `json:"defaultValue,optional"`
	// IsFilterable
	IsFilterable uint32 `json:"isFilterable,optional"`
	// IsSearchable
	IsSearchable uint32 `json:"isSearchable,optional"`
	// IsRequired
	IsRequired uint32 `json:"isRequired,optional"`
	// RequiredValidateClass
	RequiredValidateClass string `json:"requiredValidateClass,optional"`
}

// The response data of attribute list | Attribute列表数据
// swagger:model AttributeListResp
type AttributeListResp struct {
	BaseDataInfo
	// Attribute list data | Attribute列表数据
	Data AttributeListInfo `json:"data"`
}

// Attribute list data | Attribute列表数据
// swagger:model AttributeListInfo
type AttributeListInfo struct {
	BaseListInfo
	// The API list data | Attribute列表数据
	Data []AttributeInfo `json:"data"`
}

// Get attribute list request params | Attribute列表请求参数
// swagger:model AttributeListReq
type AttributeListReq struct {
	PageInfo
	// AttributeCode
	AttributeCode string `json:"attributeCode,optional"`
	// BackendClass
	BackendClass string `json:"backendClass,optional"`
	// BackendType
	BackendType string `json:"backendType,optional"`
}

// Attribute information response | Attribute信息返回体
// swagger:model AttributeInfoResp
type AttributeInfoResp struct {
	BaseDataInfo
	// Attribute information | Attribute数据
	Data AttributeInfo `json:"data"`
}
