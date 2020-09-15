package com.tfm.data.vacuna;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Vacuna {

	private int idPerro;
	private int idVeterinario;
	private long codVacuna;
	private String fechaBaja;
	private int idProteccion;
	private String fechaBajaProteccion;

	public Vacuna(int idPerro, int idVeterinario, long codVacuna, String fechaBaja, int idProteccion,
			String fechaBajaProteccion) {
		this.idPerro = idPerro;
		this.idVeterinario = idVeterinario;
		this.codVacuna = codVacuna;
		this.fechaBaja = fechaBaja;
		this.idProteccion = idProteccion;
		this.fechaBajaProteccion = fechaBajaProteccion;
	}

	public Vacuna() {
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

	public long getCodVacuna() {
		return codVacuna;
	}

	public void setCodVacuna(long codVacuna) {
		this.codVacuna = codVacuna;
	}

	public String getFechaBaja() {
		return fechaBaja;
	}

	public void setFechaBaja(String fechaBaja) {
		this.fechaBaja = fechaBaja;
	}

	public int getIdProteccion() {
		return idProteccion;
	}

	public void setIdProteccion(int idProteccion) {
		this.idProteccion = idProteccion;
	}

	public String getFechaBajaProteccion() {
		return fechaBajaProteccion;
	}

	public void setFechaBajaProteccion(String fechaBajaProteccion) {
		this.fechaBajaProteccion = fechaBajaProteccion;
	}

}
