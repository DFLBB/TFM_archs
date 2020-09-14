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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._JSON_AFIJOS
DELIMITER //
CREATE PROCEDURE `_JSON_AFIJOS`()
SELECT CONCAT(	'['
				, GROUP_CONCAT( '{'  
									, CONCAT( '"', 'docType',		'":', '"', 	'AFIJOS'						, '"'	, ',')
									, CONCAT( '"', 'IDAfijo', 		'":',		 	ID								 		, ',')
									, CONCAT( '"', 'Nombre', 		'":', '"', 	NOMBRE						, '"'	, ',')
									, CONCAT( '"', 'FechaAlta', 	'":', '"', 	IFNULL(FECHA_ALTA, '')	, '"'	, ',')
									, CONCAT( '"', 'FechaBaja', 	'":', '"',	IFNULL(FECHA_BAJA, '')	, '"'	)
									,'}\n'
									)
				,']') AS JSON_AFIJOS
FROM	AFIJOS//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
