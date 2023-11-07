/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.2.159
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.2.159:3307
 Source Schema         : trade

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 29/10/2023 21:26:54
*/
CREATE DATABASE IF NOT EXISTS trade  default character set utf8mb4 collate utf8mb4_unicode_ci;

USE trade;

SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for asset
-- ----------------------------
DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) NOT NULL COMMENT '用户ID',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `coin_id` mediumint(0) NOT NULL COMMENT '数字货币ID',
  `coin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数字货币名称',
  `available_qty` decimal(40, 18) NOT NULL COMMENT '可用余额',
  `frozen_qty` decimal(40, 18) NOT NULL COMMENT '冻结金额',
  `created_at` bigint(0) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(0) NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for entrust_order
-- ----------------------------
DROP TABLE IF EXISTS `entrust_order`;
CREATE TABLE `entrust_order`  (
  `id` bigint(0) NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
  `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
  `user_id` bigint(0) NOT NULL COMMENT '用户id',
  `symbol_id` mediumint(0) NOT NULL COMMENT '交易对ID',
  `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
  `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
  `price` decimal(40, 18) NOT NULL COMMENT '价格',
  `side` tinyint(0) NOT NULL COMMENT '方向1买 2卖',
  `amount` decimal(40, 18) NOT NULL COMMENT '金额',
  `status` tinyint(0) NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
  `order_type` tinyint(0) NOT NULL COMMENT '订单类型1市价单2限价单',
  `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
  `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
  `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
  `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
  `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
  `created_at` bigint(0) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(0) NOT NULL COMMENT '修改时间',
  `deleted_at` bigint(0) NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for kline
-- ----------------------------
DROP TABLE IF EXISTS `kline`;
CREATE TABLE `kline`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `start_time` bigint(0) NOT NULL DEFAULT 0 COMMENT 'k线开始时间',
  `end_time` bigint(0) NOT NULL DEFAULT 0 COMMENT 'k线结束时间',
  `symbol` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易对',
  `symbol_id` smallint(0) NOT NULL DEFAULT 0 COMMENT '交易对id',
  `kline_type` tinyint(0) NOT NULL DEFAULT 0 COMMENT 'k线类型1分钟 5分钟',
  `open` decimal(40, 18) UNSIGNED NOT NULL COMMENT '开盘价',
  `high` decimal(40, 18) UNSIGNED NOT NULL COMMENT 'k线内最高价',
  `low` decimal(40, 18) UNSIGNED NOT NULL COMMENT 'k线内最低价',
  `close` decimal(40, 18) UNSIGNED NOT NULL COMMENT '收盘价',
  `amount` decimal(40, 18) UNSIGNED NOT NULL COMMENT '成交量(基础币数量)',
  `volume` decimal(40, 18) UNSIGNED NOT NULL COMMENT '成交额(计价币数量)',
  `range` decimal(40, 18) NOT NULL COMMENT '涨跌幅',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_symbol_kt_open`(`symbol`, `kline_type`, `start_time`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for matched_order
-- ----------------------------
DROP TABLE IF EXISTS `matched_order`;
CREATE TABLE `matched_order`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '雪花算法id',
  `match_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '撮合id',
  `match_sub_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '本次匹配的id，一次撮合会多次匹配',
  `symbol_id` mediumint(0) NOT NULL DEFAULT 0 COMMENT '交易对id',
  `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易对名称',
  `taker_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'taker订单id',
  `maker_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'maker订单id',
  `taker_is_buyer` tinyint(0) NOT NULL DEFAULT 2 COMMENT 'taker是否是买单 1是 2否',
  `price` decimal(40, 18) NOT NULL COMMENT '价格',
  `qty` decimal(40, 18) NOT NULL COMMENT '数量(基础币)',
  `amount` decimal(40, 18) NOT NULL COMMENT '金额（计价币）',
  `match_time` bigint(0) NOT NULL DEFAULT 0 COMMENT '撮合时间',
  `created_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` bigint(0) NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `phone_number` bigint(0) NOT NULL COMMENT '手机号',
  `status` int(0) NOT NULL COMMENT '用户状态，1正常2锁定',
  `created_at` bigint(0) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;


create database if not exists dtm_barrier
/*!40100 DEFAULT CHARACTER SET utf8mb4 */
;
drop table if exists dtm_barrier.barrier;
create table if not exists dtm_barrier.barrier(
    id bigint(22) PRIMARY KEY AUTO_INCREMENT,
    trans_type varchar(45) default '',
    gid varchar(128) default '',
    branch_id varchar(128) default '',
    op varchar(45) default '',
    barrier_id varchar(45) default '',
    reason varchar(45) default '' comment 'the branch type who insert this record',
    create_time datetime DEFAULT now(),
    update_time datetime DEFAULT now(),
    key(create_time),
    key(update_time),
    UNIQUE key(gid, branch_id, op, barrier_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;




INSERT INTO `trade`.`asset`(`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (1, 1, 'test1', 1, 'BTC', 100000.000000000000000000, 0.000000000000000000, 1699151196, 1699151196);
INSERT INTO `trade`.`asset`(`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (2, 1, 'test1', 2, 'USDT', 100000.000000000000000000, 0.000000000000000000, 1699151196, 1699151196);
INSERT INTO `trade`.`asset`(`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (3, 2, 'test2', 1, 'BTC', 100000.000000000000000000, 0.000000000000000000, 1699151196, 1699151196);
INSERT INTO `trade`.`asset`(`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (4, 2, 'test2', 2, 'USDT', 100000.000000000000000000, 0.000000000000000000, 1699151196, 1699151196);


INSERT INTO `trade`.`user`(`id`, `username`, `password`, `phone_number`, `status`, `created_at`, `updated_at`) VALUES (1, 'test1', '$2a$10$d6nu5vdV05v7EtxAstPAi.2FPKqFoUWSLhiRjHnzeRR.y5pR.qiRC', 2, 1, 1699151196, 1699151196);
INSERT INTO `trade`.`user`(`id`, `username`, `password`, `phone_number`, `status`, `created_at`, `updated_at`) VALUES (2, 'test2', '$2a$10$d6nu5vdV05v7EtxAstPAi.2FPKqFoUWSLhiRjHnzeRR.y5pR.qiRC', 2, 1, 1699151196, 1699151196);
