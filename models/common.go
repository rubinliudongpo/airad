package models

// Predefined model error codes.
const (
	ErrDatabase = -1
	ErrSystem   = -2
	ErrDupRows  = -3
	ErrNotFound = -4
)

type BaseModel struct {
	CodeInfo CodeInfo
}

// CodeInfo definition.
type CodeInfo struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}

// NewErrorInfo return a CodeInfo represents error.
func NewErrorInfo(info string) *CodeInfo {
	return &CodeInfo{-1, info}
}

// NewSuccessInfo return a CodeInfo represents OK.
func NewSuccessInfo(info string) *CodeInfo {
	return &CodeInfo{0, info}
}
