package com.tfm.data.concurso;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Concurso {

	private int id;
	private String nombre;
	private int federacion;
	private String patrocinadores;
	private String pais;

	public Concurso(int id, String nombre, int federacion, String patrocinadores, String pais) {
		this.id = id;
		this.nombre = nombre;
		this.federacion = federacion;
		this.patrocinadores = patrocinadores;
		this.pais = pais;
	}

	public Concurso(String nombre, int federacion, String patrocinadores, String pais) {
		this.nombre = nombre;
		this.federacion = federacion;
		this.pais = pais;
		this.patrocinadores = patrocinadores;
	}

	public Concurso() {

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

	public int getFederacion() {
		return federacion;
	}

	public void setFederacion(int federacion) {
		this.federacion = federacion;
	}

	public String getPatrocinadores() {
		return patrocinadores;
	}

	public void setPatrocinadores(String patrocinadores) {
		this.patrocinadores = patrocinadores;
	}

	public String getPais() {
		return pais;
	}

	public void setPais(String pais) {
		this.pais = pais;
	}

}
