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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._JSON_PERROS
DELIMITER //
CREATE PROCEDURE `_JSON_PERROS`()
BEGIN
	SELECT CONCAT(	'['
						, GROUP_CONCAT( '{'  
											, CONCAT( '"', 'docType',			  		'":', '"', 	'PERROS'									, '"'	, ',')
											, CONCAT( '"', 'IDPerro', 					'":',		 	ID													, ',')
											, CONCAT( '"', 'Nombre', 					'":', '"', 	NOMBRE									, '"'	, ',')
											, CONCAT( '"', 'IDAfijo',					'":', 	 	ID_AFIJO											, ',')
											, CONCAT( '"', 'IDSexo',					'":', 	 	( CASE SEXO WHEN "HEMBRA" THEN 0 WHEN "MACHO" THEN 1 ELSE -1 END ) , ',')
											, CONCAT( '"', 'IDPerroMadre',			'":', 	 	ID_MADRE											, ',')
											, CONCAT( '"', 'IDPerroPadre',			'":', 	 	ID_PADRE											, ',')
											, CONCAT( '"', 'IDRaza', 					'":', 	 	ID_RAZA											, ',')
											, CONCAT( '"', 'FechaNacimiento', 		'":', '"', 	IFNULL(FECHA_NACIMIENTO, '')		, '"'	, ',')
											, CONCAT( '"', 'FechaDefuncion', 		'":', '"',	IFNULL(FECHA_DEFUNCION, '')		, '"'	, ',')	
											, CONCAT( '"', 'FechaAlta', 				'":', '"', 	IFNULL(FECHA_NACIMIENTO + INTERVAL ( FLOOR(1 + (RAND() * 30))) DAY, '')		, '"'	, ',')
											, CONCAT( '"', 'FechaBaja', 				'":', '"',	IFNULL(FECHA_DEFUNCION + INTERVAL ( FLOOR(1 + (RAND() * 15))) DAY, '')		, '"'	)
										
										,'}\n'
											)
						,']') AS JSON_PERROS
		FROM	PERROS;
END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
