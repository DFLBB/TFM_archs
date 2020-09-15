package com.tfm.data.federacion;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "federaciones")
public class FederacionLista {
	//
	private List<Federacion> lista;

	/**
	 * 
	 * @param lista
	 */
	public FederacionLista(List<Federacion> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public FederacionLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "federacion")
	public List<Federacion> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Federacion> lista) {
		this.lista = lista;
	}
}
