
# patent-admin-plat

## 🎬 Quick Start

#### 初始用户：

超管：admin 123456

普通用户：user 123456

#### 编译：

```bash
$ bash build_linux.sh
```

#### 初始化数据库：

```bash
$ create database dbname default charset utf8
$ bin/PatentAdminPlat-linux-amd64 migrate -c config/settings_dev.yml
```

#### 启动：

```bash
$ bin/PatentAdminPlat-linux-amd64 server -c config/settings_dev.yml
```

#### 生成数据库更新文件并更新

```bash
$ bin/PatentAdminPlat-linux-amd64 migrate -g -c config/settings_dev.yml
# 更新结束后
$ bin/PatentAdminPlat-linux-amd64 migrate -c config/settings_dev.yml
```

## ✨ Swagger

#### 生成swagger代码

```bash
$ swag init
```

## 🏦 DataBase
###  用户表（sys_user）

| 字段名    | 描述     | 类型   |
| --------- | -------- | ------ |
| UserId    | 用户ID   | Int    |
| Username  | 用户名   | String |
| Password  | 密码     | string |
| NickName  | 昵称     | string |
| Phone     | 手机号   | string |
| RoleId    | 角色ID   | int    |
| Salt      | 加盐     | string |
| Avatar    | 头像     | string |
| Sex       | 性别     | string |
| Email     | 邮箱     | string |
| Remark    | 备注     | string |
| Status    | 状态     | string |
| Departure | 单位名   | string |
| Bio       | 个人简介 | string |
| Interest  | 科研兴趣 | string |

###  角色表（sys_role）

| 字段名    | 描述       | 类型   |
| --------- | ---------- | ------ |
| RoleId    | 角色ID     | int    |
| RoleName  | 角色名称   | string |
| Status    | 状态       | string |
| RoleKey   | 角色代码   | string |
| RoleSort  | 角色排序   | int    |
| Flag      | 标志位     | string |
| Remark    | 备注       | string |
| Admin     | 是否是超管 | bool   |
| DataScope | 无         | string |

### 专利表（patent）

| 字段名   | 描述     | 类型   |
| -------- | -------- | ------ |
| PatentId | 专利ID   | int    |
| TI       | 专利名   | string |
| PNM      | 申请号   | string |
| AD       | 申请日   | string |
| PD       | 公开日   | string |
| CL       | 简介     | string |
| PA       | 申请单位 | string |
| AR       | 地址     | string |
| INN      | 申请人   | string |

### 用户-专利关系表（user_patent）

| 字段名   | 描述                  | 类型   |
| -------- | --------------------- | ------ |
| ID       | ID                    | int    |
| UserId   | 用户ID                | int    |
| PatentId | 专利ID                | int    |
| Type     | 关系类型（关注/认领） | string |

### 标签表（tag）

| 字段名  | 描述   | 类型   |
| ------- | ------ | ------ |
| TagId   |        | int    |
| TagName | 标签名 | string |
| Desc    | 描述   | string |

### 用户-标签关系表（user-tag）

| 字段名 | 描述 | 类型 |
| ------ | ---- | ---- |
| ID     |      | int  |
| UserId |      | int  |
| TagId  |      | int  |

### 专利-标签关系表（patent-tag）

| 字段名   | 描述 | 类型 |
| -------- | ---- | ---- |
| ID       |      | int  |
| PatentId |      | int  |
| TagId    |      | int  |

### 专利包表（package）

| 字段名      | 描述   | 类型   |
| ----------- | ------ | ------ |
| PackageId   |        | int    |
| PackageName | 专利包 | string |
| Desc        | 描述   | string |

### 用户-专利包关系表（user-package）

| 字段名    | 描述 | 类型 |
| --------- | ---- | ---- |
| ID        |      | int  |
| UserId    |      | int  |
| PackageId |      | int  |

### 专利-专利包关系表（patent-package）

| 字段名    | 描述 | 类型 |
| --------- | ---- | ---- |
| ID        |      | int  |
| PatentId  |      | int  |
| PackageId |      | int  |
