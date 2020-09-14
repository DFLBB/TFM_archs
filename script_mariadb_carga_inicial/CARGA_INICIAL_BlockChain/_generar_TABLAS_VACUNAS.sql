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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._generar_TABLAS_VACUNAS
DELIMITER //
CREATE PROCEDURE `_generar_TABLAS_VACUNAS`()
BEGIN
	DECLARE done INT DEFAULT FALSE;

	DECLARE v_id_perro INT;	
	DECLARE v_fecha_nacimiento DATE;
	DECLARE v_numero_vacunas INT;
	
	DECLARE v_fecha_alta DATE;
	DECLARE v_max_numero_protecciones INT; 
	DECLARE v_numero_protecciones INT;

	DECLARE cursor_perros CURSOR FOR 	
		SELECT 	ID,
					FECHA_NACIMIENTO, 
					YEAR(IFNULL(FECHA_DEFUNCION, NOW()))-YEAR(FECHA_NACIMIENTO)-1	 
		FROM PERROS
		ORDER BY 1
		LIMIT 10;
				
	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done=TRUE;


	TRUNCATE TABLE VACUNAS_perros;
	TRUNCATE TABLE VACUNAS_perros_proteccion;
	ALTER TABLE VACUNAS_perros_proteccion AUTO_INCREMENT = 1;
		
	OPEN cursor_perros;

	SET @rownum = 0;

	loop_perros: LOOP
	
		FETCH cursor_perros INTO v_id_perro, v_fecha_nacimiento,v_numero_vacunas;
		IF done THEN
			LEAVE loop_perros;
		END IF;
		
		SET v_max_numero_protecciones = ( SELECT COUNT(*) FROM VACUNAS_proteccion);
				
		FOR numero_vacuna IN 1..v_numero_vacunas
			DO
			
				SET v_fecha_alta = v_fecha_nacimiento +  INTERVAL ( numero_vacuna  * 365) DAY + INTERVAL ( 30 * RAND()) DAY;
							
				INSERT INTO VACUNAS_perros (ID, ID_PERRO, ID_VETERINARIO, COD_VACUNA, FECHA_ALTA, FECHA_BAJA)
				SELECT @rownum:=@rownum+1,
						 v_id_perro,
						 ( SELECT FLOOR(100 + (RAND() * 20))) AS ID_VETERINARIO,
						 ( SELECT LPAD (FLOOR(1 + (RAND() * 1000000000000000)),15,'0') ),
						 v_fecha_alta,
						 v_fecha_alta + INTERVAL ( 365 ) DAY;	
						 
				SET v_numero_protecciones = ( SELECT FLOOR(1 + (RAND() * v_max_numero_protecciones)));
											
				INSERT INTO VACUNAS_perros_proteccion (ID_VACUNA_PERRO, ID_VACUNA_PROTECCION, FECHA_ALTA, FECHA_BAJA)
				SELECT @rownum,
						 ID,
						 v_fecha_alta,
						 v_fecha_alta + INTERVAL ( 365 ) DAY
				FROM 	VACUNAS_proteccion
				ORDER BY RAND() LIMIT v_numero_protecciones;	
		END FOR;		
		
  	END LOOP;
	CLOSE cursor_perros; 
	
	SELECT * FROM VACUNAS_perros ORDER BY ID_PERRO, FECHA_ALTA; 
	SELECT * FROM VACUNAS_perros_proteccion ORDER BY ID_VACUNA_PERRO, FECHA_ALTA; 


END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
