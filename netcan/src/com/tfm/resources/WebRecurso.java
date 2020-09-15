package com.tfm.resources;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.UriInfo;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;

import org.apache.log4j.Logger;

import com.tfm.dao.Netcan;
import com.tfm.data.web.WebFederacion;
import com.tfm.data.web.WebPropietario;
import com.tfm.data.web.WebVeterinario;
import com.tfm.security.Secured;

@Path("web")
public class WebRecurso {

	private static final Logger LOGGER = Logger.getLogger(WebRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@GET
	@Secured
	@Produces(MediaType.APPLICATION_JSON)
	public Response web(@QueryParam("id") String id, @QueryParam("tipo") String tipo) {

		switch (tipo) {
		case "P":
			LOGGER.info("Web de tipo propietario");
			WebPropietario webPropietario = Netcan.getInstance().webPropietario(id);
			return Response.ok(webPropietario).build();
		case "V":
			LOGGER.info("Web de tipo veterinario");
			WebVeterinario webVeterinario = Netcan.getInstance().webVeterinario(id);
			return Response.ok(webVeterinario).build();
		case "F":
			LOGGER.info("Web de tipo federacion");
			WebFederacion webfederacion = Netcan.getInstance().webFederacion(id);
			return Response.ok(webfederacion).build();
		default:
			break;
		}

		return Response.status(Status.INTERNAL_SERVER_ERROR).build();
	}
}
