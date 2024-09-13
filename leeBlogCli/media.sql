/*
 Navicat Premium Data Transfer

 Source Server         : AliyunMySQL
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : rm-2zet44i534j5xmc669o.mysql.rds.aliyuncs.com:3306
 Source Schema         : media

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 13/09/2024 09:18:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `ctg_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '类别id，自增',
  `ctg_name` varchar(8) NOT NULL COMMENT '类别名称，限制长度',
  PRIMARY KEY (`ctg_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='类别表：用于文章大方向的分类（技术、生活等）';

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, 'AV');
COMMIT;

-- ----------------------------
-- Table structure for media
-- ----------------------------
DROP TABLE IF EXISTS `media`;
CREATE TABLE `media` (
  `media_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频id，自增',
  `media_ctg` int(1) unsigned DEFAULT NULL COMMENT '视频分类',
  `media_file_name` varchar(255) NOT NULL COMMENT '视频文件名称',
  `media_file_path` varchar(255) DEFAULT NULL COMMENT '视频文件路径',
  `media_code` varchar(255) DEFAULT NULL COMMENT '视频番号',
  `media_file_ext` varchar(255) DEFAULT NULL COMMENT '视频文件格式',
  `media_createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `media_updatetime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  PRIMARY KEY (`media_id`) USING BTREE,
  KEY `fk_article_ctg` (`media_ctg`),
  KEY `article_updatetime` (`media_updatetime`),
  KEY `article_title` (`media_file_name`(191)),
  FULLTEXT KEY `ft_index` (`media_file_name`,`media_file_path`) /*!50100 WITH PARSER `ngram` */ ,
  CONSTRAINT `fk_article_ctg` FOREIGN KEY (`media_ctg`) REFERENCES `category` (`ctg_id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=167 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of media
-- ----------------------------
BEGIN;
INSERT INTO `media` VALUES (2, 1, 'clot-010', 'H:\\AV\\其他\\肉丝\\clot-010.mp4', '', 'mp4', '2021-09-14 20:56:11', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (3, 1, 'ATID-397', 'H:\\AV\\其他\\肉丝\\ATID-397.mp4', '', 'mp4', '2021-09-14 22:44:41', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (4, 1, 'GENM-033', 'H:\\AV\\其他\\肉丝\\GENM-033.mp4', '', 'mp4', '2021-09-14 22:45:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (5, 1, 'DPMI-045', 'H:\\AV\\其他\\肉丝\\DPMI-045.mp4', '', 'mp4', '2021-09-14 22:56:31', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (6, 1, 'ELO-289', 'H:\\AV\\其他\\肉丝\\ELO-289.mp4', '', 'mp4', '2021-09-20 01:53:23', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (7, 1, 'EMBZ-199 ', 'H:\\AV\\其他\\肉丝\\EMBZ-199 .mp4', '', 'mp4', '2021-09-20 01:54:23', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (8, 1, 'GOJU-119', 'H:\\AV\\其他\\肉丝\\GOJU-119.mp4', '', 'mp4', '2021-09-20 01:56:19', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (9, 1, 'jufe-041', 'H:\\AV\\其他\\肉丝\\jufe-041.mp4', '', 'mp4', '2021-09-20 01:58:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (10, 1, 'ORE-344', 'H:\\AV\\其他\\肉丝\\ORE-344.mp4', '', 'mp4', '2021-09-20 02:08:16', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (11, 1, 'sw-473', 'H:\\AV\\其他\\肉丝\\sw-473.mp4', '', 'mp4', '2021-09-20 02:11:10', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (12, 1, 'sw-506', 'H:\\AV\\其他\\肉丝\\sw-506.mp4', '', 'mp4', '2021-09-20 02:12:47', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (13, 1, 'sw-562', 'H:\\AV\\其他\\肉丝\\sw-562.mp4', '', 'mp4', '2021-09-20 02:13:38', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (14, 1, 'sw-656', 'H:\\AV\\其他\\肉丝\\sw-656.mp4', '', 'mp4', '2021-09-20 02:14:35', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (15, 1, 'taak-006', 'H:\\AV\\其他\\肉丝\\taak-006.mp4', '', 'mp4', '2021-09-20 02:14:54', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (16, 1, 'taak-007', 'H:\\AV\\其他\\肉丝\\taak-007.mp4', '', 'mp4', '2021-09-20 02:15:05', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (17, 1, 'taak-010', 'H:\\AV\\其他\\肉丝\\taak-010.mp4', '', 'mp4', '2021-09-20 02:15:14', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (18, 1, 'taak-017', 'H:\\AV\\其他\\肉丝\\taak-017.mp4', '', 'mp4', '2021-09-20 02:16:14', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (19, 1, 'taak-023', 'H:\\AV\\其他\\肉丝\\taak-023.mp4', '', 'mp4', '2021-09-20 02:17:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (20, 1, 'TAAK-024', 'H:\\AV\\其他\\肉丝\\TAAK-024.mp4', '', 'mp4', '2021-09-20 02:17:45', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (21, 1, 'URPS-019', 'H:\\AV\\其他\\肉丝\\URPS-019.wmv', '', 'wmv', '2021-09-20 02:19:14', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (22, 1, 'ZMEN-004', 'H:\\AV\\其他\\肉丝\\ZMEN-004.mp4', '', 'mp4', '2021-09-20 02:19:53', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (23, 1, 'zmen-051', 'H:\\AV\\其他\\肉丝\\zmen-051.mp4', '', 'mp4', '2021-09-20 02:20:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (24, 1, 'zmen-057', 'H:\\AV\\其他\\肉丝\\zmen-057.mp4', '', 'mp4', '2021-09-20 02:20:41', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (25, 1, 'ZMEN-065', 'H:\\AV\\其他\\肉丝\\ZMEN-065.mp4', '', 'mp4', '2021-09-20 02:21:27', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (26, 1, 'ZMEN-022', 'H:\\AV\\其他\\肉丝\\ZMEN-022.mp4', '', 'mp4', '2021-09-20 02:21:51', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (27, 1, 'zmen-029', 'H:\\AV\\其他\\肉丝\\zmen-029.mp4', '', 'mp4', '2021-09-20 02:22:35', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (28, 1, 'zmen-024', 'H:\\AV\\其他\\肉丝\\zmen-024.mp4', '', 'mp4', '2021-09-20 02:22:57', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (29, 1, 'ZMEN-076', 'H:\\AV\\其他\\肉丝\\ZMEN-076.mp4', '', 'mp4', '2021-09-20 02:23:21', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (30, 1, 'ZMEN-079', 'H:\\AV\\其他\\肉丝\\ZMEN-079.mp4', '', 'mp4', '2021-09-20 02:23:36', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (31, 1, 'zmen-005', 'H:\\AV\\其他\\肉丝\\zmen-005.mp4', '', 'mp4', '2021-09-20 02:23:59', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (32, 1, 'DFET-015A', 'H:\\AV\\其他\\肉丝\\DFET-015A.mp4', '', 'mp4', '2021-09-20 02:24:38', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (33, 1, 'DFET-015B', 'H:\\AV\\其他\\肉丝\\DFET-015B.mp4', '', 'mp4', '2021-09-20 02:24:41', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (34, 1, 'elo-319-1', 'H:\\AV\\其他\\肉丝\\elo-319-1.wmv', '', 'wmv', '2021-09-20 02:25:18', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (35, 1, 'ETQR-082', 'H:\\AV\\其他\\肉丝\\ETQR-082.mp4', '', 'mp4', '2021-09-20 02:25:36', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (36, 1, 'GCF-010', 'H:\\AV\\其他\\肉丝\\GCF-010.mp4', '', 'mp4', '2021-09-20 02:25:59', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (37, 1, 'HZMEN-003-1', 'H:\\AV\\其他\\肉丝\\HZMEN-003-1.mp4', '', 'mp4', '2021-09-20 02:26:18', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (38, 1, 'hzmen-008', 'H:\\AV\\其他\\肉丝\\hzmen-008.mp4', '', 'mp4', '2021-09-20 02:26:35', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (39, 1, 'JUFD-471', 'H:\\AV\\其他\\肉丝\\JUFD-471.mp4', '', 'mp4', '2021-09-20 02:26:53', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (40, 1, 'SW-445A', 'H:\\AV\\其他\\肉丝\\SW-445A.mp4', '', 'mp4', '2021-09-20 02:29:20', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (41, 1, 'SW-445B', 'H:\\AV\\其他\\肉丝\\SW-445B.mp4', '', 'mp4', '2021-09-20 02:29:22', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (42, 1, 'ZMEN-025', 'H:\\AV\\其他\\肉丝\\ZMEN-025.mp4', '', 'mp4', '2021-09-20 02:29:38', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (43, 1, 'ZMEN-080', 'H:\\AV\\其他\\肉丝\\ZMEN-080.mp4', '', 'mp4', '2021-09-20 02:30:00', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (44, 1, 'IPX-231', 'H:\\AV\\其他\\巨尻\\IPX-231.mp4', '', 'mp4', '2021-09-21 23:45:52', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (45, 1, 'GS-285', 'H:\\AV\\其他\\巨尻\\GS-285.mp4', '', 'mp4', '2021-09-21 23:48:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (46, 1, 'DVDMS379', 'H:\\AV\\其他\\巨尻\\DVDMS379.mp4', '', 'mp4', '2021-09-21 23:49:34', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (47, 1, 'dvdms-0461', 'H:\\AV\\其他\\巨尻\\dvdms-0461.mp4', '', 'mp4', '2021-09-21 23:49:51', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (48, 1, 'DOCP-157', 'H:\\AV\\其他\\巨尻\\DOCP-157.mp4', '', 'mp4', '2021-09-21 23:50:34', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (49, 1, 'GNAB-038', 'H:\\AV\\其他\\巨尻\\GNAB-038.mp4', '', 'mp4', '2021-09-21 23:54:20', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (50, 1, 'LULU-042', 'H:\\AV\\其他\\巨尻\\LULU-042.mp4', '', 'mp4', '2021-09-21 23:56:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (51, 1, 'SNIS-497', 'H:\\AV\\其他\\巨尻\\SNIS-497.mkv', '', 'mkv', '2021-09-21 23:57:26', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (52, 1, 'gnab-021', 'H:\\AV\\其他\\巨尻\\gnab-021.mp4', '', 'mp4', '2021-09-21 23:58:02', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (53, 1, 'SW-316', 'H:\\AV\\其他\\巨尻\\SW-316.mp4', '', 'mp4', '2021-09-21 23:58:19', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (54, 1, 'HUNTA-847', 'H:\\AV\\其他\\巨尻\\HUNTA-847.mp4', '', 'mp4', '2021-09-21 23:59:25', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (55, 1, 'WANZ-976', 'H:\\AV\\其他\\巨尻\\WANZ-976.mp4', '', 'mp4', '2021-09-21 23:59:49', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (56, 1, 'juy-051', 'H:\\AV\\其他\\巨尻\\juy-051.mp4', '', 'mp4', '2021-09-22 00:00:55', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (57, 1, 'GOAL-038', 'H:\\AV\\其他\\巨尻\\GOAL-038.mp4', '', 'mp4', '2021-09-22 00:01:06', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (58, 1, 'GVG-873', 'H:\\AV\\其他\\巨尻\\GVG-873.mp4', '', 'mp4', '2021-09-22 00:01:23', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (59, 1, 'JUL-499', 'H:\\AV\\其他\\巨尻\\JUL-499.mp4', '', 'mp4', '2021-09-22 00:01:52', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (60, 1, 'MIAA-105', 'H:\\AV\\其他\\巨尻\\MIAA-105.mp4', '', 'mp4', '2021-09-22 00:02:41', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (61, 1, 'vrtm-259', 'H:\\AV\\其他\\巨尻\\vrtm-259.dcv', '', 'dcv', '2021-09-22 00:03:26', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (62, 1, 'vrtm-330', 'H:\\AV\\其他\\巨尻\\vrtm-330.dcv', '', 'dcv', '2021-09-22 00:03:32', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (63, 1, 'DVDMS-086', 'H:\\AV\\其他\\巨尻\\DVDMS-086.mp4', '', 'mp4', '2021-09-22 00:04:01', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (64, 1, 'dvdms-207', 'H:\\AV\\其他\\巨尻\\dvdms-207.mp4', '', 'mp4', '2021-09-22 00:04:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (65, 1, 'DVDMS-517', 'H:\\AV\\其他\\巨尻\\DVDMS-517.mp4', '', 'mp4', '2021-09-22 00:05:18', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (66, 1, 'DVDMS-575', 'H:\\AV\\其他\\巨尻\\DVDMS-575.mp4', '', 'mp4', '2021-09-22 00:05:35', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (67, 1, 'DVDMS-623', 'H:\\AV\\其他\\巨尻\\DVDMS-623.mp4', '', 'mp4', '2021-09-22 00:07:22', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (68, 1, 'fset-884', 'H:\\AV\\其他\\巨尻\\fset-884.mp4', '', 'mp4', '2021-09-22 00:07:40', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (69, 1, 'gnab-042', 'H:\\AV\\其他\\巨尻\\gnab-042.mp4', '', 'mp4', '2021-09-22 00:08:08', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (70, 1, 'LULU-071', 'H:\\AV\\其他\\巨尻\\LULU-071.mp4', '', 'mp4', '2021-09-22 00:08:21', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (71, 1, 'NHDTA-611', 'H:\\AV\\其他\\巨尻\\淫荡老婆\\NHDTA-611.mp4', '', 'mp4', '2021-09-22 00:09:31', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (72, 1, 'NHDTB-111', 'H:\\AV\\其他\\巨尻\\淫荡老婆\\NHDTB-111.mp4', '', 'mp4', '2021-09-22 00:09:43', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (73, 1, 'NQI1KELO161', 'H:\\AV\\其他\\巨尻\\淫荡老婆\\NQI1KELO161.mp4', '', 'mp4', '2021-09-22 00:11:03', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (74, 1, 'bf-551', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\bf-551.mp4', '', 'mp4', '2021-09-22 00:11:29', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (75, 1, 'MIAA-114', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\MIAA-114.mp4', '', 'mp4', '2021-09-22 00:12:40', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (76, 1, 'MIAA-078', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\MIAA-078.mp4', '', 'mp4', '2021-09-22 00:13:02', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (77, 1, 'LULU-005', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\LULU-005.mp4', '', 'mp4', '2021-09-22 00:13:23', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (78, 1, 'WANZ869C', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\WANZ869C.mp4', '', 'mp4', '2021-09-22 00:13:49', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (79, 1, 'DASD-747', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\DASD-747.mp4', '', 'mp4', '2021-09-22 00:14:10', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (80, 1, 'juy-536', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\juy-536.mp4', '', 'mp4', '2021-09-22 00:14:46', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (81, 1, 'MIAA-369', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\MIAA-369.mp4', '', 'mp4', '2021-09-22 00:15:17', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (82, 1, 'pred-265', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\pred-265.mp4', '', 'mp4', '2021-09-22 00:15:45', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (83, 1, 'ymdd-218', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\ymdd-218.mp4', '', 'mp4', '2021-09-22 00:16:12', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (84, 1, 'wanz-636', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\wanz-636.mp4', '', 'mp4', '2021-09-22 00:17:10', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (85, 1, 'lulu-065', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\lulu-065.mp4', '', 'mp4', '2021-09-22 00:17:24', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (86, 1, 'PFES-009', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\PFES-009.mp4', '', 'mp4', '2021-09-22 00:17:39', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (87, 1, 'mmtk-001', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\mmtk-001.mp4', '', 'mp4', '2021-09-22 00:18:03', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (88, 1, 'DVDMS-690', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\DVDMS-690.mp4', '', 'mp4', '2021-09-22 00:18:29', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (89, 1, 'GVH-272', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\GVH-272.mp4', '', 'mp4', '2021-09-22 00:18:59', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (90, 1, 'mimk-093', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\mimk-093.mp4', '', 'mp4', '2021-09-22 00:19:22', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (91, 1, 'MISG-003', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\MISG-003.mp4', '', 'mp4', '2021-09-22 00:19:58', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (92, 1, 'nkkd-216', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\nkkd-216.mp4', '', 'mp4', '2021-09-22 00:20:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (93, 1, 'pred-333', 'H:\\AV\\其他\\巨尻\\篠田ゆう\\pred-333.mp4', '', 'mp4', '2021-09-22 00:20:55', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (94, 1, '美艳单身姐姐的翘臀被我爆肏05180316', 'H:\\AV\\其他\\巨尻\\短裤\\美艳单身姐姐的翘臀被我爆肏05180316.mp4', '', 'mp4', '2021-09-22 00:21:21', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (95, 1, 'DOCP-189', 'H:\\AV\\其他\\巨尻\\短裤\\DOCP-189.mp4', '', 'mp4', '2021-09-22 00:21:31', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (96, 1, 'awd-112', 'H:\\AV\\其他\\巨尻\\awd-112.mp4', '', 'mp4', '2021-09-22 00:21:44', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (97, 1, 'dvdms-174', 'H:\\AV\\其他\\巨尻\\dvdms-174.mp4', '', 'mp4', '2021-09-22 00:21:59', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (98, 1, 'dvdms-226', 'H:\\AV\\其他\\巨尻\\dvdms-226.mp4', '', 'mp4', '2021-09-22 00:22:11', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (99, 1, 'dvdms-251', 'H:\\AV\\其他\\巨尻\\dvdms-251.mp4', '', 'mp4', '2021-09-22 00:22:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (100, 1, 'dvdms-273', 'H:\\AV\\其他\\巨尻\\dvdms-273.mp4', '', 'mp4', '2021-09-22 00:22:39', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (101, 1, 'dvdms-351', 'H:\\AV\\其他\\巨尻\\dvdms-351.mp4', '', 'mp4', '2021-09-22 00:22:52', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (102, 1, 'dvdms-368', 'H:\\AV\\其他\\巨尻\\dvdms-368.mp4', '', 'mp4', '2021-09-22 00:23:02', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (103, 1, 'DVDMS-660', 'H:\\AV\\其他\\巨尻\\DVDMS-660.mp4', '', 'mp4', '2021-09-22 00:23:10', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (104, 1, 'GVH-276', 'H:\\AV\\其他\\巨尻\\GVH-276.mp4', '', 'mp4', '2021-09-22 00:23:24', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (105, 1, 'HUNTA-813', 'H:\\AV\\其他\\巨尻\\HUNTA-813.mp4', '', 'mp4', '2021-09-22 00:23:33', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (106, 1, 'HUNTB-011', 'H:\\AV\\其他\\巨尻\\HUNTB-011.mp4', '', 'mp4', '2021-09-22 00:24:05', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (107, 1, 'miaa-084', 'H:\\AV\\其他\\巨尻\\miaa-084.mp4', '', 'mp4', '2021-09-22 00:24:17', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (108, 1, 'misg-001', 'H:\\AV\\其他\\巨尻\\misg-001.mp4', '', 'mp4', '2021-09-22 00:24:38', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (109, 1, 'ORE-412 Miu', 'H:\\AV\\其他\\巨尻\\ORE-412 Miu.mp4', '', 'mp4', '2021-09-22 00:24:57', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (110, 1, 'ROYD-031', 'H:\\AV\\其他\\巨尻\\ROYD-031.mp4', '', 'mp4', '2021-09-22 00:25:21', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (111, 1, 'royd-064', 'H:\\AV\\其他\\巨尻\\royd-064.mp4', '', 'mp4', '2021-09-22 00:26:08', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (112, 1, 'cesd-821', 'H:\\AV\\其他\\飯岡かなこ\\cesd-821.mp4', '', 'mp4', '2021-09-25 11:40:15', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (113, 1, 'ovg-146', 'H:\\AV\\其他\\飯岡かなこ\\ovg-146.mp4', '', 'mp4', '2021-09-25 11:40:55', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (114, 1, 'SPRD-1178', 'H:\\AV\\其他\\飯岡かなこ\\SPRD-1178.mp4', '', 'mp4', '2021-09-25 11:41:40', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (115, 1, 'venu-831', 'H:\\AV\\其他\\飯岡かなこ\\venu-831.mp4', '', 'mp4', '2021-09-25 11:42:03', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (116, 1, 'VENU-964', 'H:\\AV\\其他\\飯岡かなこ\\VENU-964.mp4', '', 'mp4', '2021-09-25 11:42:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (117, 1, 'vrtm-469', 'H:\\AV\\其他\\飯岡かなこ\\vrtm-469.mp4', '', 'mp4', '2021-09-25 11:43:03', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (118, 1, 'SPRD-1445', 'H:\\AV\\其他\\飯岡かなこ\\SPRD-1445.mp4', '', 'mp4', '2021-09-25 11:43:51', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (119, 1, 'venx-059', 'H:\\AV\\其他\\飯岡かなこ\\venx-059.mp4', '', 'mp4', '2021-09-25 11:44:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (120, 1, 'juy-329', 'H:\\AV\\其他\\水川かえで\\juy-329.mp4', '', 'mp4', '2021-09-25 11:49:36', '2021-09-25 12:23:46');
INSERT INTO `media` VALUES (121, 1, 'juy-383', 'H:\\AV\\其他\\老婆淫叫\\juy-383.mp4', '', 'mp4', '2021-09-25 11:50:14', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (122, 1, 'SORA012C', 'H:\\AV\\其他\\老婆淫叫\\SORA012C.mp4', '', 'mp4', '2021-09-25 11:50:54', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (123, 1, 'HTMS-121', 'H:\\AV\\其他\\老婆淫叫\\HTMS-121.mkv', '', 'mkv', '2021-09-25 11:51:26', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (124, 1, 'ZMEN-074', 'H:\\AV\\其他\\老婆淫叫\\ZMEN-074.mp4', '', 'mp4', '2021-09-25 11:51:53', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (125, 1, 'FSDSS-092', 'H:\\AV\\其他\\老婆淫叫\\FSDSS-092.mp4', '', 'mp4', '2021-09-25 11:52:25', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (126, 1, 'HUNTA-920', 'H:\\AV\\其他\\老婆淫叫\\HUNTA-920.mp4', '', 'mp4', '2021-09-25 11:52:52', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (127, 1, 'VOSS-160', 'H:\\AV\\其他\\老婆淫叫\\VOSS-160.mp4', '', 'mp4', '2021-09-25 11:54:35', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (128, 1, 'AKDL-086', 'H:\\AV\\其他\\老婆淫叫\\AKDL-086.mp4', '', 'mp4', '2021-09-25 11:55:12', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (129, 1, 'DOCP-279', 'H:\\AV\\其他\\老婆淫叫\\DOCP-279.mp4', '', 'mp4', '2021-09-25 11:55:40', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (130, 1, 'ATFB-247', 'H:\\AV\\其他\\老婆淫叫\\水菜丽\\ATFB-247.mkv', '', 'mkv', '2021-09-25 11:58:03', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (131, 1, 'GAOR-088', 'H:\\AV\\其他\\老婆淫叫\\老婆\\GAOR-088.mp4', '', 'mp4', '2021-09-25 11:59:54', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (132, 1, 'GAOR-095', 'H:\\AV\\其他\\老婆淫叫\\老婆\\GAOR-095.mp4', '', 'mp4', '2021-09-25 12:00:09', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (133, 1, 'PGD-784.1080p', 'H:\\AV\\其他\\老婆淫叫\\老婆\\PGD-784.1080p.mkv', '', 'mkv', '2021-09-25 12:00:40', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (134, 1, 'ssni-148', 'H:\\AV\\其他\\老婆淫叫\\老婆\\ssni-148.mp4', '', 'mp4', '2021-09-25 12:00:59', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (135, 1, 'emot-007', 'H:\\AV\\其他\\老婆淫叫\\老婆\\emot-007.mp4', '', 'mp4', '2021-09-25 12:01:19', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (136, 1, 'GASO-0014', 'H:\\AV\\其他\\老婆淫叫\\老婆\\GASO-0014.mp4', '', 'mp4', '2021-09-25 12:01:27', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (137, 1, 'ABP-994', 'H:\\AV\\其他\\老婆淫叫\\老婆\\ABP-994.mp4', '', 'mp4', '2021-09-25 12:01:43', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (138, 1, 'emot-009', 'H:\\AV\\其他\\老婆淫叫\\老婆\\emot-009.mp4', '', 'mp4', '2021-09-25 12:01:55', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (139, 1, 'GASO-0072', 'H:\\AV\\其他\\老婆淫叫\\老婆\\GASO-0072.mp4', '', 'mp4', '2021-09-25 12:02:17', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (140, 1, 'HQK0LEF065', 'H:\\AV\\其他\\老婆淫叫\\大声淫叫\\HQK0LEF065.mp4', '', 'mp4', '2021-09-25 12:04:00', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (141, 1, 'mtes008c', 'H:\\AV\\其他\\老婆淫叫\\大声淫叫\\mtes008c.mp4', '', 'mp4', '2021-09-25 12:06:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (142, 1, '成熟老婆被我爆肏呻吟10142303 ', 'H:\\AV\\其他\\老婆淫叫\\大声淫叫\\成熟老婆被我爆肏呻吟10142303 .mp4', '', 'mp4', '2021-09-25 12:06:52', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (143, 1, 'GVG-884', 'H:\\AV\\其他\\老婆淫叫\\大声淫叫\\GVG-884.mp4', '', 'mp4', '2021-09-25 12:10:51', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (144, 1, 'SQIS010', 'H:\\AV\\其他\\老婆淫叫\\大声淫叫\\SQIS010.mp4', '', 'mp4', '2021-09-25 12:10:54', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (145, 1, 'HUNTA-828', 'H:\\AV\\其他\\老婆淫叫\\娇嗔女儿\\HUNTA-828.mp4', '', 'mp4', '2021-09-25 12:12:08', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (146, 1, '美艳老婆被我爆肏12182342', 'H:\\AV\\其他\\老婆淫叫\\妈妈淫叫\\美艳老婆被我爆肏12182342.mkv', '', 'mkv', '2021-09-25 12:12:43', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (147, 1, 'vec-438', 'H:\\AV\\其他\\紗々原ゆり\\vec-438.mp4', '', 'mp4', '2021-09-25 12:14:10', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (148, 1, 'zmen-073', 'H:\\AV\\其他\\紗々原ゆり\\zmen-073.mp4', '', 'mp4', '2021-09-25 12:14:37', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (149, 1, 'IPX-344', 'H:\\AV\\其他\\明里つむぎ\\IPX-344.mp4', '', 'mp4', '2021-09-25 12:15:43', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (150, 1, 'ipx-404', 'H:\\AV\\其他\\明里つむぎ\\ipx-404.mp4', '', 'mp4', '2021-09-25 12:16:06', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (151, 1, 'IPX-530', 'H:\\AV\\其他\\明里つむぎ\\IPX-530.mp4', '', 'mp4', '2021-09-25 12:17:28', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (152, 1, 'IPX-540', 'H:\\AV\\其他\\明里つむぎ\\IPX-540.mp4', '', 'mp4', '2021-09-25 12:18:05', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (153, 1, 'ipz-985', 'H:\\AV\\其他\\明里つむぎ\\ipz-985.wmv', '', 'wmv', '2021-09-25 12:18:19', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (154, 1, 'IPX-586', 'H:\\AV\\其他\\明里つむぎ\\IPX-586.mp4', '', 'mp4', '2021-09-25 12:19:22', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (155, 1, 'SOAV-043', 'H:\\AV\\其他\\水川かえで\\SOAV-043.mp4', '', 'mp4', '2021-09-25 12:24:46', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (156, 1, 'lulu-045', 'H:\\AV\\其他\\藤森里穂\\lulu-045.mp4', '', 'mp4', '2021-09-25 12:25:31', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (157, 1, 'MIAA-470', 'H:\\AV\\其他\\藤森里穂\\MIAA-470.mp4', '', 'mp4', '2021-09-25 12:25:56', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (158, 1, 'LULU-054', 'H:\\AV\\其他\\舞原聖\\LULU-054.mp4', '', 'mp4', '2021-09-25 12:26:20', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (159, 1, 'LULU-071', 'H:\\AV\\其他\\舞原聖\\LULU-071.mp4', '', 'mp4', '2021-09-25 12:26:37', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (160, 1, 'jul-642', 'H:\\AV\\其他\\舞原聖\\肉丝\\jul-642.mp4', '', 'mp4', '2021-09-25 12:26:51', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (161, 1, 'KSBJ-119', 'H:\\AV\\其他\\夏希まろん\\KSBJ-119.mp4', '', 'mp4', '2021-09-25 12:27:26', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (162, 1, 'PPPD-876', 'H:\\AV\\其他\\夏希まろん\\PPPD-876.mp4', '', 'mp4', '2021-09-25 12:27:57', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (163, 1, 'SNIS-153', 'H:\\AV\\其他\\小島みなみ\\SNIS-153.mp4', '', 'mp4', '2021-09-25 12:28:15', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (164, 1, 'ssni-556', 'H:\\AV\\其他\\小島みなみ\\ssni-556.mp4', '', 'mp4', '2021-09-25 12:28:30', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (165, 1, 'ssni-978', 'H:\\AV\\其他\\小島みなみ\\ssni-978.mp4', '', 'mp4', '2021-09-25 12:28:47', '0000-00-00 00:00:00');
INSERT INTO `media` VALUES (166, 1, 'SNIS-518', 'H:\\AV\\其他\\小島みなみ\\SNIS-518.mp4', '', 'mp4', '2021-09-25 12:29:12', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for media_tags_relation
-- ----------------------------
DROP TABLE IF EXISTS `media_tags_relation`;
CREATE TABLE `media_tags_relation` (
  `relation_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '关系id，自增',
  `relation_media` int(1) unsigned DEFAULT NULL COMMENT '文章id，外键',
  `relation_tag` int(1) unsigned DEFAULT NULL COMMENT '标签id，外键',
  PRIMARY KEY (`relation_id`),
  KEY `fk_relation_article` (`relation_media`),
  KEY `fk_relation_tag` (`relation_tag`),
  CONSTRAINT `fk_relation_media` FOREIGN KEY (`relation_media`) REFERENCES `media` (`media_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_relation_tag` FOREIGN KEY (`relation_tag`) REFERENCES `tag` (`tag_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=775 DEFAULT CHARSET=utf8mb4 COMMENT='文章和标签的关系表：用于存储文章和标签的多对多关系';

-- ----------------------------
-- Records of media_tags_relation
-- ----------------------------
BEGIN;
INSERT INTO `media_tags_relation` VALUES (471, 2, 2);
INSERT INTO `media_tags_relation` VALUES (473, 4, 2);
INSERT INTO `media_tags_relation` VALUES (474, 3, 2);
INSERT INTO `media_tags_relation` VALUES (475, 3, 55);
INSERT INTO `media_tags_relation` VALUES (476, 5, 2);
INSERT INTO `media_tags_relation` VALUES (477, 6, 2);
INSERT INTO `media_tags_relation` VALUES (478, 6, 57);
INSERT INTO `media_tags_relation` VALUES (479, 6, 58);
INSERT INTO `media_tags_relation` VALUES (480, 7, 2);
INSERT INTO `media_tags_relation` VALUES (481, 7, 57);
INSERT INTO `media_tags_relation` VALUES (482, 7, 58);
INSERT INTO `media_tags_relation` VALUES (483, 8, 2);
INSERT INTO `media_tags_relation` VALUES (484, 8, 57);
INSERT INTO `media_tags_relation` VALUES (485, 9, 2);
INSERT INTO `media_tags_relation` VALUES (486, 10, 2);
INSERT INTO `media_tags_relation` VALUES (487, 10, 59);
INSERT INTO `media_tags_relation` VALUES (488, 11, 2);
INSERT INTO `media_tags_relation` VALUES (489, 11, 60);
INSERT INTO `media_tags_relation` VALUES (490, 12, 2);
INSERT INTO `media_tags_relation` VALUES (491, 12, 59);
INSERT INTO `media_tags_relation` VALUES (492, 12, 61);
INSERT INTO `media_tags_relation` VALUES (493, 13, 2);
INSERT INTO `media_tags_relation` VALUES (494, 14, 2);
INSERT INTO `media_tags_relation` VALUES (495, 15, 2);
INSERT INTO `media_tags_relation` VALUES (496, 16, 2);
INSERT INTO `media_tags_relation` VALUES (497, 17, 2);
INSERT INTO `media_tags_relation` VALUES (498, 18, 2);
INSERT INTO `media_tags_relation` VALUES (499, 19, 2);
INSERT INTO `media_tags_relation` VALUES (500, 19, 62);
INSERT INTO `media_tags_relation` VALUES (501, 20, 2);
INSERT INTO `media_tags_relation` VALUES (502, 20, 62);
INSERT INTO `media_tags_relation` VALUES (503, 21, 2);
INSERT INTO `media_tags_relation` VALUES (504, 21, 63);
INSERT INTO `media_tags_relation` VALUES (505, 21, 58);
INSERT INTO `media_tags_relation` VALUES (506, 22, 2);
INSERT INTO `media_tags_relation` VALUES (507, 22, 58);
INSERT INTO `media_tags_relation` VALUES (508, 23, 2);
INSERT INTO `media_tags_relation` VALUES (509, 24, 2);
INSERT INTO `media_tags_relation` VALUES (510, 25, 2);
INSERT INTO `media_tags_relation` VALUES (511, 26, 2);
INSERT INTO `media_tags_relation` VALUES (512, 27, 2);
INSERT INTO `media_tags_relation` VALUES (513, 27, 58);
INSERT INTO `media_tags_relation` VALUES (514, 27, 64);
INSERT INTO `media_tags_relation` VALUES (515, 28, 2);
INSERT INTO `media_tags_relation` VALUES (516, 28, 64);
INSERT INTO `media_tags_relation` VALUES (517, 29, 2);
INSERT INTO `media_tags_relation` VALUES (518, 30, 2);
INSERT INTO `media_tags_relation` VALUES (519, 31, 2);
INSERT INTO `media_tags_relation` VALUES (520, 32, 2);
INSERT INTO `media_tags_relation` VALUES (521, 32, 1);
INSERT INTO `media_tags_relation` VALUES (522, 33, 2);
INSERT INTO `media_tags_relation` VALUES (523, 33, 1);
INSERT INTO `media_tags_relation` VALUES (524, 34, 2);
INSERT INTO `media_tags_relation` VALUES (525, 34, 57);
INSERT INTO `media_tags_relation` VALUES (526, 35, 2);
INSERT INTO `media_tags_relation` VALUES (527, 36, 2);
INSERT INTO `media_tags_relation` VALUES (528, 37, 2);
INSERT INTO `media_tags_relation` VALUES (529, 37, 1);
INSERT INTO `media_tags_relation` VALUES (530, 38, 2);
INSERT INTO `media_tags_relation` VALUES (531, 39, 2);
INSERT INTO `media_tags_relation` VALUES (532, 40, 2);
INSERT INTO `media_tags_relation` VALUES (533, 41, 2);
INSERT INTO `media_tags_relation` VALUES (534, 42, 2);
INSERT INTO `media_tags_relation` VALUES (535, 42, 59);
INSERT INTO `media_tags_relation` VALUES (536, 43, 2);
INSERT INTO `media_tags_relation` VALUES (537, 44, 1);
INSERT INTO `media_tags_relation` VALUES (538, 45, 1);
INSERT INTO `media_tags_relation` VALUES (539, 45, 57);
INSERT INTO `media_tags_relation` VALUES (540, 45, 58);
INSERT INTO `media_tags_relation` VALUES (541, 46, 1);
INSERT INTO `media_tags_relation` VALUES (542, 47, 1);
INSERT INTO `media_tags_relation` VALUES (543, 48, 1);
INSERT INTO `media_tags_relation` VALUES (544, 49, 1);
INSERT INTO `media_tags_relation` VALUES (545, 50, 1);
INSERT INTO `media_tags_relation` VALUES (546, 51, 1);
INSERT INTO `media_tags_relation` VALUES (547, 52, 1);
INSERT INTO `media_tags_relation` VALUES (548, 53, 1);
INSERT INTO `media_tags_relation` VALUES (549, 54, 1);
INSERT INTO `media_tags_relation` VALUES (550, 55, 1);
INSERT INTO `media_tags_relation` VALUES (551, 56, 1);
INSERT INTO `media_tags_relation` VALUES (552, 57, 1);
INSERT INTO `media_tags_relation` VALUES (553, 58, 1);
INSERT INTO `media_tags_relation` VALUES (554, 58, 57);
INSERT INTO `media_tags_relation` VALUES (555, 59, 1);
INSERT INTO `media_tags_relation` VALUES (556, 59, 55);
INSERT INTO `media_tags_relation` VALUES (557, 60, 1);
INSERT INTO `media_tags_relation` VALUES (558, 60, 55);
INSERT INTO `media_tags_relation` VALUES (559, 61, 1);
INSERT INTO `media_tags_relation` VALUES (560, 62, 1);
INSERT INTO `media_tags_relation` VALUES (561, 63, 1);
INSERT INTO `media_tags_relation` VALUES (562, 63, 55);
INSERT INTO `media_tags_relation` VALUES (563, 64, 1);
INSERT INTO `media_tags_relation` VALUES (564, 65, 1);
INSERT INTO `media_tags_relation` VALUES (565, 66, 1);
INSERT INTO `media_tags_relation` VALUES (566, 67, 1);
INSERT INTO `media_tags_relation` VALUES (567, 68, 1);
INSERT INTO `media_tags_relation` VALUES (568, 69, 1);
INSERT INTO `media_tags_relation` VALUES (569, 70, 1);
INSERT INTO `media_tags_relation` VALUES (570, 71, 1);
INSERT INTO `media_tags_relation` VALUES (571, 72, 1);
INSERT INTO `media_tags_relation` VALUES (572, 73, 1);
INSERT INTO `media_tags_relation` VALUES (573, 74, 1);
INSERT INTO `media_tags_relation` VALUES (574, 74, 55);
INSERT INTO `media_tags_relation` VALUES (575, 75, 1);
INSERT INTO `media_tags_relation` VALUES (576, 75, 55);
INSERT INTO `media_tags_relation` VALUES (577, 76, 1);
INSERT INTO `media_tags_relation` VALUES (578, 76, 55);
INSERT INTO `media_tags_relation` VALUES (579, 77, 1);
INSERT INTO `media_tags_relation` VALUES (580, 77, 55);
INSERT INTO `media_tags_relation` VALUES (581, 78, 1);
INSERT INTO `media_tags_relation` VALUES (582, 78, 55);
INSERT INTO `media_tags_relation` VALUES (583, 79, 1);
INSERT INTO `media_tags_relation` VALUES (584, 79, 55);
INSERT INTO `media_tags_relation` VALUES (585, 80, 1);
INSERT INTO `media_tags_relation` VALUES (586, 80, 55);
INSERT INTO `media_tags_relation` VALUES (587, 81, 1);
INSERT INTO `media_tags_relation` VALUES (588, 81, 55);
INSERT INTO `media_tags_relation` VALUES (589, 82, 1);
INSERT INTO `media_tags_relation` VALUES (590, 82, 55);
INSERT INTO `media_tags_relation` VALUES (591, 82, 63);
INSERT INTO `media_tags_relation` VALUES (592, 83, 1);
INSERT INTO `media_tags_relation` VALUES (593, 83, 55);
INSERT INTO `media_tags_relation` VALUES (594, 84, 1);
INSERT INTO `media_tags_relation` VALUES (595, 84, 55);
INSERT INTO `media_tags_relation` VALUES (596, 85, 1);
INSERT INTO `media_tags_relation` VALUES (597, 85, 55);
INSERT INTO `media_tags_relation` VALUES (598, 86, 1);
INSERT INTO `media_tags_relation` VALUES (599, 86, 55);
INSERT INTO `media_tags_relation` VALUES (600, 87, 1);
INSERT INTO `media_tags_relation` VALUES (601, 87, 55);
INSERT INTO `media_tags_relation` VALUES (602, 88, 1);
INSERT INTO `media_tags_relation` VALUES (603, 88, 55);
INSERT INTO `media_tags_relation` VALUES (604, 89, 1);
INSERT INTO `media_tags_relation` VALUES (605, 89, 55);
INSERT INTO `media_tags_relation` VALUES (606, 89, 63);
INSERT INTO `media_tags_relation` VALUES (607, 90, 1);
INSERT INTO `media_tags_relation` VALUES (608, 90, 55);
INSERT INTO `media_tags_relation` VALUES (609, 91, 1);
INSERT INTO `media_tags_relation` VALUES (610, 91, 55);
INSERT INTO `media_tags_relation` VALUES (611, 92, 1);
INSERT INTO `media_tags_relation` VALUES (612, 92, 55);
INSERT INTO `media_tags_relation` VALUES (613, 93, 1);
INSERT INTO `media_tags_relation` VALUES (614, 93, 55);
INSERT INTO `media_tags_relation` VALUES (615, 94, 1);
INSERT INTO `media_tags_relation` VALUES (616, 95, 1);
INSERT INTO `media_tags_relation` VALUES (617, 96, 1);
INSERT INTO `media_tags_relation` VALUES (618, 97, 1);
INSERT INTO `media_tags_relation` VALUES (619, 98, 1);
INSERT INTO `media_tags_relation` VALUES (620, 99, 1);
INSERT INTO `media_tags_relation` VALUES (621, 100, 1);
INSERT INTO `media_tags_relation` VALUES (622, 101, 1);
INSERT INTO `media_tags_relation` VALUES (623, 102, 1);
INSERT INTO `media_tags_relation` VALUES (624, 103, 1);
INSERT INTO `media_tags_relation` VALUES (625, 104, 1);
INSERT INTO `media_tags_relation` VALUES (626, 105, 1);
INSERT INTO `media_tags_relation` VALUES (627, 106, 1);
INSERT INTO `media_tags_relation` VALUES (628, 107, 1);
INSERT INTO `media_tags_relation` VALUES (629, 108, 1);
INSERT INTO `media_tags_relation` VALUES (630, 109, 1);
INSERT INTO `media_tags_relation` VALUES (631, 110, 1);
INSERT INTO `media_tags_relation` VALUES (632, 111, 1);
INSERT INTO `media_tags_relation` VALUES (633, 112, 58);
INSERT INTO `media_tags_relation` VALUES (634, 112, 65);
INSERT INTO `media_tags_relation` VALUES (635, 113, 65);
INSERT INTO `media_tags_relation` VALUES (636, 113, 64);
INSERT INTO `media_tags_relation` VALUES (637, 114, 65);
INSERT INTO `media_tags_relation` VALUES (638, 114, 58);
INSERT INTO `media_tags_relation` VALUES (639, 115, 65);
INSERT INTO `media_tags_relation` VALUES (640, 115, 58);
INSERT INTO `media_tags_relation` VALUES (641, 115, 63);
INSERT INTO `media_tags_relation` VALUES (642, 116, 65);
INSERT INTO `media_tags_relation` VALUES (643, 116, 58);
INSERT INTO `media_tags_relation` VALUES (644, 116, 63);
INSERT INTO `media_tags_relation` VALUES (645, 117, 65);
INSERT INTO `media_tags_relation` VALUES (646, 117, 58);
INSERT INTO `media_tags_relation` VALUES (647, 117, 3);
INSERT INTO `media_tags_relation` VALUES (648, 118, 65);
INSERT INTO `media_tags_relation` VALUES (649, 118, 58);
INSERT INTO `media_tags_relation` VALUES (650, 118, 57);
INSERT INTO `media_tags_relation` VALUES (651, 119, 65);
INSERT INTO `media_tags_relation` VALUES (652, 119, 58);
INSERT INTO `media_tags_relation` VALUES (653, 119, 63);
INSERT INTO `media_tags_relation` VALUES (654, 119, 2);
INSERT INTO `media_tags_relation` VALUES (655, 120, 58);
INSERT INTO `media_tags_relation` VALUES (656, 120, 63);
INSERT INTO `media_tags_relation` VALUES (657, 121, 58);
INSERT INTO `media_tags_relation` VALUES (658, 121, 63);
INSERT INTO `media_tags_relation` VALUES (659, 122, 58);
INSERT INTO `media_tags_relation` VALUES (660, 122, 63);
INSERT INTO `media_tags_relation` VALUES (661, 123, 58);
INSERT INTO `media_tags_relation` VALUES (662, 124, 63);
INSERT INTO `media_tags_relation` VALUES (663, 125, 63);
INSERT INTO `media_tags_relation` VALUES (664, 125, 3);
INSERT INTO `media_tags_relation` VALUES (665, 125, 58);
INSERT INTO `media_tags_relation` VALUES (666, 126, 63);
INSERT INTO `media_tags_relation` VALUES (667, 126, 58);
INSERT INTO `media_tags_relation` VALUES (668, 127, 63);
INSERT INTO `media_tags_relation` VALUES (669, 127, 58);
INSERT INTO `media_tags_relation` VALUES (670, 128, 63);
INSERT INTO `media_tags_relation` VALUES (671, 128, 58);
INSERT INTO `media_tags_relation` VALUES (672, 129, 63);
INSERT INTO `media_tags_relation` VALUES (673, 129, 58);
INSERT INTO `media_tags_relation` VALUES (674, 129, 3);
INSERT INTO `media_tags_relation` VALUES (675, 130, 2);
INSERT INTO `media_tags_relation` VALUES (676, 130, 3);
INSERT INTO `media_tags_relation` VALUES (677, 130, 58);
INSERT INTO `media_tags_relation` VALUES (678, 130, 63);
INSERT INTO `media_tags_relation` VALUES (679, 130, 66);
INSERT INTO `media_tags_relation` VALUES (680, 131, 58);
INSERT INTO `media_tags_relation` VALUES (681, 131, 63);
INSERT INTO `media_tags_relation` VALUES (682, 131, 67);
INSERT INTO `media_tags_relation` VALUES (683, 132, 58);
INSERT INTO `media_tags_relation` VALUES (684, 132, 63);
INSERT INTO `media_tags_relation` VALUES (685, 132, 67);
INSERT INTO `media_tags_relation` VALUES (686, 133, 58);
INSERT INTO `media_tags_relation` VALUES (687, 133, 63);
INSERT INTO `media_tags_relation` VALUES (688, 133, 67);
INSERT INTO `media_tags_relation` VALUES (689, 134, 58);
INSERT INTO `media_tags_relation` VALUES (690, 134, 63);
INSERT INTO `media_tags_relation` VALUES (691, 134, 67);
INSERT INTO `media_tags_relation` VALUES (692, 135, 58);
INSERT INTO `media_tags_relation` VALUES (693, 135, 63);
INSERT INTO `media_tags_relation` VALUES (694, 135, 67);
INSERT INTO `media_tags_relation` VALUES (695, 136, 58);
INSERT INTO `media_tags_relation` VALUES (696, 136, 63);
INSERT INTO `media_tags_relation` VALUES (697, 136, 67);
INSERT INTO `media_tags_relation` VALUES (698, 137, 58);
INSERT INTO `media_tags_relation` VALUES (699, 137, 63);
INSERT INTO `media_tags_relation` VALUES (700, 137, 67);
INSERT INTO `media_tags_relation` VALUES (701, 138, 58);
INSERT INTO `media_tags_relation` VALUES (702, 138, 63);
INSERT INTO `media_tags_relation` VALUES (703, 138, 67);
INSERT INTO `media_tags_relation` VALUES (704, 139, 58);
INSERT INTO `media_tags_relation` VALUES (705, 139, 63);
INSERT INTO `media_tags_relation` VALUES (706, 139, 67);
INSERT INTO `media_tags_relation` VALUES (707, 140, 57);
INSERT INTO `media_tags_relation` VALUES (708, 140, 58);
INSERT INTO `media_tags_relation` VALUES (709, 140, 68);
INSERT INTO `media_tags_relation` VALUES (710, 141, 57);
INSERT INTO `media_tags_relation` VALUES (711, 141, 58);
INSERT INTO `media_tags_relation` VALUES (712, 142, 57);
INSERT INTO `media_tags_relation` VALUES (713, 142, 58);
INSERT INTO `media_tags_relation` VALUES (714, 143, 57);
INSERT INTO `media_tags_relation` VALUES (715, 143, 58);
INSERT INTO `media_tags_relation` VALUES (716, 144, 57);
INSERT INTO `media_tags_relation` VALUES (717, 144, 58);
INSERT INTO `media_tags_relation` VALUES (718, 145, 58);
INSERT INTO `media_tags_relation` VALUES (719, 145, 63);
INSERT INTO `media_tags_relation` VALUES (720, 146, 58);
INSERT INTO `media_tags_relation` VALUES (721, 146, 63);
INSERT INTO `media_tags_relation` VALUES (722, 146, 68);
INSERT INTO `media_tags_relation` VALUES (723, 147, 59);
INSERT INTO `media_tags_relation` VALUES (724, 148, 59);
INSERT INTO `media_tags_relation` VALUES (725, 148, 63);
INSERT INTO `media_tags_relation` VALUES (726, 149, 69);
INSERT INTO `media_tags_relation` VALUES (727, 149, 67);
INSERT INTO `media_tags_relation` VALUES (728, 149, 63);
INSERT INTO `media_tags_relation` VALUES (729, 150, 69);
INSERT INTO `media_tags_relation` VALUES (730, 150, 67);
INSERT INTO `media_tags_relation` VALUES (731, 150, 63);
INSERT INTO `media_tags_relation` VALUES (732, 151, 69);
INSERT INTO `media_tags_relation` VALUES (733, 151, 63);
INSERT INTO `media_tags_relation` VALUES (734, 151, 3);
INSERT INTO `media_tags_relation` VALUES (735, 152, 69);
INSERT INTO `media_tags_relation` VALUES (736, 152, 63);
INSERT INTO `media_tags_relation` VALUES (737, 152, 1);
INSERT INTO `media_tags_relation` VALUES (738, 153, 69);
INSERT INTO `media_tags_relation` VALUES (739, 153, 63);
INSERT INTO `media_tags_relation` VALUES (740, 153, 58);
INSERT INTO `media_tags_relation` VALUES (741, 154, 69);
INSERT INTO `media_tags_relation` VALUES (742, 154, 63);
INSERT INTO `media_tags_relation` VALUES (743, 154, 58);
INSERT INTO `media_tags_relation` VALUES (744, 154, 3);
INSERT INTO `media_tags_relation` VALUES (745, 155, 58);
INSERT INTO `media_tags_relation` VALUES (746, 155, 57);
INSERT INTO `media_tags_relation` VALUES (747, 155, 63);
INSERT INTO `media_tags_relation` VALUES (748, 156, 58);
INSERT INTO `media_tags_relation` VALUES (749, 156, 3);
INSERT INTO `media_tags_relation` VALUES (750, 156, 1);
INSERT INTO `media_tags_relation` VALUES (751, 157, 58);
INSERT INTO `media_tags_relation` VALUES (752, 157, 3);
INSERT INTO `media_tags_relation` VALUES (753, 157, 1);
INSERT INTO `media_tags_relation` VALUES (754, 157, 2);
INSERT INTO `media_tags_relation` VALUES (755, 158, 58);
INSERT INTO `media_tags_relation` VALUES (756, 158, 1);
INSERT INTO `media_tags_relation` VALUES (757, 159, 58);
INSERT INTO `media_tags_relation` VALUES (758, 159, 1);
INSERT INTO `media_tags_relation` VALUES (759, 160, 58);
INSERT INTO `media_tags_relation` VALUES (760, 160, 1);
INSERT INTO `media_tags_relation` VALUES (761, 160, 2);
INSERT INTO `media_tags_relation` VALUES (762, 161, 58);
INSERT INTO `media_tags_relation` VALUES (763, 161, 63);
INSERT INTO `media_tags_relation` VALUES (764, 162, 58);
INSERT INTO `media_tags_relation` VALUES (765, 162, 63);
INSERT INTO `media_tags_relation` VALUES (766, 163, 58);
INSERT INTO `media_tags_relation` VALUES (767, 163, 63);
INSERT INTO `media_tags_relation` VALUES (768, 164, 58);
INSERT INTO `media_tags_relation` VALUES (769, 164, 63);
INSERT INTO `media_tags_relation` VALUES (770, 164, 67);
INSERT INTO `media_tags_relation` VALUES (771, 165, 58);
INSERT INTO `media_tags_relation` VALUES (772, 165, 63);
INSERT INTO `media_tags_relation` VALUES (773, 166, 63);
INSERT INTO `media_tags_relation` VALUES (774, 166, 1);
COMMIT;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` int(1) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签id，自增',
  `tag_category` int(1) unsigned DEFAULT '1' COMMENT '外键，标签所属类型的id',
  `tag_name` varchar(30) NOT NULL COMMENT '标签名字',
  PRIMARY KEY (`tag_id`),
  KEY `index_ctg_name` (`tag_category`,`tag_name`),
  CONSTRAINT `fk_tag_ctg` FOREIGN KEY (`tag_category`) REFERENCES `category` (`ctg_id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COMMENT='标签表：博客的所有标签（1、标签名称：tag_name；2、标签创建时间：tag_createtime;3、标签所属类别：tag_category）';

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
INSERT INTO `tag` VALUES (3, 1, '丝袜');
INSERT INTO `tag` VALUES (66, 1, '假鸡巴自慰');
INSERT INTO `tag` VALUES (60, 1, '大槻響');
INSERT INTO `tag` VALUES (68, 1, '妈妈');
INSERT INTO `tag` VALUES (69, 1, '明里つむぎ');
INSERT INTO `tag` VALUES (58, 1, '淫叫');
INSERT INTO `tag` VALUES (57, 1, '熟女');
INSERT INTO `tag` VALUES (67, 1, '第一视角');
INSERT INTO `tag` VALUES (55, 1, '篠田ゆう');
INSERT INTO `tag` VALUES (59, 1, '紗々原ゆり');
INSERT INTO `tag` VALUES (1, 1, '美臀');
INSERT INTO `tag` VALUES (63, 1, '老婆');
INSERT INTO `tag` VALUES (2, 1, '肉丝');
INSERT INTO `tag` VALUES (62, 1, '自慰摸逼');
INSERT INTO `tag` VALUES (61, 1, '花咲一杏');
INSERT INTO `tag` VALUES (65, 1, '飯岡かなこ');
INSERT INTO `tag` VALUES (64, 1, '黑丝');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
