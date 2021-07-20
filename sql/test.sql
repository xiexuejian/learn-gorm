/*
SQLyog Ultimate v11.25 (64 bit)
MySQL - 8.0.25 : Database - test
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `test`;

/*Table structure for table `score` */

DROP TABLE IF EXISTS `score`;

CREATE TABLE `score` (
  `id` int NOT NULL AUTO_INCREMENT,
  `stu_id` int NOT NULL,
  `c_name` varchar(20) DEFAULT NULL,
  `grade` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `score` */

insert  into `score`(`id`,`stu_id`,`c_name`,`grade`) values (1,901,'计算机',98),(2,901,'英语',80),(3,902,'计算机',65),(4,902,'中文',88),(5,903,'中文',95),(6,904,'计算机',70),(7,904,'英语',92),(8,905,'英语',94),(9,906,'计算机',90),(10,906,'英语',85);

/*Table structure for table `student` */

DROP TABLE IF EXISTS `student`;

CREATE TABLE `student` (
  `id` int NOT NULL,
  `name` varchar(20) NOT NULL,
  `sex` varchar(4) DEFAULT NULL,
  `birth` year DEFAULT NULL,
  `department` varchar(20) DEFAULT NULL,
  `address` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `student` */

insert  into `student`(`id`,`name`,`sex`,`birth`,`department`,`address`) values (1,'王二','男',1997,'计算机系','河南省郑州市'),(901,'张老大','男',1985,'计算机系','北京市海淀区'),(902,'张老二','男',1986,'中文系','北京市昌平区'),(903,'张三','女',1990,'中文系','湖南省永州市'),(904,'李四','男',1990,'英语系','辽宁省阜新市'),(905,'王五','女',1991,'英语系','福建省厦门市'),(906,'王六','男',1988,'计算机系','湖南省衡阳市');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
