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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._JSON_VACUNAS_perros_1000
DELIMITER //
CREATE PROCEDURE `_JSON_VACUNAS_perros_1000`()
BEGIN
	DECLARE numero_ficheros INT; 
	DECLARE numero_registrosXfichero INT; 
	
	SET numero_registrosXfichero = 1000; 
	SET numero_ficheros = (SELECT CEIL(COUNT(*)/numero_registrosXfichero) FROM VACUNAS_perros); 
	
	FOR fichero IN 1..numero_ficheros DO

		SELECT CONCAT(	'['
							, GROUP_CONCAT( '{'  
												, CONCAT( '"', 'docType',					'":', '"', 	   'VACUNAS_PERROS'				, '"'	, ',')
												, CONCAT( '"', 'IDVacunaPerro', 			'":',		 		ID											, ',')
												, CONCAT( '"', 'IDPerro', 					'":',		 		ID_PERRO									, ',')
												, CONCAT( '"', 'IDPersonaVeterinario', '":', 	 		ID_VETERINARIO							, ',')
												, CONCAT( '"', 'CODVacuna', 				'":', '"', 		COD_VACUNA						, '"'	, ',')
												, CONCAT( '"', 'FechaAlta', 				'":', '"', 		IFNULL(FECHA_ALTA, '')		, '"'	, ',')
												, CONCAT( '"', 'FechaBaja', 				'":', '"',		IFNULL(FECHA_BAJA, '')		, '"'	)
												,'}\n'
												)
							,']') AS JSON_VACUNAS_perros
			FROM	VACUNAS_perros
			WHERE ID BETWEEN (1 + ((fichero - 1) * numero_registrosXfichero)) AND (fichero * numero_registrosXfichero);
			
	END FOR;
END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
