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

-- Volcando estructura para tabla CARGA_INICIAL_BlockChain.GRUPOS
CREATE TABLE IF NOT EXISTS `GRUPOS` (
  `ID_GRUPO` int(11) NOT NULL,
  `NOMBRE_GRUPO` varchar(50) COLLATE latin1_spanish_ci NOT NULL,
  `FECHA_ALTA` date NOT NULL DEFAULT curdate(),
  `FECHA_BAJA` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- Volcando datos para la tabla CARGA_INICIAL_BlockChain.GRUPOS: ~12 rows (aproximadamente)
/*!40000 ALTER TABLE `GRUPOS` DISABLE KEYS */;
INSERT INTO `GRUPOS` (`ID_GRUPO`, `NOMBRE_GRUPO`, `FECHA_ALTA`, `FECHA_BAJA`) VALUES
	(0, 'MESTIZOS', '2020-08-18', NULL),
	(1, 'GRUPO 1', '2020-08-18', NULL),
	(2, 'GRUPO 2', '2020-08-18', NULL),
	(3, 'GRUPO 3', '2020-08-18', NULL),
	(4, 'GRUPO 4', '2020-08-18', NULL),
	(5, 'GRUPO 5', '2020-08-18', NULL),
	(6, 'GRUPO 6', '2020-08-18', NULL),
	(7, 'GRUPO 7', '2020-08-18', NULL),
	(8, 'GRUPO 8', '2020-08-18', NULL),
	(9, 'GRUPO 9', '2020-08-18', NULL),
	(10, 'GRUPO 10', '2020-08-18', NULL),
	(999, 'GRUPO RAZAS ACEPTADAS PROVISIONALMENTE', '2020-08-18', NULL);
/*!40000 ALTER TABLE `GRUPOS` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
