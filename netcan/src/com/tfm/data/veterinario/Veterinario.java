package com.tfm.data.veterinario;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Veterinario {

	private int colegiado;
	private String nombre;
	private String apellido1;
	private String apellido2;
	private String clinicaVeterinaria;
	private String direccion;
	private String pais;
	private String ciudad;
	private String fechaBaja;
	private int idAlta;

	public Veterinario(int colegiado, String nombre, String apellido1, String apellido2, String clinicaVeterinaria,
			String direccion, String pais, String ciudad, String fechaBaja, int idAlta) {
		this.colegiado = colegiado;
		this.nombre = nombre;
		this.apellido1 = apellido1;
		this.apellido2 = apellido2;
		this.clinicaVeterinaria = clinicaVeterinaria;
		this.direccion = direccion;
		this.pais = pais;
		this.ciudad = ciudad;
		this.fechaBaja = fechaBaja;
		this.idAlta = idAlta;
	}

	public Veterinario(int colegiado, String nombre, String apellido1, String apellido2, String clinicaVeterinaria,
			String direccion, String pais, String ciudad, int idAlta) {
		this.colegiado = colegiado;
		this.nombre = nombre;
		this.apellido1 = apellido1;
		this.apellido2 = apellido2;
		this.clinicaVeterinaria = clinicaVeterinaria;
		this.direccion = direccion;
		this.pais = pais;
		this.ciudad = ciudad;
		this.idAlta = idAlta;
	}

	public Veterinario() {

	}

	public String getFechaBaja() {
		return fechaBaja;
	}

	public void setFechaBaja(String fechaBaja) {
		this.fechaBaja = fechaBaja;
	}

	public int getColegiado() {
		return colegiado;
	}

	public void setColegiado(int colegiado) {
		this.colegiado = colegiado;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public String getApellido1() {
		return apellido1;
	}

	public void setApellido1(String apellido1) {
		this.apellido1 = apellido1;
	}

	public String getApellido2() {
		return apellido2;
	}

	public void setApellido2(String apellido2) {
		this.apellido2 = apellido2;
	}

	public String getClinicaVeterinaria() {
		return clinicaVeterinaria;
	}

	public void setClinicaVeterinaria(String clinicaVeterinaria) {
		this.clinicaVeterinaria = clinicaVeterinaria;
	}

	public String getDireccion() {
		return direccion;
	}

	public void setDireccion(String direccion) {
		this.direccion = direccion;
	}

	public String getPais() {
		return pais;
	}

	public void setPais(String pais) {
		this.pais = pais;
	}

	public String getCiudad() {
		return ciudad;
	}

	public void setCiudad(String ciudad) {
		this.ciudad = ciudad;
	}

	public int getIdAlta() {
		return idAlta;
	}

	public void setIdAlta(int idAlta) {
		this.idAlta = idAlta;
	}

}
