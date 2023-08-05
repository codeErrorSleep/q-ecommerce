CREATE TABLE `t_orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `shop_id` varchar(64) NOT NULL DEFAULT '' COMMENT '店铺id',
  `user_id` varchar(64) NOT NULL DEFAULT '' COMMENT '用户id',
  `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
  `goods_id` varchar(64) NOT NULL DEFAULT '' COMMENT '商品id',
  `money` int NOT NULL COMMENT '订单金额',
  `num` int NOT NULL DEFAULT '1' COMMENT '购买数量',
  `status` int NOT NULL DEFAULT '1' COMMENT '订单状态 1:待支付,2:已支付,3:待退款,4:已退款',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除：0-正常，1-删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单表'