package com.tfm.data.vacunacion;

import java.util.Date;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Vacunacion {

	private String perroNombre;
	private String afijo;
	private String vacuna;
	private Date fecha;

	public Vacunacion(String perroNombre, String afijo, String vacuna, Date fecha) {
		super();
		this.perroNombre = perroNombre;
		this.afijo = afijo;
		this.vacuna = vacuna;
		this.fecha = fecha;
	}

	public Vacunacion() {

	}

	public String getPerroNombre() {
		return perroNombre;
	}

	public void setPerroNombre(String perroNombre) {
		this.perroNombre = perroNombre;
	}

	public String getAfijo() {
		return afijo;
	}

	public void setAfijo(String afijo) {
		this.afijo = afijo;
	}

	public String getVacuna() {
		return vacuna;
	}

	public void setVacuna(String vacuna) {
		this.vacuna = vacuna;
	}

	public Date getFecha() {
		return fecha;
	}

	public void setFecha(Date fecha) {
		this.fecha = fecha;
	}

}
