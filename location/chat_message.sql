/*
 Navicat MySQL Data Transfer

 Source Server         : 聊天测试
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : chat

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 22/04/2022 18:10:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat_message
-- ----------------------------
DROP TABLE IF EXISTS `chat_message`;
CREATE TABLE `chat_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `send_account` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `recipient_account` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `message_type` int(11) NOT NULL,
  `content` varchar(11) COLLATE utf8_unicode_ci NOT NULL,
  `time` int(11) NOT NULL,
  `time_stamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
