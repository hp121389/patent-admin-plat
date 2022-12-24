-- 开始初始化数据 ;
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

INSERT INTO sys_casbin_rule VALUES (1, 'p', 'user', '/api/v1/user-agent/*', '*', '', '', '', '', '');

INSERT INTO sys_user VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO sys_user VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);

-- 数据完成 ;