

CREATE DATABASE IF NOT EXISTS `passwd`;

CREATE TABLE IF NOT EXISTS `passwd_user` (
  `user_id` integer unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(16) NOT NULL,
  `password` varchar(32) NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `share_mode` tinyint NOT NULL,
  `role` tinyint NOT NULL,
  `profile_img_url` varchar(128),
  `description` varchar(50),
  `sex` tinyint,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `deleted_at`  datetime,
  `is_deleted` boolean NOT NULL,
  PRIMARY KEY (`user_id`)
) CHARSET=utf8;



CREATE TABLE IF NOT EXISTS `passwd_platform` (
  `platform_id` integer unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(16) NOT NULL,
  `name` varchar(16) NOT NULL,
  `abbr` varchar(16) NOT NULL,
  `description` varchar(32),
  `img_url` varchar(128),
  `login_url` varchar(128),
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `is_deleted` boolean NOT NULL,
  `deleted_at`  datetime,
  PRIMARY KEY (`platform_id`)
) CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `passwd_user_account` (
  `user_account_id` integer unsigned NOT NULL AUTO_INCREMENT,
  `user_id` integer unsigned NOT NULL,
  `platform_id` integer unsigned NOT NULL,
  `password` varchar(32) NOT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `is_deleted` boolean NOT NULL,
  `deleted_at`  datetime,
  PRIMARY KEY (`user_account_id`)
) CHARSET=utf8;