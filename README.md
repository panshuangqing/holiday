# holiday
项目需要 计算工作日时间， 排除周末和 法定节假日， 需要现将  每个节假日 先 按照开始和结束时间 写入带数据库， 
然后按照 线段原理将 时间点 排序到 线段上 计算每一段的时间 diff
粒度是到 天， 数据库的 创建表的语句 
CREATE TABLE `holiday` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标示',
  `time_stamp` int(11) DEFAULT NULL COMMENT '时间戳',
  `is_begin` tinyint(4) DEFAULT '0' COMMENT '是否是假期的开始， 每次在插入假期时是开始和结束时间 一起插入相差86400 秒',
  `is_delete` tinyint(4) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8

一般构造数据的 是使用  'http://timor.tech/api/holiday/year/2020 具体构造 内部有个脚本下次一起上传
计算工作日时间
