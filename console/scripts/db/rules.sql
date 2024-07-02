CREATE TABLE `rules` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
                         `name` varchar(64) NOT NULL COMMENT '名称',
                         `desc` varchar(255) NOT NULL COMMENT '备注',
                         `tags` json NOT NULL COMMENT '所属标签',
                         `groupId` bigint unsigned NOT NULL COMMENT '所属分组',
                         `interface` varchar(32) NOT NULL COMMENT '网卡',
                         `sip` varchar(32) NOT NULL COMMENT '源IP',
                         `sport` varchar(255) NOT NULL COMMENT '源端口',
                         `dip` varchar(32) NOT NULL COMMENT '目的IP',
                         `dport` varchar(255) NOT NULL COMMENT '目的端口',
                         `option` tinyint unsigned NOT NULL COMMENT '规则动作：accept/drop',
                         `enable` tinyint unsigned NOT NULL COMMENT '启用状态',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `name` (`name`),
                         KEY `groupId` (`groupId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='规则'