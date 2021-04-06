DROP TABLE IF EXISTS `t_user_info`;
CREATE TABLE `t_user_info` (
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `username` varchar(64) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '用户名',
  `phone` varchar(11) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '手机号',
  `gender` char(1) COLLATE utf8_bin NOT NULL COMMENT '性别 male 1； female 2；un_known or hide 0',
  `avatar` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '头像',
  `email` varchar(32) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '邮箱',
  `job` varchar(32) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '职位',
  `address` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '地址',
  `description` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '描述',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uniq_phone` (`phone`),
  UNIQUE KEY `uniq_user_id` (`user_id`),
  UNIQUE KEY `uniq_username` (`username`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_bin COMMENT = '用户信息表';
DROP TABLE IF EXISTS `t_password`;
CREATE TABLE `t_password` (
  `pwd_id` bigint(20) NOT NULL COMMENT '密码ID',
  `user_id` bigint(20) NOT NULL COMMENT '账号ID',
  `pwd` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '密码',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`pwd_id`),
  UNIQUE KEY `uniq_password_id` (`pwd_id`),
  UNIQUE KEY `uniq_user_id` (`user_id`),
  KEY `t_user_password_ibfk_1` (`user_id`),
  KEY `idx_password_id` (`pwd_id`),
  CONSTRAINT `t_user_password_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `t_user_info` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COLLATE = utf8_bin COMMENT = '密码表';
¸