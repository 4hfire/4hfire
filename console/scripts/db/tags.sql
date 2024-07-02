CREATE TABLE `tags` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '标签id',
                        `name` varchar(64) NOT NULL COMMENT '名称',
                        `desc` varchar(255) NOT NULL COMMENT '备注',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='标签'