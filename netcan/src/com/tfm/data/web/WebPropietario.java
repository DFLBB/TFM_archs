package com.tfm.data.web;

import javax.xml.bind.annotation.XmlRootElement;

import com.tfm.data.perro.PerroLista;
import com.tfm.data.propietario.Propietario;

@XmlRootElement
public class WebPropietario {

	private Propietario propietario;
	private PerroLista listaPerros;

	public WebPropietario(Propietario propietario, PerroLista listaPerros) {
		this.propietario = propietario;
		this.listaPerros = listaPerros;
	}

	public WebPropietario(Propietario propietario) {
		this.propietario = propietario;
	}

	public WebPropietario() {

	}

	public Propietario getPropietario() {
		return propietario;
	}

	public void setPropietario(Propietario propietario) {
		this.propietario = propietario;
	}

	public PerroLista getListaPerros() {
		return listaPerros;
	}

	public void setListaPerros(PerroLista listaPerros) {
		this.listaPerros = listaPerros;
	}

}
