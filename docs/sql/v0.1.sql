DROP DATABASE blog;
CREATE DATABASE blog;

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `uid` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`uid`)
);

DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags` (
  `article_id` int UNSIGNED NOT NULL,
  `tag_id` tinyint UNSIGNED NOT NULL,
  PRIMARY KEY (`gid`, `article_id`, `tag_id`)
);

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `article_id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `desc` varchar(255) NULL,
  `content_path` varchar(255) NOT NULL COMMENT '文章内容路径',
  `created_time` datetime NOT NULL,
  `modified_time` datetime NULL,
  `deleted_time` datetime NULL DEFAULT 0,
  `state` int NOT NULL DEFAULT 0,
  `cover_image_url` varchar(255) NOT NULL,
  `views` int NOT NULL DEFAULT 0,
  `effective_views` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`article_id`)
);

DROP TABLE IF EXISTS `blog_info`;
CREATE TABLE `blog_info` (
  `blog_name` varchar(50) NULL,
  `site_icon_url` varchar(255) NULL,
  `blog_avatar_url` varchar(255) NULL
);

DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends` (
  `fid` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `link` varchar(255) NOT NULL,
  `avatar_url` varchar(255) NOT NULL,
  `desc` varchar(255) NULL,
  `created_time` datetime(0) NOT NULL,
  `state` int NULL DEFAULT 0,
  PRIMARY KEY (`fid`)
);

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `tag_id` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_time` datetime NOT NULL,
  PRIMARY KEY (`tag_id`)
);

ALTER TABLE `article_tags` DROP FOREIGN KEY `fk__article_tags__tags`;
ALTER TABLE `articles` DROP FOREIGN KEY `fk__articles__article_tags`;
ALTER TABLE `article_tags`
ADD CONSTRAINT `fk__article_tags__tags` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`tag_id`) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE `articles`
ADD CONSTRAINT `fk__articles__article_tags` FOREIGN KEY (`article_id`) REFERENCES `article_tags` (`article_id`) ON DELETE RESTRICT ON UPDATE CASCADE;
