package com.tfm.data.veterinario;

import java.util.List;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement(name = "veterinarios")
public class VeterinarioLista {
	

	private List<Veterinario> lista;

	public VeterinarioLista(List<Veterinario> lista) {
		this.lista = lista;
	}

	public VeterinarioLista() {

	}

	@XmlElement(name = "veterinario")
	public List<Veterinario> getLista() {
		return lista;
	}

	public void setLista(List<Veterinario> lista) {
		this.lista = lista;
	}

}
