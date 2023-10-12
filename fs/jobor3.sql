-- SET NAMES utf8mb4;
-- SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
-- DROP TABLE IF EXISTS `api`;
CREATE TABLE IF NOT EXISTS `api` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `name` varchar(158) NOT NULL COMMENT '名称',
    `title` varchar(158) DEFAULT NULL COMMENT '标题',
    `dom` varchar(128) DEFAULT 'default' COMMENT '租户、域',
    `path` varchar(158) DEFAULT NULL COMMENT '路由路径',
    `method` varchar(64) DEFAULT NULL COMMENT '请求方法',
    `enabled` tinyint(1) DEFAULT '1' COMMENT '启用',
    `group` varchar(64) DEFAULT 'default' COMMENT '接口分组',
    `enabled_audit` tinyint(1) DEFAULT '0' COMMENT '启用审计否',
    `updater` varchar(64) DEFAULT NULL COMMENT '更新人',
    PRIMARY KEY (`id`),
    KEY `name` (`name`),
    KEY `title` (`title`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of api
-- ----------------------------
-- BEGIN;
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (1, 'get:jobor:dashboard', '查看dashboard', 'jobor', '/api/v1/jobor/dashboard', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (2, 'get:jobor:api', '查看api', 'jobor', '/api/v1/jobor/api', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (3, 'get:jobor:api-auto-update', '查看api-auto-update', 'jobor', '/api/v1/jobor/api-auto-update', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (4, 'get:jobor:apis', '查看apis', 'jobor', '/api/v1/jobor/apis', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (5, 'get:jobor:audit-log', '查看audit-log', 'jobor', '/api/v1/jobor/audit-log', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (6, 'get:jobor:log', '查看log', 'jobor', '/api/v1/jobor/log', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (7, 'get:jobor:log::id', '查看log', 'jobor', '/api/v1/jobor/log/:id', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (8, 'get:jobor:worker', '查看worker', 'jobor', '/api/v1/jobor/worker', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (9, 'get:jobor:worker-routing-key', '查看worker-routing-key', 'jobor', '/api/v1/jobor/worker-routing-key', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (10, 'get:jobor:workers', '查看workers', 'jobor', '/api/v1/jobor/workers', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (11, 'get:jobor:worker::id', '查看worker', 'jobor', '/api/v1/jobor/worker/:id', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (12, 'get:jobor:task', '查看task', 'jobor', '/api/v1/jobor/task', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (13, 'get:jobor:tasks', '查看tasks', 'jobor', '/api/v1/jobor/tasks', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (14, 'get:jobor:task::id', '查看task', 'jobor', '/api/v1/jobor/task/:id', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (15, 'get:jobor:user', '查看user', 'jobor', '/api/v1/jobor/user', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (16, 'get:jobor:user-self', '查看user-self', 'jobor', '/api/v1/jobor/user-self', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (17, 'get:jobor:user-sync', '查看user-sync', 'jobor', '/api/v1/jobor/user-sync', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (18, 'get:jobor:user-switch::user_id', '查看user-switch', 'jobor', '/api/v1/jobor/user-switch/:user_id', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (19, 'get:jobor:users', '查看users', 'jobor', '/api/v1/jobor/users', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (20, 'get:jobor:user::id', '查看user', 'jobor', '/api/v1/jobor/user/:id', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (21, 'get:jobor:oidc:redirect', '查看oidc', 'jobor', '/api/v1/jobor/oidc/redirect', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (22, 'get:jobor:oidc:callback', '查看oidc', 'jobor', '/api/v1/jobor/oidc/callback', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (23, 'get:jobor:migrate', '查看migrate', 'jobor', '/api/v1/jobor/migrate', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (24, 'get:jobor:state-code', '查看state-code', 'jobor', '/api/v1/jobor/state-code', 'GET', 1, 'jobor', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (25, 'get:sys:role', '查看role', 'jobor', '/api/v1/sys/role', 'GET', 1, 'sys', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (26, 'get:sys:role-tree', '查看role-tree', 'jobor', '/api/v1/sys/role-tree', 'GET', 1, 'sys', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (27, 'get:sys:roles', '查看roles', 'jobor', '/api/v1/sys/roles', 'GET', 1, 'sys', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (28, 'get:login', '查看login', 'jobor', '/api/v1/login', 'GET', 1, 'login', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (29, 'get:logout', '查看logout', 'jobor', '/api/v1/logout', 'GET', 1, 'logout', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (30, 'get:refresh_token', '查看refresh_token', 'jobor', '/api/v1/refresh_token', 'GET', 1, 'refresh_token', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (31, 'get:ping', '查看ping', 'jobor', '/auth/ping', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (32, 'get:ping', '查看ping', 'jobor', '/ping', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (33, 'get:panic', '查看panic', 'jobor', '/panic', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (34, 'get:routes', '查看routes', 'jobor', '/routes', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (35, 'get:*any', '查看*any', 'jobor', '/swagger/*any', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (36, 'get:*any', '查看*any', 'jobor', '/jobor/*any', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (37, 'get:login', '查看login', 'jobor', '/login', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (38, 'get:*filepath', '查看*filepath', 'jobor', '/*filepath', 'GET', 1, 'default', 0, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (39, 'put:sys:role::id', '修改role', 'jobor', '/api/v1/sys/role/:id', 'PUT', 1, 'sys', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (40, 'put:jobor:api::id', '修改api', 'jobor', '/api/v1/jobor/api/:id', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (41, 'put:jobor:worker::id', '修改worker', 'jobor', '/api/v1/jobor/worker/:id', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (42, 'put:jobor:task::id', '修改task', 'jobor', '/api/v1/jobor/task/:id', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (43, 'put:jobor:user:pass-reset', '修改user', 'jobor', '/api/v1/jobor/user/pass-reset', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (44, 'put:jobor:user:password', '修改user', 'jobor', '/api/v1/jobor/user/password', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (45, 'put:jobor:user:profile', '修改user', 'jobor', '/api/v1/jobor/user/profile', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (46, 'put:jobor:user::id', '修改user', 'jobor', '/api/v1/jobor/user/:id', 'PUT', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (47, 'delete:sys:role::id', '删除role', 'jobor', '/api/v1/sys/role/:id', 'DELETE', 1, 'sys', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (48, 'delete:jobor:api::id', '删除api', 'jobor', '/api/v1/jobor/api/:id', 'DELETE', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (49, 'delete:jobor:worker::id', '删除worker', 'jobor', '/api/v1/jobor/worker/:id', 'DELETE', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (50, 'delete:jobor:task::id', '删除task', 'jobor', '/api/v1/jobor/task/:id', 'DELETE', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (51, 'delete:jobor:user::id', '删除user', 'jobor', '/api/v1/jobor/user/:id', 'DELETE', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (52, 'post:sys:role', '创建role', 'jobor', '/api/v1/sys/role', 'POST', 1, 'sys', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (53, 'post:jobor:api', '创建api', 'jobor', '/api/v1/jobor/api', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (54, 'post:jobor:log::id', '创建log', 'jobor', '/api/v1/jobor/log/:id/abort', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (55, 'post:jobor:worker', '创建worker', 'jobor', '/api/v1/jobor/worker', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (56, 'post:jobor:task', '创建task', 'jobor', '/api/v1/jobor/task', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (57, 'post:jobor:task::id', '创建task', 'jobor', '/api/v1/jobor/task/:id/run', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (58, 'post:jobor:user', '创建user', 'jobor', '/api/v1/jobor/user', 'POST', 1, 'jobor', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (59, 'post:login', '创建login', 'jobor', '/api/v1/login', 'POST', 1, 'login', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (60, 'post:logout', '创建logout', 'jobor', '/api/v1/logout', 'POST', 1, 'logout', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (61, 'post:refresh_token', '创建refresh_token', 'jobor', '/api/v1/refresh_token', 'POST', 1, 'refresh_token', 1, '');
INSERT IGNORE INTO `api` (`id`, `name`, `title`, `dom`, `path`, `method`, `enabled`, `group`, `enabled_audit`, `updater`) VALUES (62, 'head:*filepath', '*filepath', 'jobor', '/*filepath', 'HEAD', 1, 'default', 1, '');
-- COMMIT;

-- ----------------------------
-- Table structure for audit_log
-- ----------------------------
-- DROP TABLE IF EXISTS `audit_log`;
CREATE TABLE IF NOT EXISTS `audit_log` (
    `user_id` bigint(20) DEFAULT NULL COMMENT '关联用户',
    `nickname` varchar(128) DEFAULT NULL COMMENT '显示名',
    `user` varchar(128) DEFAULT NULL COMMENT '用户名',
    `action` varchar(256) DEFAULT NULL COMMENT '具体方法描述',
    `business` varchar(128) DEFAULT NULL COMMENT '业务',
    `method` varchar(8) DEFAULT NULL COMMENT '请求方法',
    `handler` varchar(256) DEFAULT NULL COMMENT '处理程序/接口',
    `user_ip` varchar(128) DEFAULT NULL COMMENT '用户源IP',
    `http_code` bigint(20) DEFAULT NULL COMMENT '返回状态',
    `body` longtext COMMENT '请求Body',
    `resp_body` longtext COMMENT '响应Body',
    `dom` varchar(128) DEFAULT 'default' COMMENT '所属域',
    `obj_id` bigint(20) DEFAULT NULL COMMENT '关联变更对象ID',
    `cost_time` longtext COMMENT '耗时',
    `status` bigint(20) DEFAULT NULL COMMENT '状态[0: error, 1: success, 2: forbidden, 3: exception]',
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of audit_log
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
-- DROP TABLE IF EXISTS `casbin_rule`;
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
    UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/api', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (24, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/api', 'POST', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (2, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/api-auto-update', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (21, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/api/:id', 'DELETE', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (15, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/api/:id', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (3, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/apis', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (4, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/audit-log', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (5, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (25, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user', 'POST', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (6, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user-self', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (8, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user-switch/:user_id', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (7, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user-sync', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (22, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/:id', 'DELETE', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (10, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/:id', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (19, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/:id', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (16, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/pass-reset', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (17, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/password', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (18, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/user/profile', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (9, 'p', 'sys_rw', 'jobor', '/api/v1/jobor/users', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (11, 'p', 'sys_rw', 'jobor', '/api/v1/sys/role', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (23, 'p', 'sys_rw', 'jobor', '/api/v1/sys/role', 'POST', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (12, 'p', 'sys_rw', 'jobor', '/api/v1/sys/role-tree', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (20, 'p', 'sys_rw', 'jobor', '/api/v1/sys/role/:id', 'DELETE', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (14, 'p', 'sys_rw', 'jobor', '/api/v1/sys/role/:id', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (13, 'p', 'sys_rw', 'jobor', '/api/v1/sys/roles', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (26, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (32, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker', 'POST', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (27, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker-routing-key', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (31, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker/:id', 'DELETE', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (29, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker/:id', 'GET', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (30, 'p', 'task_rw', 'jobor', '/api/v1/jobor/worker/:id', 'PUT', '', '');
INSERT IGNORE INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (28, 'p', 'task_rw', 'jobor', '/api/v1/jobor/workers', 'GET', '', '');
COMMIT;

-- ----------------------------
-- Table structure for jobor_log
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_log`;
CREATE TABLE IF NOT EXISTS `jobor_log` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `name` varchar(128) DEFAULT NULL COMMENT '任务名',
    `lang` varchar(16) NOT NULL COMMENT '任务类型:[shell,api,python,golang,e.g.]',
    `task_id` bigint(20) DEFAULT NULL COMMENT '关联任务id',
    `trigger_method` varchar(16) DEFAULT NULL COMMENT '任务触发方式：auto,manual',
    `expr` varchar(32) NOT NULL COMMENT '定时任务表达式：0/1 * * ? * * * 秒分时天月星期',
    `data` mediumtext NOT NULL COMMENT '任务执行详细，格式：json',
    `resp` mediumtext COMMENT '任务返回结果',
    `cost_time` float DEFAULT '0' COMMENT '执行耗时(毫秒)',
    `result` varchar(16) DEFAULT NULL COMMENT '任务结果: success,failed',
    `err_code` bigint(20) DEFAULT NULL COMMENT '''错误码''',
    `err_msg` mediumtext COMMENT '错误信息',
    `addr` varchar(64) NOT NULL COMMENT '任务运行的log节点',
    `start_time` datetime DEFAULT NULL COMMENT '开始时间',
    `end_time` datetime DEFAULT NULL COMMENT '结束时间',
    `idempotent` varchar(156) DEFAULT NULL COMMENT '幂等标志',
    `by_task_log_id` bigint(20) DEFAULT NULL COMMENT '关联任务log id',
    `expect_content` varchar(500) DEFAULT NULL COMMENT '期望任务返回结果',
    `expect_code` bigint(20) DEFAULT '0' COMMENT '期望任务状态码',
    `executor` varchar(152) DEFAULT 'timed' COMMENT '任务执行人',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_code` (`name`,`lang`),
    KEY `task_id` (`task_id`),
    KEY `by_task_log_id` (`by_task_log_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jobor_log_childs
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_log_childs`;
CREATE TABLE IF NOT EXISTS `jobor_log_childs` (
    `jobor_log_id` bigint(20) NOT NULL COMMENT '主键id',
    `child_id` bigint(20) NOT NULL COMMENT '主键id',
    PRIMARY KEY (`jobor_log_id`,`child_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_log_childs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jobor_log_parents
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_log_parents`;
CREATE TABLE IF NOT EXISTS `jobor_log_parents` (
    `jobor_log_id` bigint(20) NOT NULL COMMENT '主键id',
    `parent_id` bigint(20) NOT NULL COMMENT '主键id',
    PRIMARY KEY (`jobor_log_id`,`parent_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_log_parents
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jobor_task
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_task`;
CREATE TABLE IF NOT EXISTS `jobor_task` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `name` varchar(128) NOT NULL COMMENT '任务名',
    `description` mediumtext COMMENT '任务描述',
    `lang` varchar(16) NOT NULL COMMENT '任务类型:[shell,api,python,golang,e.g.]',
    `data` mediumtext NOT NULL COMMENT '任务执行详细，格式：json',
    `notify` mediumtext COMMENT '告警通知，格式：json',
    `user_id` bigint(20) DEFAULT NULL COMMENT '关联用户id',
    `user` varchar(191) DEFAULT NULL COMMENT '关联用户',
    `count` bigint(20) DEFAULT NULL COMMENT '执行次数',
    `parent_task_ids` varchar(128) DEFAULT NULL COMMENT '父任务ID，最多20个',
    `parent_run_parallel` tinyint(1) DEFAULT '0' COMMENT '父任务是否并行运行',
    `child_task_ids` varchar(128) DEFAULT NULL COMMENT '子任务ID，最多20个',
    `child_run_parallel` tinyint(1) DEFAULT '0' COMMENT '子任务是否并行运行',
    `expr` varchar(32) NOT NULL COMMENT '定时任务表达式：0/1 * * ? * * * 秒分时天月星期',
    `timeout` bigint(20) DEFAULT '-1' COMMENT '超时时间',
    `route_policy` bigint(20) DEFAULT '1' COMMENT '路由策略 1:Random 2:RoundRobin 3:Weight 4:LeastTask',
    `routing_key` varchar(32) DEFAULT 'default' COMMENT '执行worker路由标识',
    `routing_keys` varchar(255) DEFAULT '["default"]' COMMENT '任务标识，多选',
    `status` varchar(32) DEFAULT 'running' COMMENT '定时任务状态: running,stop',
    `alarm_policy` bigint(20) DEFAULT '2' COMMENT '告警策略：{0:never,1:always,2:failed,3:success}',
    `expect_content` varchar(500) DEFAULT NULL COMMENT '期望任务返回结果',
    `expect_code` bigint(20) DEFAULT '0' COMMENT '期望任务状态码',
    `retry` bigint(20) DEFAULT '0' COMMENT '重试次数',
    `prev` datetime DEFAULT NULL COMMENT '上次执行时间',
    `next` datetime DEFAULT NULL COMMENT '''下次执行时间''',
    `updater` varchar(156) DEFAULT NULL,
    `deleted` tinyint(1) DEFAULT '0' COMMENT '逻辑删除',
    `keep_log` text COMMENT '保留日志数type: count、day',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE,
    KEY `idx_code` (`lang`),
    KEY `user_id` (`user_id`),
    KEY `user` (`user`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_task
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `jobor_task` (`id`, `name`, `description`, `lang`, `data`, `notify`, `user_id`, `user`, `count`, `parent_task_ids`, `parent_run_parallel`, `child_task_ids`, `child_run_parallel`, `expr`, `timeout`, `route_policy`, `routing_key`, `routing_keys`, `status`, `alarm_policy`, `expect_content`, `expect_code`, `retry`, `prev`, `next`, `updater`, `deleted`, `keep_log`) VALUES (1, 'ShellDemo', 'ShellDemo', 'shell', '{\"content\":\"#!/bin/bash \\necho `date`\\nsleep 10\\necho \\\"wellcome use jobor cron task!\\\"\",\"api\":{\"url\":\"\",\"method\":\"get\",\"content_type\":\"json\",\"payload\":\"\",\"body\":\"\",\"auth_method\":\"no_auth\",\"auth_data\":{\"username\":\"\",\"password\":\"\"},\"header_list\":[],\"form_data_list\":[],\"www_form_list\":[],\"header\":null,\"forms\":null}}', '{\"email\":{},\"webhook\":{},\"lark\":{},\"dingding\":{},\"wechat\":{}}', 0, '', 0, NULL, 0, NULL, 0, '0 */2 * * * *', 120, 1, 'default', '[\"default\"]', 'stop', 2, 'wellcome use jobor', 1, 0, '2023-09-13 13:30:28', '0001-01-01 00:00:00', '', 0, '{\"val\":180,\"keep_type\":\"count\"}');
INSERT IGNORE INTO `jobor_task` (`id`, `name`, `description`, `lang`, `data`, `notify`, `user_id`, `user`, `count`, `parent_task_ids`, `parent_run_parallel`, `child_task_ids`, `child_run_parallel`, `expr`, `timeout`, `route_policy`, `routing_key`, `routing_keys`, `status`, `alarm_policy`, `expect_content`, `expect_code`, `retry`, `prev`, `next`, `updater`, `deleted`, `keep_log`) VALUES (2, 'ApiDemo', 'ApiDemo', 'api', '{\"content\":\"#!/bin/bash \\necho `date`\\nsleep 60\\necho \\\"wellcome use jobor cron task!\\\"\",\"api\":{\"url\":\"http://127.0.0.1:5002/ping\",\"method\":\"GET\",\"content_type\":\"json\",\"payload\":\"\",\"body\":\"\",\"auth_method\":\"no_auth\",\"auth_data\":{\"username\":\"\",\"password\":\"\"},\"header_list\":[{\"key\":\"abc\",\"value\":\"abc\"}],\"form_data_list\":[],\"www_form_list\":[],\"header\":null,\"forms\":null}}', '{\"email\":{},\"webhook\":{},\"lark\":{},\"dingding\":{},\"wechat\":{}}', 0, '', 0, NULL, 0, NULL, 0, '0 */2 * * * *', 120, 1, 'default', '[\"default\"]', 'stop', 2, '', 200, 0, '2023-09-13 13:30:30', '0001-01-01 00:00:00', '', 0, '{\"val\":180,\"keep_type\":\"count\"}');
INSERT IGNORE INTO `jobor_task` (`id`, `name`, `description`, `lang`, `data`, `notify`, `user_id`, `user`, `count`, `parent_task_ids`, `parent_run_parallel`, `child_task_ids`, `child_run_parallel`, `expr`, `timeout`, `route_policy`, `routing_key`, `routing_keys`, `status`, `alarm_policy`, `expect_content`, `expect_code`, `retry`, `prev`, `next`, `updater`, `deleted`, `keep_log`) VALUES (3, 'sshModeDemo', '', 'shell', '{\"content\":\"#!/bin/bash\\nexport PATH=/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/data/consul/bin:/data/jdk1.8.0_261/bin:/data/jdk1.8.0_261/jre/bin:/data/nginx/sbin:/home/ops/.local/bin:/home/ops/bin\\n\\nifconfig\\n\\n# echo `ip addr`\\n\\necho \\\"wellcome use jobor worker ssh mode!\\\"\\n\",\"api\":{\"url\":\"\",\"method\":\"GET\",\"content_type\":\"json\",\"payload\":\"\",\"body\":\"\",\"auth_method\":\"no_auth\",\"auth_data\":{\"username\":\"\",\"password\":\"\"},\"header_list\":[],\"form_data_list\":[],\"www_form_list\":[],\"header\":null,\"forms\":null}}', '{\"email\":{\"receivers\":[\"\"]},\"webhook\":{},\"lark\":{},\"dingding\":{},\"wechat\":{}}', 0, '', 0, '[]', 0, '[]', 0, '0 0 10 * * *', 60, 1, 'ssh', '[\"ssh\"]', 'stop', 2, '', 0, 0, '2023-09-19 13:04:07', '0001-01-01 00:00:00', '', 0, '{\"val\":180,\"keep_type\":\"count\"}');
INSERT IGNORE INTO `jobor_task` (`id`, `name`, `description`, `lang`, `data`, `notify`, `user_id`, `user`, `count`, `parent_task_ids`, `parent_run_parallel`, `child_task_ids`, `child_run_parallel`, `expr`, `timeout`, `route_policy`, `routing_key`, `routing_keys`, `status`, `alarm_policy`, `expect_content`, `expect_code`, `retry`, `prev`, `next`, `updater`, `deleted`, `keep_log`) VALUES (4, 'pyDemo', '', 'python3', '{\"pre_cmd\":\"pip3 install requests\",\"content\":\"def a():\\n  print(\\\"python3:\\\")\\n  \\na()\",\"api\":{\"url\":\"\",\"method\":\"GET\",\"content_type\":\"json\",\"payload\":\"\",\"body\":\"\",\"auth_method\":\"no_auth\",\"auth_data\":{\"username\":\"\",\"password\":\"\"},\"header_list\":[],\"form_data_list\":[],\"www_form_list\":[],\"header\":null,\"forms\":null}}', '{\"email\":{},\"webhook\":{},\"lark\":{},\"dingding\":{},\"wechat\":{}}', 0, '', 0, '[]', 0, '[]', 0, '0 0 10 * * *', 60, 1, 'ssh', '[\"python3.7\"]', 'stop', 2, '', 0, 0, '2023-09-20 20:09:10', '0001-01-01 00:00:00', '', 0, '{\"val\":180,\"keep_type\":\"count\"}');
INSERT IGNORE INTO `jobor_task` (`id`, `name`, `description`, `lang`, `data`, `notify`, `user_id`, `user`, `count`, `parent_task_ids`, `parent_run_parallel`, `child_task_ids`, `child_run_parallel`, `expr`, `timeout`, `route_policy`, `routing_key`, `routing_keys`, `status`, `alarm_policy`, `expect_content`, `expect_code`, `retry`, `prev`, `next`, `updater`, `deleted`, `keep_log`) VALUES (5, 'goDemo', '', 'golang', '{\"content\":\"package main\\n\\nimport (\\n\\t\\\"fmt\\\"\\n)\\n\\nfunc main(){\\n  fmt.Println(\\\"hellow go task\\\")\\n}\",\"api\":{\"url\":\"\",\"method\":\"GET\",\"content_type\":\"json\",\"payload\":\"\",\"body\":\"\",\"auth_method\":\"no_auth\",\"auth_data\":{\"username\":\"\",\"password\":\"\"},\"header_list\":[],\"form_data_list\":[],\"www_form_list\":[],\"header\":null,\"forms\":null}}', '{\"email\":{},\"webhook\":{},\"lark\":{},\"dingding\":{},\"wechat\":{}}', 0, '', 0, '[]', 0, '[]', 0, '0 0 10 * * *', 60, 1, 'ssh', '[\"go1.20\"]', 'stop', 2, '', 0, 0, '2023-09-20 20:09:11', '0001-01-01 00:00:00', '', 0, '{\"val\":180,\"keep_type\":\"count\"}');
COMMIT;

-- ----------------------------
-- Table structure for jobor_task_owners
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_task_owners`;
CREATE TABLE IF NOT EXISTS `jobor_task_owners` (
    `task_id` bigint(20) NOT NULL COMMENT '主键id',
    `user_id` bigint(20) NOT NULL COMMENT '主键id',
    PRIMARY KEY (`task_id`,`user_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_task_owners
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jobor_worker
-- ----------------------------
-- DROP TABLE IF EXISTS `jobor_worker`;
CREATE TABLE IF NOT EXISTS `jobor_worker` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `hostname` varchar(128) NOT NULL COMMENT 'worker hostname',
    `ip` varchar(128) NOT NULL COMMENT 'worker ip',
    `addr` varchar(64) NOT NULL COMMENT 'worker ip:port',
    `version` varchar(32) DEFAULT NULL COMMENT '''版本''',
    `routing_key` varchar(64) DEFAULT 'default' COMMENT 'routing_key',
    `weight` int(11) DEFAULT NULL COMMENT '权重',
    `lease_update` bigint(20) DEFAULT NULL COMMENT '租约时间更新',
    `status` varchar(32) DEFAULT 'offline' COMMENT '状态：running,stop,offline',
    `port` int(11) DEFAULT '22' COMMENT '端口',
    `mode` varchar(8) DEFAULT 'agent' COMMENT '模式[agent,ssh]',
    `auth_mode` varchar(8) DEFAULT 'key' COMMENT '认证模式[password,pub_key]',
    `username` varchar(152) DEFAULT NULL COMMENT '认证用户',
    `password` varchar(255) DEFAULT NULL COMMENT '认证密码',
    `rsa` text  COMMENT 'SSH认证私钥',
    `private` varchar(255) DEFAULT NULL COMMENT 'SSH key秘钥',
    `description` mediumtext  COMMENT '描述',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ip` (`ip`),
    KEY `idx_code` (`hostname`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of jobor_worker
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `jobor_worker` (`id`,`hostname`, `ip`, `addr`, `version`, `routing_key`, `weight`, `lease_update`, `status`, `port`, `mode`, `auth_mode`, `username`, `password`, `rsa`, `private`, `description`) VALUES (1, 'sshmode-demo', '10.1.1.1', '10.1.1.1:22', '', 'ssh', 100, 0, 'running', 22, 'ssh', 'pub_key', 'ops', '', 'xxxx', '', 'sshmode-demo');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
-- DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `title` varchar(128) NOT NULL COMMENT '名称',
    `name` varchar(128) NOT NULL COMMENT '名称',
    `description` longtext COMMENT '描述',
    `parent_id` bigint(20) DEFAULT NULL COMMENT '父节点',
    `path` text COMMENT '节点路径',
    `sort` bigint(20) DEFAULT '1000' COMMENT '''排序''',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`),
    KEY `parent_id` (`parent_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `role` (`id`, `title`, `name`, `description`, `parent_id`, `path`, `sort`) VALUES (1, '超级管理员', 'admin', '超级管理员', 0, '[1]', 100);
INSERT IGNORE INTO `role` (`id`, `title`, `name`, `description`, `parent_id`, `path`, `sort`) VALUES (2, '系统管理员', 'sys_rw', '系统管理员: 管理用户、角色、接口、审计日志', 0, '[2]', 100);
INSERT IGNORE INTO `role` (`id`, `title`, `name`, `description`, `parent_id`, `path`, `sort`) VALUES (3, '定时任务管理员', 'task_rw', '定时任务管理员: 管理执行节点、定时任务、定时任务执行记录', 0, '[3]', 100);
COMMIT;

-- ----------------------------
-- Table structure for role_api
-- ----------------------------
-- DROP TABLE IF EXISTS `role_api`;
CREATE TABLE IF NOT EXISTS `role_api` (
    `role_id` bigint(20) NOT NULL COMMENT '主键id',
    `api_id` bigint(20) NOT NULL COMMENT '主键id',
    PRIMARY KEY (`role_id`,`api_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role_api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
-- DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `nickname` varchar(128) DEFAULT NULL COMMENT '显示名',
    `username` varchar(128) DEFAULT NULL COMMENT '用户名',
    `password` varchar(256) DEFAULT NULL COMMENT '密码',
    `empno` varchar(156) DEFAULT NULL COMMENT '员工号',
    `userid` varchar(156) DEFAULT NULL COMMENT 'UserId',
    `email` varchar(156) DEFAULT NULL COMMENT '邮箱',
    `phone` varchar(64) DEFAULT NULL COMMENT '电话',
    `user_type` varchar(16) DEFAULT 'local' COMMENT '用户类型:local,ldap,sso',
    `path` text COMMENT '节点路径',
    `avatar` varchar(64) DEFAULT 'default' COMMENT '用户头像',
    `status` varchar(32) DEFAULT '1' COMMENT '状态',
    `by_update` varchar(64) DEFAULT NULL COMMENT '更新人',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_name_code` (`nickname`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `user` (`id`, `nickname`, `username`, `password`, `empno`, `userid`, `email`, `phone`, `user_type`, `path`, `avatar`, `status`) VALUES (1, 'admin', 'admin', '$2a$10$KPq2KbyMbsBzcFRxuC7XI.6ic5XGALvBDUpCT23HxXClmKXlbFMjG', '', '', 'admin@example.com', '', 'local', NULL, 'default', '1');
INSERT IGNORE INTO `user` (`id`, `nickname`, `username`, `password`, `empno`, `userid`, `email`, `phone`, `user_type`, `path`, `avatar`, `status`) VALUES (2, 'root', 'root', '$2a$10$O9GeHneK.st.qXkrZhX5auwWUxbJ7hWOJSN7zQsD.nhlQSECYYYQG', 'root@example.cn', '', NULL, '', 'local', '[[1]]', '', '1');
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
-- DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE IF NOT EXISTS `user_roles` (
    `user_id` bigint(20) NOT NULL COMMENT '主键id',
    `role_id` bigint(20) NOT NULL COMMENT '主键id',
    PRIMARY KEY (`user_id`,`role_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT IGNORE INTO `user_roles` (`user_id`, `role_id`) VALUES (1, 1);
INSERT IGNORE INTO `user_roles` (`user_id`, `role_id`) VALUES (2, 1);
COMMIT;

-- SET FOREIGN_KEY_CHECKS = 1;
