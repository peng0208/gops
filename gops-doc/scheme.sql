/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : 127.0.0.1
 Source Database       : opsadmin

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : utf-8

 Date: 10/30/2017 20:08:36 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `conf_tag`
-- ----------------------------
DROP TABLE IF EXISTS `conf_tag`;
CREATE TABLE `conf_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tagname` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tagname` (`tagname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Table structure for `cron_tag`
-- ----------------------------
DROP TABLE IF EXISTS `cron_tag`;
CREATE TABLE `cron_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tagname` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tagname` (`tagname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Table structure for `cron_task`
-- ----------------------------
DROP TABLE IF EXISTS `cron_task`;
CREATE TABLE `cron_task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `env` varchar(255) NOT NULL,
  `args` varchar(255) DEFAULT NULL,
  `schedule` varchar(255) NOT NULL,
  `enable` int(11) DEFAULT NULL,
  `ctime` bigint(20) DEFAULT NULL,
  `mtime` bigint(20) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `nickname` varchar(32) COLLATE utf8_bin DEFAULT NULL,
  `password` varbinary(128) DEFAULT NULL,
  `ctime` bigint(20) DEFAULT NULL,
  `mtime` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
--  Records of `user`
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('1', 'admin', '管理员', 0xb630a0fd82898c31390d180d5d81d9ac9708e3a08ac85654e797747b69a7790f, '1509365086', null);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
