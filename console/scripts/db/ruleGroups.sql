CREATE TABLE `ruleGroups` (
                                       `id` BIGINT UNSIGNED   NOT NULL     ,
                                       `name` VARCHAR(64)   NOT NULL    COMMENT '名称' ,
                                       `desc` VARCHAR(255)   NOT NULL    COMMENT '备注' ,
                                       `tags` JSON   NOT NULL    COMMENT '所属标签' ,
                                       UNIQUE INDEX `name` (`name`)  ,
                                       PRIMARY KEY  (`id`)
) COMMENT='规则分组';