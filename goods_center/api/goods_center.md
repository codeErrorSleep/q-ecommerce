# api模块

## 库表设计
商品表分成3个维度

商品->规格->属性

spu->sku->attr

每个规格都有一个对应的属性

sku->stock

每个商品又有对应的类别,基本都是三四层的树形结构





# 库表设计

spu表,存储商品的基本信息,添加字段前需要思考当前字段是否合适,是否应该存到该表上

因为该表的读取量会非常大.

## 商品spu表

```sql
// 建库
CREATE DATABASE db_q_goods_center;
// 建表
CREATE TABLE `t_spu` (
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `app_id`            varchar(64)     NOT NULL COMMENT '店铺id',
    `spu_id`            varchar(64)     NOT NULL COMMENT '商品id',
    `spu_type`          varchar(128)    NOT NULL DEFAULT '' COMMENT '商品类型(关联t_spu_type)',
    `goods_category_id` varchar(64)     NOT NULL DEFAULT '' COMMENT '商品分类id',
    `goods_name`        varchar(256)    NOT NULL DEFAULT '' COMMENT '商品名称',
    `goods_img`         varchar(1024)   NOT NULL DEFAULT '' COMMENT '商品封面图（默认封面图）',
    `price`             int             NOT NULL DEFAULT '0' COMMENT '商品最低价',
    `price_line`        int             NOT NULL DEFAULT '0' COMMENT '划线价',
    `goods_tag`         varchar(258)    NOT NULL DEFAULT '' COMMENT '商品标签',
    `sale_status`       tinyint(1)      NOT NULL DEFAULT '1' COMMENT '上架状态： 0下架 1上架 2（定时上架还未上架阶段）待上架 4强制下架 5封禁',
    `sale_at`           timestamp       NOT NULL COMMENT '实际上架的时间',
    `is_display`        tinyint(1)      NOT NULL DEFAULT '0' COMMENT '是否显示：0否(隐藏状态) 1是(显示状态)',
    `limit_purchase`    int             NOT NULL DEFAULT '0' COMMENT '限购数量',
    `stock_deduct_mode` tinyint(1)      NOT NULL DEFAULT '0' COMMENT '扣库存方式：0付款减库存 1拍下减库存',
    `is_deleted`        tinyint         NOT NULL DEFAULT '0' COMMENT '0正常 1已删除',
    `created_at`        timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，有修改自动更新',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unx_app_spu` (`app_id`, `spu_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE = InnoDB
```

