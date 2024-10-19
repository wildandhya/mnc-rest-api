-- MySQL dump 10.13  Distrib 8.0.33, for Linux (aarch64)
--
-- Host: localhost    Database: example
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `payment_id` longtext,
  `user_id` longtext,
  `amount` bigint DEFAULT NULL,
  `remarks` longtext,
  `balance_before` bigint DEFAULT NULL,
  `balance_after` bigint DEFAULT NULL,
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES ('7e22726b-5647-4491-a4da-ac3b5ee88b2f','ae38adfc-6646-4633-b742-e6557155e4d1',100000,'Pulsa Telkomsel 100k',40000,-60000,'2024-10-19 17:16:15.542','2024-10-19 17:16:15.542'),('83a3cc80-8336-48f0-ab2a-6db672eebdd5','ae38adfc-6646-4633-b742-e6557155e4d1',20000,'Pulsa Telkomsel 100k',40000,20000,'2024-10-19 17:18:00.385','2024-10-19 17:18:00.385'),('577c8a81-24f0-47ef-8284-6ddd8926a209','ae38adfc-6646-4633-b742-e6557155e4d1',20000,'Pulsa Telkomsel 100k',20000,0,'2024-10-19 17:18:58.115','2024-10-19 17:18:58.115'),('2ba2b88e-8286-4e7c-b29d-43b1324f8e5d','ae38adfc-6646-4633-b742-e6557155e4d1',20000,'Pulsa Telkomsel 100k',100000,80000,'2024-10-19 17:19:38.789','2024-10-19 17:19:38.789'),('180040bb-0ea3-4253-817c-b8cd0eed8945','ae38adfc-6646-4633-b742-e6557155e4d1',20000,'Pulsa Telkomsel 100k',80000,60000,'2024-10-19 17:20:46.584','2024-10-19 17:20:46.584');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `top_ups`
--

DROP TABLE IF EXISTS `top_ups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `top_ups` (
  `top_up_id` longtext,
  `user_id` longtext,
  `amount` bigint DEFAULT NULL,
  `balance_before` bigint DEFAULT NULL,
  `balance_after` bigint DEFAULT NULL,
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `top_ups`
--

LOCK TABLES `top_ups` WRITE;
/*!40000 ALTER TABLE `top_ups` DISABLE KEYS */;
INSERT INTO `top_ups` VALUES ('16a578e6-d082-4620-8771-ea263500cb6d','ae38adfc-6646-4633-b742-e6557155e4d1',20000,0,20000,'2024-10-19 17:06:34.110','2024-10-19 17:06:34.110'),('000fc9b0-e54f-4683-859a-062189f5628c','ae38adfc-6646-4633-b742-e6557155e4d1',20000,20000,40000,'2024-10-19 17:07:02.550','2024-10-19 17:07:02.550'),('05732d60-cf78-4366-9546-7495263f4b43','ae38adfc-6646-4633-b742-e6557155e4d1',100000,0,100000,'2024-10-19 17:19:34.367','2024-10-19 17:19:34.367'),('cca48876-3757-4318-ae6b-5ec15f26289f','ae38adfc-6646-4633-b742-e6557155e4d1',100000,0,100000,'2024-10-19 18:09:50.207','2024-10-19 18:09:50.207');
/*!40000 ALTER TABLE `top_ups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transfers`
--

DROP TABLE IF EXISTS `transfers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transfers` (
  `transfer_id` longtext,
  `user_id` longtext,
  `target_user_id` longtext,
  `status` longtext,
  `transaction_type` longtext,
  `amount` bigint DEFAULT NULL,
  `remarks` longtext,
  `balance_before` bigint DEFAULT NULL,
  `balance_after` bigint DEFAULT NULL,
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transfers`
--

LOCK TABLES `transfers` WRITE;
/*!40000 ALTER TABLE `transfers` DISABLE KEYS */;
INSERT INTO `transfers` VALUES ('69376840-3297-45c0-8ccd-e4036b9b2d86','ae38adfc-6646-4633-b742-e6557155e4d1','f6e2060f-913c-4f8e-ba2a-30bc9c1bd6a4','','',20000,'Pulsa Telkomsel 100k',60000,40000,'2024-10-19 18:07:38.098','2024-10-19 18:07:38.098'),('ecac7072-2668-451a-bb3b-59707b1e9192','ae38adfc-6646-4633-b742-e6557155e4d1','f6e2060f-913c-4f8e-ba2a-30bc9c1bd6a44','','',20000,'Pulsa Telkomsel 100k',40000,20000,'2024-10-19 18:08:05.803','2024-10-19 18:08:05.803'),('8154eb40-d9d5-4644-8b07-09aee4a88a53','ae38adfc-6646-4633-b742-e6557155e4d1','f6e2060f-913c-4f8e-ba2a-30bc9c1bd6a44','','',20000,'Pulsa Telkomsel 100k',20000,0,'2024-10-19 18:09:34.085','2024-10-19 18:09:34.085'),('360f4370-67b8-4f00-9adc-a799117307b4','ae38adfc-6646-4633-b742-e6557155e4d1','f6e2060f-913c-4f8e-ba2a-30bc9c1bd6a4','','',20000,'Pulsa Telkomsel 100k',100000,80000,'2024-10-19 18:09:59.492','2024-10-19 18:09:59.492');
/*!40000 ALTER TABLE `transfers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` varchar(191) NOT NULL,
  `first_name` longtext,
  `last_name` longtext,
  `phone_number` longtext,
  `address` longtext,
  `pin` longtext,
  `balance` bigint DEFAULT NULL,
  `created_date` datetime(3) DEFAULT NULL,
  `updated_date` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('ae38adfc-6646-4633-b742-e6557155e4d1','budi','salim','081293366701','','123452',80000,'2024-10-19 17:00:24.829','2024-10-19 17:00:24.829'),('f6e2060f-913c-4f8e-ba2a-30bc9c1bd6a4','toni','san','08129336671','','123452',40000,'2024-10-19 18:03:37.553','2024-10-19 18:03:37.553');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-19 11:34:32
