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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._generar_TABLA_AFIJOS_propietarios
DELIMITER //
CREATE PROCEDURE `_generar_TABLA_AFIJOS_propietarios`()
BEGIN

	DECLARE numero_afijos INT;
	SET numero_afijos = ( SELECT COUNT(*) FROM AFIJOS );

	TRUNCATE TABLE AFIJOS_propietarios;
	 
	SET @id_afijo=0; 
	INSERT INTO AFIJOS_propietarios ( ID_AFIJO, ID_PROPIETARIO, FECHA_ALTA, FECHA_BAJA )
	SELECT AFIJOS.ID AS ID_AFIJO,
			 AFIJO_PROPIETARIO.ID_PROPIETARIO, 
			 AFIJOS.FECHA_ALTA,
			 AFIJOS.FECHA_BAJA AS FECHA_BAJA
	FROM 	 AFIJOS,
			 ( SELECT @rownum:=@rownum+1 AS ID_AFIJO, 
			 			 PERSONAS_ALEATORIAS.ID AS ID_PROPIETARIO
				FROM 	(SELECT @rownum:=0) AS ORDEN, 
						( SELECT ID FROM PERSONAS ORDER BY RAND() LIMIT numero_afijos ) AS PERSONAS_ALEATORIAS
			 ) AS AFIJO_PROPIETARIO
		WHERE	AFIJOS.ID = AFIJO_PROPIETARIO.ID_AFIJO	;
	
	SELECT * FROM AFIJOS_propietarios;

END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
