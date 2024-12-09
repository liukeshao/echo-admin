package types

type OrgCreateInput struct {
	// 组织名称
	Name string `json:"name,omitempty"`
	// 显示名称
	DisplayName string `json:"display_name,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// logo
	Logo string `json:"logo,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
}
