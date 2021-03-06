CREATE TABLE `sys_settings`  (
	`id` BIGINT auto_increment COMMENT '编号',
  `name`  VARCHAR ( 50 ) NOT NULL DEFAULT NULL COMMENT '分类',
  `classify` BIGINT NOT NULL DEFAULT NULL COMMENT '分类',
  `content` json NULL COMMENT '内容',
	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '系统配置';
