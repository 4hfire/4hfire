CREATE TABLE `admins` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                          `uid` varchar(64) NOT NULL COMMENT 'uid',
                          `name` varchar(64) NOT NULL COMMENT '用户名称',
                          `account` varchar(32) NOT NULL COMMENT '账号',
                          `password` varchar(64) NOT NULL COMMENT '密码',
                          `phone` varchar(32) NOT NULL COMMENT '手机号',
                          `isDisable` tinyint unsigned NOT NULL COMMENT '状态',
                          `email` varchar(32) NOT NULL COMMENT '邮箱',
                          `otp` tinyint NOT NULL COMMENT 'otp 状态',
                          `code` varchar(64) NOT NULL COMMENT 'otp code',
                          `createdAt` int NOT NULL COMMENT '创建时间',
                          `updatedAt` int NOT NULL COMMENT '更新时间',
                          `lastLoginTime` int NOT NULL COMMENT '最后登陆时间',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uid` (`uid`),
                          UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员用户表'