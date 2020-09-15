package com.tfm.data.chip;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "chips")
public class ChipLista {

	//
	private List<Chip> lista;

	/**
	 * 
	 * @param lista
	 */
	public ChipLista(List<Chip> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public ChipLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "chip")
	public List<Chip> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Chip> lista) {
		this.lista = lista;
	}

}