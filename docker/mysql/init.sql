DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Users';

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

INSERT INTO `users` (`id`, `name`, `email`) VALUES
(1, 'foo', 'foo@example.com'),
(2, 'bar', 'bar@example.com'),
(3, 'hoge', 'hoge@example.com');