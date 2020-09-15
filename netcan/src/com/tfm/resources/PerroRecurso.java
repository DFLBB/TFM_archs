package com.tfm.resources;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.UriInfo;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.core.Response;

import org.apache.log4j.Logger;

import com.tfm.dao.Netcan;
import com.tfm.data.perro.Perro;

@Path("/perros")
public class PerroRecurso {

	private static final Logger LOGGER = Logger.getLogger(PerroRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@POST
	@Consumes(MediaType.APPLICATION_JSON)
	public Response insertaPerro(Perro perro) {
		LOGGER.info("Peticion de registrar perro en el sistema");

		if (Netcan.getInstance().insertaPerro(perro)) {
			return Response.status(Status.CREATED).build();
		}

		return Response.status(Status.INTERNAL_SERVER_ERROR).build();

	}
	
	@Path("baja")
	@POST
	@Consumes(MediaType.APPLICATION_JSON)
	public Response bajaPerro(Perro perro) {
		LOGGER.info("Peticion de registrar perro en el sistema");

		if (Netcan.getInstance().defuncionPerro(perro)) {
			return Response.status(Status.CREATED).build();
		}

		return Response.status(Status.INTERNAL_SERVER_ERROR).build();

	}

}
