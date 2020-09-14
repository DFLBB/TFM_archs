-- --------------------------------------------------------
-- Host:                         tfm.cnhygfwtvffy.eu-west-3.rds.amazonaws.com
-- Versión del servidor:         10.4.8-MariaDB-log - Source distribution
-- SO del servidor:              Linux
-- HeidiSQL Versión:             11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Volcando estructura para tabla CARGA_INICIAL_BlockChain.AFIJOS_propietarios
CREATE TABLE IF NOT EXISTS `AFIJOS_propietarios` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ID_AFIJO` int(11) NOT NULL DEFAULT 0,
  `ID_PROPIETARIO` int(11) NOT NULL DEFAULT 0,
  `FECHA_ALTA` date DEFAULT NULL,
  `FECHA_BAJA` date DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_AFIJO_ID_PROPIETARIO_FECHA_ALTA` (`ID_AFIJO`,`ID_PROPIETARIO`,`FECHA_ALTA`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- Volcando datos para la tabla CARGA_INICIAL_BlockChain.AFIJOS_propietarios: ~50 rows (aproximadamente)
/*!40000 ALTER TABLE `AFIJOS_propietarios` DISABLE KEYS */;
INSERT INTO `AFIJOS_propietarios` (`ID`, `ID_AFIJO`, `ID_PROPIETARIO`, `FECHA_ALTA`, `FECHA_BAJA`) VALUES
	(1, 1, 6, '2003-08-18', NULL),
	(2, 2, 5, '2003-07-22', NULL),
	(3, 3, 4, '2002-11-17', NULL),
	(4, 4, 3, '2003-05-14', NULL),
	(5, 5, 2, '2003-06-17', NULL),
	(6, 6, 1, '1994-02-18', NULL),
	(7, 7, 929, '1995-04-21', NULL),
	(8, 8, 324, '2000-11-10', NULL),
	(9, 9, 234, '1997-10-08', NULL),
	(10, 10, 194, '2020-05-24', NULL),
	(11, 11, 648, '2005-11-25', NULL),
	(12, 12, 869, '2002-06-14', NULL),
	(13, 13, 267, '2001-04-19', NULL),
	(14, 14, 261, '2005-11-21', NULL),
	(15, 15, 257, '2004-12-06', NULL),
	(16, 16, 439, '2013-09-26', NULL),
	(17, 17, 8, '2005-11-17', NULL),
	(18, 18, 618, '1994-12-08', NULL),
	(19, 19, 226, '2018-07-21', NULL),
	(20, 20, 255, '2005-03-12', NULL),
	(21, 21, 500, '2004-06-08', NULL),
	(22, 22, 216, '2013-05-10', NULL),
	(23, 23, 21, '2005-06-17', NULL),
	(24, 24, 684, '1993-12-26', NULL),
	(25, 25, 260, '2015-01-20', NULL),
	(26, 26, 426, '2017-12-06', NULL),
	(27, 27, 578, '1996-06-28', NULL),
	(28, 28, 773, '2017-07-16', NULL),
	(29, 29, 584, '1995-06-20', NULL),
	(30, 30, 489, '2013-08-07', NULL),
	(31, 31, 16, '2006-03-19', NULL),
	(32, 32, 52, '1997-01-06', NULL),
	(33, 33, 669, '2000-07-28', NULL),
	(34, 34, 881, '2018-07-28', NULL),
	(35, 35, 938, '2015-10-01', NULL),
	(36, 36, 410, '2002-05-29', NULL),
	(37, 37, 955, '1998-11-29', NULL),
	(38, 38, 586, '1994-02-02', NULL),
	(39, 39, 401, '2007-11-06', NULL),
	(40, 40, 435, '2008-12-18', NULL),
	(41, 41, 634, '2000-08-25', NULL),
	(42, 42, 868, '2010-06-25', NULL),
	(43, 43, 874, '2002-06-14', NULL),
	(44, 44, 427, '2014-12-10', NULL),
	(45, 45, 276, '2019-05-06', NULL),
	(46, 46, 612, '2003-11-21', NULL),
	(47, 47, 710, '1995-07-20', NULL),
	(48, 48, 639, '2000-03-14', NULL),
	(49, 49, 859, '1993-09-24', NULL),
	(50, 50, 259, '2002-03-08', NULL);
/*!40000 ALTER TABLE `AFIJOS_propietarios` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
