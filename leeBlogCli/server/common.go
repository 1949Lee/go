package server

// 生成新的文章ID。目前是找到数据库中最大的id并加1。分布式之后可以改成snowflake算法生成ID
func (b *Blog) GenerateArticleID() (ID int32) {
	ID = b.Dao.NextArticleID()
	return ID
}
