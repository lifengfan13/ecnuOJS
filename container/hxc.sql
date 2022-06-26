create database `paper` default character set utf8 collate utf8_general_ci;

use paper;

DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `submission`;
DROP TABLE IF EXISTS `reviewer_submission`;
DROP TABLE IF EXISTS `key_word`;

CREATE TABLE `user` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `created_at` DATE NOT NULL,
    `updated_at` DATE NOT NULL,
    `deleted_at` DATE,
    `is_del` tinyint(1) NOT NULL,
    `username` varchar(100) NOT NULL,
    `email` varchar(100),
    `telephone` varchar(100) NOT NULL,
    `organization` varchar(100),
    `address` varchar(100),
    `postcode` varchar(100),
    `role` tinyint(3) NOT NULL,
    `password_hash` varchar(200) NOT NULL,
    -- 超级用户0，投稿者1，审稿者2
    primary key(id)
);

CREATE TABLE `submission` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `created_at` DATE NOT NULL,
    `updated_at` DATE NOT NULL,
    `deleted_at` DATE,
    `is_del` tinyint(1) NOT NULL,
    `name` varchar(100) NOT NULL,
    `first_author` int(10) NOT NULL,
    `corresponding_author` int(10) NOT NULL,
    `second_author` int(10),
    `third_author` int(10),
    `forth_author` int(10),
    `topic` int(10) NOT NULL,
    `status` int(5) NOT NULL,
    `file_path` varchar(100) NOT NULL,
    primary key(id)
);


CREATE TABLE `reviewer_submission` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `created_at` DATE NOT NULL,
    `updated_at` DATE NOT NULL,
    `deleted_at` DATE,
    `reviewer_id` int(10) NOT NULL,
    `submission_id` int(10) NOT NULL,
    `status` int(5) NOT NULL,
    `review_comment` varchar(200),
    primary key(id)
);

CREATE TABLE `key_word` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `created_at` DATE NOT NULL,
    `updated_at` DATE NOT NULL,
    `deleted_at` DATE,
    `is_del` tinyint(1) NOT NULL,
    `name` varchar(100) NOT NULL,
    `submission_id` int(10) NOT NULL,
    primary key(id)
);

INSERT INTO `user` (`id`,`created_at`,`updated_at`,`deleted_at`,`is_del`,`username`,`email`,`telephone`,`role`,`password_hash`)
VALUE(0,CURDATE(),CURDATE(),NULL,0,'admin','houxc_mail@163.com','18019134181',0,'e10adc3949ba59abbe56e057f20f883e');