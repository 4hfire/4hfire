CREATE TABLE `certs` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                         `name` varchar(32) NOT NULL COMMENT '名称',
                         `account` varchar(64) NOT NULL COMMENT '账号',
                         `password` varchar(64) NOT NULL COMMENT '密码',
                         `secret` varchar(1024) NOT NULL COMMENT '私钥',
                         `desc` varchar(255) NOT NULL COMMENT '备注',
                         `type` tinyint unsigned NOT NULL COMMENT '类型',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='凭证'