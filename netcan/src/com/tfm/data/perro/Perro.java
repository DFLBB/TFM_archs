package com.tfm.data.perro;

import javax.xml.bind.annotation.XmlRootElement;

@XmlRootElement
public class Perro {

	private int id;
	private String nombre;
	private int sexo;
	private int idPadre;
	private int idMadre;
	private String fechaNacimiento;
	private String fechaDefuncion;
	private int idRaza;
	private String justificacion;

	public Perro(String nombre, int sexo, int idPadre, int idMadre, String fechaNacimiento) {
		this.nombre = nombre;
		this.sexo = sexo;
		this.idPadre = idPadre;
		this.idMadre = idMadre;
		this.fechaNacimiento = fechaNacimiento;
	}

	public Perro(int id, String fechaDefuncion) {
		this.id = id;
		this.fechaDefuncion = fechaDefuncion;
	}

	public Perro(int id, int idRaza, String justificacion) {
		this.id = id;
		this.idRaza = idRaza;
		this.justificacion = justificacion;
	}

	public Perro() {

	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getFechaDefuncion() {
		return fechaDefuncion;
	}

	public void setFechaDefuncion(String fechaDefuncion) {
		this.fechaDefuncion = fechaDefuncion;
	}

	public int getIdRaza() {
		return idRaza;
	}

	public void setIdRaza(int idRaza) {
		this.idRaza = idRaza;
	}

	public String getJustificacion() {
		return justificacion;
	}

	public void setJustificacion(String justificacion) {
		this.justificacion = justificacion;
	}

	public String getNombre() {
		return nombre;
	}

	public void setNombre(String nombre) {
		this.nombre = nombre;
	}

	public int getSexo() {
		return sexo;
	}

	public void setSexo(int sexo) {
		this.sexo = sexo;
	}

	public int getIdPadre() {
		return idPadre;
	}

	public void setIdPadre(int idPadre) {
		this.idPadre = idPadre;
	}

	public int getIdMadre() {
		return idMadre;
	}

	public void setIdMadre(int idMadre) {
		this.idMadre = idMadre;
	}

	public String getFechaNacimiento() {
		return fechaNacimiento;
	}

	public void setFechaNacimiento(String fechaNacimiento) {
		this.fechaNacimiento = fechaNacimiento;
	}

	// Codigo para web mas visible
	// pero no compatible con lo actual
	/**
	 * private int id; private String nombre; private String afijo; private double
	 * peso; private double altura; private String raza; private Chip chip; private
	 * String sexo; private Date fechaNacimiento; private Date fechaDefuncion;
	 * private int padre; private int madre; private String enlaceImagen; private
	 * VacunaLista vacunas; private ParticipacionLista concursos;
	 * 
	 * public Perro(int id, String nombre, String afijo, double peso, double altura,
	 * String raza, Chip chip, String sexo, Date fechaNacimiento, Date
	 * fechaDefuncion, int padre, int madre, String enlaceImagen) { this.id = id;
	 * this.nombre = nombre; this.afijo = afijo; this.peso = peso; this.altura =
	 * altura; this.raza = raza; this.chip = chip; this.sexo = sexo;
	 * this.fechaNacimiento = fechaNacimiento; this.fechaDefuncion = fechaDefuncion;
	 * this.padre = padre; this.madre = madre; this.enlaceImagen = enlaceImagen; }
	 * 
	 * public Perro(int id, String nombre, String afijo, double peso, double altura,
	 * String raza, Chip chip, String sexo, Date fechaNacimiento, Date
	 * fechaDefuncion, String enlaceImagen, VacunaLista vacunas, ParticipacionLista
	 * concursos) { this.id = id; this.nombre = nombre; this.afijo = afijo;
	 * this.peso = peso; this.altura = altura; this.raza = raza; this.chip = chip;
	 * this.sexo = sexo; this.fechaNacimiento = fechaNacimiento; this.fechaDefuncion
	 * = fechaDefuncion; this.enlaceImagen = enlaceImagen; this.vacunas = vacunas;
	 * this.concursos = concursos; }
	 * 
	 * public Perro() {
	 * 
	 * }
	 * 
	 * public ParticipacionLista getConcursos() { return concursos; }
	 * 
	 * public void setConcursos(ParticipacionLista concursos) { this.concursos =
	 * concursos; }
	 * 
	 * public VacunaLista getVacunas() { return vacunas; }
	 * 
	 * public void setVacunas(VacunaLista vacunas) { this.vacunas = vacunas; }
	 * 
	 * public int getId() { return id; }
	 * 
	 * public void setId(int id) { this.id = id; }
	 * 
	 * public String getNombre() { return nombre; }
	 * 
	 * public void setNombre(String nombre) { this.nombre = nombre; }
	 * 
	 * public String getAfijo() { return afijo; }
	 * 
	 * public void setAfijo(String afijo) { this.afijo = afijo; }
	 * 
	 * public double getPeso() { return peso; }
	 * 
	 * public void setPeso(double peso) { this.peso = peso; }
	 * 
	 * public double getAltura() { return altura; }
	 * 
	 * public void setAltura(double altura) { this.altura = altura; }
	 * 
	 * public String getRaza() { return raza; }
	 * 
	 * public void setRaza(String raza) { this.raza = raza; }
	 * 
	 * public Chip getChip() { return chip; }
	 * 
	 * public void setChip(Chip chip) { this.chip = chip; }
	 * 
	 * public String getSexo() { return sexo; }
	 * 
	 * public void setSexo(String sexo) { this.sexo = sexo; }
	 * 
	 * public Date getFechaNacimiento() { return fechaNacimiento; }
	 * 
	 * public void setFechaNacimiento(Date fechaNacimiento) { this.fechaNacimiento =
	 * fechaNacimiento; }
	 * 
	 * public Date getFechaDefuncion() { return fechaDefuncion; }
	 * 
	 * public void setFechaDefuncion(Date fechaDefuncion) { this.fechaDefuncion =
	 * fechaDefuncion; }
	 * 
	 * public int getPadre() { return padre; }
	 * 
	 * public void setPadre(int padre) { this.padre = padre; }
	 * 
	 * public int getMadre() { return madre; }
	 * 
	 * public void setMadre(int madre) { this.madre = madre; }
	 * 
	 * public String getEnlaceImagen() { return enlaceImagen; }
	 * 
	 * public void setEnlaceImagen(String enlaceImagen) { this.enlaceImagen =
	 * enlaceImagen; }
	 */
}
