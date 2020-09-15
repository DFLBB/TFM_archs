package com.tfm.data.vacuna;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;


@XmlRootElement(name = "vacunas")
public class VacunaLista {
	//
	private List<Vacuna> lista;

	/**
	 * 
	 * @param lista
	 */
	public VacunaLista(List<Vacuna> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public VacunaLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "vacuna")
	public List<Vacuna> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Vacuna> lista) {
		this.lista = lista;
	}

}
