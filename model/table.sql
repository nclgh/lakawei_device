CREATE TABLE `device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `model` varchar(255) NOT NULL DEFAULT '',
  `brand` varchar(255) NOT NULL DEFAULT '',
  `tag_code` varchar(255) NOT NULL DEFAULT '',
  `department_id` bigint(20) unsigned NOT NULL DEFAULT 0,
  `manufacturer_id` bigint(20) unsigned NOT NULL DEFAULT 0,
  `manufacturer_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `rent_status` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `description` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_device_code` (`code`),
  KEY `idx_device_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `manufacturer` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `device_rent` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_id` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(1) unsigned NOT NULL DEFAULT 0,
  `borrower_member_id` bigint(20) NOT NULL DEFAULT '',
  `borrower_department_id` bigint(20) unsigned NOT NULL DEFAULT 0,
  `borrow_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `expect_return_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `returner_member_id` bigint(20) NOT NULL DEFAULT '',
  `returner_department_id` bigint(20) unsigned NOT NULL DEFAULT 0,
  `real_return_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `borrow_remark` text NOT NULL,
  `return_remark` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `achievement` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_id` bigint(20) NOT NULL DEFAULT 0,
  `member_id` bigint(20) NOT NULL DEFAULT 0,
  `department_id` bigint(20) NOT NULL DEFAULT 0,
  `achievement_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `achievement_description` text NOT NULL,
  `achievement_remark` text NOT NULL,
  `patent_description` text NOT NULL,
  `paper_description` text NOT NULL,
  `competition_description` text NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;