package com.tfm.data.participacion;

import java.util.Date;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Participacion {

	private String nombreConcurso;
	private Date fecha;
	private String posicion;
	private double premio;

	public Participacion(String nombreConcurso, Date fecha, String posicion, double premio) {
		this.nombreConcurso = nombreConcurso;
		this.fecha = fecha;
		this.posicion = posicion;
		this.premio = premio;
	}

	public Participacion() {

	}

	public String getNombreConcurso() {
		return nombreConcurso;
	}

	public void setNombreConcurso(String nombreConcurso) {
		this.nombreConcurso = nombreConcurso;
	}

	public Date getFecha() {
		return fecha;
	}

	public void setFecha(Date fecha) {
		this.fecha = fecha;
	}

	public String getPosicion() {
		return posicion;
	}

	public void setPosicion(String posicion) {
		this.posicion = posicion;
	}

	public double getPremio() {
		return premio;
	}

	public void setPremio(double premio) {
		this.premio = premio;
	}

}
