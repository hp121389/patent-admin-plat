/*
 Navicat Premium Data Transfer

 Source Server         : 10.112.138.178
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : dbname

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 12/01/2023 17:06:00
*/
USE dbname;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '团队ID(主键)',
  `dept_name` longtext COMMENT '团队名称',
  `dept_properties` longtext COMMENT '团队详情',
  `research_interest` longtext COMMENT '研究方向',
  `created_at` longtext COMMENT '创建时间',
  `updated_at` longtext COMMENT '最后更新时间',
  `dept_status` longtext COMMENT '团队状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`dept_id`),
  KEY `idx_dept_create_by` (`create_by`),
  KEY `idx_dept_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (1, '北邮网安六组', '北京邮电大学网络空间安全学院六组', '隐写分析，加密算法，NLP', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);
INSERT INTO `dept` (`dept_id`, `dept_name`, `dept_properties`, `research_interest`, `created_at`, `updated_at`, `dept_status`, `create_by`, `update_by`) VALUES (2, '北邮网安七组', '北京邮电大学网络空间安全学院六组', '网络攻防', '2022-10-18 18:49:24', NULL, '存在', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for dept_rela
-- ----------------------------
DROP TABLE IF EXISTS `dept_rela`;
CREATE TABLE `dept_rela` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint DEFAULT NULL COMMENT '成员ID',
  `dept_id` bigint DEFAULT NULL COMMENT '部门ID',
  `mem_type` longtext,
  `created_at` longtext COMMENT '创建时间',
  `updated_at` longtext COMMENT '最后更新时间',
  `mem_status` longtext COMMENT '成员状态',
  `examine_status` longtext COMMENT '审核状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_dept_rela_create_by` (`create_by`),
  KEY `idx_dept_rela_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of dept_rela
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for package
-- ----------------------------
DROP TABLE IF EXISTS `package`;
CREATE TABLE `package` (
  `package_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
  `package_name` varchar(128) DEFAULT NULL COMMENT '专利包',
  `desc` varchar(128) DEFAULT NULL COMMENT '描述',
  `files` longtext COMMENT '专利包附件',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`package_id`),
  KEY `idx_package_create_by` (`create_by`),
  KEY `idx_package_update_by` (`update_by`),
  KEY `idx_package_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of package
-- ----------------------------
BEGIN;
INSERT INTO `package` (`package_id`, `package_name`, `desc`, `files`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '工艺包测试', '', '[{\"FileName\":\"基于生成对抗网络的系统日志级异常检测算法_夏彬.pdf\",\"FilePath\":\"10.112.138.178:8000/static/uploadfile/e8f027fe-c441-4f11-bb2b-f262d9e8c2bd.基于生成对抗网络的系统日志级异常检测算法_夏彬.pdf\"},{\"FileName\":\"(CNN)Detecting Anomaly in Big Data System Logs Using Convolutional Neural Network.pdf\",\"FilePath\":\"10.112.138.178:8000/static/uploadfile/d652a0ec-0df9-4672-906a-3fc1263a594e.(CNN)Detecting Anomaly in Big Data System Logs Using Convolutional Neural Network.pdf\"}]', 1, 0, '2023-01-12 16:44:39.252', '2023-01-12 16:48:48.233', NULL);
INSERT INTO `package` (`package_id`, `package_name`, `desc`, `files`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '工艺包测试B', '', '[{\"FileName\":\"Anomaly Detection in Operating System Logs with Deep Learning-based Sentiment Analysis.pdf\",\"FilePath\":\"10.112.138.178:8000/static/uploadfile/ae03fd26-4757-4f49-b695-fbe8a5becf79.Anomaly Detection in Operating System Logs with Deep Learning-based Sentiment Analysis.pdf\"},{\"FileName\":\"面向云数据中心多语法日志通用异常检测机制_张圣林.pdf\",\"FilePath\":\"10.112.138.178:8000/static/uploadfile/18f1e8b7-a5eb-49b1-a4eb-137170d73ada.面向云数据中心多语法日志通用异常检测机制_张圣林.pdf\"}]', 1, 0, '2023-01-12 16:45:37.083', '2023-01-12 16:49:05.725', NULL);
COMMIT;

-- ----------------------------
-- Table structure for patent
-- ----------------------------
DROP TABLE IF EXISTS `patent`;
CREATE TABLE `patent` (
  `patent_id` bigint NOT NULL AUTO_INCREMENT COMMENT '专利ID(主键)',
  `pnm` varchar(128) DEFAULT NULL COMMENT '申请号',
  `patent_properties` longtext COMMENT '专利详情',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`patent_id`),
  KEY `idx_patent_create_by` (`create_by`),
  KEY `idx_patent_update_by` (`update_by`),
  KEY `idx_patent_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of patent
-- ----------------------------
BEGIN;
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'CN114662096A', '{\"patentId\":0,\"TI\":\"一种基于图核聚类的威胁狩猎方法\",\"PNM\":\"CN114662096A\",\"AD\":\"2022.03.25\",\"PD\":\"2022.06.24\",\"CL\":\" 1.一种基于图核聚类的威胁狩猎方法，其特征在于，包括：A、构建行为依赖图：从审计日志中抽取实体和实体之间的关系构建行为依赖图，并基于密度对长时间运行的进程进行分区，其中行为依赖图是带标签的有向图，节点表示系统层实体，有向边表示实体之间的操作关系；B、相似度计算与聚类：针对行为依赖图的特点，设计图核方法将行为依赖图嵌入到高维空间，并在高维空间中计算图之间的相似度得到图核矩阵，其中，图核矩阵可以看作相似度矩阵，再利用聚类方法对图核矩阵进行分析，实现正常行为与异常行为的分离；C、威胁评估：通过类簇中图的数量判断哪些类簇包含异常行为，对异常行为图进行威胁评估量化，计算异常行为图的威胁值，最终实现威胁狩猎。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"李家威\",\"CLS\":\"公开\",\"INN\":\"李家威;程杰;张茹;刘建毅;高雅婷;王婵;夏昂;崔博;孔汉章;LI JIAWEI;CHENG JIE;ZHANG RU;LIU JIANYI;GAO YATING;WANG CHAN;XIA ANG;CUI BO;KONG HANZHANG\",\"IDX\":39,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:01.579', '2023-01-12 16:42:01.579', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'CN114238977A', '{\"patentId\":0,\"TI\":\"一种融合efficient-net和1d-cnn的恶意代码家族分类方法\",\"PNM\":\"CN114238977A\",\"AD\":\"2021.12.23\",\"PD\":\"2022.03.25\",\"CL\":\" 1.一种融合efficient-net和1d-cnn的恶意代码家族分类方法，其特征在于，包括：A、恶意代码宽度自适应图像化：对恶意代码PE文件进行反汇编得到asm文件，提取.text字段和.rdata字段，采用一种自适应宽度的方法进行图像化处理；B、将恶意代码图像输入基于efficient-net网络的训练模型：将恶意代码图像缩放到固定尺寸的大小，输入efficient-net网络；C、将恶意代码图像输入基于1d-cnn的训练模型：输入网络的图像被降维为1维后，利用嵌入算法拓展维度，输入1d-cnn网络；D、将两个模型进行特征融合：将efficient-net各个阶段的特征进行处理后，与1d-cnn的不同特征图进行融合，联合这些特征进行家族分类。\",\"PA\":\"国家电网有限公司\",\"AR\":\"北京市西城区西长安街86号\",\"PINN\":\"刘莹\",\"CLS\":\"公开\",\"INN\":\"刘莹;胡威;黄星杰;高雅婷;李显旭;黄华;盛华;张茹;种旭磊;刘建毅;LIU YING;HU WEI;HUANG XINGJIE;GAO YATING;LI XIANXU;HUANG HUA;SHENG HUA;ZHANG RU;CHONG XU;LIU JIANYI\",\"IDX\":39,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:03.144', '2023-01-12 16:42:03.144', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'CN114239737A', '{\"patentId\":0,\"TI\":\"一种基于时空特征与双层注意力的加密恶意流量检测方法\",\"PNM\":\"CN114239737A\",\"AD\":\"2021.12.21\",\"PD\":\"2022.03.25\",\"CL\":\" 1.一种基于时空特征与双层注意力的加密恶意流量检测方法，其特征在于，包括以下步骤：A、加密流量的提取与预处理：收集网卡节点处的原始流量，从原始pcap文件中提取加密流量双向会话流，经过删除以太网头、mask IP地址、传输层包头对齐、对齐数据包等操作，得到流量矩阵X；B、数据包内恶意特征提取层：通过一维卷积神经网络学习每个加密网络流量包的空间特征A，再通过基于数据包字段的注意力机制层，在每个数据包内提取重要的恶意特征P；C、数据流内恶意特征提取层：通过BiGRU学习数据包之间前向和后向的时间关联特征H，再通过基于流中数据包的注意力机制层，通过软注意力机制提取数据流中的重要的恶意特征F，通过上述步骤提取的最终特征通过分类器，得到最终的恶意检测结果。\",\"PA\":\"国家电网有限公司信息通信分公司\",\"AR\":\"北京市西城区白广路二条一号\",\"PINN\":\"胡威\",\"CLS\":\"公开\",\"INN\":\"胡威;庞进;王景初;张亚昊;尹红珊;张茹;王岚婷;刘建毅;陈连栋;程凯;HU WEI;PANG JIN;WANG JINGCHU;ZHANG YAHAO;YIN HONGSHAN;ZHANG RU;WANG LANTING;LIU JIANYI;CHEN LIANDONG;CHENG KAI\",\"IDX\":48,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:05.354', '2023-01-12 16:42:05.354', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'CN114241581A', '{\"patentId\":0,\"TI\":\"一种基于语义特征的人脸生成式图像隐写术\",\"PNM\":\"CN114241581A\",\"AD\":\"2021.12.31\",\"PD\":\"2022.03.25\",\"CL\":\" 1.一种基于语义特征的人脸生成式图像隐写术，其特征在于，包括：A、生成载密语义特征：人脸语义特征生成网络由全连接网络实现，该网络负责由高斯噪声生成人脸语义特征S＝{S 1 ,S 2 ,...S n }，将秘密信息s＝{s 1 ,s 2 ,...s m }采用采用奇偶特征差值算法嵌入生成的语义特征，并得到载密语义特征S′＝{S′ 1 ,S′ 2 ,...S′ n }；B、生成载密图像：人脸生成模型采用StarGAN v2网络实现，负责将嵌入秘密信息的语义特征S′＝{S′ 1 ,S′ 2 ，...S′ n }输入人脸合成模型中得到载密人脸图像；C、提取秘密信息：语义提取网络由卷积神经网络实现，负责提取载密人脸图像中的语义特征S″＝{S″ 1 ,S″ 2 ,...S″ n }，采用奇偶特征差值算法从语义特征中恢复秘密信息；D、提取网络鲁棒训练：通过图像变换载密数据集对提取网络进行训练，提升特征提取模型的鲁棒性。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"孙佑强\",\"CLS\":\"公开\",\"INN\":\"孙佑强;张茹;刘建毅;唐球;SUN YOUQIANG;ZHANG RU;LIU JIANYI;TANG QIU\",\"IDX\":39,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:07.921', '2023-01-12 16:42:07.921', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'CN113486190A', '{\"patentId\":0,\"TI\":\"一种融合实体图像信息和实体类别信息的多模态知识表示方法\",\"PNM\":\"CN113486190A\",\"AD\":\"2021.06.21\",\"PD\":\"2021.10.08\",\"CL\":\" 1.一种融合实体图像信息和实体类别信息的多模态知识表示方法，其特征在于，包括：A、实体图像信息的嵌入方法：通过设计图编码器来完成实体图像特征信息的抽取以及从图像空间到知识空间的转换，利用注意力机制来对图像信息进行筛选组合，使用图像特征和实体以及对应关系特征的相关性大小作为注意力分数计算依据，构建实体基于图像的表示；B、实体类别信息的嵌入方法：通过注意力机制对实体类别和对应三元组关系的语义联系进行建模，构建实体基于类别下的表示；C、融合实体图像信息和实体类别信息的多模态图注意力网络知识表示：将实体类别信息，将实体结构特征、实体图像特征和实体类别特征结合起来，使用GAT模型进行训练，实现多模态知识表示模型的构建。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"刘建毅\",\"CLS\":\"实审\",\"INN\":\"刘建毅;张茹;李萌;吕智帅;LIU JIANYI;ZHANG RU;LI MENG;Lv Zhishuai\",\"IDX\":48,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:09.659', '2023-01-12 16:42:09.659', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'CN113486932A', '{\"patentId\":0,\"TI\":\"一种面向卷积神经网络隐写分析的优化方法\",\"PNM\":\"CN113486932A\",\"AD\":\"2021.06.21\",\"PD\":\"2021.10.08\",\"CL\":\" 1.一种面向卷积神经网络隐写分析的优化方法，其特征在于，包括：A、采用非线性的t-sne降维算法可视化同类隐写样本的类内聚集度，即高维点集：X{x 1 ,x2,……,x n }映射到低维点集Y{y 1 ,y 2 ,……,y n }；B、采用平均变异系数衡量卷积神经网络隐写分析算法特征表达能力；C、计算各种算法的不同特征的变异系数值，根据变异系数调整特征集，优化隐写分析性能。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"实审\",\"INN\":\"张茹;邹盛;刘建毅;田思远;ZHANG RU;ZOU SHENG;LIU JIANYI;TIAN SIYUAN\",\"IDX\":39,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:12.115', '2023-01-12 16:42:12.115', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'CN113158390A', '{\"patentId\":0,\"TI\":\"一种基于辅助分类式生成对抗网络的网络攻击流量生成方法\",\"PNM\":\"CN113158390A\",\"AD\":\"2021.04.29\",\"PD\":\"2021.07.23\",\"CL\":\" 1.一种基于辅助分类式生成对抗网络的网络攻击流量生成方法，其特征在于，包括：A、多源异构网络流量融合：将不同的网络包格式如PCAP格式，NETFLOW格式，CFLOW格式，JFLOW格式及SFLOW格式的数据文件进行特征提取和统一的定义与标注来定义一种通用的数据格式，将统一格式后的数据用于生成模型的训练与数据的生成；B、网络攻击流量生成模型训练：定义辅助分类式生成对抗网络所需的流量生成器与流量判别器的网络结构以及辅助分类式生成对抗网络所需的生成损失函数与分类损失函数及训练方法；C、生成模型的分类微调：对上一步生成的网络攻击流量样本进行进一步的验证与微调，以此来调试生成模型生成特定攻击类型流量样本的性能。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"实审\",\"INN\":\"张茹;吕智帅;刘建毅;胡威;李静;曲延盛;王婵;ZHANG RU;Lv Zhishuai;LIU JIANYI;HU WEI;LI JING;QU YANSHENG;WANG CHAN\",\"IDX\":54,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:14.146', '2023-01-12 16:42:14.146', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'CN113132414A', '{\"patentId\":0,\"TI\":\"一种多步攻击模式挖掘方法\",\"PNM\":\"CN113132414A\",\"AD\":\"2021.05.08\",\"PD\":\"2021.07.16\",\"CL\":\" 1.一种多步攻击模式挖掘方法，其特征在于，包括：A、敏感信息流量筛选与数据归一化：从海量流量数据中基于spark框架筛选敏感信息并根据杀伤链模型进行归一化；B、敏感信息与告警日志融合算法：针对告警日志的误报和漏报性质，将从流量数据中筛选出的敏感信息和告警日志通过IP相似度聚簇、攻击簇内合并和过滤、攻击簇间筛选三种算法进行融合；C、多步攻击模型：多步攻击模型定义如下 其中N表示某类攻击的实际攻击过程步数，ABC代表多步攻击中每一个单步攻击的属性特征值；D、启发式多步攻击模型生成和攻击预测算法：通过图的概率匹配达到针对多步攻击的预测，步骤包括匹配对应点、计算概率值、生成多步攻击图模型、衡量转换。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"刘建毅\",\"CLS\":\"授权\",\"INN\":\"刘建毅;田思远;张茹;胡威;程杰;陈连栋;高雅婷;LIU JIANYI;TIAN SIYUAN;ZHANG RU;HU WEI;CHENG JIE;CHEN LIANDONG;GAO YATING\",\"IDX\":50,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:16.127', '2023-01-12 16:42:16.127', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'CN113094860A', '{\"patentId\":0,\"TI\":\"一种基于注意力机制的工控网络流量建模方法\",\"PNM\":\"CN113094860A\",\"AD\":\"2021.04.29\",\"PD\":\"2021.07.09\",\"CL\":\" 1.一种基于注意力机制的工控网络流量建模方法，其特征在于，包括以下步骤：A、数据采集以及预处理：针对工控网络中上位机与下位机之间的通信，即过程监控层包括工程师站、操作员站、数据服务器等设备，与现场控制层设备控制器之间的通信并进行流量采集工作，生成工控网络流量矩阵M，并进行归一化处理。B、模型构建：通过二维卷积神经网络提取工控网络流量中的空间特征M s ，再通过BiLSTM提取时间特征M t ，最后通过多头注意力机制提取工控网络流量的重要特征M i ，通过全连接层完成对工控网络流量的建模。C、迭代和训练：使用归一化均方根误差(RNMSE)作为损失函数，并通过计算梯度，反向传播更新参数，反复迭代和训练得到基于注意力机制的深度学习模型。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"实审\",\"INN\":\"张茹;王岚婷;刘建毅;胡威;庞进;袁国泉;史睿;ZHANG RU;WANG LANTING;LIU JIANYI;HU WEI;PANG JIN;YUAN GUOQUAN;SHI RUI\",\"IDX\":54,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:18.282', '2023-01-12 16:42:18.282', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'CN112073362A', '{\"patentId\":0,\"TI\":\"一种基于流量特征的APT组织流量识别方法\",\"PNM\":\"CN112073362A\",\"AD\":\"2020.06.19\",\"PD\":\"2020.12.11\",\"CL\":\" 1.一种基于流量特征的APT组织流量识别方法，其特征在于，包括：A、提取常见特征：以时间窗口为截断周期，记录该周期内最早出现的同IP、同域名查询包时间，作为输出DNS流量特征序列的时间戳，提取Alexa排名(Alexa_score)、端口异常(Port_abnormal)等DNS常见特征；利用会话窗口聚类提取上行包数目(Upload_num)、上行包负载(Upload_load)等TCP、HTTP/HTTPS常用特征。B、针对DNS提取APT组织流量特征：提出并定义APT组织特征Response_type和包负载波动特征C2Load_fluct。Response_type特征通过离散DNS响应包中的record type字段获取，C2Load_fluct特征利用定义公式 计算时间窗口内同一源IP、域名的流量包簇在单位域名下的平均负载量。C、针对TCP、HTTP/HTTPS提取APT组织流量特征：提出并定义TCP、HTTP/HTTPS流量中乱序包和重传包相似性特征Bad_rate，该特征可以反应APT组织恶意流量产生时的网络状态，利用公式 计算得到。D、利用组织特征训练分类模型：利用组织特征和样本数据，通过标记样本，根据标记结果更新样本权重，并重构基学习器再次标注，迭代训练，从而训练基于决策树模型的AdaBoost分类器。E、使用分类器进行分类：使用训练好的AdaBoost分类器对网络流量数据进行分类，生成带有分类标识信息、带标签的流量特征序列，实现APT组织流量的识别。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"刘建毅\",\"CLS\":\"授权\",\"INN\":\"刘建毅;张茹;李静;程杰;王婵;郭邯;孙文新;闫晓帆;LIU JIANYI;ZHANG RU;LI JING;CHENG JIE;WANG CHAN;GUO HAN;SUN WENXIN;YAN XIAOFAN\",\"IDX\":54,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:21.482', '2023-01-12 16:42:21.482', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'CN111881935A', '{\"patentId\":0,\"TI\":\"一种基于内容感知GAN的对抗样本生成方法\",\"PNM\":\"CN111881935A\",\"AD\":\"2020.06.19\",\"PD\":\"2020.11.03\",\"CL\":\" 1.一种基于内容感知GAN的对抗样本生成方法，其特征在于，包括：A、基于WGAN_GP的生成对抗网络进行对抗样本的生成工作，通过使用两个不同目标的无监督训练阶段，正常训练阶段学习正常样本分布，对抗性训练部分学习对抗样本的分布，使GAN模型能够从随机噪声中学习对抗样本的分布，批量生成不受限的对抗样本，对目标模型进行对抗攻击；B、正常训练部分：使用噪声z作为生成器输入，生成样本G(z)和真实样本x作为判别器输入，初始化生成器G和判别器D，使用WGAN_GP原始损失函数L GAN 作为目标函数，每轮训练完成后更新生成器G和判别器D的参数，获得学习到正常样本分布的生成器和判别器；C、对抗性训练部分：在正常训练部分得到的生成器和判别器的基础上，使生成器能够从噪声z中学习到对抗样本的分布，在继续优化WGAN_GP损失L GAN 的前提下，新增加目标模型f、扰动评估部分以及特征提取网络N feature ，组成模型的对抗性训练结构，在生成对抗样本时保持内容特征尽可能不变；D、通过内容特征约束来生成高质量的对抗样本，定义图像x和图像的内容特征x content ，基于CNN的特征提取能力，借助内容特征提取网络N feature 对生成样本的语义信息进行约束，引入了新的样本质量约束损失函数L content ，改进基础攻击模型的对抗性训练流程，在不影响攻击效果的前提下提高对抗样本的质量，降低对人类的可感知程度。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"刘建毅\",\"CLS\":\"实审\",\"INN\":\"刘建毅;张茹;田宇;李娟;李婧雯;LIU JIANYI;ZHANG RU;TIAN YU;LI JUAN;李娟�;LI JINGWEN\",\"IDX\":67,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:28.798', '2023-01-12 16:42:28.798', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 'CN111861844A', '{\"patentId\":0,\"TI\":\"一种基于图像分块认证的可逆水印方法\",\"PNM\":\"CN111861844A\",\"AD\":\"2020.06.19\",\"PD\":\"2020.10.30\",\"CL\":\" 1.一种基于图像分块认证的可逆水印方法，其特征在于，包括：A、像素预测：将载体图像的最高有效位MSB位视为冗余空间，使用梯度调整算法(GAP)遍历图像找到可用的嵌入像素位，并将预测错误的像素位置写入误差矩阵；B、秘密信息嵌入：待嵌入的秘密信息经隐藏密钥加密，将预测误差位的信息与秘密信息组合并构造新的嵌入信息，将新的信息嵌入载体图像中，生成隐藏有秘密信息的载密图像，其大小与原始图像相同；C、脆弱水印嵌入：将基于频域的自恢复水印嵌入到隐藏有秘密信息的载密图像中，得到既包含秘密信息又包含脆弱水印的双模水印图片；D、篡改定位及还原：接收端收到载密图像后，提取脆弱水印的信息，判断图像是否发生篡改并进行恢复；E、秘密信息提取：载密图像恢复过后，提取嵌入的秘密信息，进行秘密信息的恢复；F、载体重建：使用恢复的预测错误像素位信息，结合预测可逆地恢复原始载体图像。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"实审\",\"INN\":\"张茹;刘建毅;胡威;郭邯;黄星杰;尚智婕;赵凯峰;李依;ZHANG RU;LIU JIANYI;HU WEI;GUO HAN;HUANG XINGJIE;SHANG ZHIJIE;ZHAO KAIFENG;LI YI\",\"IDX\":63,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:32.940', '2023-01-12 16:42:32.940', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'CN111861845A', '{\"patentId\":0,\"TI\":\"一种基于阈值分割和直方图均衡的可逆水印方法\",\"PNM\":\"CN111861845A\",\"AD\":\"2020.06.19\",\"PD\":\"2020.10.30\",\"CL\":\" 1.一种基于阈值分割和直方图均衡的可逆水印方法，其特征在于，包括：A、水印嵌入阶段：使用阈值分割算法对图像的梯度幅值矩阵进行处理，从而获取图像前景区的轮廓，并辅以横纵双向扫描确定整个前景区域，将其作为水印嵌入区域，使用直方图移位技术，对前景区的直方图中像素值介于峰值点和零值点之间的直方柱进行移位，实现直方图均衡，在移位的同时，进行水印嵌入，为了在之后实现信息提取和载体无失真恢复，少量的附加信息随水印一同嵌入，随着嵌入循环次数的增加，每次循环中嵌入的附加信息逐渐增长，当图像像素为峰值点的数量小于附加信息的长度时，嵌入循环结束；B、水印提取、载体无失真恢复阶段：由载密图像获取水印提取、载体无失真阶段首轮循环的峰值点和零值点，利用峰值点和零值点从载密图像的前景区提取信息，利用提取到的信息进行图像恢复，提取到的信息除去附加信息便是本轮所提取的水印，进行多次循环提取，当某一次提取信息的附加信息显示为最后一轮，提取终止。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"授权\",\"INN\":\"张茹;刘建毅;程杰;尚智婕;庞进;王婵;王悦;李晓丽;ZHANG RU;LIU JIANYI;CHENG JIE;SHANG ZHIJIE;PANG JIN;WANG CHAN;WANG YUE;LI XIAOLI\",\"IDX\":52,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:35.515', '2023-01-12 16:42:35.515', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 'CN110278210A', '{\"patentId\":0,\"TI\":\"一种云环境下基于属性的可扩展视频数据访问控制方法\",\"PNM\":\"CN110278210A\",\"AD\":\"2019.06.24\",\"PD\":\"2019.09.24\",\"CL\":\" 1.一种云环境下基于属性的可扩展视频数据访问控制方法，其特征在于：包括如下步骤，第一步，密钥机构系统初始化：密钥机构运行Setup算法，构造阶为素数p的双线性群 和 的生成元为g，对应的双线性映射为e: 定义散列函数H 1 : 随机选择α,β, 其中 表示模p的整数集合；公开发布系统公钥 然后，密钥机构生成系统主密钥MK＝(g α ,β,λ)并秘密保存；第二步，密钥生成：密钥机构运行KeyGen算法，选择随机数 并为属性集合S中的每个属性x∈S选择随机数 生成私钥 第三步，数据加密：对于SVC编码视频m中的每一视频层w i,j ，数据贡献者选择随机的对称密钥sk i,j ，先对该视频层运行对称加密SE算法，得到密文头部Hdr i,j ＝SE(w i,j ,sk i,j )；然后，数据贡献者运行加密算法Encrypt算法，使用访问策略T加密对称密钥sk i,j ；第四步，数据解密；云媒体中心将加密的视频数据分发给不同的用户。如果用户的属性满足相应的访问树，则可以运行解密算法Decrypt算法来解密出一定数量的视频层。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"黄勤龙\",\"CLS\":\"授权\",\"INN\":\"黄勤龙;张志成;杨义先;HUANG QINLONG;ZHANG ZHICHENG;YANG YIXIAN;Huang qin long;Zhang zhi cheng;Yang Yi Xian\",\"IDX\":60,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:37.988', '2023-01-12 16:42:37.988', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 'CN109617682A', '{\"patentId\":0,\"TI\":\"一种基于直方图左右移位的密文域可逆信息隐藏方法\",\"PNM\":\"CN109617682A\",\"AD\":\"2018.12.12\",\"PD\":\"2019.04.12\",\"CL\":\" 1.一种基于直方图左右移位的密文域可逆信息隐藏方法，其特征在于，所述方法包括：A、加密阶段：内容拥有者根据加密密钥对载体进行加性同态加密和置乱加密，生成密文载体并发送到信息隐藏者；B、信息隐藏阶段：信息隐藏者根据嵌入密钥，通过预测误差直方图左右移位进行信息嵌入，生成载密密文载体并传送给有效接收者；C、信息提取和载体解密阶段：存在两种情况，情况一：有效接收者使用加密密钥对载密密文载体直接解密，得到与原始载体高度相似的载体；情况二：有效接收者使用加密密钥和嵌入密钥，无失真获取密文信息和原始载体。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"授权\",\"INN\":\"张茹;刘建毅;卢春景;ZHANG RU;LIU JIANYI;LU CHUNJING\",\"IDX\":69,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:40.380', '2023-01-12 16:42:40.380', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 'CN109587372A', '{\"patentId\":0,\"TI\":\"一种基于生成对抗网络的不可见图像隐写术\",\"PNM\":\"CN109587372A\",\"AD\":\"2018.12.11\",\"PD\":\"2019.04.05\",\"CL\":\" 1.一种基于生成对抗网络的不可见图像隐写术，其特征在于，包括：A、生成载密图像：编码器网络由卷积神经网络实现，负责将灰度秘密图像嵌入到彩色载体图像中，生成彩色载密图像；B、恢复出秘密图像：解码器网络由全卷积神经网络实现，负责从彩色载密图像中恢复出灰度秘密图像；C、使用判别器网络进行隐写分析：判别器网络对输入的自然图像或编码器网络生成的载密图像进行隐写分析；D、参数更新与迭代训练：使用符合损失函数计算损失值，计算梯度，更新参数；E、验证模型性能与泛化训练：通过结构相似度指标验证模型性能，使用多尺度样本进行泛化训练。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"张茹\",\"CLS\":\"授权\",\"INN\":\"张茹;刘建毅;董士琪;ZHANG RU;LIU JIANYI;DONG SHIQI\",\"IDX\":73,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:42:42.928', '2023-01-12 16:42:42.928', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 'CN114896595A', '{\"patentId\":0,\"TI\":\"一种针对处理器指令集安全缺陷的隐藏指令检测技术\",\"PNM\":\"CN114896595A\",\"AD\":\"2022.04.19\",\"PD\":\"2022.08.12\",\"CL\":\" 1.权利要求1：一种针对处理器指令集安全缺陷的隐藏指令检测技术，其特征在于，包括以下步骤：步骤S1：基于指令格式分析的隐藏指令搜索算法。根据待测指令集的指令格式分析指令集中冗余的指令空间，然后对其进行剪枝。在内存执行的过程中，在确定待测指令的长度后，通过对执行环境的恢复，完成单一指令的执行；步骤S2：基于操作系统信号与反汇编器的指令分析。基于步骤S1中提出的单一指令执行，每条指令会得到来自操作系统的信号。与此同时将指令放入反汇编器中，并以两者的数据，判断其是否为隐藏指令；步骤S3：基于静态分析与动态执行的指令筛选机制。在步骤S2得到隐藏指令的基础上，针对其数据量大、结果冗余和误判的问题，提出了基于静态分析与动态执行的指令筛选机制，其中静态分析解决的是数据量大和结果冗余的问题，动态执行解决的是误判的问题。通过此次指令筛选可以将检测出的指令再次精确的分类；步骤S4：面向处理器全架构的隐藏指令检测方法。结合S1、S2、S3中的方法，建立面向处理器全架构的隐藏指令检测方法，并将最终的分析结果反馈至待测指令集。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"崔宝江\",\"CLS\":\"实审\",\"INN\":\"崔宝江;董任海;孙溢;CUI BAOJIANG;Dong Renhai;SUN YI\",\"IDX\":37,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:43:51.265', '2023-01-12 16:43:51.265', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 'CN115225551A', '{\"patentId\":0,\"TI\":\"一种模糊测试方法、装置、设备及存储介质\",\"PNM\":\"CN115225551A\",\"AD\":\"2022.07.14\",\"PD\":\"2022.10.21\",\"CL\":\" 1.一种模糊测试方法，其特征在于，应用于计算机设备，所述方法包括：基于遗传算法生成模糊测试用例，并执行所述模糊测试用例，以触发被监测目标程序的运行；获取所述被监测目标程序运行后的反馈结果，所述反馈结果包括：测试结果满足停止条件以及测试结果不满足停止条件；根据所述反馈结果确定目标测试用例；基于所述目标测试用例对待测程序进行测试。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"杨俊\",\"CLS\":\"实审\",\"INN\":\"杨俊;崔宝江;于博;王炳铨;巫俊杰;池晓峰;韩春阳;YANG JUN;杨俊�;CUI BAOJIANG;YU BO;WANG BINGQUAN;WU JUNJIE;CHI XIAOFENG;HAN CHUNYANG\",\"IDX\":44,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:43:53.394', '2023-01-12 16:43:53.394', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 'CN114679261A', '{\"patentId\":0,\"TI\":\"基于密钥派生算法的链上匿名通信方法和系统\",\"PNM\":\"CN114679261A\",\"AD\":\"2021.12.22\",\"PD\":\"2022.06.28\",\"CL\":\" 1.一种基于密钥派生算法的链上匿名通信方法，其特征在于，应用于区块链网络中任意两个拥有初始密钥的节点，基于密钥衍生算法，将消息加密上链，以去中心化的区块链作为中继，实现节点到区块链再到节点的消息安全传输的过程，所述方法包括：在第一客户端和第二客户端向认证中心进行身份认证通过后，所述认证中心生成所述第一客户端对应的第一初始私钥、第一初始公钥，以及生成所述第二客户端对应的第二初始私钥、第二初始公钥；所述第一客户端根据第一初始公钥确定第一初始地址，第二客户端根据第二初始公钥确定第二初始地址；所述第一客户端和所述第二客户端根据所述第一初始私钥、第一初始公钥、第一初始地址、第二初始私钥、第二初始公钥和第二初始地址完成链上好友确认和种子密钥的生成；所述第一客户端根据所述第一初始私钥、所述第二初始公钥、所述种子密钥和通信时间戳确定第一派生私钥、第二派生公钥和第二派生地址；所述第二客户端根据所述第二初始私钥、所述第一初始公钥、所述种子密钥和所述通信时间戳确定第二派生私钥、第一派生公钥和第一派生地址；所述第一客户端和所述第二客户端根据所述第一派生私钥、所述第一派生公钥、所述第一派生地址、所述第二派生私钥、所述第二派生公钥和所述第二派生地址完成所述第一客户端和所述第二客户端之间消息的匿名发送。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"徐洁\",\"CLS\":\"实审\",\"INN\":\"徐洁;宋绪言;崔宝江;陈思源;付俊松;XU JIE;SONG XUYAN;CUI BAOJIANG;CHEN SIYUAN;FU JUNSONG\",\"IDX\":40,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:43:55.597', '2023-01-12 16:43:55.597', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 'CN114423007A', '{\"patentId\":0,\"TI\":\"终端接入点的确定方法、确定装置、电子设备和存储介质\",\"PNM\":\"CN114423007A\",\"AD\":\"2022.01.25\",\"PD\":\"2022.04.29\",\"CL\":\" 1.一种终端接入点的确定方法，其特征在于，所述确定方法包括：接收目标用户通过用户终端发送的业务访问需求；其中，所述业务访问需求包括目标业务需求、安全性需求和服务质量需求；基于所述目标用户的位置信息，确定所述位置信息所对应的覆盖范围内的至少一个接入点；利用预先构建好的接入态势预测模型对每个接入点的接入态势进行预测，得到每个接入点对应的预测态势信息；根据所述业务访问需求和预先构建好的用户信任评估模型，对所述目标用户的信任等级进行评估，得到所述目标用户的信任等级信息；根据所述目标用户的信任等级信息以及每个接入点的预测态势信息，从至少一个接入点中确定出目标接入点，并将所述目标接入点发送至所述用户终端，以使所述用户终端通过所述目标接入点访问微服务应用。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"崔宝江\",\"CLS\":\"实审\",\"INN\":\"崔宝江;王子奇;徐洁;CUI BAOJIANG;WANG ZIQI;XU JIE\",\"IDX\":43,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:43:57.925', '2023-01-12 16:43:57.925', NULL);
INSERT INTO `patent` (`patent_id`, `pnm`, `patent_properties`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 'CN114138671A', '{\"patentId\":0,\"TI\":\"一种协议的测试方法、装置、电子设备及存储介质\",\"PNM\":\"CN114138671A\",\"AD\":\"2021.12.13\",\"PD\":\"2022.03.04\",\"CL\":\" 1.一种协议的测试方法，其特征在于，包括：获取测试样本；利用强化学习模型对所述测试样本进行处理，获取反馈结果和模糊测试结果；将所述反馈结果输入预设的有限状态自动机模型，获取统计结果，所述统计结果用于检验所述模糊测试结果。\",\"PA\":\"北京邮电大学\",\"AR\":\"北京市海淀区西土城路10号\",\"PINN\":\"崔宝江\",\"CLS\":\"实审\",\"INN\":\"崔宝江;侯晓庚;吴启凡;陈晨;李帅;毛立涛;王灏宇;CUI BAOJIANG;HOU XIAOGENG;WU QIFAN;CHEN CHEN;LI SHUAI;MAO LITAO;WANG HAOYU\",\"IDX\":54,\"desc\":\"\",\"CreateBy\":1,\"UpdateBy\":0}', 1, 0, '2023-01-12 16:44:00.531', '2023-01-12 16:44:00.531', NULL);
COMMIT;

-- ----------------------------
-- Table structure for patent_package
-- ----------------------------
DROP TABLE IF EXISTS `patent_package`;
CREATE TABLE `patent_package` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `patent_id` bigint DEFAULT NULL COMMENT '专利Id',
  `package_id` bigint DEFAULT NULL COMMENT '专利包ID',
  `pnm` varchar(128) DEFAULT NULL COMMENT '申请号',
  `desc` varchar(128) DEFAULT NULL COMMENT '描述',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_patent_package_create_by` (`create_by`),
  KEY `idx_patent_package_update_by` (`update_by`),
  KEY `idx_patent_package_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of patent_package
-- ----------------------------
BEGIN;
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 1, 1, 'CN114662096A', '', 1, 0, '2023-01-12 16:44:46.758', '2023-01-12 16:44:46.758', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 2, 1, 'CN114238977A', '', 1, 0, '2023-01-12 16:44:49.795', '2023-01-12 16:44:49.795', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 3, 1, 'CN114239737A', '', 1, 0, '2023-01-12 16:44:52.501', '2023-01-12 16:44:52.501', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 4, 1, 'CN114241581A', '', 1, 0, '2023-01-12 16:44:55.525', '2023-01-12 16:44:55.525', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 5, 1, 'CN113486190A', '', 1, 0, '2023-01-12 16:44:58.543', '2023-01-12 16:44:58.543', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 6, 1, 'CN113486932A', '', 1, 0, '2023-01-12 16:45:02.724', '2023-01-12 16:45:02.724', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 7, 1, 'CN113158390A', '', 1, 0, '2023-01-12 16:45:05.254', '2023-01-12 16:45:05.254', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 8, 1, 'CN113132414A', '', 1, 0, '2023-01-12 16:45:07.951', '2023-01-12 16:45:07.951', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 9, 1, 'CN113094860A', '', 1, 0, '2023-01-12 16:45:11.247', '2023-01-12 16:45:11.247', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 10, 2, 'CN112073362A', '', 1, 0, '2023-01-12 16:45:44.056', '2023-01-12 16:45:44.056', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 11, 2, 'CN111881935A', '', 1, 0, '2023-01-12 16:45:47.014', '2023-01-12 16:45:47.014', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 12, 2, 'CN111861844A', '', 1, 0, '2023-01-12 16:45:50.440', '2023-01-12 16:45:50.440', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 13, 2, 'CN111861845A', '', 1, 0, '2023-01-12 16:45:53.918', '2023-01-12 16:45:53.918', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 14, 2, 'CN110278210A', '', 1, 0, '2023-01-12 16:45:57.292', '2023-01-12 16:45:57.292', NULL);
INSERT INTO `patent_package` (`id`, `patent_id`, `package_id`, `pnm`, `desc`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 16, 2, 'CN109587372A', '', 1, 0, '2023-01-12 16:46:03.156', '2023-01-12 16:46:03.156', NULL);
COMMIT;

-- ----------------------------
-- Table structure for patent_tag
-- ----------------------------
DROP TABLE IF EXISTS `patent_tag`;
CREATE TABLE `patent_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `patent_id` bigint DEFAULT NULL COMMENT '专利Id',
  `tag_id` bigint DEFAULT NULL COMMENT '标签ID',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_patent_tag_create_by` (`create_by`),
  KEY `idx_patent_tag_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of patent_tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for report
-- ----------------------------
DROP TABLE IF EXISTS `report`;
CREATE TABLE `report` (
  `report_id` bigint NOT NULL AUTO_INCREMENT COMMENT '报告ID(主键)',
  `report_name` longtext COMMENT '报告名称',
  `report_properties` longtext COMMENT '报告详情',
  `type` varchar(64) DEFAULT NULL COMMENT '报告类型（侵权/估值）',
  `reject_tag` varchar(8) DEFAULT NULL COMMENT '驳回标签(null:未审核/reject/upload)',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `files` longtext,
  PRIMARY KEY (`report_id`),
  KEY `idx_report_create_by` (`create_by`),
  KEY `idx_report_update_by` (`update_by`),
  KEY `idx_report_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of report
-- ----------------------------
BEGIN;
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (1, 'infringement.defaultName.7b541439-f975-4e88-bcf0-a023777e9aff', '', 'infringement', '未审核', 1, 0, '2023-01-12 16:51:37.305', '2023-01-12 16:51:37.305', NULL, '');
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (2, 'valuation.defaultName.3b5c4c04-7c03-47f5-83e7-1932db615bc6', '', 'valuation', '未审核', 1, 0, '2023-01-12 16:51:41.172', '2023-01-12 16:51:41.172', NULL, '');
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (3, '使用无人机数据产生用于自主车辆导航的高清地图 查新报告', '', 'novelty', '', 0, 0, '2023-01-12 16:52:17.854', '2023-01-12 16:52:17.854', NULL, '');
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (4, '使用无人机数据产生用于自主车辆导航的高清地图 查新报告', '', 'novelty', '', 0, 0, '2023-01-12 16:55:17.738', '2023-01-12 16:55:17.738', NULL, '');
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (5, '使用无人机数据产生用于自主车辆导航的高清地图 查新报告', '', 'novelty', '', 0, 0, '2023-01-12 17:00:45.295', '2023-01-12 17:00:45.295', NULL, '');
INSERT INTO `report` (`report_id`, `report_name`, `report_properties`, `type`, `reject_tag`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`, `files`) VALUES (6, '使用无人机数据产生用于自主车辆导航的高清地图 查新报告', '', 'novelty', '', 0, 0, '2023-01-12 17:02:21.591', '2023-01-12 17:02:21.591', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for report_rela
-- ----------------------------
DROP TABLE IF EXISTS `report_rela`;
CREATE TABLE `report_rela` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `patent_id` bigint DEFAULT NULL COMMENT '专利ID',
  `report_id` bigint DEFAULT NULL COMMENT '报告ID',
  `user_id` bigint DEFAULT NULL COMMENT '用户ID',
  `type` longtext,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` longtext COMMENT '创建时间',
  `updated_at` longtext COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_report_rela_create_by` (`create_by`),
  KEY `idx_report_rela_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of report_rela
-- ----------------------------
BEGIN;
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (1, 1, 1, 1, 'infringement', 1, 0, '2023-01-12 16:51:37', '');
INSERT INTO `report_rela` (`id`, `patent_id`, `report_id`, `user_id`, `type`, `create_by`, `update_by`, `created_at`, `updated_at`) VALUES (2, 2, 2, 1, 'valuation', 1, 0, '2023-01-12 16:51:41', '');
COMMIT;

-- ----------------------------
-- Table structure for stored_query
-- ----------------------------
DROP TABLE IF EXISTS `stored_query`;
CREATE TABLE `stored_query` (
  `query_id` bigint NOT NULL AUTO_INCREMENT COMMENT '检索表达式ID(主键)',
  `name` varchar(128) DEFAULT NULL COMMENT '名称',
  `desc` varchar(1024) DEFAULT NULL COMMENT '描述',
  `expression` longtext COMMENT '检索表达式',
  `db` longtext COMMENT '检索数据库',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`query_id`),
  KEY `idx_stored_query_deleted_at` (`deleted_at`),
  KEY `idx_stored_query_create_by` (`create_by`),
  KEY `idx_stored_query_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of stored_query
-- ----------------------------
BEGIN;
INSERT INTO `stored_query` (`query_id`, `name`, `desc`, `expression`, `db`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '测试检索', '这是一个测试用检索表达式', '刘建毅 北京邮电大学', 'wgzl,syxx,fmzl', 1, 1, '2023-01-12 16:50:14.315', '2023-01-12 16:50:14.315', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_casbin_rule`;
CREATE TABLE `sys_casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `v6` varchar(25) DEFAULT NULL,
  `v7` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sys_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`,`v6`,`v7`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of sys_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `sys_casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) VALUES (1, 'p', 'user', '/apis/v1/user-agent/*', '*', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `username` varchar(128) DEFAULT NULL COMMENT '用户名',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) DEFAULT NULL COMMENT '固件',
  `login_time` datetime(3) DEFAULT NULL COMMENT '登录时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `msg` varchar(255) DEFAULT NULL COMMENT '信息',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_sys_login_log_create_by` (`create_by`),
  KEY `idx_sys_login_log_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_migration
-- ----------------------------
DROP TABLE IF EXISTS `sys_migration`;
CREATE TABLE `sys_migration` (
  `version` varchar(191) NOT NULL,
  `apply_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of sys_migration
-- ----------------------------
BEGIN;
INSERT INTO `sys_migration` (`version`, `apply_time`) VALUES ('1599190683659', '2023-01-12 16:41:10.615');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` varchar(4) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` bigint DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` tinyint(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`),
  KEY `idx_sys_role_create_by` (`create_by`),
  KEY `idx_sys_role_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
INSERT INTO `sys_role` (`role_id`, `role_name`, `status`, `role_key`, `role_sort`, `flag`, `remark`, `admin`, `data_scope`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '普通用户', '2', 'user', 0, '', '', 0, '', 0, 0, '2021-05-13 19:56:37.913', '2021-05-13 19:56:37.913', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编码',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `password` varchar(128) DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(128) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `role_id` mediumint DEFAULT NULL COMMENT '角色ID',
  `salt` varchar(255) DEFAULT NULL COMMENT '加盐',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `sex` varchar(255) DEFAULT NULL COMMENT '性别',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` varchar(4) DEFAULT NULL COMMENT '状态',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_update_by` (`update_by`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'admin', '13818888888', 1, '', '', '1', '1@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'user', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user', '13818888888', 2, '', '', '1', '1@qq.com', '', '2', 0, 0, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'user2', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user2', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'user3', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user3', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` (`user_id`, `username`, `password`, `nick_name`, `phone`, `role_id`, `salt`, `avatar`, `sex`, `email`, `remark`, `status`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'user4', '$2a$10$/Glr4g9Svr6O0kvjsRJCXu3f0W8/dsP3XZyVNi1019ratWpSPMyw.', 'user4', '13818888888', 2, NULL, NULL, '1', '1@qq.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` bigint NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(128) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`tag_id`),
  KEY `idx_tag_create_by` (`create_by`),
  KEY `idx_tag_update_by` (`update_by`),
  KEY `idx_tag_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for trace_log
-- ----------------------------
DROP TABLE IF EXISTS `trace_log`;
CREATE TABLE `trace_log` (
  `trace_id` bigint NOT NULL AUTO_INCREMENT COMMENT 'traceID(主键)',
  `action` varchar(128) DEFAULT NULL COMMENT '操作',
  `desc` varchar(1024) DEFAULT NULL COMMENT '描述',
  `request` varchar(1024) DEFAULT NULL COMMENT '用户请求',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`trace_id`),
  KEY `idx_trace_log_deleted_at` (`deleted_at`),
  KEY `idx_trace_log_create_by` (`create_by`),
  KEY `idx_trace_log_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of trace_log
-- ----------------------------
BEGIN;
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'Search', '查询操作，表达式：网络安全', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:41:31.106', '2023-01-12 16:41:31.106', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'Search', '查询操作，表达式：刘建毅 北京邮电大学', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:41:58.041', '2023-01-12 16:41:58.041', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'Search', '查询操作，表达式：刘建毅 北京邮电大学', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:42:23.859', '2023-01-12 16:42:23.859', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'Search', '查询操作，表达式：崔宝江', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:43:13.041', '2023-01-12 16:43:13.041', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'Search', '查询操作，表达式：崔宝江 北京邮电大学', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:43:36.749', '2023-01-12 16:43:36.749', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'Search', '查询操作，表达式：测试', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:43:55.853', '2023-01-12 16:43:55.853', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'Search', '查询操作，表达式：测试', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:44:00.985', '2023-01-12 16:44:00.985', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'Search', '查询操作，表达式：刘建', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:46:17.011', '2023-01-12 16:46:17.011', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'Search', '查询操作，表达式：刘建', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:46:19.374', '2023-01-12 16:46:19.374', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'Search', '查询操作，表达式：崔', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:46:26.037', '2023-01-12 16:46:26.037', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'Search', '查询操作，表达式：刘建毅 北京邮电大学', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:49:54.202', '2023-01-12 16:49:54.202', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 'Search', '查询操作，表达式： PNM=\'123\'', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:50:26.516', '2023-01-12 16:50:26.516', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'Search', '查询操作，表达式：PNM=\'123\'', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:50:32.600', '2023-01-12 16:50:32.600', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 'Search', '查询操作，表达式：DESCR=\'123\' and PINN=\'213\'', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:50:40.008', '2023-01-12 16:50:40.008', NULL);
INSERT INTO `trace_log` (`trace_id`, `action`, `desc`, `request`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 'Search', '查询操作，表达式：刘建毅 北京邮电大学', '请求URI：/api/v1/user-agent/auth-search', 1, 0, '2023-01-12 16:50:56.071', '2023-01-12 16:50:56.071', NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_patent
-- ----------------------------
DROP TABLE IF EXISTS `user_patent`;
CREATE TABLE `user_patent` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `patent_id` bigint DEFAULT NULL COMMENT 'PatentId',
  `user_id` bigint DEFAULT NULL COMMENT '用户ID',
  `pnm` varchar(128) DEFAULT NULL COMMENT '申请号',
  `type` varchar(64) DEFAULT NULL COMMENT '关系类型（关注/认领）',
  `desc` varchar(128) DEFAULT NULL COMMENT '描述',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`),
  KEY `idx_user_patent_deleted_at` (`deleted_at`),
  KEY `idx_user_patent_create_by` (`create_by`),
  KEY `idx_user_patent_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_patent
-- ----------------------------
BEGIN;
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (1, 1, 1, 'CN114662096A', '认领', '', '2023-01-12 16:42:01.593', '2023-01-12 16:42:01.593', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (2, 2, 1, 'CN114238977A', '认领', '', '2023-01-12 16:42:03.159', '2023-01-12 16:42:03.159', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (3, 3, 1, 'CN114239737A', '认领', '', '2023-01-12 16:42:05.371', '2023-01-12 16:42:05.371', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (4, 4, 1, 'CN114241581A', '认领', '', '2023-01-12 16:42:07.935', '2023-01-12 16:42:07.935', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (5, 5, 1, 'CN113486190A', '认领', '', '2023-01-12 16:42:09.677', '2023-01-12 16:42:09.677', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (6, 6, 1, 'CN113486932A', '认领', '', '2023-01-12 16:42:12.129', '2023-01-12 16:42:12.129', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (7, 7, 1, 'CN113158390A', '认领', '', '2023-01-12 16:42:14.167', '2023-01-12 16:42:14.167', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (8, 8, 1, 'CN113132414A', '认领', '', '2023-01-12 16:42:16.134', '2023-01-12 16:42:16.134', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (9, 9, 1, 'CN113094860A', '认领', '', '2023-01-12 16:42:18.297', '2023-01-12 16:42:18.297', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (10, 10, 1, 'CN112073362A', '认领', '', '2023-01-12 16:42:21.496', '2023-01-12 16:42:21.496', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (11, 11, 1, 'CN111881935A', '认领', '', '2023-01-12 16:42:28.819', '2023-01-12 16:42:28.819', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (12, 12, 1, 'CN111861844A', '认领', '', '2023-01-12 16:42:32.954', '2023-01-12 16:42:32.954', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (13, 13, 1, 'CN111861845A', '认领', '', '2023-01-12 16:42:35.529', '2023-01-12 16:42:35.529', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (14, 14, 1, 'CN110278210A', '认领', '', '2023-01-12 16:42:38.006', '2023-01-12 16:42:38.006', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (15, 15, 1, 'CN109617682A', '认领', '', '2023-01-12 16:42:40.398', '2023-01-12 16:42:40.398', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (16, 16, 1, 'CN109587372A', '认领', '', '2023-01-12 16:42:42.934', '2023-01-12 16:42:42.934', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (17, 17, 1, 'CN114896595A', '关注', '', '2023-01-12 16:43:51.279', '2023-01-12 16:43:51.279', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (18, 18, 1, 'CN115225551A', '关注', '', '2023-01-12 16:43:53.408', '2023-01-12 16:43:53.408', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (19, 19, 1, 'CN114679261A', '关注', '', '2023-01-12 16:43:55.611', '2023-01-12 16:43:55.611', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (20, 20, 1, 'CN114423007A', '关注', '', '2023-01-12 16:43:57.941', '2023-01-12 16:43:57.941', NULL, 0, 0);
INSERT INTO `user_patent` (`id`, `patent_id`, `user_id`, `pnm`, `type`, `desc`, `created_at`, `updated_at`, `deleted_at`, `create_by`, `update_by`) VALUES (21, 21, 1, 'CN114138671A', '关注', '', '2023-01-12 16:44:00.545', '2023-01-12 16:44:00.545', '2023-01-12 16:49:28.391', 0, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
