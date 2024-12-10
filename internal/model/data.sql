CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '用户密码（加密存储）',
  `created_at` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint DEFAULT '0' COMMENT '更新时间',
  `portrait_url` varchar(1024) DEFAULT NULL COMMENT '头像链接',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4  COMMENT='用户表'



CREATE TABLE `feed` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'auto increment primary key',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `title` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '投稿标题',
  `tag` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '标签[]',
  `description` text NOT NULL COMMENT '投稿内容',
  `vote_count` int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `share_count` int unsigned NOT NULL DEFAULT '0' COMMENT '分享数量',
  `common_count` int unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
  `report_count` int unsigned NOT NULL DEFAULT '0' COMMENT '举报数量',
  `image_list` varchar(2000) NOT NULL DEFAULT '' COMMENT 'multi json',
  `operate_name` varchar(128) NOT NULL DEFAULT '' COMMENT '后台操作 人',
  `delete_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '删除状态，0=正常',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '投稿类型,1为普通投稿,2图片投稿3为视频投稿',
   `db_time` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `db_time` (`db_time`),
  KEY `tag` (`tag`),
  KEY `idx_type_status` (`user_id`,`delete_status`,`type`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户投稿表'
