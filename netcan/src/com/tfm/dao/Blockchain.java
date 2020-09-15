package com.tfm.dao;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintStream;

import org.apache.log4j.Logger;

import com.jcraft.jsch.*;
import com.tfm.utils.Constantes;

public class Blockchain {

//	private static final Logger LOGGER = Logger.getLogger(Blockchain.class);
//
//	private static Blockchain myInstance;
//	// private static Session session = null;
//
//	public Blockchain() {
//
//		JSch jsch = new JSch();
//
//		/*
//		 * try { jsch.addIdentity(Constantes.BK_pem_CONN); session =
//		 * jsch.getSession(Constantes.BK_user_CONN, Constantes.BK_IP_CONN,
//		 * Constantes.BK_port_CONN); } catch (JSchException e) {
//		 * LOGGER.error("Error al crear la sesion a las maquinas de EC2 ", e); session =
//		 * null; }
//		 */
//
//	}
//
//	public static Blockchain getInstance() {
//		/*
//		 * if (myInstance == null || session == null) { myInstance = new Blockchain(); }
//		 */
//		return myInstance;
//	}
//
//	public static String runComando(String comando) {
//		String caFile = "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem";
//		String canal = "netcanchannel";
//
//		String prueba = "docker exec cli peer chaincode invoke -n afijos -c '{\"function\":\"registrarAfijo\",\"Args\":[\"{\\\"Nombre\\\":\\\"NUEVO AFIJO\\\","
//				+ "\\\"Propietarios\\\":[{\\\"IDPersona\\\":700},{\\\"IDPersona\\\":701},{\\\"IDPersona\\\":702}]}\", \"{\\\"IDPersona\\\":6}\"]}' "
//				+ "-o" + Constantes.BK_orderer + " --tls --cafile " + caFile + " -C " + canal;	
//		String ls = "ls";
//		String sudo = "sudo su";
//		String cd = "cd ..";
//		String docker = "docker ps";
//		StringBuilder outputBuffer = new StringBuilder();
//		try {
//			System.out.println("hola");
//			JSch jsch = new JSch();
//			jsch.addIdentity(Constantes.BK_pem_CONN);
//			java.util.Properties config = new java.util.Properties();
//			config.put("StrictHostKeyChecking", "no");
//			Session session = jsch.getSession(Constantes.BK_user_CONN, Constantes.BK_IP_CONN, Constantes.BK_port_CONN);
//			session.setConfig(config);
//			session.connect();
//			Channel channel = session.openChannel("shell");
//			System.out.println(sudo + ";" + cd + ";" + ls);
//			OutputStream ops = channel.getOutputStream();
//			PrintStream ps = new PrintStream(ops, true);
//			channel.connect();
//			InputStream input = channel.getInputStream();
//			
//			ps.println(sudo);
//			ps.println(cd);
//			ps.println(ls);
//			ps.println(docker);
//			ps.println("echo 'acabe'");
//			ps.close();
//			
//			printResult(input, channel);
//			channel.disconnect();
//			session.disconnect();
//		}catch(Exception e) {
//			System.out.println(e);
//		}
//
//		return "hola";
//	}
//
//	private static void printResult(InputStream input, Channel channel) throws Exception {
//		int SIZE = 1024;
//		byte[] tmp = new byte[SIZE];
//		String line;
//		try(BufferedReader input2 = new BufferedReader(new InputStreamReader(input))) {
//		    while ((line = input2.readLine()) != null) {
//		    	line = line.trim();
//		    	if(line.equals("acabe")) {
//		    		break;
//		    	}
//		        System.out.println(line);
//		    }
//		    System.out.println("****"); 
//		} catch (IOException e) {
//		    e.printStackTrace();
//		}
//		/*
//		while (true) {
//			while (input.available() > 0) {
//				int i = input.read(tmp, 0, SIZE);
//				if (i < 0)
//					break;
//				
//				System.out.print(new String(tmp, 0, i));
//			}
//			if(input.available() < 0) {
//				break;
//			}
//			if (channel.isClosed()) {
//				System.out.println("exit-status: " + channel.getExitStatus());
//			}
//			try {
//				Thread.sleep(300);
//			} catch (Exception ee) {
//			}
//			System.out.println("salgooo");
//		}*/
//	}
//
//	public static void main(String[] args) {
//		runComando("prueba");
//	}
//
//	/**
//	 * //TODO: quiza haya que quitar el IDPersona //TODO: Ademas con el IDPERSONA la
//	 * query esta mal jaja public String comandoRegistrarPersona(String nombre,
//	 * String apellido1, String apellido2, String tipoDocumento, String
//	 * identificadorDocumento, String paisEmisor, int IDPersona) {
//	 * 
//	 * String comando = "'{\"function\":\"registrarPersona\"," +
//	 * "\"Args\":[\"{\\\"Nombre\\\":\\\"" + nombre + "\\\",\\\"Apellido1\\\":\\\"" +
//	 * apellido1 + "\\\",\\\"Apellido2\\\":\\\"" + apellido2 +
//	 * "\\\",\\\"TipoDocumento\\\":\\\"" + tipoDocumento +
//	 * "\\\",\\\"IdentificadorDocumento\\\":\\\"" + identificadorDocumento +
//	 * "\\\",\\\"PaisEmisor\\\":\\\"" + paisEmisor + "\\\"}\"," + " \"{\\\"" +
//	 * IDPersona + "\\\":6}\"]}'" + " -o $ORDENER_URL --tls --cafile " +
//	 * Constantes.BK_caFile + " -C " + Constantes.BK_canal; return
//	 * runComando(comando);
//	 * 
//	 * }
//	 * 
//	 * 
//	 * public String comandoRegistrarAfijo(int idPersona1, int idPersona2, int
//	 * idPersona3,) {
//	 * 
//	 * String comando = "'{\"function\":\"registrarPersona\"," +
//	 * "\"Args\":[\"{\\\"Nombre\\\":\\\"" + nombre + "\\\",\\\"Apellido1\\\":\\\"" +
//	 * apellido1 + "\\\",\\\"Apellido2\\\":\\\"" + apellido2 +
//	 * "\\\",\\\"TipoDocumento\\\":\\\"" + tipoDocumento +
//	 * "\\\",\\\"IdentificadorDocumento\\\":\\\"" + identificadorDocumento +
//	 * "\\\",\\\"PaisEmisor\\\":\\\"" + paisEmisor + "\\\"}\"," + " \"{\\\"" +
//	 * IDPersona + "\\\":6}\"]}'" + " -o $ORDENER_URL --tls --cafile " +
//	 * Constantes.BK_caFile + " -C " + Constantes.BK_canal; return
//	 * runComando(comando);
//	 * 
//	 * }
//	 */

}
