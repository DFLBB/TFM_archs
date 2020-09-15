package com.tfm.data.web;

import javax.xml.bind.annotation.XmlRootElement;

import com.tfm.data.chip.ChipLista;
import com.tfm.data.vacunacion.VacunacionLista;
import com.tfm.data.veterinario.Veterinario;

@XmlRootElement
public class WebVeterinario {

	private Veterinario veterinario;
	private ChipLista listaChips;
	private VacunacionLista vacunaciones;

	public WebVeterinario(Veterinario veterinario, ChipLista listaChips, VacunacionLista vacunaciones) {
		this.veterinario = veterinario;
		this.listaChips = listaChips;
		this.vacunaciones = vacunaciones;
	}

	public WebVeterinario(Veterinario veterinario) {
		this.veterinario = veterinario;
	}

	public WebVeterinario() {

	}

	public Veterinario getVeterinario() {
		return veterinario;
	}

	public void setVeterinario(Veterinario veterinario) {
		this.veterinario = veterinario;
	}

	public ChipLista getListaChips() {
		return listaChips;
	}

	public void setListaChips(ChipLista listaChips) {
		this.listaChips = listaChips;
	}

	public VacunacionLista getVacunaciones() {
		return vacunaciones;
	}

	public void setVacunaciones(VacunacionLista vacunaciones) {
		this.vacunaciones = vacunaciones;
	}

}
