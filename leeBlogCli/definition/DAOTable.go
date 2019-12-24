package definition

// MySQL数据结构体：作者
type Author struct {
	ID       int    `db:"author_id"`
	Nickname string `db:"author_nickname"`
	Email    string `db:"author_email"`
	Password string `db:"author_password"`
	Motto    string `db:"author_motto"`
	IsActive byte   `db:"author_is_active"`
}

// MySQL数据结构体：文章
type Article struct {
	ID         int    `db:"article_id"`
	CategoryID int    `db:"article_ctg"`
	Title      string `db:"article_title"`
	Tags       string `db:"article_tags"`
	TagsID     string `db:"article_tags_id"`
	Summary    string `db:"article_summary"`
	Content    string `db:"article_content"`
	CreateTime int    `db:"article_createtime"`
	UpdateTime int    `db:"article_updatetime"`
}

// MySQL数据结构体：文章分类（大分类，如：技术、生活等）
type Category struct {
	ID   int    `db:"ctg_id"`
	Name string `db:"ctg_name"`
}

// MySQL数据结构体：文章标签
type Tag struct {
	// 标签ID
	ID int32 `json:"id" db:"tag_id"`

	// 标签名称
	Name string `json:"name" db:"tag_name"`

	// 标签所属分类ID
	CategoryID int32 `json:"categoryID" db:"tag_category"`
}
