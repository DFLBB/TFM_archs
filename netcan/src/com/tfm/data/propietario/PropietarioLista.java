package com.tfm.data.propietario;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

/**
 * 
 * @author
 *
 */
@XmlRootElement(name = "propietarios")
public class PropietarioLista {

	//
	private List<Propietario> lista;

	/**
	 * 
	 * @param lista
	 */
	public PropietarioLista(List<Propietario> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public PropietarioLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "propietario")
	public List<Propietario> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Propietario> lista) {
		this.lista = lista;
	}

}
