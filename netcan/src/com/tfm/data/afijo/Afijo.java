package com.tfm.data.afijo;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Afijo {

	private String nombre;
	private int idPersona1;
	private int idPersona2;
	private int idPersona3;
	private int id;

	public Afijo(String nombre, int idPersona1, int idPersona2, int idPersona3) {
		this.nombre = nombre;
		this.idPersona1 = idPersona1;
		this.idPersona2 = idPersona2;
		this.idPersona3 = idPersona3;
	}

	public Afijo(String nombre, int id) {
		this.nombre = nombre;
		this.id = id;
	}

	public Afijo() {

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

	public int getIdPersona1() {
		return idPersona1;
	}

	public void setIdPersona1(int idPersona1) {
		this.idPersona1 = idPersona1;
	}

	public int getIdPersona2() {
		return idPersona2;
	}

	public void setIdPersona2(int idPersona2) {
		this.idPersona2 = idPersona2;
	}

	public int getIdPersona3() {
		return idPersona3;
	}

	public void setIdPersona3(int idPersona3) {
		this.idPersona3 = idPersona3;
	}

}
