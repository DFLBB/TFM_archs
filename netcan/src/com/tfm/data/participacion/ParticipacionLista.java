package com.tfm.data.participacion;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "participaciones")
public class ParticipacionLista {

	//
	private List<Participacion> lista;

	/**
	 * 
	 * @param lista
	 */
	public ParticipacionLista(List<Participacion> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public ParticipacionLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "participacion")
	public List<Participacion> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Participacion> lista) {
		this.lista = lista;
	}

}
