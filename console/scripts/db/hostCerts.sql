CREATE TABLE `hostCerts` (
                             `id` bigint unsigned NOT NULL COMMENT '主机ID',
                             `account` varchar(64) NOT NULL COMMENT '账号',
                             `password` varchar(64) NOT NULL COMMENT '密码',
                             `secret` varchar(1024) NOT NULL COMMENT '私钥',
                             `certId` bigint unsigned NOT NULL COMMENT '凭证ID',
                             `type` tinyint unsigned NOT NULL COMMENT '类型',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='主机凭证'