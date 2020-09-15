package com.tfm.data.concurso;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "concursos")
public class ConcursoLista {

	//
	private List<Concurso> lista;

	/**
	 * 
	 * @param lista
	 */
	public ConcursoLista(List<Concurso> lista) {
		this.lista = lista;
	}

	/**
	 * 
	 */
	public ConcursoLista() {

	}

	/**
	 * 
	 */
	@XmlElement(name = "concurso")
	public List<Concurso> getLista() {
		return lista;
	}

	/**
	 * 
	 */
	public void setLista(List<Concurso> lista) {
		this.lista = lista;
	}

}
