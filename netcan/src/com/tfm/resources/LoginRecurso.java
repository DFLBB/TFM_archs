package com.tfm.resources;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.UriInfo;
import javax.ws.rs.core.Response;

import org.apache.log4j.Logger;

import com.tfm.dao.Netcan;
import com.tfm.data.login.Login;

@Path("login")
public class LoginRecurso {

	private static final Logger LOGGER = Logger.getLogger(LoginRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@POST
	@Consumes(MediaType.APPLICATION_JSON)
	public Response nuevoUsuario(Login login) {

		LOGGER.info("Peticion de nuevo usuario, procedemos a insertar en la tabla altas");
		// Suponemos que el tipo va ser correcto siempre porque viene capado de la
		// pagina web
		int id = Netcan.getInstance().insertaAlta(login);
		System.out.println(id);
		if (id != 0) {

			boolean check2;

			switch (login.getTipo()) {
			case "P":
				LOGGER.info("Nuevo usuario de tipo propietario");
				check2 = Netcan.getInstance().insertaPropietario(login.getPropietario(), id);
				break;
			case "V":
				LOGGER.info("Nuevo usuario de tipo veterinario");
				check2 = Netcan.getInstance().insertaVeterinario(login.getVeterinario(), id);
				break;
			case "F":
				LOGGER.info("Nuevo usuario de tipo federacion");
				check2 = Netcan.getInstance().insertaFederacion(login.getFederacion(), id);
				break;
			default:
				check2 = false;
				break;
			}
			System.out.println(check2);
			if (check2) {
				return Response.status(Response.Status.CREATED).build();
			}
		}

		return Response.status(Response.Status.INTERNAL_SERVER_ERROR).build();

	}

	@GET
	@Produces(MediaType.APPLICATION_JSON)
	@Consumes(MediaType.APPLICATION_JSON)
	public Response login(@QueryParam("nickname") String nickname, @QueryParam("email") String email,
			@QueryParam("password") String pass) {

		LOGGER.info(
				"Peticion a login con parametros: nickname: " + nickname + " email: " + email + " password: " + pass);

		Login user = Netcan.getInstance().login(nickname, email, pass);

		if (user != null && user.getId() != 0) {

			return Response.ok(user).build();

		} else {
			return Response.status(Response.Status.NO_CONTENT).build();
		}

	}

}
