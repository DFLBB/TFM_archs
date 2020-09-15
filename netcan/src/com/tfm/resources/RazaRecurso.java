package com.tfm.resources;

//@Path("/razas")
public class RazaRecurso {
	/**
	 * private static final Logger LOGGER = Logger.getLogger(RazaRecurso.class);
	 * 
	 * @Context UriInfo uri;
	 * @Context Request req;
	 * 
	 * @GET // @Secured
	 * @Produces(MediaType.APPLICATION_JSON) public Response dameRazas() {
	 * 
	 *                                       RazaLista objRazaLista = new
	 *                                       RazaLista(Netcan.getInstance().dameRazas());
	 * 
	 *                                       if (objRazaLista.getLista() == null) {
	 *                                       return
	 *                                       Response.status(Response.Status.INTERNAL_SERVER_ERROR).build();
	 *                                       }
	 * 
	 *                                       return
	 *                                       Response.ok(objRazaLista).build();
	 * 
	 *                                       }
	 * 
	 * @POST // @Secured
	 * @Consumes(MediaType.APPLICATION_JSON) @Produces(MediaType.APPLICATION_JSON)
	 *                                       public Response insertaRaza(Raza raza)
	 *                                       {
	 * 
	 *                                       boolean check =
	 *                                       Netcan.getInstance().insertaRaza(raza);
	 * 
	 *                                       if (check) { try { return
	 *                                       Response.created(new
	 *                                       URI(uri.getAbsolutePath().toString() +
	 *                                       "/")).build(); } catch
	 *                                       (URISyntaxException e) { // TODO
	 *                                       Auto-generated catch block
	 *                                       e.printStackTrace(); } }
	 * 
	 *                                       return
	 *                                       Response.status(Response.Status.INTERNAL_SERVER_ERROR).build();
	 * 
	 *                                       }
	 */

}
