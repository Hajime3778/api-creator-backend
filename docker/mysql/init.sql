DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis` (
  `id` varchar(36) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `url` varchar(64) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='APIs';

SET @users_api_id = UUID();
SET @posts_api_id = UUID();
SET @photos_api_id = UUID();

INSERT INTO `apis` (`id`, `name`, `url`, `description`) VALUES
(@users_api_id, 'Users', 'my-project/api/users', 'ユーザーに関する操作をするAPIです'),
(@posts_api_id, 'Posts', 'my-project/api/posts', '投稿に関する操作をするAPIです'),
(@photos_api_id, 'Photos', 'my-project/api/photos', '写真に関する操作をするAPIです');

DROP TABLE IF EXISTS `methods`;
CREATE TABLE `methods` (
  `id` varchar(36) NOT NULL,
  `api_id` varchar(36) NOT NULL,
  `type` varchar(8) NOT NULL DEFAULT '',
  `url` varchar(64) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `request_parameter` varchar(64) NOT NULL DEFAULT '',
  `request_model_id` varchar(36) NOT NULL DEFAULT '',
  `response_model_id` varchar(36) NOT NULL DEFAULT '',
  `is_array` boolean,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (api_id) 
    REFERENCES apis(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Methods';

SET @users_getall_id = UUID();
SET @users_getbyid_id = UUID();
SET @users_create_id = UUID();
SET @users_update_id = UUID();
SET @users_delete_id = UUID();

SET @posts_getall_id = UUID();
SET @posts_getbyid_id = UUID();
SET @posts_create_id = UUID();
SET @posts_update_id = UUID();
SET @posts_delete_id = UUID();

SET @photos_getall_id = UUID();
SET @photos_getbyid_id = UUID();
SET @photos_create_id = UUID();
SET @photos_update_id = UUID();
SET @photos_delete_id = UUID();

SET @users_model_id = UUID();
SET @posts_model_id = UUID();
SET @photos_model_id = UUID();

INSERT INTO `methods` (`id`, `api_id`, `type`, `url`, `description`, `request_parameter`, `request_model_id`, `response_model_id`, `is_array`) VALUES
(@users_getall_id, @users_api_id, 'GET', '', 'すべてのユーザーを取得します。', '', '', @users_model_id, true),
(@users_getbyid_id, @users_api_id, 'GET', '/{id}', 'idから1件のユーザーを取得します。', 'id', '', @users_model_id, false),
(@users_create_id, @users_api_id, 'POST', '', 'ユーザーを1件作成します。', '', @users_model_id, '', false),
(@users_update_id, @users_api_id, 'PUT', '', 'ユーザーを1件更新します。', '', @users_model_id, '', false),
(@users_delete_id, @users_api_id, 'DELETE', '/{id}', 'ユーザーを1件削除します。', 'id', '', '', false),
(@posts_getall_id, @posts_api_id, 'GET', '', 'すべての投稿を取得します。', '', '', @posts_model_id, true),
(@posts_getbyid_id, @posts_api_id, 'GET', '/{id}', 'idから1件の投稿を取得します。', 'id', '', @posts_model_id, false),
(@posts_create_id, @posts_api_id, 'POST', '', '投稿を1件作成します。', '', @posts_model_id, '', false),
(@posts_update_id, @posts_api_id, 'PUT', '', '投稿を1件更新します。', '', @posts_model_id, '', false),
(@posts_delete_id, @posts_api_id, 'DELETE', '/{id}', '投稿を1件削除します。', 'id', '', '', false),
(@photos_getall_id, @photos_api_id, 'GET', '', 'すべての写真を取得します。', '', '', @photos_model_id, true),
(@photos_getbyid_id, @photos_api_id, 'GET', '/{id}', 'idから1件の写真を取得します。', 'id', '', @photos_model_id, false),
(@photos_create_id, @photos_api_id, 'POST', '', '写真を1件作成します。', '', @photos_model_id, '', false),
(@photos_update_id, @photos_api_id, 'PUT', '', '写真を1件更新します。', '', @photos_model_id, '', false),
(@photos_delete_id, @photos_api_id, 'DELETE', '/{id}', '写真を1件削除します。', 'id', '', '', false);

DROP TABLE IF EXISTS `models`;
CREATE TABLE `models` (
  `id` varchar(36) NOT NULL,
  `api_id` varchar(36) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `description` varchar(255) NOT NULL DEFAULT '',
  `schema` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Models';

INSERT INTO `models` (`id`, `api_id`, `name`, `description`, `schema`) VALUES
(@users_model_id, @users_api_id, 'User', 'ユーザーを定義するモデルです。', '{\n    "type": "object",\n    "additionalProperties": false,\n    "keys": ["id"],\n    "properties": {\n        "id": {\n            "type": "string",\n            "description": "ID"\n        },\n        "name": {\n            "type": "string",\n            "description": "名前"\n        },\n        "email": {\n            "type": "string",\n            "description": "メールアドレス"\n        },\n        "description": {\n            "type": "string",\n            "description": "説明"\n        }\n    },\n    "required": [\n        "id",\n        "name",\n        "email",\n        "description"\n    ]\n}'),
(@posts_model_id, @posts_api_id, 'Post', '投稿を定義するモデルです。', '{\n    "type": "object",\n    "additionalProperties": false,\n    "keys": ["id"],\n    "properties": {\n        "id": {\n            "type": "string",\n            "description": "ID"\n        },\n        "name": {\n            "type": "string",\n            "description": "投稿名"\n        },\n       "body": {\n            "type": "string",\n            "description": "投稿内容"\n        },\n        "postedDate": {\n            "type": "string",\n            "description": "投稿日"\n        },\n        "postedUserId": {\n            "type": "string",\n            "description": "投稿者ID"\n        }\n    },\n    "required": [\n        "id",\n        "name",\n        "body",\n        "name",\n        "postedDate",\n        "postedUserId"\n    ]\n}'),
(@photos_model_id, @photos_api_id, 'Photo', '写真を定義するモデルです。', '{\n    "type": "object",\n    "additionalProperties": false,\n    "keys": ["id"],\n    "properties": {\n        "id": {\n            "type": "string",\n            "description": "ID"\n        },\n        "name": {\n            "type": "string",\n            "description": "写真名"\n        },\n       "url": {\n            "type": "string",\n            "description": "写真のURL"\n        }\n    },\n    "required": [\n        "id",\n        "name",\n        "url"\n    ]\n}');
