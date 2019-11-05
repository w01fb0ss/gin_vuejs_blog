CREATE TABLE `blog_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_user` (`id`, `username`, `email`, `password`) VALUES (null, 'test', 'test@gmail.com', 'test123456');