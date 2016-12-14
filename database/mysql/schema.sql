/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50711
Source Host           : localhost:3306
Source Database       : dish

Target Server Type    : MYSQL
Target Server Version : 50711
File Encoding         : 65001

Date: 2016-04-05 21:18:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) NOT NULL COMMENT '业务类型，1游戏、0非游戏',
  `application_name` varchar(255) NOT NULL COMMENT '业务名称，唯一',
  `life_cycle` varchar(255) NOT NULL COMMENT '生命周期',
  `level` tinyint(1) NOT NULL,
  `owner_id` int(11) unsigned NOT NULL,
  `default` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `application_name` (`application_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4064 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `model` varchar(32) NOT NULL default '',
  `cpu` int(3) NOT NULL default '0',
  `memory` int(3) NOT NULL default '0',
  `host_name` varchar(255) DEFAULT NULL,
  `inner_ip` varchar(32) NOT NULL COMMENT '内网IP地址',
  `inner_gate` varchar(32) DEFAULT NULL,
  `inner_interface` varchar(255) DEFAULT NULL,
  `bgp_ip` varchar(255) DEFAULT NULL,
  `bgp_gate` varchar(32) DEFAULT NULL,
  `bgp_interface` varchar(255) DEFAULT NULL,
  `outer_ip` varchar(32) DEFAULT NULL COMMENT '外网IP地址',
  `outer_gate` varchar(32) DEFAULT NULL,
  `outer_interface` varchar(255) DEFAULT NULL,
  `ilo_ip` varchar(255) DEFAULT NULL,
  `source` tinyint(1) unsigned NOT NULL COMMENT '来源，1腾讯云，3其他云',
  `module_id` int(11) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL,
  `set_id` int(11) DEFAULT NULL,
  `set_name` varchar(255) DEFAULT NULL,
  `application_id` int(11) DEFAULT NULL,
  `application_name` varchar(255) DEFAULT NULL,
  `owner` varchar(255) DEFAULT NULL,
  `checked` varchar(255) DEFAULT NULL,
  `is_distributed` tinyint(1) unsigned NOT NULL COMMENT '是否已被分配，0未分配，1已分配',
  PRIMARY KEY (`id`),
  UNIQUE KEY `inner_ip` (`inner_ip`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for module
-- ----------------------------
DROP TABLE IF EXISTS `module`;
CREATE TABLE `module` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `module_name` varchar(255) NOT NULL COMMENT '模块名称、唯一',
  `operator` int(11) unsigned NOT NULL COMMENT '维护人',
  `bak_operator` int(11) unsigned NOT NULL COMMENT '备份维护人',
  `application_id` int(10) unsigned NOT NULL COMMENT '所属业务',
  `owner` int(11) DEFAULT NULL,
  `set_id` int(11) DEFAULT NULL COMMENT '所属集群',
  PRIMARY KEY (`id`),
  UNIQUE KEY `module_name` (`module_name`,`set_id`) USING BTREE,
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for set
-- ----------------------------
DROP TABLE IF EXISTS `set`;
CREATE TABLE `set` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `application_id` int(11) unsigned NOT NULL COMMENT '所属业务',
  `capacity` int(11) NOT NULL DEFAULT '0' COMMENT '集群容量',
  `chn_name` varchar(255) DEFAULT NULL COMMENT '中文名称',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `default` tinyint(1) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL COMMENT '集群描述信息',
  `envi_type` tinyint(1) DEFAULT NULL COMMENT '环境类型',
  `last_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `open_status` varchar(255) DEFAULT '0' COMMENT '集群状态',
  `parent_id` int(11) DEFAULT NULL,
  `service_status` tinyint(4) DEFAULT NULL COMMENT '服务状态',
  `set_name` varchar(255) NOT NULL COMMENT '集群名称',
  `owner` int(11) unsigned NOT NULL COMMENT '所属开发商',
  PRIMARY KEY (`id`),
  UNIQUE KEY `set_name` (`set_name`,`application_id`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，0正常 -1禁用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', 'admin@example.com', '7fef6171469e80d32c0559f88b377245', 'This is a Secret KEY', '1458897589', '127.0.0.1', '0');