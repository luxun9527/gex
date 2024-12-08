

DROP DATABASE IF EXISTS admin;
CREATE DATABASE admin CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

use admin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `coin`;
CREATE TABLE `coin`  (
                         `id` smallint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                         `coin_id` smallint NOT NULL DEFAULT 0 COMMENT '币种ID',
                         `coin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '币种名称',
                         `prec` tinyint NOT NULL DEFAULT 0 COMMENT '币种精度，小数位保留多少',
                         `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
                         `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
                         `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `uni_coin_name`(`coin_name` ASC) USING BTREE,
                         UNIQUE INDEX `uni_coin_id`(`coin_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `error_code`;
CREATE TABLE `error_code`  (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `error_code_id` int NOT NULL DEFAULT 0,
                               `error_code_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                               `language` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                               `created_at` int UNSIGNED NOT NULL DEFAULT 0,
                               `updated_at` int UNSIGNED NOT NULL DEFAULT 0,
                               `deleted_at` int UNSIGNED NOT NULL DEFAULT 0,
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `uni_error_code_id`(`error_code_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `symbol`;
CREATE TABLE `symbol`  (
                           `id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
                           `symbol_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易对名称',
                           `symbol_id` smallint NOT NULL DEFAULT 0 COMMENT '交易对id',
                           `base_coin_id` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '基础币ID',
                           `base_coin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '基础币名称',
                           `base_coin_prec` tinyint NOT NULL DEFAULT 0 COMMENT '基础币精度',
                           `quote_coin_id` smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '计价币ID',
                           `quote_coin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '计价币名称',
                           `quote_coin_prec` tinyint NOT NULL DEFAULT 0 COMMENT '计价币精度',
                           `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
                           `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
                           `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
                           PRIMARY KEY (`id`) USING BTREE,
                           UNIQUE INDEX `uni_symbol_name`(`symbol_name` ASC) USING BTREE,
                           UNIQUE INDEX `uni_symbol_id`(`symbol_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                         `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
                         `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
                         `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `created_at` int UNSIGNED NOT NULL DEFAULT 0,
                         `updated_at` int UNSIGNED NOT NULL DEFAULT 0,
                         `deleted_at` int UNSIGNED NOT NULL DEFAULT 0,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;





DROP DATABASE IF EXISTS trade;
CREATE DATABASE trade CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

use trade;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset`  (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `user_id` bigint NOT NULL COMMENT '用户ID',
                          `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
                          `coin_id` mediumint NOT NULL COMMENT '数字货币ID',
                          `coin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数字货币名称',
                          `available_qty` decimal(40, 18) NOT NULL COMMENT '可用余额',
                          `frozen_qty` decimal(40, 18) NOT NULL COMMENT '冻结金额',
                          `created_at` bigint NOT NULL COMMENT '创建时间',
                          `updated_at` bigint NOT NULL COMMENT '修改时间',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 208 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_00`;
CREATE TABLE `entrust_order_00`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_01`;
CREATE TABLE `entrust_order_01`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_02`;
CREATE TABLE `entrust_order_02`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_03`;
CREATE TABLE `entrust_order_03`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_04`;
CREATE TABLE `entrust_order_04`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_05`;
CREATE TABLE `entrust_order_05`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_06`;
CREATE TABLE `entrust_order_06`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_07`;
CREATE TABLE `entrust_order_07`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_08`;
CREATE TABLE `entrust_order_08`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `entrust_order_09`;
CREATE TABLE `entrust_order_09`  (
                                     `id` bigint NOT NULL COMMENT '序号 主键 雪花算法生成，递增',
                                     `order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
                                     `user_id` bigint NOT NULL COMMENT '用户id',
                                     `symbol_id` mediumint NOT NULL COMMENT '交易对ID',
                                     `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易对名称',
                                     `qty` decimal(40, 18) NOT NULL COMMENT '下单数量',
                                     `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                     `side` tinyint NOT NULL COMMENT '方向1买 2卖',
                                     `amount` decimal(40, 18) NOT NULL COMMENT '金额',
                                     `status` tinyint NOT NULL COMMENT '状态1新订单2部分成交 3全部成交，4撤销，5无效订单',
                                     `order_type` tinyint NOT NULL COMMENT '订单类型1市价单2限价单',
                                     `filled_qty` decimal(40, 18) NOT NULL COMMENT '成交数量',
                                     `un_filled_qty` decimal(40, 18) NOT NULL COMMENT '未成交数量',
                                     `filled_avg_price` decimal(40, 18) NOT NULL COMMENT '成交均价',
                                     `filled_amount` decimal(40, 18) NOT NULL COMMENT '成交金额',
                                     `un_filled_amount` decimal(40, 18) NOT NULL COMMENT '未成交金额',
                                     `created_at` bigint NOT NULL COMMENT '创建时间',
                                     `updated_at` bigint NOT NULL COMMENT '修改时间',
                                     `deleted_at` bigint NOT NULL COMMENT '删除时间',
                                     PRIMARY KEY (`id`) USING BTREE,
                                     INDEX `idx_user_id_status`(`user_id` ASC, `status` ASC) USING BTREE,
                                     INDEX `uni_order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `kline`;
CREATE TABLE `kline`  (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `start_time` bigint NOT NULL DEFAULT 0 COMMENT 'k线开始时间',
                          `end_time` bigint NOT NULL DEFAULT 0 COMMENT 'k线结束时间',
                          `symbol` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易对',
                          `symbol_id` smallint NOT NULL DEFAULT 0 COMMENT '交易对id',
                          `kline_type` tinyint NOT NULL DEFAULT 0 COMMENT 'k线类型1分钟 5分钟',
                          `open` decimal(40, 18) UNSIGNED NOT NULL COMMENT '开盘价',
                          `high` decimal(40, 18) UNSIGNED NOT NULL COMMENT 'k线内最高价',
                          `low` decimal(40, 18) UNSIGNED NOT NULL COMMENT 'k线内最低价',
                          `close` decimal(40, 18) UNSIGNED NOT NULL COMMENT '收盘价',
                          `amount` decimal(40, 18) UNSIGNED NOT NULL COMMENT '成交量(基础币数量)',
                          `volume` decimal(40, 18) UNSIGNED NOT NULL COMMENT '成交额(计价币数量)',
                          `range` decimal(40, 18) NOT NULL COMMENT '涨跌幅',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `uni_symbol_kt_open`(`symbol` ASC, `kline_type` ASC, `start_time` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80243 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `matched_order`;
CREATE TABLE `matched_order`  (
                                  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '雪花算法id',
                                  `match_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '撮合id',
                                  `match_sub_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '本次匹配的id，一次撮合会多次匹配',
                                  `symbol_id` mediumint NOT NULL DEFAULT 0 COMMENT '交易对id',
                                  `symbol_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易对名称',
                                  `taker_user_id` int NOT NULL DEFAULT 0 COMMENT 'taker用户id',
                                  `taker_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'taker订单id',
                                  `maker_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'maker订单id',
                                  `maker_user_id` int NOT NULL DEFAULT 0 COMMENT 'maker用户id',
                                  `taker_is_buyer` tinyint NOT NULL DEFAULT 2 COMMENT 'taker是否是买单 1是 2否',
                                  `price` decimal(40, 18) NOT NULL COMMENT '价格',
                                  `qty` decimal(40, 18) NOT NULL COMMENT '数量(基础币)',
                                  `amount` decimal(40, 18) NOT NULL COMMENT '金额（计价币）',
                                  `match_time` bigint NOT NULL DEFAULT 0 COMMENT '撮合时间',
                                  `created_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
                                  `updated_at` bigint NOT NULL DEFAULT 0 COMMENT '修改时间',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE INDEX `unqi_match_sub_id`(`match_sub_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 332 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `username` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
                         `password` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                         `phone_number` bigint NOT NULL COMMENT '手机号',
                         `status` int NOT NULL DEFAULT 1 COMMENT '用户状态，1正常2锁定',
                         `created_at` bigint NOT NULL COMMENT '创建时间',
                         `updated_at` bigint NOT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `uni_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 112 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;


create database if not exists dtm_barrier;

use dtm_barrier;
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




use trade;
INSERT INTO `trade`.`user` (`id`, `username`, `password`, `phone_number`, `status`, `created_at`, `updated_at`) VALUES (3, 'lisi', '$2a$10$9RMgCUfhSur5Gdcf9lFK/OzH8lDfpy95h829TsP14WKeOUIdcZboa', 0, 1, 1709361748, 1709361748);


INSERT INTO `trade`.`asset` (`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (1, 3, 'lisilisi', 10001, 'IKUN', 100000.000000000000000000, 0.000000000000000000, 1699151196, 1717842556);
INSERT INTO `trade`.`asset` (`id`, `user_id`, `username`, `coin_id`, `coin_name`, `available_qty`, `frozen_qty`, `created_at`, `updated_at`) VALUES (2, 3, 'lisilisi', 10002, 'USDT', 1000000.000000000000000000, 0.000000000000000000, 1699151196, 1717842556);

use admin;
INSERT INTO `admin`.`user` (`id`, `nickname`, `username`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '', 'test', '$2a$10$JlwKAMWRujhfry1WjQGGYO8a/LbkSmb0L/NJxReNBqdexYE697Gv6', 1707199053, 1707199053, 0);

INSERT INTO `admin`.`coin` (`id`, `coin_id`, `coin_name`, `prec`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 10002, 'USDT', 5, 1709478663, 1717332751, 0);
INSERT INTO `admin`.`coin` (`id`, `coin_id`, `coin_name`, `prec`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 10001, 'IKUN', 3, 1717430307, 1717430307, 0);

INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, 100001, '内部错误', 'zh-CN', 1707961978, 1707961978, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 100002, '内部错误', 'zh-CN', 1707962048, 1707962048, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 100003, '内部错误', 'zh-CN', 1707962054, 1707962054, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (31, 100004, '参数错误', 'zh-CN', 1707962068, 1707962068, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 100005, '记录未找到', 'zh-CN', 1707962099, 1707962099, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (33, 100006, '重复数据', 'zh-CN', 1707962110, 1707962110, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (34, 100007, '内部错误', 'zh-CN', 1707962118, 1707962118, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (35, 100009, '内部错误', 'zh-CN', 1707962125, 1707962125, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 100010, '内部错误', 'zh-CN', 1707962131, 1707962131, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (37, 100011, '内部错误', 'zh-CN', 1707962135, 1707962135, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (38, 200001, '用户不存在', 'zh-CN', 1707962179, 1707962179, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (39, 200002, '用户余额不足', 'zh-CN', 1707962186, 1707962186, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 200003, 'token验证失败', 'zh-CN', 1707962200, 1707962200, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 200004, 'token到期', 'zh-CN', 1707962210, 1707962210, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (42, 200005, '账户密码验证失败1', 'zh-CN', 1707962221, 1709366796, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (43, 500001, '订单未找到', 'zh-CN', 1707962248, 1707962248, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (44, 500002, '订单已经成交获取已经取消', 'zh-CN', 1707962273, 1707962273, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (45, 500003, '市价单不允许手动取消', 'zh-CN', 1707962297, 1707962297, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (46, 500004, '订单簿没有买单', 'zh-CN', 1707962337, 1707962337, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (47, 500005, '订单簿没有卖单', 'zh-CN', 1707962347, 1707972800, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (54, 11212, 'test', 'zh-CN', 1709477053, 1715176545, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (55, 50006, '超过最小精度11', 'zh-CN', 1715610797, 1717333078, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (58, 500006, '超过币种最小精度', 'zh-CN', 1715611690, 1717333116, 0);
INSERT INTO `admin`.`error_code` (`id`, `error_code_id`, `error_code_name`, `language`, `created_at`, `updated_at`, `deleted_at`) VALUES (59, 100012, '验证码错误', 'zh-CN', 1717341916, 1717341916, 0);


INSERT INTO `admin`.`symbol` (`id`, `symbol_name`, `symbol_id`, `base_coin_id`, `base_coin_name`, `base_coin_prec`, `quote_coin_id`, `quote_coin_name`, `quote_coin_prec`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'IKUN_USDT', 1, 10001, 'IKUN', 3, 10002, 'USDT', 5, 1717851844, 1717851844, 0);


