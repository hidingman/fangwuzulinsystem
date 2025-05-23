/*
 Navicat Premium Data Transfer

 Source Server         : 47.113.218.133_3306
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44-log)
 Source Host           : 47.113.218.133:3306
 Source Schema         : fangwuzulinsystem

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44-log)
 File Encoding         : 65001

 Date: 11/05/2025 13:00:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chat
-- ----------------------------
DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `adminid` bigint(20) NULL DEFAULT NULL COMMENT '管理员id',
  `ask` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '提问',
  `reply` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '回复',
  `isreply` int(11) NULL DEFAULT NULL COMMENT '是否回复',
  `isread` int(11) NULL DEFAULT 0 COMMENT '已读/未读(1:已读,0:未读)',
  `uname` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户头像',
  `uimage` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '用户名',
  `type` int(11) NULL DEFAULT 1 COMMENT '内容类型(1:文本,2:图片,3:视频,4:文件,5:表情)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '智能客服' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chat
-- ----------------------------
INSERT INTO `chat` VALUES (1, '2025-04-15 00:34:18', 1, 1, '提问1', '回复1', 1, 1, '用户头像1', 'upload/chat_uimage1.jpg,upload/chat_uimage2.jpg,upload/chat_uimage3.jpg', 1);
INSERT INTO `chat` VALUES (2, '2025-04-15 00:34:18', 2, 1, '提问2', '回复2', 2, 2, '用户头像2', 'upload/chat_uimage2.jpg,upload/chat_uimage3.jpg,upload/chat_uimage4.jpg', 2);
INSERT INTO `chat` VALUES (3, '2025-04-15 00:34:18', 3, 1, '提问3', '回复3', 3, 3, '用户头像3', 'upload/chat_uimage3.jpg,upload/chat_uimage4.jpg,upload/chat_uimage5.jpg', 3);
INSERT INTO `chat` VALUES (4, '2025-04-15 00:34:18', 4, 1, '提问4', '回复4', 4, 4, '用户头像4', 'upload/chat_uimage4.jpg,upload/chat_uimage5.jpg,upload/chat_uimage6.jpg', 4);
INSERT INTO `chat` VALUES (5, '2025-04-15 00:34:18', 5, 1, '提问5', '回复5', 5, 5, '用户头像5', 'upload/chat_uimage5.jpg,upload/chat_uimage6.jpg,upload/chat_uimage7.jpg', 5);
INSERT INTO `chat` VALUES (6, '2025-04-15 00:34:18', 6, 1, '提问6', '回复6', 6, 6, '用户头像6', 'upload/chat_uimage6.jpg,upload/chat_uimage7.jpg,upload/chat_uimage8.jpg', 6);
INSERT INTO `chat` VALUES (7, '2025-04-15 00:34:18', 7, 1, '提问7', '回复7', 7, 7, '用户头像7', 'upload/chat_uimage7.jpg,upload/chat_uimage8.jpg,upload/chat_uimage1.jpg', 7);
INSERT INTO `chat` VALUES (8, '2025-04-15 00:34:18', 8, 1, '提问8', '回复8', 8, 8, '用户头像8', 'upload/chat_uimage8.jpg,upload/chat_uimage1.jpg,upload/chat_uimage2.jpg', 8);

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '配置参数名称',
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '配置参数值',
  `url` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'url',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '配置文件' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES (1, 'picture1', 'upload/picture1.jpg', NULL);
INSERT INTO `config` VALUES (2, 'picture2', 'upload/picture2.jpg', NULL);
INSERT INTO `config` VALUES (3, 'picture3', 'upload/picture3.jpg', NULL);

-- ----------------------------
-- Table structure for discussfangyuanxinxi
-- ----------------------------
DROP TABLE IF EXISTS `discussfangyuanxinxi`;
CREATE TABLE `discussfangyuanxinxi`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `refid` bigint(20) NOT NULL COMMENT '关联表id',
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `avatarurl` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '头像',
  `nickname` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `reply` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '回复内容',
  `thumbsupnum` int(11) NULL DEFAULT 0 COMMENT '赞',
  `crazilynum` int(11) NULL DEFAULT 0 COMMENT '踩',
  `istop` int(11) NULL DEFAULT 0 COMMENT '置顶(1:置顶,0:非置顶)',
  `tuserids` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '赞用户ids',
  `cuserids` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '踩用户ids',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `refid`(`refid`) USING BTREE,
  INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `discussfangyuanxinxi_ibfk_1` FOREIGN KEY (`refid`) REFERENCES `fangyuanxinxi` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `discussfangyuanxinxi_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `guke` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '房源信息评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of discussfangyuanxinxi
-- ----------------------------

-- ----------------------------
-- Table structure for fangwuhuxing
-- ----------------------------
DROP TABLE IF EXISTS `fangwuhuxing`;
CREATE TABLE `fangwuhuxing`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `fangwuhuxing`(`fangwuhuxing`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '房屋户型' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fangwuhuxing
-- ----------------------------
INSERT INTO `fangwuhuxing` VALUES (1, '2025-04-15 00:34:18', '复式');
INSERT INTO `fangwuhuxing` VALUES (2, '2025-04-15 00:34:18', '五房两厅');
INSERT INTO `fangwuhuxing` VALUES (3, '2025-04-15 00:34:18', '两房一厅');
INSERT INTO `fangwuhuxing` VALUES (4, '2025-04-15 00:34:18', '一房一厅');
INSERT INTO `fangwuhuxing` VALUES (5, '2025-04-15 00:34:18', '单间');
INSERT INTO `fangwuhuxing` VALUES (6, '2025-04-15 00:34:18', '四房两厅');
INSERT INTO `fangwuhuxing` VALUES (7, '2025-04-15 00:34:18', '三房一厅');
INSERT INTO `fangwuhuxing` VALUES (8, '2025-04-15 00:34:18', '三房两厅');

-- ----------------------------
-- Table structure for fangwuzulin
-- ----------------------------
DROP TABLE IF EXISTS `fangwuzulin`;
CREATE TABLE `fangwuzulin`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `qishishijian` datetime NULL DEFAULT NULL COMMENT '起始时间',
  `jieshushijian` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `zulinshijian` int(11) NULL DEFAULT NULL COMMENT '租赁时间',
  `zujinjiage` double NULL DEFAULT NULL COMMENT '租金价格',
  `zongji` double NULL DEFAULT NULL COMMENT '总计',
  `yajin` double NULL DEFAULT NULL COMMENT '押金',
  `jifen` int(11) NULL DEFAULT NULL COMMENT '积分',
  `zulinxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '租赁详情',
  `zulinbeizhu` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '租赁备注',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  `ispay` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '未支付' COMMENT '是否支付',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  INDEX `idx_time_range`(`qishishijian`, `jieshushijian`) USING BTREE,
  CONSTRAINT `fangwuzulin_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fangwuzulin_ibfk_2` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '房屋租赁' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fangwuzulin
-- ----------------------------
INSERT INTO `fangwuzulin` VALUES (8, '2025-04-15 00:34:18', '碧桂园3期', 'upload/1746902148.png', '	 三房一厅', '2025-04-15 00:34:18', '2025-04-15 00:34:18', 1, 3500, 7000, 3500, 8, '', '', '123888', '黄民杰', '123456', '齐上', '未支付');

-- ----------------------------
-- Table structure for fangyuanxinxi
-- ----------------------------
DROP TABLE IF EXISTS `fangyuanxinxi`;
CREATE TABLE `fangyuanxinxi`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `fangwudizhi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋地址',
  `mianji` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '面积',
  `zujinjiage` double NULL DEFAULT NULL COMMENT '租金价格',
  `yajinfangshi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '押金方式',
  `fukuanfangshi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '付款方式',
  `fangwuchaoxiang` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋朝向',
  `yajin` double NULL DEFAULT NULL COMMENT '押金',
  `xiaoqu` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '小区',
  `fangwushipin` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋视频',
  `loudongdanyuan` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '楼栋单元',
  `fanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房号',
  `jifen` int(11) NULL DEFAULT NULL COMMENT '积分',
  `fangwujiegou` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋结构',
  `fangwuzhuangtai` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋状态',
  `fangchanzhengbianhao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房产证编号',
  `fangchanzhengzhaopian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房产证照片',
  `fangzhuxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房主姓名',
  `fangzhushenfenzheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房主身份证',
  `fangzhudianhua` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房主电话',
  `fangwuxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋详情',
  `fabushijian` datetime NULL DEFAULT NULL COMMENT '发布时间',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `thumbsupnum` int(11) NULL DEFAULT 0 COMMENT '赞',
  `crazilynum` int(11) NULL DEFAULT 0 COMMENT '踩',
  `clicktime` datetime NULL DEFAULT NULL COMMENT '最近点击时间',
  `clicknum` int(11) NULL DEFAULT 0 COMMENT '点击次数',
  `discussnum` int(11) NULL DEFAULT 0 COMMENT '评论数',
  `storeupnum` int(11) NULL DEFAULT 0 COMMENT '收藏数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `idx_search`(`zujinjiage`, `mianji`, `fangwuzhuangtai`) USING BTREE,
  CONSTRAINT `fangyuanxinxi_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '房源信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fangyuanxinxi
-- ----------------------------
INSERT INTO `fangyuanxinxi` VALUES (1, '2025-04-15 00:34:17', '招商海月花园', 'upload/1746901822.png', '五房两厅', '广东省深圳市南山区招商海月花园2栋301', '180', 9800, '年付', '转账', '朝东', 9800, '招商海月花园', '', '2栋', '301', 20, '中间承重', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian1.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian2.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian3.jpg', '杨幂', '440300199101010001', '13823888881', '<h3>小区基本信息</h3><ul><li>位置：位于深圳市南山区蛇口后海海月路43号。</li><li>开发商：深圳招商房地产有限公司</li><li>。</li><li>物业公司：深圳招商局物业管理有限公司</li><li>。</li><li>建成年代：1999-2009年</li><li>。</li><li>总户数：约3400户</li><li>。</li><li>容积率：1.67</li><li>。</li><li>绿化率：35%-45%</li><li>。</li><li>物业费：3.3元/月/㎡</li><li>。</li></ul><h3>户型与建筑</h3><ul><li>户型设计：户型方正，实用率高，主力户型为3房2厅</li><li>。小区内有多种户型可供选择，包括2房、3房、4房等</li><li>。</li><li>建筑风格：小区建筑风格现代，楼间距合理，采光和通风条件良好</li><li>。</li></ul><h3>周边配套</h3><ul><li>交通：</li><li>地铁：距离2号线海月站仅280米，出行十分便利</li><li class=\"ql-indent-1\">。</li><li class=\"ql-indent-1\">公交：附近有多个公交站，如登良地铁站、兰园等。</li><li>教育：</li><li>幼儿园：小区内有配套的海月谷双语幼儿园</li><li class=\"ql-indent-1\">。</li><li>小学：周边有后海小学</li><li class=\"ql-indent-1\">。</li><li>中学：附近有育才三中</li><li class=\"ql-indent-1\">。</li><li>商业：</li><li class=\"ql-indent-1\">购物中心：附近有宝能太古城购物中心。</li><li>超市：周边有沃尔玛超市</li><li class=\"ql-indent-1\">。</li><li>医疗：附近有蛇口医院、联合医院</li><li>。</li><li>休闲：</li><li>公园：周边有四海公园</li><li class=\"ql-indent-1\">。</li><li>运动设施：小区内设有网球场、泳池、健身会所等</li><li class=\"ql-indent-1\">。</li></ul><h3>小区特色</h3><ul><li>社区配套：小区内有社区体育馆、招商会但昭义钢琴中心等特色设施</li><li>。</li><li>人文生态：小区依托蛇口丰富的服务资源，构建了一个滨海、成熟、高尚的人文生态社区</li><li>。</li><li>居住品质：小区环境优美，绿化率高，居住品质高</li><li>。</li></ul>', '2025-04-15 00:34:17', '123888', '黄民杰', 1, 1, '2025-04-15 00:34:17', 1, 0, 2);
INSERT INTO `fangyuanxinxi` VALUES (2, '2025-04-15 00:34:17', '文竹园', 'upload/fangyuanxinxi_fangwutupian2.jpg,upload/fangyuanxinxi_fangwutupian3.jpg,upload/fangyuanxinxi_fangwutupian4.jpg', '单间', '广东省深圳市南山区文竹园1栋102', '30', 2200, '押一付一', '转账', '朝南', 2200, '文竹园', '', '1栋', '102', 2, '单间', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian2.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian3.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian4.jpg', '黄兴从', '440300199202020002', '13823888882', '<ul><li>位置：文竹园位于深圳市南山区蛇口工业八路171号，东起后海大道，西靠公园路，南临工业七路，北至工业八路</li><li>。</li><li>开发商：由招商局地产开发有限公司开发建设</li><li>。</li><li>竣工时间：1992年</li><li>。</li><li>物业类型：普通住宅</li><li>。</li><li>总户数：约680户。</li><li>建筑面积：约112000平方米。</li><li>容积率：1.85</li><li>。</li><li>绿化率：35%-50%</li><li>。</li><li>物业公司：早期由招商其乐物业管理，目前为业主自管</li><li>。</li><li>物业费：0.8元/平方米·月</li><li>。</li><li>停车位：约300个</li><li>。</li></ul><h3>户型与建筑</h3><ul><li>户型：小区户型多样，以小户型为主，常见的有58.5平米的2房、68平米的3房、88平米的3房等</li><li>。</li><li>建筑结构：板楼设计，多为7层高的多层建筑</li><li>。</li></ul><h3>周边配套</h3><ul><li>交通：</li><li>地铁：距离2号线（8号线）海月站约600米</li><li class=\"ql-indent-1\">。</li><li>公交：附近有文竹园公交站，多条公交线路经过，如B737路、B817路等</li><li class=\"ql-indent-1\">。</li><li>教育：</li><li>幼儿园：小区内有配套幼儿园</li><li class=\"ql-indent-1\">。</li><li>小学：距离小区最近的小学是育才第二小学</li><li class=\"ql-indent-1\">。</li><li>中学：周边还有育才二中等学校</li><li class=\"ql-indent-1\">。</li><li>商业：</li><li class=\"ql-indent-1\">购物中心：附近有宝能·allcity购物中心-南区、宝能allcity购物中心、宝能太古城花园购物中心。</li><li class=\"ql-indent-1\">超市：200米内有华城超市，附近还有后海市场、沃尔玛等。</li><li>休闲：</li><li>公园：周边有四海公园，环境优雅</li><li class=\"ql-indent-1\">。</li><li>运动设施：小区内有花园、健身设施、儿童游乐设施等</li><li class=\"ql-indent-1\">。</li></ul><p><br></p>', '2025-04-15 00:34:17', '123888', '黄民杰', 2, 2, '2025-04-15 00:34:17', 2, 0, 2);
INSERT INTO `fangyuanxinxi` VALUES (3, '2025-04-15 00:34:17', '君汇新天', 'upload/fangyuanxinxi_fangwutupian3.jpg,upload/fangyuanxinxi_fangwutupian4.jpg,upload/fangyuanxinxi_fangwutupian5.jpg', '三房两厅', '广东省深圳市南山区君汇新天3栋2201', '120', 4500, '半年付', '转账', '朝南', 45000, '君汇新天', '', '3栋', '2201', 8, '中间承重', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian3.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian4.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian5.jpg', '黄芩为', '440300199303030003', '13823888883', '<h3>小区基本信息</h3><ul><li>位置：君汇新天位于深圳市南山区中心路（深圳湾段）2269号（东滨路和中心路交汇处西南角）。</li><li>开发商：深圳市新天时代投资有限公司。</li><li>物业公司：深圳市新天时代物业管理有限公司。</li><li>建成时间：2009年12月10日。</li><li>容积率：2.0。</li><li>绿化率：30%。</li><li>总户数：744户。</li><li>车位数：地上424个，地下1780个。</li></ul><h3>3栋2201户型特点</h3><ul><li>户型面积：君汇新天的户型以大户型为主，常见的户型面积包括156-283平方米的四房和五房。3栋2201的具体户型面积未明确，但大概率属于大户型。</li><li>户型设计：小区户型设计注重南北通透，采光和通风条件良好</li><li>。3栋2201大概率也具备这一特点。</li><li>层高：小区层高为3.4米</li><li>，这使得居住空间更加宽敞舒适。</li></ul><h3>周边配套</h3><ul><li>交通配套：</li><li class=\"ql-indent-1\">地铁：距离地铁2号线海月站仅约2分钟步行距离，出行非常便利。</li><li class=\"ql-indent-1\">公交：附近有多个公交站点，如后海滨路口、曦湾华府等，多条公交线路经过。</li><li>教育配套：</li><li class=\"ql-indent-1\">幼儿园：小区内部有配套幼儿园。</li><li>中小学：周边有北师大南山附中、后海小学、育才小学、育才一中、二中、蛇口中学、蛇口小学等</li><li class=\"ql-indent-1\">。</li><li>商业配套：</li><li>购物中心：附近有宝能太古城、宝能allcity购物中心（南山店）-南区、保利文化广场等</li><li class=\"ql-indent-1\">。</li><li>超市：周边有华润万家便利超市（蔚蓝海岸店）、华润万家便利超市（支六路店）、金汇城百货（蛇口店）、永辉超市（深圳来福士广场店）等</li><li class=\"ql-indent-1\">。</li><li>医疗配套：周边有蛇口人民医院深圳湾社康中心、小珂丽格医疗美容、深圳蒳美迩医疗美容门诊部、ISYOU医疗美容、蛇口人民医院等。</li><li>休闲配套：</li><li>公园：周边有15公里滨海生态长廊、四海公园等</li><li class=\"ql-indent-1\">。</li><li>体育设施：小区内部有健身房、游泳池、乒乓球、羽毛球场、跑道、运动中心等</li><li class=\"ql-indent-1\">。</li></ul><h3>小区其他特点</h3><ul><li>建筑风格：由法国欧博米歇尔大师率领国际团队设计，结合国际豪宅理念打造而成的滨海纯大户豪宅。</li><li>园林景观：近4万平米园林由法国欧博精心设计，现代自然风格，注重自然、舒适空间的景观诉求。项目为双泳池设计，中心无边际泳池约600平米，儿童泳池约100平米。</li><li>物业费：4.3-4.3元/㎡</li><li>。</li></ul><p><br></p>', '2025-04-15 00:34:17', '123888', '黄民杰', 3, 3, '2025-04-15 00:34:17', 4, 0, 3);
INSERT INTO `fangyuanxinxi` VALUES (4, '2025-04-15 00:34:17', '宏丰新城', 'upload/1746901158.png', '一房一厅', '广东省茂名市茂南区宏丰新城3栋1202', '40', 2400, '押一付三', '转账', '朝北', 2400, '宏丰新城', '', '3栋', '1202', 4, '南北通透', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian4.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian5.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian6.jpg', '黄子涵', '440300199404040004', '13823888884', '<h3>基本信息</h3><ul><li>地理位置 ：位于茂名市茂南区油城十路与茂名大道交汇处，处于茂名市城区未来城央行政商务中心，西靠茂名大道交通大动脉，与政务、商务、文化休闲高度集中的繁华油城九路相连</li><li>。</li><li>开发商 ：茂名市宏丰长湖房地产开发有限公司</li><li>。</li><li>物业类型 ：包括住宅、公寓</li><li>。</li><li>占地面积与建筑面积 ：占地面积 176551㎡，建筑面积 1073805.57㎡</li><li>。</li><li>容积率与绿化率 ：容积率为 4.9，绿化率为 26%</li><li>。</li><li>规划户数 ：总规划户数为 4001户</li><li>。</li><li>车位信息 ：共有 8354 个机动车位，其中地上 425 个，地下 7928 个。</li><li>产权年限 ：普通住宅产权为 70 年</li><li>。</li></ul><h3>小区配套</h3><ul><li>教育资源 ：小区内部有配套幼儿园，周边有茂名市区优质的九年制学校如茂名市育才、文悦、祥和学校，还有茂名市第九小学、新世纪学校等，且靠近茂名市第一中学附属学校</li><li>。</li><li>商业配套 ：北有大型购物中心东汇城，自身规划有 1.3 公里商业街、星级酒店、写字楼、公寓等设施，周边还有沃尔玛广场、文化广场、宏丰新城商业广场以及万达广场等</li><li>。</li><li>医疗配套 ：周边有茂名康复医院、石化医院等医疗机构</li><li>。</li><li>休闲娱乐 ：小区自身建设有游泳池、儿童乐园、图书馆、乒乓球室、棋牌室及大型市政休闲广场，附近还有文化广场</li><li>。</li><li>交通配套 ：临近茂名大道等交通干线，一公里以内有 6 个公交站，多路公交车经过，交通十分便利</li><li>。</li></ul>', '2025-04-15 00:34:17', '123888', '黄民杰', 4, 4, '2025-04-15 00:34:17', 4, 0, 4);
INSERT INTO `fangyuanxinxi` VALUES (5, '2025-04-15 00:34:18', '名雅世家', 'upload/1746900822.png', '两房一厅', '广东省茂名市茂南区名雅世家1栋502', '60', 2000, '押一付三', '转账', '朝南', 2000, '名雅世家', '', '1栋', '502', 5, '侧边承重', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian5.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian6.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian7.jpg', '黄子涵', '440300199505050005', '13823888885', '<h3>周边配套</h3><ul><li>交通：小区周边交通便利，官山五路、荔晶新城等公交站点途经多条公交线路，如5路、13路等。</li><li>购物：附近有沃尔玛超级市场、三有家居广场等大型购物场所，祥和社区肉菜市场也步行可达。</li><li>教育：教育资源丰富，名雅幼儿园、文化广场实验幼儿园距离小区很近</li><li>。茂名市一级学校市九小、全国示范性普通中学广东省重点中学茂名市一中也近在咫尺</li><li>。</li><li>医疗：车行10分钟内可到达茂名市石化医院、茂名市康复医院等医疗机构</li><li>。</li><li>休闲：小区内部有38000平方米东南亚风情园林，还有文化广场等休闲场所</li><li>。周边3公里内有多个公园</li><li>。</li></ul>', '2025-04-15 00:34:18', '12888', '黄子涵', 5, 5, '2025-04-15 00:34:18', 5, 0, 5);
INSERT INTO `fangyuanxinxi` VALUES (6, '2025-04-15 00:34:18', '保利东湾三期', 'upload/1746900500.png', '一房一厅', '广东省茂名市茂南区保利东湾三期5栋1003', '45', 1200, '押一付一', '转账', '朝南', 1200, '保利东湾三期', '', '5栋', '1003', 2, '中间承重', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian6.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian7.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian8.jpg', '黄子涵', '440300199606060006', '13823888886', '<h3>交通与配套</h3><ul><li>交通便利：保利东湾三期周边交通网络发达，大园桥、大园路、保利大道等主干道贯通，轻松畅达茂名全城</li><li>。</li><li>教育资源：小区内部配套有九年一贯制公立学校——愉园东湾学校以及市第一幼儿园东湾分园，周边还有茂名市第一中学等优质教育资源</li><li>。</li><li>商业配套：小区内建有约2万㎡的风情商业街和2000㎡的鲜菜市场，引进了肯德基、嘉荣超市等品牌商家。</li><li>休闲景观：小区内设有七大主题园林，包括儿童乐园、550m星际跑道等。外部拥有小东江十里滨江景观及2.1公里滨江生态公园，还有奥林匹克公园等七大公园环绕</li><li>。</li><li>医疗资源：距离小区2公里内有茂名中医院、妇幼保健院等三甲医院</li><li>。</li></ul><h3>小区环境与服务</h3><ul><li>绿化与景观：小区绿化率为33.03%，配备约8万方的三大中芯园林，居民下楼即可享受江园双景。</li><li>物业服务：保利东湾三期的物业服务由保利物业管理有限公司提供，物业费为1.8元/㎡·月，能够为业主提供良好的居住体验</li><li>。</li></ul>', '2025-04-15 00:34:18', '123888', '黄民杰', 6, 6, '2025-04-15 00:34:18', 6, 0, 6);
INSERT INTO `fangyuanxinxi` VALUES (7, '2025-04-15 00:34:18', '粤西明珠', 'upload/1746900184.png', '复式', '广东省茂名市茂南区粤西明珠2栋2301', '60', 1500, '押一付三', '转账', '朝东', 2500, '九龙湾一号', '', '二栋', '2301', 7, '中梁支撑', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian7.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian8.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian1.jpg', '黄天气', '440300199707070007', '13823888887', '<h3>小区配套</h3><ul><li>交通：小区周边交通便利，附近有多个公交站，如**粤西明珠(市十中)**站，距离小区仅64米，有11路、17路、210路等多条公交线路。</li><li>教育：周边教育资源丰富，包括粤西明珠幼儿园、桥北小学（茂名市第十九小学）、茂名市第十中学等</li><li>。</li><li>购物：附近有1+1购物广场、兴发购物广场（官渡店）等购物中心</li><li>，购物便利。</li><li>医疗：周边3公里内有4个医院</li><li>，能够满足居民的基本医疗需求。</li></ul><h3>小区环境与服务</h3><ul><li>绿化与环境：小区绿化率为54.5%，容积率为2.64</li><li>，整体环境较好。</li><li>物业服务：物业由茂名市天华房地产开发有限公司提供，物业费为0.85元/㎡·月</li></ul>', '2025-04-15 00:34:18', '13888', '黄七楼', 7, 7, '2025-04-15 00:34:18', 7, 0, 7);
INSERT INTO `fangyuanxinxi` VALUES (8, '2025-04-15 00:34:18', '碧桂园3期', 'upload/1746899736.png', '三房一厅', '广东省茂名市茂南区碧桂园3期3栋1203', '120', 3500, '半年付', '转账', '坐北朝南', 3500, '碧桂园3期', 'upload/1746899502.mp4', '3栋', '房号8', 8, '大厅承重', '未租', '房产证编号', 'upload/fangyuanxinxi_fangchanzhengzhaopian8.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian1.jpg,upload/fangyuanxinxi_fangchanzhengzhaopian2.jpg', '黄民杰', '440300199808080008', '13823888888', '<h3>位置与周边</h3><ul><li>小区位置：该房子位于茂名市茂南区进园路北侧</li><li>，地处茂名城市的“休闲度假板块”，连接茂名唯一的省级森林公园，周边环境宁静，空气清新，居住舒适</li><li>。</li><li>交通：周边有环市西路等主要道路，205、206、207路等多条公交线路可直达社区</li><li>，交通便利。</li><li>商业：附近有中港城超市等商业配套</li><li>，可满足日常购物需求。</li></ul><h3>小区与房屋整体情况</h3><ul><li>小区概况：茂名碧桂园由茂名市碧桂园房地产开发有限公司开发，小区建筑类型包括板楼、塔楼、板塔结合</li><li>，总规划户数为2366户</li><li>，绿化率为35%</li><li>，容积率为0.79</li><li>，居住氛围较好。</li><li>房屋户型：三房一厅的户型，南北朝向，通常具有较好的通风和采光条件，能保证室内阳光充足，空气流通，居住舒适度较高</li><li>。</li></ul><h3>房屋价格</h3><ul><li>截至2025年4月，茂名碧桂园的参考均价为5929元/㎡</li></ul>', '2025-04-15 00:34:18', '123888', '黄民杰', 9, 8, '2025-04-15 00:34:18', 10, 0, 8);

-- ----------------------------
-- Table structure for guke
-- ----------------------------
DROP TABLE IF EXISTS `guke`;
CREATE TABLE `guke`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `mima` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '姓名',
  `xingbie` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '性别',
  `shouji` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机',
  `touxiang` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '头像',
  `shenfenzheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '身份证',
  `minzu` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '民族',
  `jiguan` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '籍贯',
  `chushengriqi` date NULL DEFAULT NULL COMMENT '出生日期',
  `nianling` int(11) NULL DEFAULT NULL COMMENT '年龄',
  `xueli` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '学历',
  `zhuzhi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '住址',
  `xuqiu` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '需求',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `zhanghao`(`zhanghao`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '顾客' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of guke
-- ----------------------------
INSERT INTO `guke` VALUES (19, '2025-05-11 02:37:33', '123456', '123456', '齐上', '', '', '', '', '', '', NULL, 0, '', '', '');
INSERT INTO `guke` VALUES (20, '2025-05-11 02:38:34', '123456789', '123456789', '黄琦琪', '', '', '', '', '', '', NULL, 0, '', '', '');

-- ----------------------------
-- Table structure for kaoqindaka
-- ----------------------------
DROP TABLE IF EXISTS `kaoqindaka`;
CREATE TABLE `kaoqindaka`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `dakaleixing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '打卡类型',
  `dakashijian` datetime NULL DEFAULT NULL COMMENT '打卡时间',
  `dakaxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '打卡详情',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `touxiang` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '头像',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  CONSTRAINT `kaoqindaka_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '考勤打卡' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of kaoqindaka
-- ----------------------------

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `introduction` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '简介',
  `typename` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '分类名称',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '发布人',
  `headportrait` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '头像',
  `clicknum` int(11) NULL DEFAULT 0 COMMENT '点击次数',
  `clicktime` datetime NULL DEFAULT NULL COMMENT '最近点击时间',
  `thumbsupnum` int(11) NULL DEFAULT 0 COMMENT '赞',
  `crazilynum` int(11) NULL DEFAULT 0 COMMENT '踩',
  `storeupnum` int(11) NULL DEFAULT 0 COMMENT '收藏数',
  `picture` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `typename`(`typename`) USING BTREE,
  CONSTRAINT `news_ibfk_1` FOREIGN KEY (`typename`) REFERENCES `newstype` (`typename`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '公告资讯' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news` VALUES (9, '2025-05-11 02:41:46', '深圳市南山区招商海月花园优质房源出租公告', '', '咨询类', '黄民杰', '', 0, '2025-05-11 02:41:46', 0, 0, 0, 'upload/1746902503.png', '<h2>一、房源信息</h2><ul><li>小区名称：招商海月花园</li><li>位置：深圳市南山区蛇口后海海月路 43 号</li><li>户型：3 房 2 厅 2 卫，建筑面积约 120 平方米</li><li>楼层：总楼层 18 层，出租房源位于第 10 层，视野开阔，采光充足</li><li>装修情况：精装修，家电家具齐全，可拎包入住。房屋内部装修风格现代简约，温馨舒适，配备品牌家电，包括空调、冰箱、洗衣机、电视等，满足您的日常居住需求。</li></ul><h2>二、租金与付款方式</h2><ul><li>租金：9800 元 / 月，价格面议，性价比高，适合追求高品质居住环境的租客。</li><li>付款方式：押二付三，即支付两个月租金作为押金，首次支付三个月租金。</li></ul><h2>三、房屋优势</h2><ul><li>交通便利：距离地铁 2 号线海月站仅 280 米，步行 3 - 5 分钟即可到达，周边公交线路丰富，出行十分便捷。</li><li>教育资源丰富：小区内配套有海月谷双语幼儿园，周边还有后海小学、育才三中等优质学校，为有孩子的家庭提供良好的教育环境。</li><li>商业配套完善：附近有宝能太古城购物中心、沃尔玛超市等，满足您的购物、餐饮、娱乐等需求。</li><li>医疗资源近在咫尺：周边有蛇口医院、联合医院等医疗机构，为您的健康保驾护航。</li><li>休闲设施齐全：小区内设有网球场、泳池、健身会所等运动设施，周边还有四海公园，让您在繁忙的生活中也能享受休闲时光。</li><li>社区环境优美：小区绿化率高，环境安静舒适，是您放松身心的理想居所。</li></ul><h2>四、联系方式</h2><ul><li>联系人：[房东姓名]</li><li>联系电话：[房东电话]</li><li>看房时间：周一至周日，上午 9:00 - 12:00，下午 14:00 - 18:00，可提前预约看房时间。</li></ul><h2>五、温馨提示</h2><ul><li>本房源不接受宠物入住，希望租客能够爱护房屋设施，保持室内整洁。</li><li>租赁期间，水电费、物业费等由租客承担，物业费为 3.3 元 / 月 / ㎡。</li><li>为保障双方权益，签订租赁合同前，请仔细阅读合同条款，如有疑问可随时咨询。</li></ul><p><br></p>');

-- ----------------------------
-- Table structure for newstype
-- ----------------------------
DROP TABLE IF EXISTS `newstype`;
CREATE TABLE `newstype`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `typename` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `newstype_jw39`(`typename`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '公告资讯分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of newstype
-- ----------------------------
INSERT INTO `newstype` VALUES (9, '2025-05-11 02:39:40', '咨询类');
INSERT INTO `newstype` VALUES (10, '2025-05-11 02:39:59', '通知');

-- ----------------------------
-- Table structure for popupremind
-- ----------------------------
DROP TABLE IF EXISTS `popupremind`;
CREATE TABLE `popupremind`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `userid` bigint(20) NOT NULL COMMENT '发布人id',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '个人' COMMENT '类型',
  `brief` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '简介',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `remindtime` datetime NULL DEFAULT NULL COMMENT '提醒时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '弹窗提醒' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of popupremind
-- ----------------------------
INSERT INTO `popupremind` VALUES (1, '2025-04-15 00:34:18', 1, '标题1', '个人', '简介1', '内容1', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (2, '2025-04-15 00:34:18', 2, '标题2', '个人', '简介2', '内容2', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (3, '2025-04-15 00:34:18', 3, '标题3', '个人', '简介3', '内容3', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (4, '2025-04-15 00:34:18', 4, '标题4', '个人', '简介4', '内容4', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (5, '2025-04-15 00:34:18', 5, '标题5', '个人', '简介5', '内容5', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (6, '2025-04-15 00:34:18', 6, '标题6', '个人', '简介6', '内容6', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (7, '2025-04-15 00:34:18', 7, '标题7', '个人', '简介7', '内容7', '2025-04-15 00:34:18');
INSERT INTO `popupremind` VALUES (8, '2025-04-15 00:34:18', 8, '标题8', '个人', '简介8', '内容8', '2025-04-15 00:34:18');

-- ----------------------------
-- Table structure for storeup
-- ----------------------------
DROP TABLE IF EXISTS `storeup`;
CREATE TABLE `storeup`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `userid` bigint(20) NOT NULL COMMENT '用户id',
  `refid` bigint(20) NULL DEFAULT NULL COMMENT '商品id',
  `tablename` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '表名',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `picture` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '图片',
  `type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '1' COMMENT '类型',
  `inteltype` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '推荐类型',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `refid`(`refid`) USING BTREE,
  INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `storeup_ibfk_1` FOREIGN KEY (`refid`) REFERENCES `fangyuanxinxi` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `storeup_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `guke` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '收藏表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of storeup
-- ----------------------------
INSERT INTO `storeup` VALUES (2, '2025-05-11 02:46:36', 19, 1, 'fangyuanxinxi', '招商海月花园', 'upload/1746901822.png', '1', '', '');
INSERT INTO `storeup` VALUES (3, '2025-05-11 02:49:20', 19, 8, 'fangyuanxinxi', '碧桂园3期', 'upload/1746899736.png', '21', '', '');

-- ----------------------------
-- Table structure for tuizushenqing
-- ----------------------------
DROP TABLE IF EXISTS `tuizushenqing`;
CREATE TABLE `tuizushenqing`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `yajin` double NULL DEFAULT NULL COMMENT '押金',
  `tuizushijian` datetime NULL DEFAULT NULL COMMENT '退租时间',
  `tuizuyuanyin` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '退租原因',
  `tuizuxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '退租详情',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  `sfsh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '待审核' COMMENT '是否审核',
  `shhf` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '审核回复',
  `ispay` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '未支付' COMMENT '是否支付',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  CONSTRAINT `tuizushenqing_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `tuizushenqing_ibfk_2` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '退租申请' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tuizushenqing
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `image` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  `role` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '管理员' COMMENT '角色',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '新增时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'admin', 'admin', 'upload/image1.jpg', '管理员', '2025-04-15 00:34:18');

-- ----------------------------
-- Table structure for yuangong
-- ----------------------------
DROP TABLE IF EXISTS `yuangong`;
CREATE TABLE `yuangong`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '工号',
  `mima` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '员工姓名',
  `xingbie` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '性别',
  `touxiang` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '头像',
  `shenfenzheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '身份证',
  `dianhua` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '电话',
  `minzu` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '民族',
  `jiguan` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '籍贯',
  `chushengriyue` date NULL DEFAULT NULL COMMENT '出生日月',
  `nianling` int(11) NULL DEFAULT NULL COMMENT '年龄',
  `xueli` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '学历',
  `zhuzhi` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '住址',
  `tezhang` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '特长',
  `ziwopingjia` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '自我评价',
  `gongzuoanpai` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '工作安排',
  `gongzixinxi` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '工资信息',
  `jifen` int(11) NULL DEFAULT NULL COMMENT '积分',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `gonghao`(`gonghao`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '员工' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of yuangong
-- ----------------------------
INSERT INTO `yuangong` VALUES (29, '2025-05-11 02:31:23', '123888', 'hmj.1807446746', '黄民杰', '男', 'upload/1746901873.jpg', '440419197310311212', '18124593333', '汉', '', '1973-02-07', 0, '', '', '', '', '', '', 0);
INSERT INTO `yuangong` VALUES (30, '2025-05-11 02:32:12', '12888', 'hmj.1807446746', '黄子涵', '男', 'upload/1746901971.jpg', '', '', '', '', NULL, 0, '', '', '', '', '', '', 0);
INSERT INTO `yuangong` VALUES (31, '2025-05-11 02:33:35', '13888', 'hmj.1807446746', '黄七楼', '', 'upload/1746902013.jpg', '', '', '', '', NULL, 0, '', '', '', '', '', '', 0);

-- ----------------------------
-- Table structure for yuyuekanfang
-- ----------------------------
DROP TABLE IF EXISTS `yuyuekanfang`;
CREATE TABLE `yuyuekanfang`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `yuyueshijian` datetime NULL DEFAULT NULL COMMENT '预约时间',
  `yuyuexiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '预约详情',
  `yuyueneirong` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '预约内容',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  `shouji` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机',
  `sfsh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '待审核' COMMENT '是否审核',
  `shhf` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '审核回复',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  CONSTRAINT `yuyuekanfang_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `yuyuekanfang_ibfk_2` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '预约看房' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of yuyuekanfang
-- ----------------------------

-- ----------------------------
-- Table structure for zufangyixiang
-- ----------------------------
DROP TABLE IF EXISTS `zufangyixiang`;
CREATE TABLE `zufangyixiang`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `yixiangbiaoti` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '意向标题',
  `yixiangtupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '意向图片',
  `zufangdidian` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '租房地点',
  `lixiangjiawei` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '理想价位',
  `qiwanghuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '期望户型',
  `dengjishijian` datetime NULL DEFAULT NULL COMMENT '登记时间',
  `yixiangxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '意向详情',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  CONSTRAINT `zufangyixiang_ibfk_1` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '租房意向' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of zufangyixiang
-- ----------------------------
INSERT INTO `zufangyixiang` VALUES (9, '2025-05-11 02:48:08', '租房需求', '', '深圳市南山区四海公园', '6000-8000', '三房两厅', '2025-05-11 02:48:08', '', '123456', '齐上');

-- ----------------------------
-- Table structure for zulinhetong
-- ----------------------------
DROP TABLE IF EXISTS `zulinhetong`;
CREATE TABLE `zulinhetong`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `yajin` double NULL DEFAULT NULL COMMENT '押金',
  `jifen` int(11) NULL DEFAULT NULL COMMENT '积分',
  `hetongbianhao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '合同编号',
  `hetongfujian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '合同附件',
  `qiandingshijian` datetime NULL DEFAULT NULL COMMENT '签订时间',
  `hetongxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '合同详情',
  `hetongshixiang` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '合同事项',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  `sfsh` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '待审核' COMMENT '是否审核',
  `shhf` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '审核回复',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `hetongbianhao`(`hetongbianhao`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  CONSTRAINT `zulinhetong_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `zulinhetong_ibfk_2` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '租赁合同' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of zulinhetong
-- ----------------------------

-- ----------------------------
-- Table structure for zulinpingjia
-- ----------------------------
DROP TABLE IF EXISTS `zulinpingjia`;
CREATE TABLE `zulinpingjia`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `addtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fangwumingcheng` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋名称',
  `fangwutupian` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '房屋图片',
  `fangwuhuxing` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '房屋户型',
  `pingjiapingfen` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '评价评分',
  `gongsipingjia` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '公司评价',
  `yuangongpingjia` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '员工评价',
  `pingjiaxiangqing` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '评价详情',
  `gonghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '工号',
  `yuangongxingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '员工姓名',
  `zhanghao` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '账号',
  `xingming` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '姓名',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `gonghao`(`gonghao`) USING BTREE,
  INDEX `zhanghao`(`zhanghao`) USING BTREE,
  CONSTRAINT `zulinpingjia_ibfk_1` FOREIGN KEY (`gonghao`) REFERENCES `yuangong` (`gonghao`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `zulinpingjia_ibfk_2` FOREIGN KEY (`zhanghao`) REFERENCES `guke` (`zhanghao`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '租赁评价' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of zulinpingjia
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
