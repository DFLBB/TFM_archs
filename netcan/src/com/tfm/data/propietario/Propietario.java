package com.tfm.data.propietario;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Propietario {

	private String documento;
	private String pais;
	private String tipoDoc;
	private String paisEmisorDoc;
	private String nombre;
	private String apellido1;
	private String apellido2;
	private String ciudad;
	private String direccion;
	private int idAlta;

	public Propietario(String documento, String pais, String tipoDoc, String nombre, String apellido1, String apellido2,
			String ciudad, String direccion, int idAlta, String paisEmisorDoc) {
		this.documento = documento;
		this.pais = pais;
		this.tipoDoc = tipoDoc;
		this.nombre = nombre;
		this.apellido1 = apellido1;
		this.apellido2 = apellido2;
		this.ciudad = ciudad;
		this.direccion = direccion;
		this.idAlta = idAlta;
		this.paisEmisorDoc = paisEmisorDoc;
	}

	public Propietario() {

	}

	public String getPaisEmisorDoc() {
		return paisEmisorDoc;
	}

	public void setPaisEmisorDoc(String paisEmisorDoc) {
		this.paisEmisorDoc = paisEmisorDoc;
	}

	public String getDocumento() {
		return documento;
	}

	public void setDocumento(String documento) {
		this.documento = documento;
	}

	public String getPais() {
		return pais;
	}

	public void setPais(String pais) {
		this.pais = pais;
	}

	public String getTipoDoc() {
		return tipoDoc;
	}

	public void setTipoDoc(String tipoDoc) {
		this.tipoDoc = tipoDoc;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public String getApellido1() {
		return apellido1;
	}

	public void setApellido1(String apellido1) {
		this.apellido1 = apellido1;
	}

	public String getApellido2() {
		return apellido2;
	}

	public void setApellido2(String apellido2) {
		this.apellido2 = apellido2;
	}

	public String getCiudad() {
		return ciudad;
	}

	public void setCiudad(String ciudad) {
		this.ciudad = ciudad;
	}

	public String getDireccion() {
		return direccion;
	}

	public void setDireccion(String direccion) {
		this.direccion = direccion;
	}

	public int getIdAlta() {
		return idAlta;
	}

	public void setIdAlta(int idAlta) {
		this.idAlta = idAlta;
	}

}
