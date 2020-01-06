package definition

// MySQL数据结构体：作者
type Author struct {
	ID       int32  `db:"author_id"`
	Nickname string `db:"author_nickname"`
	Email    string `db:"author_email"`
	Password string `db:"author_password"`
	Motto    string `db:"author_motto"`
	IsActive byte   `db:"author_is_active"`
}

// MySQL数据结构体：文章
type Article struct {
	ID         int32  `json:"id" db:"article_id"`
	CategoryID int32  `json:"categoryID" db:"article_ctg"`
	Title      string `json:"title" db:"article_title"`
	AuthorID   int32  `json:"authorID" db:"article_author"`
	Summary    string `json:"summary" db:"article_summary"`
	Content    string `json:"content" db:"article_content"`
	CreateTime string `json:"createTime" db:"article_createtime"`
	UpdateTime string `json:"updateTime" db:"article_updatetime"`
}

// MySQL数据结构体：文章分类（大分类，如：技术、生活等）
type Category struct {
	ID   int32  `json:"id" db:"ctg_id"`
	Name string `json:"name" db:"ctg_name"`
}

// MySQL数据结构体：文章标签
type Tag struct {
	// 标签ID
	ID int32 `json:"id" db:"tag_id"`

	// 标签所属分类ID
	CategoryID int32 `json:"categoryID" db:"tag_category"`

	// 标签名称
	Name string `json:"name" db:"tag_name"`
}
