package com.tfm.data.afijo;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "afijos")
public class AfijoLista {

	//
	private List<Afijo> lista;

	/**
	 * 
	 * @param lista
	 */
	public AfijoLista(List<Afijo> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public AfijoLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "afijo")
	public List<Afijo> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Afijo> lista) {
		this.lista = lista;
	}

}
