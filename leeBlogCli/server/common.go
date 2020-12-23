package server

import "leeBlogCli/definition"

// 生成新的文章ID。目前是找到数据库中最大的id并加1。分布式之后可以改成snowflake算法生成ID
func (b *Blog) GenerateArticleID() (ID int32) {
	ID = b.Dao.NextArticleID()
	return ID
}

// 获取分类和分类下的所有标签
func (b *Blog) GetTagsGroupByCategory() (c []definition.CategoryWithTags) {
	var err error
	c, err = b.Dao.SelectTagsGroupByCategory()
	if err == nil {
		return c
	} else {
		return make([]definition.CategoryWithTags, 0)
	}
}

func (b *Blog) NewTag(param definition.Tag) (tag definition.Tag) {
	tag.ID = -1
	ID := b.Dao.InsertTag(param)
	if ID == -1 {
		return tag
	}
	tag.ID = ID
	tag.Name = param.Name
	tag.CategoryID = param.CategoryID
	return tag
}

func (b *Blog) DeleteTag(param definition.Tag) bool {
	ok := b.Dao.DeleteTag(param)
	if !ok {
		return false
	}
	return true
}

func (b *Blog) NewCategory(param definition.Category) (ctg definition.Category) {
	ctg.ID = -1
	ID := b.Dao.InsertCategory(param)
	if ID == -1 {
		return ctg
	}
	ctg.ID = ID
	ctg.Name = param.Name
	return ctg
}

func (b *Blog) DeleteCategory(param definition.Category) bool {
	ok := b.Dao.DeleteCategory(param)
	if !ok {
		return false
	}
	return true
}
