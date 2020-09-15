package com.tfm.data.perro;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "perros")
public class PerroLista {

	private List<Perro> lista;

	public PerroLista(List<Perro> lista) {
		this.lista = lista;
	}

	public PerroLista() {

	}

	public void setLista(List<Perro> lista) {
		this.lista = lista;
	}

	@XmlElement(name = "perro")
	public List<Perro> getLista() {
		return lista;
	}
}
