CREATE TABLE holiday (
    id int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标示',
    time_stamp int(11) DEFAULT NULL COMMENT '时间戳',
    is_begin tinyint(4) DEFAULT '0' COMMENT '是否是假期的开始， 每次在插入假期时是开始和结束时间 一起插入相差86400 秒',
    is_delete tinyint(4) DEFAULT '0' COMMENT '是否删除', PRIMARY KEY (id) )
ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8