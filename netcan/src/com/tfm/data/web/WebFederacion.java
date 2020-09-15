package com.tfm.data.web;

import javax.xml.bind.annotation.XmlRootElement;

import com.tfm.data.concurso.ConcursoLista;
import com.tfm.data.federacion.Federacion;

@XmlRootElement
public class WebFederacion {

	private Federacion federacion;
	private ConcursoLista lista;

	public WebFederacion(Federacion federacion, ConcursoLista lista) {
		this.federacion = federacion;
		this.lista = lista;
	}

	public WebFederacion(Federacion federacion) {
		this.federacion = federacion;
	}

	public WebFederacion() {
	}

	public Federacion getFederacion() {
		return federacion;
	}

	public void setFederacion(Federacion federacion) {
		this.federacion = federacion;
	}

	public ConcursoLista getLista() {
		return lista;
	}

	public void setLista(ConcursoLista lista) {
		this.lista = lista;
	}

}
