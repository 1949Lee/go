SELECT MAX(article_id) + 1 as ID FROM article;
SELECT * FROM tag where tag_category=1;SELECT * FROM tag where tag_category=2;
INSERT INTO tag (tag_name,tag_category) VALUES ('AI',1);
DELETE FROM category WHERE ctg_id=3;DELETE FROM tag WHERE tag_category IS NULL;
SELECT
	( @rownum := @rownum + 1 ) AS no,
	a.*,
	c.ctg_name,
	t.*
FROM
	article a
	LEFT JOIN category c ON a.article_ctg = c.ctg_id
	LEFT JOIN articles_tags_relation r ON r.relation_article = a.article_id
	LEFT JOIN tag t ON r.relation_tag = t.tag_id,
	(SELECT @rownum := 0) row
GROUP BY
	a.article_id
ORDER BY
	a.article_updatetime DESC;
