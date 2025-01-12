package blog

import "fmt"

// 这个类型由哪些值
const (
	// 使用0表示默认值, 默认就是草稿状态,
	STATUS_DRAFT STATUS = iota
	// 已发布状态
	STATUS_PUBLISHED
)

// 自定义类型
type STATUS int

// Marshaler is the interface implemented by types that
// // can marshal themselves into valid JSON.
//
//	type Marshaler interface {
//		MarshalJSON() ([]byte, error)
//	}
//
// 你自己定义当前类型的JSON输出, 一定要是一个合法的JSON
// "status": "xxx", "xxx"
func (s STATUS) MarshalJSON() ([]byte, error) {
	switch s {
	case STATUS_DRAFT:
		// 草稿   "草稿"
		return []byte(`"草稿"`), nil
	case STATUS_PUBLISHED:
		return []byte(`"已发布"`), nil
	}

	return []byte(fmt.Sprintf("%d", s)), nil
}

// json.Unmarshaler
// 完成 STATUS类型的自定义反序列化
func (s *STATUS) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case "草稿":
		*s = STATUS_DRAFT
		// 草稿   "草稿"
		return nil
	case "已发布":
		*s = STATUS_PUBLISHED
		return nil
	}

	return nil
}
