/*
 Navicat Premium Data Transfer

 Source Server         : sevice
 Source Server Type    : MySQL
 Source Server Version : 50739 (5.7.39-log)
 Source Host           : ip:host
 Source Schema         : sfgpl

 Target Server Type    : MySQL
 Target Server Version : 50739 (5.7.39-log)
 File Encoding         : 65001

 Date: 08/05/2023 19:34:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for actor
-- ----------------------------
DROP TABLE IF EXISTS `actor`;
CREATE TABLE `actor` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '作者id',
  `name` varchar(255) NOT NULL COMMENT '作者名',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `is_deleted` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of actor
-- ----------------------------
BEGIN;
INSERT INTO `actor` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (1, '文学家', '2023-04-27 00:04:34', '2023-04-27 00:04:34', NULL);
INSERT INTO `actor` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (2, '另一个文学家', '2023-04-27 00:08:10', '2023-04-27 00:08:10', NULL);
INSERT INTO `actor` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (3, '文学家-另一个', '2023-04-27 00:08:19', '2023-04-27 00:08:19', NULL);
INSERT INTO `actor` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (4, '历史学家', '2023-04-27 01:12:50', '2023-04-27 01:12:50', NULL);
COMMIT;

-- ----------------------------
-- Table structure for categorie
-- ----------------------------
DROP TABLE IF EXISTS `categorie`;
CREATE TABLE `categorie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `name` varchar(255) DEFAULT NULL COMMENT '分类名',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `is_deleted` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of categorie
-- ----------------------------
BEGIN;
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (1, '文学', '2023-04-25 11:24:31', '2023-04-25 11:24:34', NULL);
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (2, '科技', '2023-04-25 11:24:45', '2023-04-25 11:24:46', NULL);
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (3, '艺术', '2023-04-25 11:24:55', '2023-04-25 11:24:57', NULL);
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (4, '历史', '2023-04-25 11:25:06', '2023-04-25 11:25:08', NULL);
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (5, '语言', '2023-04-25 11:25:20', '2023-04-25 11:25:22', NULL);
INSERT INTO `categorie` (`id`, `name`, `update_time`, `create_time`, `is_deleted`) VALUES (6, '自然科学', '2023-04-25 11:25:35', '2023-04-25 11:25:36', NULL);
COMMIT;

-- ----------------------------
-- Table structure for program
-- ----------------------------
DROP TABLE IF EXISTS `program`;
CREATE TABLE `program` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '书籍id',
  `title` varchar(255) NOT NULL COMMENT '书籍名',
  `view` varchar(255) NOT NULL COMMENT '书籍简介',
  `actor_id` int(11) unsigned NOT NULL COMMENT '作者id',
  `categorie_id` int(11) unsigned NOT NULL COMMENT '分类id',
  `updeate_tine` datetime DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `is_deleted` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `categorie_id_index` (`categorie_id`),
  KEY `actor_id_key` (`actor_id`),
  CONSTRAINT `actor_id_key` FOREIGN KEY (`actor_id`) REFERENCES `actor` (`id`),
  CONSTRAINT `categorie_id_key` FOREIGN KEY (`categorie_id`) REFERENCES `categorie` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of program
-- ----------------------------
BEGIN;
INSERT INTO `program` (`id`, `title`, `view`, `actor_id`, `categorie_id`, `updeate_tine`, `create_time`, `is_deleted`) VALUES (4, '文学巨著-第三本书', '一本文学作品', 1, 1, '2023-04-27 01:15:50', '2023-04-27 01:15:50', NULL);
INSERT INTO `program` (`id`, `title`, `view`, `actor_id`, `categorie_id`, `updeate_tine`, `create_time`, `is_deleted`) VALUES (5, '文学巨著-第二本', '一本文学作品的第二本', 1, 1, '2023-04-27 00:53:02', '2023-04-27 00:53:02', '2023-04-27 01:17:37');
INSERT INTO `program` (`id`, `title`, `view`, `actor_id`, `categorie_id`, `updeate_tine`, `create_time`, `is_deleted`) VALUES (6, '文学巨著-另一本', '另外的一本文学作品', 3, 1, '2023-04-27 00:53:24', '2023-04-27 00:53:24', '2023-04-27 01:00:19');
INSERT INTO `program` (`id`, `title`, `view`, `actor_id`, `categorie_id`, `updeate_tine`, `create_time`, `is_deleted`) VALUES (7, '历史巨著', '一本历史学书', 4, 4, '2023-04-27 01:12:50', '2023-04-27 01:12:50', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
