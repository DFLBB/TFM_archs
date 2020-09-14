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

-- Volcando estructura para tabla CARGA_INICIAL_BlockChain.__NOMBRES_afijos
CREATE TABLE IF NOT EXISTS `__NOMBRES_afijos` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NOMBRE` varchar(50) COLLATE latin1_spanish_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- Volcando datos para la tabla CARGA_INICIAL_BlockChain.__NOMBRES_afijos: ~50 rows (aproximadamente)
/*!40000 ALTER TABLE `__NOMBRES_afijos` DISABLE KEYS */;
INSERT INTO `__NOMBRES_afijos` (`ID`, `NOMBRE`) VALUES
	(1, 'GLIMES DE BRABANT'),
	(2, 'CORAMONTE'),
	(3, 'LAS CIRAJAS'),
	(4, 'VILLA ROCALLA '),
	(5, 'VILLA ALKOR'),
	(6, 'OZMILION '),
	(7, 'CUMBA'),
	(8, 'LA VIRREYNA'),
	(9, 'WAYKEMOLA'),
	(10, 'KRYSTLE´S LINE'),
	(11, 'PARQUE DEL RETIRO'),
	(12, 'EL VALLE ENCANTADO'),
	(13, 'PORTITXOL'),
	(14, 'SEIKYTAS'),
	(15, 'MARVELS LUX'),
	(16, 'SHIKARAH'),
	(17, 'GOLDEN ROSE MANOR'),
	(18, 'SHERAZADE'),
	(19, 'VEGA DE LA MARTINA'),
	(20, 'DIAMOND DE VALLMASERI'),
	(21, 'TRASTUCA'),
	(22, 'MINATERRA'),
	(23, 'EL LAGO ANDARA'),
	(24, 'LAS MERINDADES'),
	(25, 'SILVER POISE'),
	(26, 'VEGA DE LA MARTINA'),
	(27, 'ALTAGIS'),
	(28, 'GUARDIANES DEL SILENCIO'),
	(29, 'SONTAY'),
	(30, 'VALESTORY'),
	(31, 'ANXOS'),
	(32, 'ORUMNILAI'),
	(33, 'YORSHICAN'),
	(34, 'DU MAS DE NAÏLYS'),
	(35, 'LA VEGA DE AZAHAR'),
	(36, 'SORIENA'),
	(37, 'EL BOSQUE DE LUGH'),
	(38, 'TARAWAY'),
	(39, 'CHOWS OF THE ISLAND'),
	(40, 'ANNANGELS'),
	(41, 'FORATA'),
	(42, 'DASHLUT'),
	(43, 'ZAFFIRO BLU'),
	(44, 'TIERRA DEL VIENTO'),
	(45, 'DELS PELUTS'),
	(46, 'FUENTE BLANCA'),
	(47, 'MANADA DE MELVIC'),
	(48, 'JUNGLA NEGRA'),
	(49, 'SOTO DE RIOFRIO'),
	(50, 'GARDELCAN');
/*!40000 ALTER TABLE `__NOMBRES_afijos` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
