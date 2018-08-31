/*
 Navicat Premium Data Transfer

 Source Server         : AliyunLee
 Source Server Type    : MySQL
 Source Server Version : 50638
 Source Host           : 39.106.17.117:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50638
 File Encoding         : 65001

 Date: 31/08/2018 10:45:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `article_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id，自增',
  `article_title` varchar(255) NOT NULL COMMENT '文章标题',
  `article_author` int(1) unsigned DEFAULT NULL COMMENT '文章作者ID，外键',
  `article_tags` varchar(255) DEFAULT NULL COMMENT '文章标签名字，逗号隔开',
  `article_tags_id` varchar(255) DEFAULT NULL COMMENT '文章标签id，逗号隔开',
  `article_summary` varchar(255) DEFAULT NULL COMMENT '文章摘要',
  `article_content` text COMMENT '文章正文',
  `article_createtime` timestamp NOT NULL COMMENT '文章创建时间',
  `article_updatetime` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '文章修改时间',
  PRIMARY KEY (`article_id`),
  KEY `fk_article_author` (`article_author`),
  CONSTRAINT `fk_article_author` FOREIGN KEY (`article_author`) REFERENCES `author` (`author_id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表：1、标签会存一组标签名字和标签id，都分别用逗号隔开 2、文章摘要用于显示文章简介';

-- ----------------------------
-- Table structure for author
-- ----------------------------
DROP TABLE IF EXISTS `author`;
CREATE TABLE `author` (
  `author_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `author_nickname` varchar(7) NOT NULL COMMENT '作者昵称，唯一',
  `author_email` varchar(25) NOT NULL COMMENT '作者的邮箱，唯一标识',
  `author_password` varchar(15) NOT NULL COMMENT '密码',
  `author_motto` varchar(25) DEFAULT NULL COMMENT '作者座右铭',
  `author_is_active` bit(1) NOT NULL DEFAULT b'1' COMMENT '账户是否启用：1启用，0不启用',
  PRIMARY KEY (`author_id`),
  UNIQUE KEY `index_id_nn_email` (`author_id`,`author_nickname`(3),`author_email`(12)) COMMENT '作者id、昵称、邮箱'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='作者表：方便以后扩展（1、作者邮箱：author_email；2、作者昵称：author_nickname；3、作者密码：author_password；4、作者座右铭：author_motto）';

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `ctg_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '类别id，自增',
  `ctg_name` varchar(8) NOT NULL COMMENT '类别名称，限制长度',
  PRIMARY KEY (`ctg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='类别表：用于文章大方向的分类（技术、生活等）';

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` int(1) NOT NULL AUTO_INCREMENT COMMENT '标签id，自增',
  `tag_category` int(1) unsigned DEFAULT '1' COMMENT '外键，标签所属类型的id',
  `tag_name` varchar(30) NOT NULL COMMENT '标签名字',
  PRIMARY KEY (`tag_id`),
  KEY `index_ctg_name` (`tag_category`,`tag_name`),
  CONSTRAINT `fk_tag_ctg` FOREIGN KEY (`tag_category`) REFERENCES `category` (`ctg_id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签表：博客的所有标签（1、标签名称：tag_name；2、标签创建时间：tag_createtime;3、标签所属类别：tag_category）';

-- ----------------------------
-- Triggers structure for table author
-- ----------------------------
DROP TRIGGER IF EXISTS `author_no_delete`;
delimiter ;;
CREATE TRIGGER `author_no_delete` BEFORE DELETE ON `author` FOR EACH ROW BEGIN
DECLARE	msg VARCHAR (255);
SET msg = "表中的记录无法删除";
SIGNAL SQLSTATE 'HY000' SET mysql_errno = 22, message_text = msg;-- HY000为系统内部错误号，22为自定义的显示错误号，msg为错误文本
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table category
-- ----------------------------
DROP TRIGGER IF EXISTS `ctg_no_delete`;
delimiter ;;
CREATE TRIGGER `ctg_no_delete` BEFORE DELETE ON `category` FOR EACH ROW BEGIN
DECLARE	msg VARCHAR (255);
SET msg = "表中的记录无法删除";
SIGNAL SQLSTATE 'HY000' SET mysql_errno = 22, message_text = msg;-- HY000为系统内部错误号，22为自定义的显示错误号，msg为错误文本
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table tag
-- ----------------------------
DROP TRIGGER IF EXISTS `tag_no_delete`;
delimiter ;;
CREATE TRIGGER `tag_no_delete` BEFORE DELETE ON `tag` FOR EACH ROW BEGIN
DECLARE	msg VARCHAR (255);
SET msg = "表中的记录无法删除";
SIGNAL SQLSTATE 'HY000' SET mysql_errno = 22, message_text = msg;-- HY000为系统内部错误号，22为自定义的显示错误号，msg为错误文本
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
