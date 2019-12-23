package definition

// 文章分类
type CategoryWithTags struct {
	// 分类ID
	ID int32 `json:"id" db:"ctg_id"`

	// 分类名称
	Name string `json:"name" db:"ctg_name"`

	// 分类下标签
	Tags []Tag `json:"tags"`
}
