package com.tfm.resources;

import java.io.IOException;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.core.UriInfo;

import org.apache.log4j.Logger;

import com.tfm.dao.Blockchain;

@Path("blockchain")
public class BlockchainPruebaRecurso {

	private static final Logger LOGGER = Logger.getLogger(BlockchainPruebaRecurso.class);

	@Context
	UriInfo uri;
	@Context
	Request req;

	@GET
	public Response prueba() {
		return null;

	}

}
