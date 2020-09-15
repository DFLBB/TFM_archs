package com.tfm.data.federacion;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Federacion {

	private int id;
	private String nombre;
	private String pais;
	private int idAlta;

	public Federacion(int id, String nombre, String pais, int idAlta) {
		this.id = id;
		this.nombre = nombre;
		this.pais = pais;
		this.idAlta = idAlta;
	}

	public Federacion() {

	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public String getPais() {
		return pais;
	}

	public void setPais(String pais) {
		this.pais = pais;
	}

	public int getIdAlta() {
		return idAlta;
	}

	public void setIdAlta(int idAlta) {
		this.idAlta = idAlta;
	}

}
