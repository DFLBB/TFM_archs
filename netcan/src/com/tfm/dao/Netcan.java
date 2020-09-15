package com.tfm.dao;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintStream;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import javax.naming.InitialContext;
import javax.naming.NamingException;
import javax.sql.DataSource;

import org.apache.log4j.Logger;
import org.apache.naming.NamingContext;

import com.jcraft.jsch.Channel;
import com.jcraft.jsch.JSch;
import com.jcraft.jsch.Session;
import com.tfm.data.afijo.Afijo;
import com.tfm.data.chip.Chip;
import com.tfm.data.federacion.Federacion;
import com.tfm.data.login.Login;
import com.tfm.data.perro.Perro;
import com.tfm.data.propietario.Propietario;
import com.tfm.data.vacuna.Vacuna;
import com.tfm.data.veterinario.Veterinario;
import com.tfm.data.web.WebFederacion;
import com.tfm.data.web.WebPropietario;
import com.tfm.data.web.WebVeterinario;
import com.tfm.utils.Constantes;

public class Netcan {

	private static final Logger LOGGER = Logger.getLogger(Netcan.class);

	private DataSource ds;
	private Connection conn;
	private static Netcan myInstance;

	/**
	 * // Mapa para cachear las razas en la primera peticion al servidor private
	 * Map<Integer, String> razas; // Mapa para cachear los afijos en la primera
	 * peticion al servidor private Map<Integer, String> afijos; // Mapa para
	 * cachear los chips en la primera peticion al servidor private Map<Integer,
	 * Chip> chips; // Mapa para cachear las vacunas en la primera peticion al
	 * servidor private Map<Integer, String> vacunas; // Mapa para cachear los
	 * concursos en la primera peticion al servidor private Map<Integer, Concurso>
	 * concursos;
	 */

	/**
	 * Constructor de la clase, tiene varias acciones:
	 * 
	 * <li>1º Genera la conexion a la base de datos
	 * <li>2º Carga todos los datos de los mapas en memoria para evitar tanto acceso
	 * a la base de datos
	 */
	public Netcan() {
		InitialContext ctx;
		try {
			ctx = new InitialContext();
			NamingContext envCtx = (NamingContext) ctx.lookup(Constantes.DB_ENV);
			ds = (DataSource) envCtx.lookup(Constantes.DB_RESOURCE);
			conn = ds.getConnection();
		} catch (NamingException | SQLException e) {
			LOGGER.error("Error en la obtencion de la base de datos", e);
		}
		/**
		 * if (razas == null) { LOGGER.info("Generando mapa de razas"); razas = new
		 * HashMap<>(); rellenaRazas(); } if (afijos == null) { LOGGER.info("Generando
		 * mapa de afijos"); afijos = new HashMap<>(); rellenaAfijos(); }
		 * 
		 * if (chips == null) { LOGGER.info("Generando mapa de chips"); chips = new
		 * HashMap<>(); rellenaChips(); }
		 * 
		 * if (vacunas == null) { LOGGER.info("Generando mapa de vacunas"); vacunas =
		 * new HashMap<>(); rellenaVacunas(); }
		 * 
		 * if (concursos == null) { LOGGER.info("Generando mapa de concursos");
		 * concursos = new HashMap<>(); rellenaConcursos(); }
		 */
	}

	/**
	 * Metodo para generar la instancia en la primera llamada, Este hara que se
	 * ejecute el constructor cargando los mapas y generando la conexion a la base
	 * de datos <br>
	 * Si la conexion existe, nos devuelve directamente la instancia del objeto
	 * 
	 * @return
	 */
	public static Netcan getInstance() {
		if (myInstance == null) {
			myInstance = new Netcan();
		}
		return myInstance;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos GENERALES
	// Todos para tener una web mas visible y aceptable
	// No compatible con lo actual
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	/**
	 * Nos traemos los datos de las razas de la base de datos <br>
	 * <br>
	 * <i> El unico objetivo de este metodo es guardar las razas en memoria en el
	 * mapa <b> razas </b> para evitar accesos multiples a la base de datos </i>
	 */
	/**
	 * public void rellenaRazas() { String query = "SELECT * FROM raza;";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { while (rs.next()) { razas.put(rs.getInt("id"),
	 * rs.getString(Constantes.NOMBRE)); }
	 * 
	 * } catch (SQLException e) { LOGGER.error(Constantes.QUERY_ERROR, e); } }
	 */

	/**
	 * Nos traemos los datos de los afijos de la base de datos <br>
	 * <br>
	 * <i> El unico objetivo de este metodo es guardar los afijos en memoria en el
	 * mapa <b> afijos </b> para evitar accesos multiples a la base de datos </i>
	 */
	/**
	 * public void rellenaAfijos() { String query = "SELECT id, nombre FROM afijo;";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { while (rs.next()) { afijos.put(rs.getInt("id"),
	 * rs.getString(Constantes.NOMBRE)); } } catch (SQLException e) {
	 * LOGGER.error(Constantes.QUERY_ERROR, e); } }
	 */

	/**
	 * Nos traemos los datos de los chips de la base de datos <br>
	 * <br>
	 * <i> El unico objetivo de este metodo es guardar los chips en memoria en el
	 * mapa <b> chips </b> para evitar accesos multiples a la base de datos </i>
	 */
	/**
	 * public void rellenaChips() { String query = "SELECT * FROM chip;";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { while (rs.next()) { chips.put(rs.getInt("codigo"),
	 * new Chip(rs.getString("marca"), rs.getInt("colegiado"),
	 * rs.getString("pais"))); } } catch (SQLException e) {
	 * LOGGER.error(Constantes.QUERY_ERROR, e); } }
	 */

	/**
	 * Nos traemos los datos de las vacunas de la base de datos <br>
	 * <br>
	 * <i> El unico objetivo de este metodo es guardar los chips en memoria en el
	 * mapa <b> vacunas </b> para evitar accesos multiples a la base de datos </i>
	 */
	/**
	 * public void rellenaVacunas() { String query = "SELECT * FROM vacuna;";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { while (rs.next()) { vacunas.put(rs.getInt("id"),
	 * rs.getString(Constantes.NOMBRE)); } } catch (SQLException e) {
	 * LOGGER.error(Constantes.QUERY_ERROR, e); } }
	 */

	/**
	 * Nos traemos los datos de los concursos de la base de datos <br>
	 * <br>
	 * <i> El unico objetivo de este metodo es guardar los chips en memoria en el
	 * mapa <b> concursos </b> para evitar accesos multiples a la base de datos </i>
	 */
	/**
	 * public void rellenaConcursos() { String query = "SELECT * FROM concurso;";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { while (rs.next()) {
	 * concursos.put(rs.getInt(Constantes.ID_CONCURSO), new
	 * Concurso(rs.getString(Constantes.NOMBRE), rs.getInt("federacion"),
	 * rs.getString("patrocinadores"), rs.getString("pais"))); } } catch
	 * (SQLException e) { LOGGER.error(Constantes.QUERY_ERROR, e); } }
	 */

	/**
	 * Metodo para cuando se quiere insertar una nueva raza, un nuevo afijo o una
	 * nueva vacuna, se buscara en el mapa, si este existe en el mapa, no se
	 * añadira, ya que ya existe en la base de datos
	 * 
	 * @param nombre nombre de la raza, afijo o vacuna que estamos buscando
	 * @param mapa   nombre del mapa en el que buscar, solo para <b> razas y afijos
	 *               </b>
	 * @return Devuelve falso si no se encuentra y verdadero si se cuentra
	 */
	/**
	 * public boolean buscaEnMapas(String nombre, String mapa) {
	 * 
	 * switch (mapa) { case "razas":
	 * 
	 * for (Entry<Integer, String> entrada : razas.entrySet()) { if
	 * (entrada.getValue().equals(nombre)) { return true; } }
	 * 
	 * return false;
	 * 
	 * case "afijos": for (Entry<Integer, String> entrada : afijos.entrySet()) { if
	 * (entrada.getValue().equals(nombre)) { return true; } }
	 * 
	 * return false; case "vacunas": for (Entry<Integer, String> entrada :
	 * vacunas.entrySet()) { if (entrada.getValue().equals(nombre)) { return true; }
	 * }
	 * 
	 * return false; default: return false;
	 * 
	 * }
	 * 
	 * }
	 */

	/**
	 * Funcion para buscar el ID de un
	 * propietario/veterinario/perro/premio/vacuna/federacion/concurso/afijo y a mi
	 * 
	 * @return
	 */
	/**
	 * public boolean buscaID(int id, String tabla) { String query = "SELECT
	 * count(*) FROM " + tabla + " WHERE id = " + "'" + id + "';";
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet rs =
	 * st.executeQuery(query)) { if (rs.getInt("count(*)") > 0) { return true; }
	 * 
	 * } catch (SQLException e) { LOGGER.error(Constantes.QUERY_ERROR, e); return
	 * false; } return false; }
	 */

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso WEB
	//
	// Estos son los metodos que nos devuelven la informacion para meterla en las
	// web
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	/**
	 * Nos devuelve la informacion del propietario y la informacion de sus perros
	 * que vamos a utilizar para mostrar en la web
	 * 
	 * Dividido en pasos para mejor entendimiento:
	 * <li>1º Obtenemos los datos del propietario
	 * <li>2º Obtenemos los datos de los perros del propietario
	 * <li>3º Obtenemos las vacunas de cada perro del propietario
	 * <li>4º Obtenemos los concursos en los que han participado cada perrete
	 * 
	 * @param id
	 * @return
	 */
	public WebPropietario webPropietario(String id) {

		// 1º paso atacar al propietario

		Propietario propietario = null;

		String queryPropietario = "SELECT * FROM propietario WHERE idAlta = '" + id + "';";

		try (Statement st = conn.createStatement(); ResultSet rs = st.executeQuery(queryPropietario);) {

			while (rs.next()) {
				propietario = new Propietario(rs.getString("documento"), rs.getString("pais"), rs.getString("tipo_doc"),
						rs.getString(Constantes.NOMBRE), rs.getString("apellido1"), rs.getString("apellido2"),
						rs.getString("ciudad"), rs.getString("direccion"), rs.getInt(Constantes.ID_ALTA),
						rs.getString("paisEmisorDoc"));
			}

		} catch (SQLException e) {
			LOGGER.error(Constantes.WEBPROPIETARIO_ERROR, e);
			return null;
		}
		// Codigo para una pagina web mucho mas visible y aceptable
		// No es compatible con lo actual
		/**
		 * // 2º paso, atacar a los perros con chip, afijo y raza
		 * 
		 * String queryPerro = "SELECT * FROM perro WHERE documento = '" +
		 * propietario.getDocumento() + "' AND pais = '" + propietario.getPais() + "';";
		 * 
		 * ArrayList<Perro> listaPerros = new ArrayList<>();
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(queryPerro);) {
		 * 
		 * while (rs.next()) { listaPerros.add(new Perro(rs.getInt("id"),
		 * rs.getString(Constantes.NOMBRE), afijos.get(rs.getInt("idAfijo")),
		 * rs.getDouble("peso"), rs.getDouble("altura"), razas.get(rs.getInt("idRaza")),
		 * chips.get(rs.getInt("idChip")), rs.getString("sexo"),
		 * rs.getDate("Fecha_Nacimiento"), rs.getDate("Fecha_Defuncion"),
		 * rs.getString("enlaceImagen"), new VacunaLista(new ArrayList<Vacuna>()), new
		 * ParticipacionLista(new ArrayList<Participacion>()))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error(Constantes.WEBPROPIETARIO_ERROR, e);
		 * return null; }
		 * 
		 * PerroLista lista = new PerroLista(listaPerros);
		 * 
		 * // 3º paso obtenemos las vacunas de cada perro
		 * 
		 * for (int i = 0; i < listaPerros.size(); i++) { String queryVacunas = "SELECT
		 * idVacuna, fecha FROM vacunado WHERE idPerro = '" + listaPerros.get(i).getId()
		 * + "';";
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(queryVacunas);) {
		 * 
		 * while (rs.next()) { listaPerros.get(i).getVacunas().getLista() .add(new
		 * Vacuna(vacunas.get(rs.getInt("idVacuna")), rs.getDate(Constantes.FECHA))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error(Constantes.WEBPROPIETARIO_ERROR, e);
		 * return null; } }
		 * 
		 * // 4º paso obtenemos los concursos de cada perro
		 * 
		 * for (int i = 0; i < listaPerros.size(); i++) { String queryVacunas = "SELECT
		 * * FROM participa WHERE idPerro = '" + listaPerros.get(i).getId() + "';";
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(queryVacunas);) {
		 * 
		 * while (rs.next()) { listaPerros.get(i).getConcursos().getLista() .add(new
		 * Participacion(concursos.get(rs.getInt(Constantes.ID_CONCURSO)).getNombre(),
		 * rs.getDate(Constantes.FECHA), rs.getString("posicion"),
		 * rs.getDouble("premio"))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error(Constantes.WEBPROPIETARIO_ERROR, e);
		 * return null; } }
		 */
		return new WebPropietario(propietario);
	}

	/**
	 * Nos devuelve la informacion del veterinario, los chips que ha puesto y las
	 * vacunaciones que ha hecho
	 * 
	 * Dividido en pasos para mejor entendimiento:
	 * <li>1º Obtenemos los datos del veterinario
	 * <li>2º Obtenemos los datos de los chips
	 * <li>3º Obtenemos las vacunas que ha puesto
	 * 
	 * @param id
	 * @return
	 */
	public WebVeterinario webVeterinario(String id) {

		// 1º paso atacar al propietario

		Veterinario veterinario = null;

		String queryVeterinario = "SELECT * FROM veterinario WHERE idAlta = '" + id + "';";

		try (Statement st = conn.createStatement(); ResultSet rs = st.executeQuery(queryVeterinario);) {

			while (rs.next()) {
				veterinario = new Veterinario(rs.getInt("colegiado"), rs.getString(Constantes.NOMBRE),
						rs.getString("apellido1"), rs.getString("apellido2"), rs.getString("clinica_veterinaria"),
						rs.getString("direccion"), rs.getString("pais"), rs.getString("ciudad"),
						rs.getInt(Constantes.ID_ALTA));
			}

		} catch (SQLException e) {
			LOGGER.error(Constantes.WEBVETERINARIO_ERROR, e);
			return null;
		}
		// Codigo para una pagina web mucho mas visible y aceptable
		// No es compatible con lo actual
		/**
		 * // 2º paso chips que ha puesto String queryChips = "SELECT codigo, marca FROM
		 * chip WHERE colegiado = '" + veterinario.getColegiado() + "' AND pais = '" +
		 * veterinario.getPais() + "';";
		 * 
		 * ArrayList<Chip> listaChips = new ArrayList<>();
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(queryChips);) {
		 * 
		 * while (rs.next()) { listaChips.add(new Chip(rs.getInt("codigo"),
		 * rs.getString("marca"))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error(Constantes.WEBVETERINARIO_ERROR, e);
		 * return null; }
		 * 
		 * // 3º paso obtener las vacunaciones String querysVacunaciones = "SELECT
		 * perro.nombre AS perroNombre, perro.idAfijo AS perroAfijo, vacuna.id AS
		 * vacunaID, " + "vacunado.fecha AS fecha FROM vacunado INNER JOIN perro on
		 * idPerro = perro.id inner join vacuna" + " on idVacuna = vacuna.id where
		 * idVeterinario = '" + veterinario.getColegiado() + "';";
		 * 
		 * ArrayList<Vacunacion> listaVacunaciones = new ArrayList<>();
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(querysVacunaciones);) {
		 * 
		 * while (rs.next()) { listaVacunaciones.add(new
		 * Vacunacion(rs.getString("perroNombre"), afijos.get(rs.getInt("perroAfijo")),
		 * vacunas.get(rs.getInt("vacunaID")), rs.getDate(Constantes.FECHA))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error(Constantes.WEBVETERINARIO_ERROR, e);
		 * return null; }
		 */

		return new WebVeterinario(veterinario);
	}

	/**
	 * Nos devuelve la informacion de la federacion y los concursos que organiza
	 * 
	 * Dividido en pasos para mejor entendimiento:
	 * <li>1º Obtenemos los datos de la federacion
	 * <li>2º Obtenemos los datos de los concursos que organiza
	 * 
	 * @param id
	 * @return
	 */
	public WebFederacion webFederacion(String id) {
		// 1º paso atacar a la federacion

		Federacion federacion = null;

		String queryFederacion = "SELECT * FROM federacion WHERE idAlta = '" + id + "';";

		try (Statement st = conn.createStatement(); ResultSet rs = st.executeQuery(queryFederacion);) {

			while (rs.next()) {
				federacion = new Federacion(rs.getInt("id"), rs.getString(Constantes.NOMBRE), rs.getString("pais"),
						rs.getInt(Constantes.ID_ALTA));
			}

		} catch (SQLException e) {
			LOGGER.error("Error en el metodo webFederacion ", e);
			return null;
		}
		// Codigo para una pagina web mucho mas visible y aceptable
		// No es compatible con lo actual
		/**
		 * // 2º paso obtenemos los concursos de la federacion
		 * 
		 * String queryConcursos = "SELECT * FROM concurso WHERE federacion = '" +
		 * federacion.getId() + "';"; ArrayList<Concurso> listaConcursos = new
		 * ArrayList<>();
		 * 
		 * try (Statement st = conn.createStatement(); ResultSet rs =
		 * st.executeQuery(queryConcursos);) {
		 * 
		 * while (rs.next()) {
		 * listaConcursos.add(concursos.get(rs.getInt(Constantes.ID_CONCURSO))); }
		 * 
		 * } catch (SQLException e) { LOGGER.error("Error en el metodo webFederacion ",
		 * e); return null; }
		 */

		return new WebFederacion(federacion);
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso LOGIN
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	/**
	 * Metodo que se utilizara al hacer login desde la pagina web, este recibe el
	 * nickname, el email/usuario y la password. Con estos 3 parametros obtendremos
	 * el id y el tipo del usuario al que se refiere
	 * 
	 * @param nickname
	 * @param usuario
	 * @param password
	 * @return
	 */
	public Login login(String nickname, String email, String password) {

		String query = "SELECT `idAlta`, `tipo` FROM altas WHERE nick = '" + nickname + "' AND email = '" + email
				+ "' AND password = '" + password + "';";

		try (Statement st = conn.createStatement(); ResultSet rs = st.executeQuery(query);) {

			while (rs.next()) {
				return new Login(rs.getInt("idAlta"), rs.getString("tipo"));
			}

		} catch (SQLException e) {
			LOGGER.error("Error en el metodo login ", e);
			return null;
		}

		return null;

	}

	/**
	 * Metodo para crear un nuevo usuario al registrar en la web (para los datos
	 * generales)
	 * 
	 * @param login
	 * @return
	 */
	public int insertaAlta(Login login) {

		int id = 0;
		String query = "INSERT INTO altas VALUES ('" + 0 + "','" + login.getNickname() + "','" + login.getTipo() + "','"
				+ login.getEmail() + "','" + login.getPassword() + "','" + login.getTelefono() + "');";

		String queryAux = "SELECT idAlta FROM altas WHERE nick = '" + login.getNickname() + "' AND email  ='"
				+ login.getEmail() + "' AND password  ='" + login.getPassword() + "';";

		try (Statement st = conn.createStatement();) {
			st.executeUpdate(query);

		} catch (SQLException e) {
			LOGGER.error("Error en la funcion insertaAlta ", e);
			return id;
		}

		try (Statement st = conn.createStatement(); ResultSet rs = st.executeQuery(queryAux);) {

			while (rs.next()) {
				return rs.getInt(Constantes.ID_ALTA);
			}
		} catch (SQLException e) {
			LOGGER.error("Error en la funcion insertaAlta ", e);
			return id;
		}

		return id;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Conexion al CLI de Blockchain
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	public String runComando(String comando, String cadenaBuscar) {

		String result = "";

		try {
			JSch jsch = new JSch();
			java.util.Properties config = new java.util.Properties();
			config.put("StrictHostKeyChecking", "no");
			Session session = jsch.getSession(Constantes.BK_USER_CONN, Constantes.BK_IP_CONN, Constantes.BK_PORT_CONN);
			session.setPassword(Constantes.BK_PASS);
			session.setConfig(config);
			session.connect();
			Channel channel = session.openChannel(Constantes.BK_SHELL);
			OutputStream ops = channel.getOutputStream();
			PrintStream ps = new PrintStream(ops, true);
			channel.connect();
			InputStream input = channel.getInputStream();

			ps.println(comando);
			ps.println(Constantes.BK_END_COMMAND);
			ps.close();

			result = printResult(input, cadenaBuscar);
			channel.disconnect();
			session.disconnect();
		} catch (Exception e) {
			LOGGER.error("Error en la conexion a la blockchain ", e);
			return null;
		}

		return result;
	}

	private String printResult(InputStream input, String cadenaBuscar) throws Exception {
		String line;
		try (BufferedReader input2 = new BufferedReader(new InputStreamReader(input))) {
			while ((line = input2.readLine()) != null) {
				System.out.println(line);
				line = line.trim();

				if (line.contains("status:200")) {
					String[] array = line.split(",");
					String id = "";
					for (int i = 0; i < array.length; i++) {
						if (array[i].contains(cadenaBuscar)) {
							id = array[i];
							break;
						}

					}
					return (id.split(":")[1]);
				}

				if (line.equals("acabe")) {
					break;
				}
			}
			System.out.println("hola");
		} catch (IOException e) {
			LOGGER.error("Error al leer los resultados de los comandos ", e);
			return null;
		}

		return null;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso PERRO
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	public boolean insertaPerro(Perro perro) {

		// Blockchain
		String block = runComando(
				"docker exec cli peer chaincode invoke -n perros -c '{\"function\":\"registrarPerro\","
						+ "\"Args\":[\"{\\\"Perros\\\":" + "[{\\\"Nombre\\\":\\\"" + perro.getNombre()
						+ "\\\",\\\"IDSexo\\\":" + perro.getSexo() + "}" + "]," + "\\\"IDPerroMadre\\\":"
						+ perro.getIdMadre() + "," + "\\\"IDPerroPadre\\\":" + perro.getIdPadre() + ","
						+ "\\\"FechaNacimiento\\\":\\\"" + perro.getFechaNacimiento() + "\\\"}\","
						+ " \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER + " --tls --cafile "
						+ Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"\\\"IDPerro\\\"");

		if (block == null) {
			return false;
		}

		// Base de datos, no pasamos por aqui
		/*
		 * String query = "INSERT INTO perro VALUES ('" + block + "','" +
		 * perro.getNombre() + "','" + user.getTipoDoc() + "','" +
		 * user.getPaisEmisorDoc() + "','" + user.getNombre() + "','" +
		 * user.getApellido1() + "','" + user.getApellido2() + "','" + user.getCiudad()
		 * + "','" + user.getDireccion() + "','" + id + "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query); }
		 * catch (SQLException e) { LOGGER.error("Error en la creacion de la query ",
		 * e); return false; }
		 */

		return true;
	}

	public boolean defuncionPerro(Perro perro) {

		// Blockchain
		String block = runComando("docker exec cli peer chaincode invoke -n perros -c "
				+ "'{\"function\":\"registrarDefuncionPerro\"," + "\"Args\":[\"{\\\"IDPerro\\\":" + perro.getId() + ","
				+ "\\\"FechaDefuncion\\\":\\\"" + perro.getFechaDefuncion() + "\\\"}\","
				+ " \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER + " --tls --cafile "
				+ Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL, "\\\"IDPerro\\\"");

		if (block == null) {
			return false;
		}

		// Base de datos, no pasamos por aqui
		/*
		 * String query = "INSERT INTO perro VALUES ('" + block + "','" +
		 * perro.getNombre() + "','" + user.getTipoDoc() + "','" +
		 * user.getPaisEmisorDoc() + "','" + user.getNombre() + "','" +
		 * user.getApellido1() + "','" + user.getApellido2() + "','" + user.getCiudad()
		 * + "','" + user.getDireccion() + "','" + id + "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query); }
		 * catch (SQLException e) { LOGGER.error("Error en la creacion de la query ",
		 * e); return false; }
		 */

		return true;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso PROPIETARIO
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	public boolean insertaPropietario(Propietario user, int id) {

		// Blockchain
		String block = runComando("docker exec cli peer chaincode invoke -n personas -c "
				+ "'{\"function\":\"registrarPersona\",\"Args\":[\"{\\\"Nombre\\\":\\\"" + user.getNombre() + "\\\","
				+ "\\\"Apellido1\\\":\\\"" + user.getApellido1() + "\\\",\\\"Apellido2\\\":\\\"" + user.getApellido2()
				+ "\\\",\\\"TipoDocumento\\\":\\\"" + user.getTipoDoc() + "\\\""
				+ ",\\\"IdentificadorDocumento\\\":\\\"" + user.getDocumento() + "\\\",\\\"PaisEmisor\\\":\\\""
				+ user.getPaisEmisorDoc() + "\\\"}\"]}'" + " -o " + Constantes.BK_ORDERER + " --tls --cafile "
				+ Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL, "\\\"IDPersona\\\"");

		if (block == null) {
			return false;
		}

		// Base de datos

		String query = "INSERT INTO propietario VALUES ('" + block + "','" + user.getDocumento() + "','"
				+ user.getPais() + "','" + user.getTipoDoc() + "','" + user.getNombre() + "','" + user.getApellido1()
				+ "','" + user.getApellido2() + "','" + user.getCiudad() + "','" + user.getDireccion() + "','"
				+ user.getPaisEmisorDoc() + "','" + id + "');";

		try (Statement st = conn.createStatement();) {
			st.executeUpdate(query);
		} catch (SQLException e) {
			LOGGER.error("Error en la creacion de la query ", e);
			return false;
		}

		return true;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso CONCURSO
	// No implementado el Blockchain
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Concurso> dameConcursos() { String query = "SELECT * FROM
	 * concurso"; List<Concurso> resultado = new ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new
	 * Concurso(result.getInt("idConcurso"), result.getString(Constantes.NOMBRE),
	 * result.getInt("federacion"), result.getString("patrocinadores"),
	 * result.getString("pais"))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameConcursos ", e); return
	 * null; }
	 * 
	 * return resultado; }
	 * 
	 * public boolean insertaConcurso(Concurso concurso) { String query = "INSERT
	 * INTO concurso VALUES ('" + concurso.getNombre() + "','" +
	 * concurso.getFederacion() + "','" + concurso.getPatrocinadores() + "');";
	 * 
	 * try (Statement st = conn.createStatement();) { st.executeUpdate(query); }
	 * catch (SQLException e) { LOGGER.error("Error en la creacion de la query ",
	 * e); return false; }
	 * 
	 * return true; }
	 */

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso VETERINARIO
	// ---------------------------------------------------------------------------------------------------------------------------------------------------

	public boolean insertaVeterinario(Veterinario veterinario, int id) {

		// Blockchain
		String block = runComando(
				"docker exec cli peer chaincode invoke -n veterinarios -c '{\"function\":\"registrarColegiaturaPersona\", \"Args\":[\" {\\\"IDPersona\\\":"
						+ id + ", \\\"CODColegiatura\\\":" + veterinario.getColegiado() + "\\\"}]}\","
						+ " \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER + " --tls --cafile "
						+ Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"");

		if (block == null) {
			return false;
		}

		// Base de datos
		String query = "INSERT INTO veterinario VALUES ('" + veterinario.getColegiado() + "','"
				+ veterinario.getNombre() + "','" + veterinario.getApellido1() + "','" + veterinario.getApellido2()
				+ "','" + veterinario.getClinicaVeterinaria() + "','" + veterinario.getDireccion() + "','"
				+ veterinario.getPais() + "','" + veterinario.getCiudad() + "','" + id + "');";

		try (Statement st = conn.createStatement();) {
			st.executeUpdate(query);
		} catch (SQLException e) {
			LOGGER.error("Error en la creacion de la query ", e);
			return false;
		}

		return true;
	}

	public boolean borraVeterinario(String id) {
		// TODO
		return false;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso RAZAS
	// Descartados
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Raza> dameRazas() {
	 * 
	 * String query = "SELECT * FROM razas"; List<Raza> resultado = new
	 * ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new Raza(result.getInt("id"),
	 * result.getString(Constantes.NOMBRE))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameRazas ", e); return null;
	 * }
	 * 
	 * return resultado; }
	 * 
	 * public boolean insertaRaza(Raza raza) {
	 * 
	 * // Buscamos si el nombre de la raza esta en el mapita, si esta no hace falta
	 * // insertalo en la bbdd if (!razas.containsValue(raza.getNombre())) {
	 * 
	 * // XXX Llamar a la Blockchain
	 * 
	 * // XXX Llamar a la blockchain // Si la blockchain acaba bien, metemos en la
	 * base de datos String query = "INSERT INTO razas VALUES ('" + 0 + "','" +
	 * raza.getNombre() + "');";
	 * 
	 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en insertaRaza ", e); return
	 * false; }
	 * 
	 * return true; }
	 * 
	 * return false;
	 * 
	 * }
	 */

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso CHIPS
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Chip> dameChips() {
	 * 
	 * String query = "SELECT * FROM chip"; List<Chip> resultado = new
	 * ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new Chip(result.getInt("codigo"),
	 * result.getString("marca"), result.getInt("colegiado"),
	 * result.getString("pais"))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameRazas ", e); return null;
	 * }
	 * 
	 * return resultado; }
	 */

	public boolean insertaChip(Chip chip) {

		// Blockchain
		String block = runComando(
				"docker exec cli peer chaincode invoke -n microchips -c '{\"function\":\"registrarMicrochipPerro\", \"Args\":[\" {\\\"IDPerro\\\":"
						+ chip.getIdPerro() + ", \\\"IDPersonaVeterinario\\\":" + chip.getIdVeterinario()
						+ ", \\\"CODMicrochip\\\":\\\"" + chip.getCodMicrochip() + "\\\"}\","
						+ " \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER + " --tls --cafile "
						+ Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"");

		if (block == null) {
			return false;
		}

		// No pasamos por la base de datos
		/**
		 * String query = "INSERT INTO chip VALUES ('" + chip.getCodigo() + "','" +
		 * chip.getMarca() + "','" + chip.getColegiado() + "','" + chip.getPais() +
		 * "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
		 * 
		 * } catch (Exception e) { LOGGER.error("Error en insertaRaza ", e); return
		 * false; }
		 */

		return true;

	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso Federaciones
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Federacion> dameFederaciones() {
	 * 
	 * String query = "SELECT * FROM federacion"; List<Federacion> resultado = new
	 * ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new Federacion(result.getInt("id"),
	 * result.getString(Constantes.NOMBRE), result.getString("pais"),
	 * result.getInt(Constantes.ID_ALTA))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameFederaciones ", e); return
	 * null; }
	 * 
	 * return resultado; }
	 */

	public boolean insertaFederacion(Federacion federacion, int idAlta) {

		String query = "INSERT INTO federacion VALUES ('" + 0 + "','" + federacion.getNombre() + "', '"
				+ federacion.getPais() + "', '" + idAlta + "');";

		try (Statement st = conn.createStatement();) {
			st.executeUpdate(query);

		} catch (Exception e) {
			LOGGER.error("Error en insertaFederacion ", e);
			return false;
		}

		return true;
	}

	/**
	 * public boolean updateFederacion(Federacion federacion) {
	 * 
	 * String query = "UPDATE INTO federacion VALUES ('" + 0 + "','" +
	 * federacion.getNombre() + "', '" + federacion.getPais() + "', '" +
	 * federacion.getIdAlta() + "');";
	 * 
	 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en insertaFederacion ", e);
	 * return false; }
	 * 
	 * return true; }
	 * 
	 * public boolean deleteFederacion() { return false; }
	 */

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso AFIJOS
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Afijo> dameAfijos() {
	 * 
	 * String query = "SELECT * FROM afijo"; List<Afijo> resultado = new
	 * ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new Afijo(result.getInt("id"),
	 * result.getString(Constantes.NOMBRE), result.getString("documento"),
	 * result.getString("pais"), result.getDate("fecha_alta"),
	 * result.getDate("fecha_baja"))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameAfijos ", e); return null;
	 * }
	 * 
	 * return resultado; }
	 */

	public boolean insertaAfijo(Afijo afijo) {

		// Blockchain
		String block = runComando(
				"docker exec cli peer chaincode invoke -n afijos -c '{\"function\":\"registrarAfijo\", \"Args\":[\" {\\\"Nombre\\\":\\\""
						+ afijo.getNombre() + "\\\", \\\"Propietarios\\\":[{\\\"IDPersona\\\":" + afijo.getIdPersona1()
						+ "},{\\\"IDPersona\\\":" + afijo.getIdPersona2() + "},{\\\"IDPersona\\\":"
						+ afijo.getIdPersona3() + "}]}\", \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER
						+ " --tls --cafile " + Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"");

		if (block == null) {
			return false;
		}

		// Base de datos
		/**
		 * String query = "INSERT INTO afijo VALUES ('" + 0 + "','" afijo.getNombre() +
		 * "','" + afijo.getDocumento() + "','" + afijo.getPais() "','" +
		 * afijo.getFechaAlta() + "','" + afijo.getFechaBaja() + "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
		 * 
		 * } catch (Exception e) { LOGGER.error("Error en insertaAfijo ", e); return
		 * false; }
		 */
		return true;
	}

	public boolean bajaAfijo(Afijo afijo) {

		// Blockchain
		String block = runComando(
				"docker exec cli peer chaincode invoke -n afijos -c '{\"function\":\"registrarCancelacionAfijo\",\"Args\":[\"{\\\"IDAfijo\\\":"
						+ afijo.getId() + "}\", \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER
						+ " --tls --cafile " + Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"");

		if (block == null) {
			return false;
		}

		// Base de datos
		/**
		 * String query = "INSERT INTO afijo VALUES ('" + 0 + "','" afijo.getNombre() +
		 * "','" + afijo.getDocumento() + "','" + afijo.getPais() "','" +
		 * afijo.getFechaAlta() + "','" + afijo.getFechaBaja() + "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
		 * 
		 * } catch (Exception e) { LOGGER.error("Error en insertaAfijo ", e); return
		 * false; }
		 */
		return true;
	}

	/**
	 * return true; }
	 * 
	 * public boolean updateAfijo(Afijo afijo) {
	 * 
	 * String query = "UPDATE INTO afijo VALUES ('" + 0 + "','" + afijo.getNombre()
	 * + "','" + afijo.getDocumento() + "','" + afijo.getPais() + "','" +
	 * afijo.getFechaAlta() + "','" + afijo.getFechaBaja() + "');";
	 * 
	 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en insertaAfijo ", e); return
	 * false; }
	 * 
	 * return true; }
	 */

	// TODO: hay un registrarCancelacionAfijo
	public boolean deleteAfijo() {
		return false;
	}

	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	// Metodos del recurso Vacunas
	// ---------------------------------------------------------------------------------------------------------------------------------------------------
	/**
	 * public List<Vacuna> dameVacunas() {
	 * 
	 * String query = "SELECT * FROM vacuna"; List<Vacuna> resultado = new
	 * ArrayList<>();
	 * 
	 * try (Statement st = conn.createStatement(); ResultSet result =
	 * st.executeQuery(query);) {
	 * 
	 * while (result.next()) { resultado.add(new Vacuna(result.getInt("id"),
	 * result.getString(Constantes.NOMBRE))); }
	 * 
	 * } catch (Exception e) { LOGGER.error("Error en dameVacunas", e); return null;
	 * }
	 * 
	 * return resultado; }
	 */

	public boolean insertaVacuna(Vacuna vacuna) {

		// Blockchain
		// XXX Mirar que coño devuelve el mensaje por si hay que pillar algo
		String block = runComando(
				"docker exec cli peer chaincode invoke -n vacunas -c '{\"function\":\"registrarVacunaPerro\", \"Args\":[\" {\\\"IDPerro\\\":"
						+ vacuna.getIdPerro() + ", \\\"IDPersonaVeterinario\\\":" + vacuna.getIdVeterinario()
						+ ", \\\"CODVacuna\\\":\\\"" + vacuna.getCodVacuna() + "\\\", \\\"FechaBaja\\\":\\\""
						+ vacuna.getFechaBaja() + "\\\", \\\"Protecciones\\\":[ {\\\"IDVacunaProteccion\\\":"
						+ vacuna.getIdProteccion() + ",\\\"FechaBaja\\\":\\\"" + vacuna.getFechaBajaProteccion()
						+ "\\\"}]}\"," + " \"{\\\"IDPersona\\\":6}\"]}'" + " -o " + Constantes.BK_ORDERER
						+ " --tls --cafile " + Constantes.BK_CAFILE + " -C " + Constantes.BK_CANAL,
				"");

		if (block == null) {
			return false;
		}

		/**
		 * String query = "INSERT INTO vacuna VALUES ('" + 0 + "','" +
		 * vacuna.getNombre() + "');";
		 * 
		 * try (Statement st = conn.createStatement();) { st.executeUpdate(query);
		 * 
		 * } catch (Exception e) { LOGGER.error("Error en insertaVacuna ", e); return
		 * false; }
		 */

		return true;
	}

}
