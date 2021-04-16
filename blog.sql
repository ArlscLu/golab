/*
 Navicat MySQL Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 16/04/2021 18:25:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '内容',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `created_on` int unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth` VALUES (1, 'test', 'test123');
COMMIT;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

-- ----------------------------
-- Table structure for t_dict
-- ----------------------------
DROP TABLE IF EXISTS `t_dict`;
CREATE TABLE `t_dict` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID,自增主键',
  `type` varchar(50) NOT NULL DEFAULT '' COMMENT '字典组,建议表名+字段组合',
  `parent_code` varchar(50) DEFAULT NULL COMMENT '父级CODE',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '字典CODE',
  `value` varchar(50) NOT NULL DEFAULT '' COMMENT '字典CODE对应值',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0-正常 1-删除',
  `create_time` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int NOT NULL DEFAULT '0' COMMENT '删除时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'mysql更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_ty_co_en` (`type`,`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=COMPACT COMMENT='字典表';

-- ----------------------------
-- Records of t_dict
-- ----------------------------
BEGIN;
INSERT INTO `t_dict` VALUES (1, 'coupon_type', NULL, '1', '满减券', '优惠券类型', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (2, 'coupon_type', NULL, '2', '折扣券', '优惠券类型\r\n', 1, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (3, 'coupon_type', NULL, '3', '随机金额券', '优惠券类型\r\n', 1, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (4, 'coupon_type', NULL, '4', '兑换券', '优惠券类型', 1, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (5, 'coupon_type', NULL, '5', '运费券', '优惠券类型', 1, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (6, 'use_platform', NULL, '1', 'app', '优惠券使用平台', 0, 0, 0, 0, '2020-07-02 11:44:59');
INSERT INTO `t_dict` VALUES (7, 'use_platform', NULL, '2', 'h5', '优惠券使用平台', 0, 0, 0, 0, '2020-07-02 10:24:04');
INSERT INTO `t_dict` VALUES (8, 'use_platform', NULL, '3', '微信小程序', '优惠券使用平台', 0, 0, 0, 0, '2020-07-02 11:45:05');
INSERT INTO `t_dict` VALUES (9, 'use_platform', NULL, '4', '小店', '优惠券使用平台', 1, 0, 0, 0, '2020-07-02 10:24:04');
INSERT INTO `t_dict` VALUES (10, 'coupon_status', NULL, '0', '创建中', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (11, 'coupon_status', NULL, '1', '待提交审核', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (12, 'coupon_status', NULL, '2', '待审核', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (13, 'coupon_status', NULL, '3', '审核不通过', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (14, 'coupon_status', NULL, '4', '审核通过', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (15, 'coupon_status', NULL, '5', '未开始', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (16, 'coupon_status', NULL, '6', '进行中', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (17, 'coupon_status', NULL, '7', '已停止', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (18, 'coupon_status', NULL, '8', '已结束', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (19, 'coupon_status', NULL, '9', '已失效', '优惠券状态', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (20, 'apply_range', NULL, '0', '平台券', '优惠券适用范围', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (21, 'apply_range', NULL, '1', '店铺券', '优惠券适用范围', 1, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (22, 'apply_range', NULL, '2', '品类券', '优惠券适用范围', 1, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (23, 'apply_range', NULL, '3', '品牌券', '优惠券适用范围', 1, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (24, 'apply_goods_type', NULL, '1', '指定商品标签可用', '优惠券适用商品类型', 0, 0, 0, 0, '2020-08-14 10:32:57');
INSERT INTO `t_dict` VALUES (25, 'apply_goods_type', NULL, '2', '指定商品不可用', '优惠券适用商品类型', 1, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (26, 'receive_user_limit', NULL, '1', '按用户等级', '领取用户限制', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (27, 'receive_user_limit', NULL, '2', '按用户标签', '领取用户限制', 0, 0, 0, 0, '2020-07-02 11:43:17');
INSERT INTO `t_dict` VALUES (28, 'sponsor', NULL, '1', '平台', '优惠券承担方', 0, 0, 0, 0, '2020-07-02 11:45:31');
INSERT INTO `t_dict` VALUES (29, 'sponsor', NULL, '2', '商家', '优惠券承担方', 0, 0, 0, 0, '2020-07-02 11:45:36');
INSERT INTO `t_dict` VALUES (30, 'sponsor', NULL, '3', '平台和商家', '优惠券承担方', 0, 0, 0, 0, '2020-07-02 11:45:36');
INSERT INTO `t_dict` VALUES (31, 'coupon_use_type', NULL, '1', '未使用', '优惠券使用状态', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (32, 'coupon_use_type', NULL, '2', '已使用', '优惠券使用状态', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (33, 'coupon_use_type', NULL, '3', '已过期', '优惠券使用状态', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (35, 'task_reward_type', NULL, '2', '积分', '任务奖励说明', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (36, 'task_reward_type', NULL, '3', '优惠券', '任务奖励说明', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (37, 'task_reward_type', NULL, '4', '幸运大转盘', '任务奖励说明', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (38, 'coupon_push_type', NULL, '1', '商品详情页', '优惠券定向推送', 0, 0, 0, 0, '2020-07-02 11:46:08');
INSERT INTO `t_dict` VALUES (39, 'coupon_push_type', NULL, '2', '购物车', '优惠券定向推送', 1, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (44, 'use_platform', NULL, '5', '微信公众号', '优惠券使用平台', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (45, 'business_type', NULL, '1', '港华精选', '业务类型', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (46, 'business_type', NULL, '2', '港华紫荆', '业务类型', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (47, 'business_type', NULL, '3', '港华到家', '业务类型', 0, 0, 0, 0, '2020-07-02 11:45:47');
INSERT INTO `t_dict` VALUES (48, 'use_platform', NULL, '6', '支付宝小程序', '优惠券使用平台', 1, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (49, 'apply_goods_type', NULL, '3', '商品分类', '优惠券适用商品类型', 1, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (54, 'card_mutex', NULL, '1', '紫荆卡', '用券卡互斥', 0, 0, 0, 0, '2020-07-02 11:43:18');
INSERT INTO `t_dict` VALUES (56, 'apply_goods_type', NULL, '4', '指定商品', '优惠券适用商品类型', 0, 0, 0, 0, '2020-08-10 15:33:18');
INSERT INTO `t_dict` VALUES (57, 'gf_base_status', NULL, '1', '待提交', '礼金卡模版状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (58, 'gf_base_status', NULL, '2', '待审核', '礼金卡模版状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (59, 'gf_base_status', NULL, '3', '已生效', '礼金卡模版状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (60, 'gf_base_status', NULL, '4', '审核不通过', '礼金卡模版状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (61, 'gf_base_status', NULL, '5', '已终止生卡', '礼金卡模版状态', 0, 0, 0, 0, '2020-07-02 11:43:16');
INSERT INTO `t_dict` VALUES (62, 'coupon_push_type', NULL, '3', '申请管家成功页', '优惠券定向推送', 0, 0, 0, 0, '2020-10-19 12:09:04');
INSERT INTO `t_dict` VALUES (63, 'gf_usage', NULL, 'Q011', '销售', '礼金卡用途备注', 0, 1605852069, 1605852069, 0, '2020-11-24 11:04:05');
INSERT INTO `t_dict` VALUES (64, 'gf_usage', NULL, 'Q012', '员工福利', '礼金卡用途备注', 0, 1605852069, 1605852069, 0, '2020-11-24 11:04:04');
INSERT INTO `t_dict` VALUES (66, 'gf_usage', NULL, 'G005', '充值1', '', 1, 1605851994, 1606715382, 1606715382, '2020-11-30 13:49:47');
INSERT INTO `t_dict` VALUES (67, 'gf_usage', NULL, 'G002', '充值', 'xxx', 0, 1605852069, 1605855045, 0, '2020-11-20 14:50:48');
INSERT INTO `t_dict` VALUES (68, 'gf_usage_parent', NULL, 'V001', '礼金卡用途', '', 0, 1605852069, 1606298788, 0, '2020-11-30 09:36:18');
INSERT INTO `t_dict` VALUES (69, 'gf_usage', NULL, '1111', '2222', 'sfdsfsdf', 0, 1606209559, 1606715359, 0, '2020-11-30 13:49:24');
INSERT INTO `t_dict` VALUES (70, 'gf_usage', NULL, 'asd', 'asdf', 'asdfasfd', 0, 1606214221, 1606715359, 0, '2020-11-30 13:49:24');
INSERT INTO `t_dict` VALUES (71, 'gf_usage', NULL, 'asdfs', 'asdfd', 'asdfasfd', 1, 1606214363, 1606297077, 1606297077, '2020-11-25 17:38:01');
INSERT INTO `t_dict` VALUES (72, 'gf_usage', NULL, 'G003', '充值', '', 1, 1606715255, 1606715255, 1606715255, '2020-11-30 13:47:41');
INSERT INTO `t_dict` VALUES (73, 'use_platform', NULL, '7', '管理后台', '优惠券使用平台', 0, 0, 0, 0, '2020-12-11 14:08:02');
INSERT INTO `t_dict` VALUES (74, 'use_platform', NULL, '8', '微信小程序直播', '优惠券使用平台', 0, 0, 0, 0, '2020-12-11 14:08:23');
INSERT INTO `t_dict` VALUES (75, 'coupon_push_type', NULL, '4', '积分兑换专区', '优惠券定向推送', 0, 0, 0, 0, '2020-10-19 12:09:04');
INSERT INTO `t_dict` VALUES (76, 'use_platform', NULL, '9', '链接', '优惠券使用平台', 0, 0, 0, 0, '2020-12-21 11:04:38');
INSERT INTO `t_dict` VALUES (79, 'vipgift_type', NULL, '1', '赠送金币', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:24:58');
INSERT INTO `t_dict` VALUES (80, 'vipgift_type', NULL, '2', '赠送商城优惠券', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:25:28');
INSERT INTO `t_dict` VALUES (81, 'vipgift_type', NULL, '3', '赠送到家优惠券', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:25:30');
INSERT INTO `t_dict` VALUES (82, 'vipgift_type', NULL, '4', '赠送燃气优惠券', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:25:31');
INSERT INTO `t_dict` VALUES (83, 'vipgift_type', NULL, '5', '营养师健康咨询', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:25:33');
INSERT INTO `t_dict` VALUES (84, 'vipgift_type', NULL, '6', '赠送商城优惠券(门店福利)', 'vip权益礼包权益类型', 0, 1617949707, 1617949707, 0, '2021-04-16 11:25:35');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `passwd` varchar(200) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'Tom', '', 11, '', '2020-08-31 15:11:37', '2020-08-31 15:11:37');
INSERT INTO `user` VALUES (2, 'Rose', '', 12, '', '2020-08-31 15:15:05', '2020-08-31 15:15:05');
INSERT INTO `user` VALUES (3, 'Jack', '', 19, '', '2020-08-31 16:42:17', '2020-08-31 16:42:17');
INSERT INTO `user` VALUES (4, 'Kitty', '', 14, '', '2020-08-31 16:49:50', '2020-08-31 16:49:50');
INSERT INTO `user` VALUES (5, 'Asp', '', 6, '', '2020-08-31 17:19:09', '2020-08-31 17:19:09');
INSERT INTO `user` VALUES (6, 'Steven', '', 18, '', '2020-08-31 17:48:20', '2020-08-31 17:48:20');
INSERT INTO `user` VALUES (7, 'Goblin', '', 16, '', '2020-08-31 18:01:42', '2020-08-31 18:01:42');
INSERT INTO `user` VALUES (8, 'Andy', '', 1, '', '2020-08-31 18:23:37', '2020-08-31 18:23:37');
INSERT INTO `user` VALUES (9, 'Gandof', '', 33, '', '2020-08-31 18:24:03', '2020-08-31 18:24:03');
INSERT INTO `user` VALUES (10, 'Tom', '', 18, '', '2020-09-01 16:34:46', '2020-09-01 16:34:46');
INSERT INTO `user` VALUES (11, 'Tom', '', 18, '', '2020-09-01 16:48:51', '2020-09-01 16:48:51');
INSERT INTO `user` VALUES (12, 'Tom', '', 18, '', '2020-09-01 16:51:15', '2020-09-01 16:51:15');
INSERT INTO `user` VALUES (13, 'Tom', '', 18, '', '2020-09-01 17:34:39', '2020-09-01 17:34:39');
INSERT INTO `user` VALUES (14, 'Tom', '', 18, '', '2020-09-02 15:38:48', '2020-09-02 15:38:48');
INSERT INTO `user` VALUES (15, 'Tom', '', 18, '', '2020-09-02 15:50:44', '2020-09-02 15:50:44');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
