package com.tfm.data.vacunacion;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "vacunaciones")
public class VacunacionLista {

	//
	private List<Vacunacion> lista;

	/**
	 * 
	 * @param lista
	 */
	public VacunacionLista(List<Vacunacion> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public VacunacionLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "vacunacion")
	public List<Vacunacion> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Vacunacion> lista) {
		this.lista = lista;
	}

}