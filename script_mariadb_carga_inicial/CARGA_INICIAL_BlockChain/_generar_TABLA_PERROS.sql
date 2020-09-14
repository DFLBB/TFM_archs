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

-- Volcando estructura para procedimiento CARGA_INICIAL_BlockChain._generar_TABLA_PERROS
DELIMITER //
CREATE PROCEDURE `_generar_TABLA_PERROS`(
	IN `p_numeroGeneraciones` INT
)
BEGIN
	DECLARE done INT DEFAULT FALSE;
		
	DECLARE numero_perros INT;	
	
	DECLARE v_id_afijo INT;
	DECLARE v_id_hembra INT;	
	
	DECLARE v_fecha_nacimiento DATE;
	DECLARE v_fecha_defuncion DATE;
	
	DECLARE id_macho INT;
	DECLARE fecha_nacimiento_camada DATE;
	DECLARE intervalo_fechas DATE;
		
	DECLARE cursor_afijos CURSOR FOR SELECT ID FROM AFIJOS LIMIT 5 ;
	DECLARE cursor_hembras_afijo CURSOR(p_id_afijo INT) FOR 
			SELECT	ID,
						FECHA_NACIMIENTO,
						FECHA_DEFUNCION
			FROM 		PERROS 
			WHERE 	ID_AFIJO = p_id_afijo 
			AND 		SEXO ="HEMBRA";
	
	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done=TRUE;
		
	# Crear tabla de Nombres caninos

	# DROP TABLE IF EXISTS __NOMBRES_perros;

	# CREATE TABLE __NOMBRES_perros as
	# SELECT NOMBRE, "MACHO" as SEXO
	# FROM __NOMBRES_machos 
	# UNION ALL
	# SELECT NOMBRE, "HEMBRA" as SEXO
	# FROM __NOMBRES_hembras;

	TRUNCATE TABLE PERROS;
	
	# Insertar perros para pruebas

	INSERT INTO PERROS (NOMBRE, ID_AFIJO, SEXO, ID_MADRE, ID_PADRE, ID_RAZA, FECHA_NACIMIENTO, FECHA_DEFUNCION )	VALUES
			( "BRISA", 			-1, "HEMBRA", 0, 0, 0, NULL, NULL ),
			( "ROBERTA I", 	-1, "HEMBRA", 0, 0, 0, NULL, NULL ),
			( "VANESSA", 		-1, "HEMBRA", 0, 0, 0, NULL, NULL ),
			( "CONDOR", 		-1, "MACHO",  0, 0, 0, NULL, NULL ),
			( "SOPHIA", 		-1, "HEMBRA", 0, 0, 0, NULL, NULL ),
			( "EL DRUIDA I", 	-1, "MACHO",  0, 0, 0, NULL, NULL ),
			( "CUSTO", 			-1, "MACHO",  0, 0, 0, NULL, NULL ),
			( "AIRE", 			-1, "HEMBRA", 0, 0, 0, NULL, NULL ),
			( "RABANNE", 		-1, "MACHO",  0, 0, 0, NULL, NULL );

	# insertar primera generacion

	OPEN cursor_afijos;
	
	loop_primera_generacion: LOOP
	
		FETCH cursor_afijos INTO v_id_afijo;
		IF done THEN
			LEAVE loop_primera_generacion;
		END IF;
		
		FOR numero_hembras IN 1..(SELECT FLOOR(1 + (RAND() * 3)))
		DO
			iNSERT INTO PERROS (NOMBRE, ID_AFIJO, SEXO, ID_MADRE, ID_PADRE, ID_RAZA, FECHA_NACIMIENTO, FECHA_DEFUNCION)
			SELECT NOMBRE,
					 v_id_afijo,
					 SEXO,
					 0,
					 0,
					 0,
					 CURDATE() + INTERVAL ( 365 * RAND()) DAY,
					 CURDATE() + INTERVAL ( 20*365 * RAND()) DAY
			FROM  __NOMBRES_perros 
			WHERE SEXO = "HEMBRA"
			AND	NOMBRE NOT IN ( SELECT NOMBRE FROM PERROS WHERE ID_AFIJO = v_id_afijo )				
			ORDER BY RAND() LIMIT 1;	
		END FOR;
		
		FOR numeroGeneracion IN 1..(SELECT FLOOR(1 + (RAND() * 5)))
		DO
			INSERT INTO PERROS (NOMBRE, ID_AFIJO, SEXO, ID_MADRE, ID_PADRE, ID_RAZA, FECHA_NACIMIENTO, FECHA_DEFUNCION)
			SELECT NOMBRE,
					 v_id_afijo,
					 SEXO,
					 0,
					 0,
					 0,
					 CURDATE() + INTERVAL ( 365 * RAND()) DAY,
					 CURDATE() + INTERVAL ( 20*365 * RAND()) DAY
			FROM  __NOMBRES_perros 
			WHERE SEXO = "MACHO" 	
			AND	NOMBRE NOT IN ( SELECT NOMBRE FROM PERROS WHERE ID_AFIJO = v_id_afijo )							
			ORDER BY RAND() LIMIT 1;	
		END FOR;
		
  	END LOOP;
	CLOSE cursor_afijos; 
	
	# insertar otras generacion
	
	FOR numeroGeneracion IN 2..p_numeroGeneraciones
	DO
	   
		SET done=FALSE;
		OPEN cursor_afijos;
		loop_otras_generacion: LOOP
			FETCH cursor_afijos INTO v_id_afijo;
			IF done THEN
				LEAVE loop_otras_generacion;
			END IF;
					
			OPEN cursor_hembras_afijo(v_id_afijo);
	
				loop_hembras_afijo: LOOP
				
					FETCH cursor_hembras_afijo INTO v_id_hembra, v_fecha_nacimiento, v_fecha_defuncion;
					IF done THEN
			      	LEAVE loop_hembras_afijo;
			    	END IF;
			   		    
			    	SET numero_perros = ( SELECT FLOOR(1 + (RAND() * 5)));
			    	SET fecha_nacimiento_camada = ( SELECT CURDATE() + INTERVAL (2*(numeroGeneracion-1)) YEAR + INTERVAL ( 365 * RAND()) DAY);
			    	SET id_macho = ( SELECT ID 
												   FROM   PERROS 
													WHERE  ID_AFIJO = v_id_afijo 
													AND    SEXO ="MACHO"
													AND    fecha_nacimiento_camada BETWEEN ( FECHA_NACIMIENTO + INTERVAL 1 YEAR ) AND FECHA_DEFUNCION 
													AND    fecha_nacimiento_camada < FECHA_NACIMIENTO + INTERVAL 12 YEAR
													ORDER BY RAND() LIMIT 1);
			    
			    	IF ( 		( fecha_nacimiento_camada <  ( v_fecha_nacimiento + INTERVAL 10 YEAR ) )
						AND	( fecha_nacimiento_camada >= ( v_fecha_nacimiento + INTERVAL 1  YEAR ) )
						AND   ( fecha_nacimiento_camada <    v_fecha_defuncion ) 
						AND   ( id_macho IS NOT NULL ) 
						) THEN
						
						FOR numero_perros_camada IN 1..( SELECT FLOOR(1 + (RAND() * 5)))
						DO
							
							INSERT INTO PERROS (NOMBRE, ID_AFIJO, SEXO, ID_MADRE, ID_PADRE, ID_RAZA, FECHA_NACIMIENTO, FECHA_DEFUNCION)
							SELECT NOMBRE,
									 v_id_afijo,
									 SEXO,
									 v_id_hembra,
									 id_macho,
									 0,
									 fecha_nacimiento_camada,
						 			 fecha_nacimiento_camada + INTERVAL ( 20*365 * RAND()) DAY
							FROM  __NOMBRES_perros 
							WHERE  NOMBRE NOT IN ( SELECT NOMBRE FROM PERROS WHERE ID_AFIJO = v_id_afijo )
							ORDER BY RAND() LIMIT 1;
							
						END FOR;
							
				   END IF;						
					  
				END LOOP;
				
			SET done=FALSE;
			CLOSE cursor_hembras_afijo;
						
	 	END LOOP;
		CLOSE cursor_afijos; 
		
	END FOR;		
	
	# Traspolacion de fechas
	
	UPDATE PERROS
   SET   FECHA_NACIMIENTO=FECHA_NACIMIENTO + INTERVAL (SELECT DATEDIFF ( CURDATE(), MAX(FECHA_NACIMIENTO)) FROM PERROS) DAY
  		 , FECHA_DEFUNCION=FECHA_DEFUNCION   + INTERVAL (SELECT DATEDIFF ( CURDATE(), MAX(FECHA_NACIMIENTO)) FROM PERROS) DAY;
	
	UPDATE AFIJOS
	INNER JOIN
	(	SELECT ID_AFIJO,
			 	 MIN(FECHA_NACIMIENTO) - INTERVAL ( 365 * RAND()) DAY  AS FECHA_ALTA
		FROM 	 PERROS
		WHERE  ID_AFIJO > 0
		GROUP BY ID_AFIJO
	) AS FECHAS_AFIJOS
	ON
	AFIJOS.ID = FECHAS_AFIJOS.ID_AFIJO
	SET AFIJOS.FECHA_ALTA = FECHAS_AFIJOS.FECHA_ALTA;
	  		
	# Insertar perros mestizos

	FOR numeroRegistros IN 1..100
	DO
		INSERT INTO PERROS (NOMBRE, ID_AFIJO, SEXO, ID_MADRE, ID_PADRE, ID_RAZA, FECHA_NACIMIENTO, FECHA_DEFUNCION)
		SELECT NOMBRE,
				 0,
				 SEXO,
				 0,
				 0,
				 0,
				 CURDATE() - INTERVAL ( 20*365 * RAND()) DAY,
				 CURDATE() - INTERVAL ( 20*365 * RAND()) DAY
		FROM __NOMBRES_perros ORDER BY RAND() LIMIT 1;	
			
	END FOR;	
	
	# Correguir fechas de defuncion no corectas
	
	UPDATE PERROS
	SET 	 FECHA_DEFUNCION = NULL
	WHERE  FECHA_DEFUNCION <= FECHA_NACIMIENTO
	OR		 FECHA_DEFUNCION >= CURDATE();
		
	# ##################
	
	SELECT * FROM PERROS;
	SELECT MAX(ID)  FROM PERROS;
	

END//
DELIMITER ;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
