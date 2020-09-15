package com.tfm.data.raza;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Raza {

	private int id;
	private String nombre;

	public Raza(int id, String nombre) {
		this.id = id;
		this.nombre = nombre;
	}

	public Raza() {

	}
	
	public int getId() {
		return this.id;
	}
	
	public void setId(int id) {
		this.id = id;
	}
	
	public String getNombre() {
		return this.nombre;
	}
	
	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

}
