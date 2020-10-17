DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Users';

INSERT INTO `users` (`id`, `name`, `email`) VALUES
(1, 'foo', 'foo@example.com'),
(2, 'bar', 'bar@example.com'),
(3, 'hoge', 'hoge@example.com');

DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis` (
  `id` varchar(36) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `url` varchar(40) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='APIs';

SET @users_api_id = UUID();
SET @posts_api_id = UUID();

INSERT INTO `apis` (`id`, `name`, `url`, `description`) VALUES
(@users_api_id, 'users', 'my-projects/api/users', 'ユーザーに関する操作をするAPIです'),
(@posts_api_id, 'posts', 'my-project/api/posts', '投稿に関する操作をするAPIです');
-- (UUID(), 'photos', 'my-project/api/photos', '写真に関する操作をするAPIです');

DROP TABLE IF EXISTS `methods`;
CREATE TABLE `methods` (
  `id` varchar(36) NOT NULL,
  `api_id` varchar(36) NOT NULL,
  `type` varchar(10) NOT NULL DEFAULT '',
  `url` varchar(40) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `request_parameter` varchar(60) NOT NULL DEFAULT '',
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

-- SET @posts_getall_id = UUID();
-- SET @posts_getbyid_id = UUID();
-- SET @posts_create_id = UUID();
-- SET @posts_update_id = UUID();
-- SET @posts_delete_id = UUID();

INSERT INTO `methods` (`id`, `api_id`, `type`, `url`, `description`, `request_parameter`, `request_model_id`, `response_model_id`, `is_array`) VALUES
(@users_getall_id, @users_api_id, 'GET', '', 'すべてのユーザーを取得します。', '', '', '', true),
(@users_getbyid_id, @users_api_id, 'GET', '/{user_id}', 'user_idから1件のユーザーを取得します。', 'user_id', '', '', false),
(@users_create_id, @users_api_id, 'POST', '', 'ユーザーを1件作成します。', '', '', '', false),
(@users_update_id, @users_api_id, 'PUT', '', 'ユーザーを1件更新します。', '', '', '', false),
(@users_delete_id, @users_api_id, 'DELETE', '/{user_id}', 'ユーザーを1件削除します。', 'user_id', '', '', false);

DROP TABLE IF EXISTS `models`;
CREATE TABLE `models` (
  `id` varchar(36) NOT NULL,
  `api_id` varchar(36) NOT NULL DEFAULT '',
  `method_id` varchar(36) NOT NULL DEFAULT '',
  `name` varchar(40) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Models';

DROP TABLE IF EXISTS `columns`;
CREATE TABLE `columns` (
  `id` varchar(36) NOT NULL,
  `model_id` varchar(36) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `type` varchar(40) NOT NULL DEFAULT '',
  `column_model_id` varchar(36) NOT NULL DEFAULT '',
  `is_array` boolean,
  `is_nullable` boolean,
  `is_unique` boolean,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (model_id) 
    REFERENCES models(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Columns';

DROP TABLE IF EXISTS `column_options`;
CREATE TABLE `column_options` (
  `id` varchar(36) NOT NULL,
  `column_id` varchar(36) NOT NULL,
  `option_key` varchar(40) NOT NULL DEFAULT '',
  `option_value` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (column_id) 
    REFERENCES columns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Column options';

