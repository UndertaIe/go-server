-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: passwd
-- ------------------------------------------------------
-- Server version	5.7.38

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `passwd_platform`
--

DROP TABLE IF EXISTS `passwd_platform`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `passwd_platform` (
  `platform_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(16) NOT NULL,
  `abbr` varchar(16) NOT NULL,
  `type` varchar(16) DEFAULT NULL,
  `description` varchar(32) DEFAULT NULL,
  `domain` varchar(128) DEFAULT NULL,
  `img_url` varchar(128) DEFAULT NULL,
  `login_url` varchar(128) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`platform_id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `abbr` (`abbr`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `passwd_platform`
--

LOCK TABLES `passwd_platform` WRITE;
/*!40000 ALTER TABLE `passwd_platform` DISABLE KEYS */;
INSERT INTO `passwd_platform` VALUES (3,'B站','bilibili','娱乐','视频传播网站','https://www.bilibili.com/','','https://www.bilibili.com/','2022-08-30 01:00:35','2022-08-30 01:00:35',0),(4,'github','github','技术','代码托管和分享平台','https://github.com/','https://github.githubassets.com/images/modules/site/icons/footer/github-mark.svg','https://github.com/login','2022-08-30 01:10:51','2022-08-30 01:10:51',0),(5,'github_delete','github_delete','技术','代码托管和分享平台','https://github.com/','https://github.githubassets.com/images/modules/site/icons/footer/github-mark.svg','https://github.com/login','2022-08-30 01:31:50','2022-08-30 01:31:50',0),(6,'github_deleted','github_deleted','技术','代码托管和分享平台','https://github.com/','https://github.githubassets.com/images/modules/site/icons/footer/github-mark.svg','https://github.com/login','2022-08-30 01:33:40','2022-08-30 01:34:34',1);
/*!40000 ALTER TABLE `passwd_platform` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `passwd_user`
--

DROP TABLE IF EXISTS `passwd_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `passwd_user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(16) NOT NULL,
  `password` varchar(32) NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `share_mode` tinyint(4) NOT NULL,
  `role` tinyint(4) NOT NULL,
  `profile_img_url` varchar(128) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `sex` tinyint(4) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `phone_number` (`phone_number`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `passwd_user`
--

LOCK TABLES `passwd_user` WRITE;
/*!40000 ALTER TABLE `passwd_user` DISABLE KEYS */;
INSERT INTO `passwd_user` VALUES (5,'user1','root','12345678910',0,0,NULL,'user5_desc',0,'2022-08-29 22:38:32','2022-08-29 22:50:43',0),(6,'user1','root','12345678911',0,0,NULL,'user6_desc',0,'2022-08-29 22:38:33','2022-08-29 22:51:17',0),(7,'user1','root','12345678912',0,0,NULL,'user7_desc',0,'2022-08-29 22:38:34','2022-08-29 22:51:22',0);
/*!40000 ALTER TABLE `passwd_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `passwd_user_account`
--

DROP TABLE IF EXISTS `passwd_user_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `passwd_user_account` (
  `user_account_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `platform_id` int(10) unsigned NOT NULL,
  `password` varchar(32) NOT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  PRIMARY KEY (`user_account_id`),
  UNIQUE KEY `user_id` (`user_id`,`platform_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `passwd_user_account`
--

LOCK TABLES `passwd_user_account` WRITE;
/*!40000 ALTER TABLE `passwd_user_account` DISABLE KEYS */;
INSERT INTO `passwd_user_account` VALUES (1,5,3,'root','2022-08-30 02:53:17','2022-08-30 04:20:37',0),(2,5,4,'user4\'s password','2022-08-30 03:47:39','2022-08-30 04:20:37',0),(3,6,4,'new pwd','2022-08-30 04:10:34','2022-08-30 04:56:03',0);
/*!40000 ALTER TABLE `passwd_user_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `time_test`
--

DROP TABLE IF EXISTS `time_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `time_test` (
  `d` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `time_test`
--

LOCK TABLES `time_test` WRITE;
/*!40000 ALTER TABLE `time_test` DISABLE KEYS */;
INSERT INTO `time_test` VALUES ('2022-08-29 21:22:40'),('2022-08-29 21:22:40');
/*!40000 ALTER TABLE `time_test` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-30  5:05:29
