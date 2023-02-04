-- 开始初始化数据 ;
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

INSERT INTO sys_casbin_rule VALUES (1, 'p', 'user', '/apis/v1/user-agent/*', '*', '', '', '', '', '');

INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'user2', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user2', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'user3', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user3', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'user4', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user4', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);

INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (1, '北邮网安六组', '北京邮电大学网络空间安全学院六组', '2022-10-18 18:49:24', NULL, NULL, NULL);
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `created_at`, `updated_at`, `create_by`, `update_by`) VALUES (2, '北邮网安七组', '北京邮电大学网络空间安全学院六组', '2022-10-18 18:49:24', NULL, NULL, NULL);

-- 数据完成 ;