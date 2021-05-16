-- MySQL Backup
-- Database: jobor
-- Backup Time: 2021-05-16 00:23:15


-- SET FOREIGN_KEY_CHECKS=0;

CREATE TABLE IF NOT EXISTS `casbin_rule` (
                               `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                               `ptype` varchar(100) DEFAULT NULL,
                               `v0` varchar(100) DEFAULT NULL,
                               `v1` varchar(100) DEFAULT NULL,
                               `v2` varchar(100) DEFAULT NULL,
                               `v3` varchar(100) DEFAULT NULL,
                               `v4` varchar(100) DEFAULT NULL,
                               `v5` varchar(100) DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`),
                               UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `jobor_log` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '''主键id''',
                             `name` varchar(128) DEFAULT NULL COMMENT '''任务名''',
                             `lang` varchar(16) NOT NULL COMMENT '''任务类型:[shell,api,python,golang,e.g.]''',
                             `task_id` bigint(20) unsigned DEFAULT NULL COMMENT '''关联任务id''',
                             `trigger_method` varchar(16) DEFAULT NULL COMMENT '''任务触发方式：auto,manual''',
                             `expr` varchar(32) NOT NULL COMMENT '''定时任务表达式：0/1 * * ? * * * 秒分时天月星期''',
                             `data` text NOT NULL COMMENT '''任务执行详细，格式：json''',
                             `resp` text COMMENT '''任务返回结果''',
                             `cost_time` double DEFAULT NULL COMMENT '''任务耗时''',
                             `result` varchar(16) DEFAULT NULL COMMENT '''任务结果: success,failed''',
                             `err_code` bigint(20) DEFAULT NULL COMMENT '''错误码''',
                             `err_msg` text COMMENT '''错误信息''',
                             `addr` varchar(64) NOT NULL COMMENT '''任务运行的worker节点''',
                             `start_time` datetime DEFAULT NULL COMMENT '''开始时间''',
                             `end_time` datetime DEFAULT NULL COMMENT '''结束时间''',
                             `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '''创建时间''',
                             `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '''更新时间''',
                             PRIMARY KEY (`id`),
                             KEY `idx_code` (`name`,`lang`),
                             KEY `task_id` (`task_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `jobor_task` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '''主键id''',
                              `name` varchar(128) NOT NULL COMMENT '''任务名''',
                              `description` varchar(512) DEFAULT NULL COMMENT '''任务描述',
                              `lang` varchar(16) NOT NULL COMMENT '''任务类型:[shell,api,python,golang,e.g.]''',
                              `data` text NOT NULL COMMENT '''任务执行详细，格式：json''',
                              `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '''关联用户id''',
                              `count` bigint(20) DEFAULT NULL COMMENT '''执行次数''',
                              `expr` varchar(32) NOT NULL COMMENT '''定时任务表达式：0/1 * * ? * * * 秒分时天月星期''',
                              `timeout` bigint(20) DEFAULT '-1' COMMENT '''超时时间''',
                              `alarm_policy` bigint(20) DEFAULT '2' COMMENT '''告警策略：{0:never,1:always,2:failed,3:success}''',
                              `expect_code` bigint(20) DEFAULT '0' COMMENT '''期望任务状态码''',
                              `expect_content` varchar(500) DEFAULT NULL COMMENT '''期望任务返回结果''',
                              `retry` bigint(20) DEFAULT '0' COMMENT '''重试次数''',
                              `route_policy` bigint(20) DEFAULT '1' COMMENT '''路由策略 1:Random 2:RoundRobin 3:Weight 4:LeastTask''',
                              `routing_key` varchar(32) DEFAULT 'default' COMMENT '''执行worker路由标识''',
                              `status` varchar(32) DEFAULT NULL COMMENT '''定时任务状态: running,stop''',
                              `by_update` varchar(64) DEFAULT NULL COMMENT '''更新人''',
                              `prev` datetime DEFAULT NULL COMMENT '''上次时间''',
                              `next` datetime DEFAULT NULL COMMENT '''下次时间''',
                              `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '''创建时间''',
                              `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '''更新时间''',
                              `d` text,
                              `notify` text COMMENT '''告警通知，格式：json''',
                              PRIMARY KEY (`id`),
                              KEY `user_id` (`user_id`),
                              KEY `idx_code` (`lang`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `jobor_worker` (
                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '''主键id''',
                                `hostname` varchar(128) NOT NULL COMMENT '''worker hostname''',
                                `addr` varchar(64) NOT NULL COMMENT '''worker ip''',
                                `version` varchar(32) DEFAULT NULL COMMENT '''版本''',
                                `routing_key` varchar(64) DEFAULT 'default' COMMENT '''routing_key''',
                                `weight` bigint(20) DEFAULT NULL COMMENT '''权重''',
                                `lease_update` bigint(20) DEFAULT NULL COMMENT '''租约时间更新''',
                                `status` varchar(32) DEFAULT 'offline' COMMENT '''状态：running,stop,offline''',
                                `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '''创建时间''',
                                `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '''更新时间''',
                                PRIMARY KEY (`id`),
                                KEY `idx_code` (`hostname`,`addr`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
CREATE TABLE IF NOT EXISTS `permission` (
                              `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                              `name` varchar(128) DEFAULT NULL,
                              `method` varchar(64) NOT NULL,
                              `path` varchar(128) NOT NULL,
                              `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`) USING BTREE,
                              UNIQUE KEY `idx_name_code` (`path`,`method`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `role` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(128) DEFAULT NULL,
                        `description` longtext,
                        `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `by_update` varchar(64) DEFAULT NULL,
                        PRIMARY KEY (`id`) USING BTREE,
                        UNIQUE KEY `idx_name_code` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `role_permissions` (
                                    `role_id` int(10) unsigned NOT NULL,
                                    `permission_id` int(10) unsigned NOT NULL,
                                    PRIMARY KEY (`role_id`,`permission_id`) USING BTREE,
                                    KEY `fk_role_permissions_permission` (`permission_id`),
                                    CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`),
                                    CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `user` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                        `nickname` varchar(128) DEFAULT NULL,
                        `username` varchar(128) DEFAULT NULL,
                        `password` varchar(128) DEFAULT NULL,
                        `email` varchar(156) DEFAULT NULL,
                        `phone` varchar(32) DEFAULT NULL,
                        `user_type_id` varchar(156) DEFAULT '1',
                        `user_type` varchar(16) DEFAULT 'local' COMMENT '用户类型:local,ldap,sso',
                        `status` varchar(32) DEFAULT '1' COMMENT '状态',
                        `by_update` varchar(64) DEFAULT NULL,
                        `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`) USING BTREE,
                        KEY `idx_name_code` (`nickname`,`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `user_roles` (
                              `user_id` int(10) unsigned NOT NULL,
                              `role_id` int(10) unsigned NOT NULL,
                              PRIMARY KEY (`user_id`,`role_id`) USING BTREE,
                              KEY `fk_user_roles_role` (`role_id`),
                              CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
                              CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `usergroup` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `name` varchar(128) DEFAULT NULL,
                             `description` longtext,
                             `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `owner_id` int(10) unsigned DEFAULT NULL,
                             `by_update` varchar(64) DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             KEY `owner_id` (`owner_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `usergroup_roles` (
                                   `usergroup_id` int(10) unsigned NOT NULL,
                                   `role_id` int(10) unsigned NOT NULL,
                                   PRIMARY KEY (`usergroup_id`,`role_id`) USING BTREE,
                                   KEY `fk_usergroup_users_user` (`role_id`),
                                   CONSTRAINT `usergroup_roles_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `user` (`id`),
                                   CONSTRAINT `usergroup_roles_ibfk_2` FOREIGN KEY (`usergroup_id`) REFERENCES `usergroup` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
CREATE TABLE IF NOT EXISTS `usergroup_users` (
                                   `usergroup_id` int(10) unsigned NOT NULL,
                                   `user_id` int(10) unsigned NOT NULL,
                                   PRIMARY KEY (`usergroup_id`,`user_id`) USING BTREE,
                                   KEY `fk_usergroup_users_user` (`user_id`),
                                   CONSTRAINT `fk_usergroup_users_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
                                   CONSTRAINT `fk_usergroup_users_usergroup` FOREIGN KEY (`usergroup_id`) REFERENCES `usergroup` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
