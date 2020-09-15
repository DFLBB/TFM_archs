package com.tfm.data.chip;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Chip {

	private int idPerro;
	private int idVeterinario;
	private long codMicrochip;

	public Chip(int idPerro, int idVeterinario, long codMicrochip) {
		this.idPerro = idPerro;
		this.idVeterinario = idVeterinario;
		this.codMicrochip = codMicrochip;
	}

	public Chip() {
	}

	public int getIdPerro() {
		return idPerro;
	}

	public void setIdPerro(int idPerro) {
		this.idPerro = idPerro;
	}

	public int getIdVeterinario() {
		return idVeterinario;
	}

	public void setIdVeterinario(int idVeterinario) {
		this.idVeterinario = idVeterinario;
	}

	public long getCodMicrochip() {
		return codMicrochip;
	}

	public void setCodMicrochip(long codMicrochip) {
		this.codMicrochip = codMicrochip;
	}

}
