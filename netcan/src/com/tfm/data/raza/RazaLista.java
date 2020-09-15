package com.tfm.data.raza;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

import com.tfm.data.propietario.Propietario;

@XmlRootElement(name = "razas")
public class RazaLista {

	//
	private List<Raza> lista;

	/**
	 * 
	 * @param lista
	 */
	public RazaLista(List<Raza> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public RazaLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "raza")
	public List<Raza> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Raza> lista) {
		this.lista = lista;
	}
}
