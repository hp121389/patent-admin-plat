-- 开始初始化数据 ;
INSERT INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO sys_role VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);

INSERT INTO sys_casbin_rule VALUES (1, 'p', 'user', '/api/v1/user-agent/*', '*', '', '', '', '', '');

INSERT INTO sys_user VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO sys_user VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);

-- user-patent
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`) VALUES (1, 1, 1, '关注', 1, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`) VALUES (2, 2, 1, '关注', 1, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `type`, `create_by`, `update_by`) VALUES (3, 3, 1, '认领', 1, 0);

-- tag
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'good', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'go', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'java', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'c', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` (`tag_id`, `tag_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'py', NULL, NULL, NULL, NULL, NULL, NULL);


-- patent-tag

INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (3, 3, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (4, 4, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (5, 5, 1, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (6, 6, 5, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (7, 7, 5, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (8, 7, 4, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (9, 1, 3, NULL, NULL);
INSERT INTO `patent_tag` (`id`, `patent_id`, `tag_id`, `create_by`, `update_by`) VALUES (10, 2, 1, 0, 0);

-- patent

INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`) VALUES (1, '101', '{\"patentId\":1,\"TI\":\"string\",\"PNM\":\"101\",\"AD\":\"20220101\",\"PD\":\"20220103\",\"CL\":\"a patent of MessageHidden\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author001\",\"CLS\":\"a patent of MessageHidden\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`) VALUES (2, '102', '{\"patentId\":2,\"TI\":\"string\",\"PNM\":\"102\",\"AD\":\"20220101\",\"PD\":\"20220104\",\"CL\":\"a patent of MemoryFramework\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author002\",\"CLS\":\"a patent of MemoryFramework\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`) VALUES (3, '103', '{\"patentId\":3,\"TI\":\"string\",\"PNM\":\"103\",\"AD\":\"20220101\",\"PD\":\"20220105\",\"CL\":\"a patent of TwoSoftware\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author003\",\"CLS\":\"a patent of TwoSoftware\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`) VALUES (4, '104', '{\"patentId\":4,\"TI\":\"string\",\"PNM\":\"104\",\"AD\":\"20220101\",\"PD\":\"20220106\",\"CL\":\"a patent of ThreeSoftware\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author004\",\"CLS\":\"a patent of ThreeSoftware\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`) VALUES (5, '105', '{\"patentId\":5,\"TI\":\"string\",\"PNM\":\"105\",\"AD\":\"20220101\",\"PD\":\"20220107\",\"CL\":\"a patent of T0Software\",\"PA\":\"BUPT\",\"AR\":\"Beijing\",\"PINN\":\"author005\",\"CLS\":\"a patent of T0Software\",\"CreateBy\":1,\"UpdateBy\":0}', 0, 0);

--package

INSERT INTO `package` (`package_id`, `package_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'golangSoftWarePatents', 'patents of Golang software', 1, 0, '2022-10-18 18:49:24.070', NULL);
INSERT INTO `package` (`package_id`, `package_name`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'CPPSoftWarePatents', 'patents of CPP software', 1, 0, '2022-10-18 18:49:53.081', '2022-10-18 18:49:53.081', NULL);

-- 数据完成 ;