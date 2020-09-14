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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._generar_TABLA_PERROS_propietarios
DELIMITER //
CREATE PROCEDURE `_generar_TABLA_PERROS_propietarios`()
BEGIN
 
 	TRUNCATE TABLE PERROS_propietarios; 	

   INSERT INTO PERROS_propietarios ( ID_PERRO, ID_PROPIETARIO, FECHA_ALTA, FECHA_BAJA )
	SELECT	  PERROS.ID AS ID_PERRO
				, AFIJOS_propietarios.ID_PROPIETARIO
				, PERROS.FECHA_NACIMIENTO AS FECHA_ALTA
				, PERROS.FECHA_DEFUNCION AS FECHA_BAJA
				
	FROM		  PERROS
				, AFIJOS_propietarios
	WHERE 	PERROS.ID_AFIJO > 0
	AND		PERROS.ID_AFIJO = AFIJOS_propietarios.ID_AFIJO
	
	UNION
	
	SELECT	  PERROS.ID AS ID_PERRO
				, ( SELECT ID FROM PERSONAS ORDER BY RAND() LIMIT 1 ) AS ID_PROPIETARIO
				, PERROS.FECHA_NACIMIENTO AS FECHA_ALTA
				, PERROS.FECHA_DEFUNCION AS FECHA_BAJA
	FROM		PERROS
	WHERE 	PERROS.ID_AFIJO = 0;
	
	SELECT * FROM PERROS_propietarios;
END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
