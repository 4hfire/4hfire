CREATE TABLE `hosts` (
                         `id` bigint unsigned NOT NULL COMMENT 'id',
                         `uuid` varchar(32) NOT NULL COMMENT '主机uuid',
                         `secret` varchar(32) NOT NULL COMMENT '主机密钥',
                         `address` varchar(32) NOT NULL COMMENT '主机地址',
                         `tags` json NOT NULL COMMENT '所属标签',
                         `name` varchar(32) NOT NULL COMMENT '主机名称',
                         `desc` varchar(255) NOT NULL COMMENT '备注',
                         `state` tinyint unsigned NOT NULL COMMENT '状态',
                         `cpu` varchar(32) NOT NULL COMMENT 'cpu信息',
                         `mem` varchar(32) NOT NULL COMMENT '内存信息',
                         `updatedAt` bigint NOT NULL COMMENT '更新时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='主机'