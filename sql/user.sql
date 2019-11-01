DROP DATABASE IF EXISTS account_db;
CREATE DATABASE account_db;
USE account_db;

CREATE TABLE `user`(
  `id` bigint(20) AUTO_INCREMENT NOT NULL,
  `user_name` varchar(50) NOT NULL,
  `phone` int(11) NOT NULL,
  `pwd` varchar(1024) NOT NULL,
  `head_url` varchar(1000) NOT NULL,
  `level` int(10) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8;