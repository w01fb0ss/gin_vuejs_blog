CREATE TABLE `blog_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `created_on` int(11) DEFAULT NULL COMMENT '创建时间',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_on` int(10) unsigned DEFAULT '0'COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_user` (`id`, `username`, `email`, `password`) VALUES (null, 'test', 'test@gmail.com', 'test123456');