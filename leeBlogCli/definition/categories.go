package definition

//GetTagsWithArticleID接口返回结果
type CategoryWithArticleIDListResult struct {
	// 分类列表
	List []CategoryWithTagsAndArticleID `json:"list"`
}

// 文章分类，但其标签，附带含标签本身的文章的ID数组
type CategoryWithTagsAndArticleID struct {
	// 分类ID
	ID int32 `json:"id" db:"ctg_id"`

	// 分类名称
	Name string `json:"name" db:"ctg_name"`

	// 分类下标签
	Tags []TagsWithArticleID `json:"tags"`
}

// 标签，附带含标签本身的文章的ID数组
type TagsWithArticleID struct {
	Tag

	//文章ID数组
	ArticleIDs []int32 `json:"article_ids"`
}
