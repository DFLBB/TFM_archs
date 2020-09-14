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

-- Volcando estructura para tabla CARGA_INICIAL_BlockChain.VACUNAS_proteccion
CREATE TABLE IF NOT EXISTS `VACUNAS_proteccion` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NOMBRE` varchar(254) COLLATE latin1_spanish_ci NOT NULL,
  `FECHA_ALTA` date DEFAULT current_timestamp(),
  `FECHA_BAJA` date DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4751 DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- Volcando datos para la tabla CARGA_INICIAL_BlockChain.VACUNAS_proteccion: ~11 rows (aproximadamente)
/*!40000 ALTER TABLE `VACUNAS_proteccion` DISABLE KEYS */;
INSERT INTO `VACUNAS_proteccion` (`ID`, `NOMBRE`, `FECHA_ALTA`, `FECHA_BAJA`) VALUES
	(1, 'Vacuna contra el moquillo canino', '2020-08-29', NULL),
	(2, 'Vacuna contra la hepatitis infecciosa', '2020-08-29', NULL),
	(3, 'Vacuna contra la leptospirosis', '2020-08-29', NULL),
	(4, 'Vacuna contra el parvovirus', '2020-08-29', NULL),
	(5, 'Vacuna contra el coronavirus', '2020-08-29', NULL),
	(6, 'Vacuna contra la rabia', '2020-08-29', NULL),
	(7, 'Vacuna contra la parainfluenza', '2020-08-29', NULL),
	(8, 'Vacuna contra Bordetella bronchiseptica', '2020-08-29', NULL),
	(9, 'Vacuna contra borreliosis o enfermedad de Lyme', '2020-08-29', NULL),
	(10, 'Vacuna contra herpesvirus canino', '2020-08-29', NULL),
	(11, 'Vacuna contra la leishmaniasis', '2020-08-29', NULL);
/*!40000 ALTER TABLE `VACUNAS_proteccion` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
