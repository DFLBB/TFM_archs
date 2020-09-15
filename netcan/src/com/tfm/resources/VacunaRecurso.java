package com.tfm.resources;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.UriInfo;
import javax.ws.rs.core.Response.Status;

import org.apache.log4j.Logger;

import com.tfm.dao.Netcan;
import com.tfm.data.vacuna.Vacuna;

@Path("vacunas")
public class VacunaRecurso {

	private static final Logger LOGGER = Logger.getLogger(VacunaRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@POST
	@Consumes(MediaType.APPLICATION_JSON)
	public Response insertaVacuna(Vacuna vacuna) {
		LOGGER.info("Peticion de registrar vacuna en el sistema");

		if (Netcan.getInstance().insertaVacuna(vacuna)) {
			return Response.status(Status.CREATED).build();
		}

		return Response.status(Status.INTERNAL_SERVER_ERROR).build();

	}

}
