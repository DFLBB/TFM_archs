package com.tfm.utils;

public class Constantes {

	/**
	 * Constantes de la conexion a la base de datos
	 */
	public static final String DB_RESOURCE = "jdbc/bbdd";
	public static final String DB_ENV = "java:comp/env";

	/**
	 * Mensajes de logs
	 */
	public static final String QUERY_ERROR = "Error en la ejecucion de la query ";
	public static final String WEBVETERINARIO_ERROR = "Error en el metodo webVeterinario ";
	public static final String WEBPROPIETARIO_ERROR = "Error en el metodo webPropietario ";

	/**
	 * Nombres de columnas de la base de datos
	 * 
	 */

	public static final String NOMBRE = "nombre";
	public static final String ID_CONCURSO = "idConcurso";
	public static final String FECHA = "fecha";
	public static final String ID_ALTA = "idAlta";

	/**
	 * Constantes para Blockchain
	 */

	public static final String BK_IP_CONN = "82.223.122.137";
	public static final String BK_USER_CONN = "root";
	//public static final String BK_PEM_CONN = "";
	public static final String BK_PASS = "*Password19!";
	public static final int BK_PORT_CONN = 22;

	//XXX hay que cambiarlo al del tls
	public static final String BK_CAFILE = "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/netcan.com/orderers/orderer.netcan.com/msp/tlscacerts/tlsca.netcan.com-cert.pem";
	public static final String BK_CANAL = "netcanchannel";
	public static final String BK_ORDERER = "orderer.netcan.com:7050";
	public static final String BK_SUDO = "sudo su";
	public static final String BK_SHELL = "shell";
	public static final String BK_END_COMMAND = "echo 'acabe'";

}
