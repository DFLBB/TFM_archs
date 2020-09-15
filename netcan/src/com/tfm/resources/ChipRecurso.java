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
import com.tfm.data.chip.Chip;

@Path("/chips")
public class ChipRecurso {

	private static final Logger LOGGER = Logger.getLogger(ChipRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@POST
	@Consumes(MediaType.APPLICATION_JSON)
	public Response insertaChip(Chip chip) {
		LOGGER.info("Peticion de registrar chip en el sistema");

		if (Netcan.getInstance().insertaChip(chip)) {
			return Response.status(Status.CREATED).build();
		}

		return Response.status(Status.INTERNAL_SERVER_ERROR).build();

	}

}
