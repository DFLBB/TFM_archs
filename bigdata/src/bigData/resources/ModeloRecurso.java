package bigData.resources;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Request;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.UriInfo;

@Path("modelos")
public class ModeloRecurso {

	@Context
	UriInfo uri;
	@Context
	Request req;
	
	@GET
	@Path("/prediccion")
	@Consumes(MediaType.MULTIPART_FORM_DATA)
	@Produces(MediaType.APPLICATION_JSON)
	public Response uploadImage(@QueryParam("path") String path) throws IOException {
		String resultado = "";
		String comando = "python /home/ec2-user/scripts/predict.py " + path;
		Process p = Runtime.getRuntime().exec(comando);
		BufferedReader in = new BufferedReader(new InputStreamReader(p.getInputStream()));
		String line;
		while ((line = in.readLine()) != null) {
			System.out.println(line);
			line = line.trim();
			resultado = resultado + line;
			if (line.contains("raza:")) {
				return Response.ok(line.split("-")[1]).build();
			}
		}
		return Response.ok(resultado).build();
	}

	
}
